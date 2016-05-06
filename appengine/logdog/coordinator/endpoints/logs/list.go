// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package logs

import (
	ds "github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/appengine/logdog/coordinator"
	"github.com/luci/luci-go/appengine/logdog/coordinator/config"
	"github.com/luci/luci-go/appengine/logdog/coordinator/hierarchy"
	"github.com/luci/luci-go/common/api/logdog_coordinator/logs/v1"
	"github.com/luci/luci-go/common/grpcutil"
	log "github.com/luci/luci-go/common/logging"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

const (
	// listResultLimit is the maximum number of log streams that will be
	// returned in a single query. If the user requests more, it will be
	// automatically truncated to this value.
	listResultLimit = 500
)

// List returns log stream paths rooted under the hierarchy.
func (s *server) List(c context.Context, req *logdog.ListRequest) (*logdog.ListResponse, error) {
	log.Fields{
		"project":       req.Project,
		"path":          req.PathBase,
		"offset":        req.Offset,
		"streamOnly":    req.StreamOnly,
		"includePurged": req.IncludePurged,
		"maxResults":    req.MaxResults,
	}.Debugf(c, "Received List request.")

	hr := hierarchy.Request{
		Project:    req.Project,
		PathBase:   req.PathBase,
		StreamOnly: req.StreamOnly,
		Limit:      s.limit(int(req.MaxResults), listResultLimit),
		Next:       req.Next,
		Skip:       int(req.Offset),
	}

	// Non-admin users may not request purged results.
	if req.IncludePurged {
		if err := coordinator.IsAdminUser(c); err != nil {
			log.Fields{
				log.ErrorKey: err,
			}.Errorf(c, "Non-superuser requested to see purged paths. Denying.")
			return nil, grpcutil.Errf(codes.PermissionDenied, "non-admin user cannot request purged log paths")
		}

		// TODO(dnj): Apply this to the hierarchy request, when purging is
		// enabled.
	}

	l, err := hierarchy.Get(c, hr)
	if err != nil {
		log.WithError(err).Errorf(c, "Failed to get hierarchy listing.")
		if err == config.ErrNoAccess {
			// User requested a project that either doesn't exist or they don't have
			// access to.
			return nil, grpcutil.NotFound
		}
		return nil, grpcutil.InvalidArgument
	}

	resp := logdog.ListResponse{
		Project:  string(l.Project),
		Next:     l.Next,
		PathBase: string(l.PathBase),
	}

	if len(l.Comp) > 0 {
		resp.Components = make([]*logdog.ListResponse_Component, len(l.Comp))

		for i, c := range l.Comp {
			comp := logdog.ListResponse_Component{
				Name: c.Name,
			}
			switch {
			case l.Project == "":
				comp.Type = logdog.ListResponse_Component_PROJECT
			case c.Stream:
				comp.Type = logdog.ListResponse_Component_STREAM
			default:
				comp.Type = logdog.ListResponse_Component_PATH
			}

			resp.Components[i] = &comp
		}
	}

	// Perform additional stream metadata fetch if state is requested. Collect
	// a list of streams to load.
	if req.State && l.Project != "" {
		c := c
		if err := coordinator.WithProjectNamespace(&c, l.Project); err != nil {
			// This should work, since the list would have rejected the namespace if
			// the user was not a member, so a failure here is an internal error.
			log.Fields{
				log.ErrorKey: err,
				"project":    l.Project,
			}.Errorf(c, "Failed to enter namespace for metadata lookup.")
			return nil, grpcutil.Internal
		}

		idxMap := make(map[int]*logdog.ListResponse_Component)
		var streams []*coordinator.LogStream

		for i, comp := range l.Comp {
			if !comp.Stream {
				continue
			}

			idxMap[len(streams)] = resp.Components[i]
			log.Fields{
				"value": l.Path(comp),
			}.Infof(c, "Loading stream.")
			streams = append(streams, coordinator.LogStreamFromPath(l.Path(comp)))
		}

		if len(streams) > 0 {
			if err := ds.Get(c).GetMulti(streams); err != nil {
				log.Fields{
					log.ErrorKey: err,
					"count":      len(streams),
				}.Errorf(c, "Failed to load stream descriptors.")
				return nil, grpcutil.Internal
			}

			for sidx, lrs := range idxMap {
				ls := streams[sidx]
				lrs.State = loadLogStreamState(ls)

				lrs.Desc, err = ls.DescriptorValue()
				if err != nil {
					log.Fields{
						log.ErrorKey: err,
						"path":       ls.Path(),
					}.Errorf(c, "Failed to unmarshal descriptor protobuf.")
					return nil, grpcutil.Internal
				}
			}
		}
	}

	log.Fields{
		"count": len(resp.Components),
	}.Infof(c, "List completed successfully.")
	return &resp, nil
}
