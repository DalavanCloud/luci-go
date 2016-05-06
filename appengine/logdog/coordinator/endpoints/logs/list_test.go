// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package logs

import (
	"fmt"
	"testing"

	ds "github.com/luci/gae/service/datastore"
	"github.com/luci/luci-go/appengine/logdog/coordinator"
	ct "github.com/luci/luci-go/appengine/logdog/coordinator/coordinatorTest"
	"github.com/luci/luci-go/appengine/logdog/coordinator/hierarchy"
	"github.com/luci/luci-go/common/api/logdog_coordinator/logs/v1"
	"github.com/luci/luci-go/common/config"
	"github.com/luci/luci-go/common/logdog/types"
	"github.com/luci/luci-go/common/proto/logdog/logpb"
	"golang.org/x/net/context"

	. "github.com/luci/luci-go/common/testing/assertions"
	. "github.com/smartystreets/goconvey/convey"
)

func listPaths(l *logdog.ListResponse) []string {
	s := make([]string, len(l.Components))
	for i, c := range l.Components {
		s[i] = c.Name

		var suffix string
		switch c.Type {
		case logdog.ListResponse_Component_STREAM:
			s[i] += "$"
		case logdog.ListResponse_Component_PATH:
			suffix = "/"
		case logdog.ListResponse_Component_PROJECT:
			suffix = "#"
		}
		s[i] += suffix
	}
	return s
}

func TestList(t *testing.T) {
	t.Parallel()

	Convey(`With a testing configuration, a List request`, t, func() {
		c, env := ct.Install()

		svc := New()

		const project = config.ProjectName("proj-foo")

		// Install a set of stock log streams to query against.
		streams := map[string]*coordinator.LogStream{}
		descs := map[string]*logpb.LogStreamDescriptor{}
		for i, v := range []types.StreamPath{
			"blargh/+/foo",
			"other/+/foo/bar",
			"other/+/baz",
			"testing/+/foo",
			"testing/+/foo/bar",
		} {
			// di is a datastore bound to the test project namespace.
			prefix, name := v.Split()
			desc := ct.TestLogStreamDescriptor(c, string(name))
			desc.Prefix = string(prefix)

			ls := ct.TestLogStream(c, desc)

			ct.WithProjectNamespace(c, project, func(c context.Context) {
				di := ds.Get(c)

				if err := di.Put(ls); err != nil {
					panic(fmt.Errorf("failed to put log stream %d: %v", i, err))
				}

				for _, c := range hierarchy.Components(ls.Path()) {
					if err := c.Put(di); err != nil {
						panic(fmt.Errorf("failed to put log component %d: %v", i, err))
					}
				}

				descs[string(v)] = desc
				streams[string(v)] = ls
			})
		}
		ds.Get(c).Testable().CatchupIndexes()

		var req logdog.ListRequest

		Convey(`A project-level list request`, func() {
			req.Project = ""

			// Only add project namespaces for some of our registered projects.
			//
			// We add a project namespace to our datastore by creating a single entity
			// within that namespace.
			addProjectNamespace := func(proj config.ProjectName) {
				c := c
				if err := coordinator.WithProjectNamespaceNoAuth(&c, proj); err != nil {
					panic(err)
				}

				entity := ds.PropertyMap{
					"$id":   []ds.Property{ds.MkProperty("woof")},
					"$kind": []ds.Property{ds.MkProperty("Dog")},
				}
				if err := ds.Get(c).Put(entity); err != nil {
					panic(err)
				}
			}
			addProjectNamespace("proj-foo")
			addProjectNamespace("proj-bar")
			addProjectNamespace("proj-exclusive")

			// Add this just to make sure we ignore namespaces without config entries.
			addProjectNamespace("proj-nonexistent")

			Convey(`Empty list request will return projects.`, func() {
				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{"proj-bar#", "proj-foo#"})
			})

			Convey(`Empty list request can skip projects.`, func() {
				req.Offset = 1

				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{"proj-foo#"})
			})

			Convey(`When authenticated, empty list includes authenticated project.`, func() {
				env.JoinGroup("auth")

				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{"proj-bar#", "proj-exclusive#", "proj-foo#"})
			})
		})

		Convey(`A project-bound list request`, func() {
			req.Project = string(project)

			Convey(`Can list within a project, when the project is empty.`, func() {
				req.PathBase = "proj-bar"

				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{})
			})

			Convey(`Empty list request will return top-level stream paths.`, func() {
				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{"blargh/", "other/", "testing/"})
			})

			Convey(`Will skip elements if requested.`, func() {
				req.Offset = 2

				l, err := svc.List(c, &req)
				So(err, ShouldBeRPCOK)
				So(listPaths(l), ShouldResemble, []string{"testing/"})
			})

			Convey(`Can list within a populated project.`, func() {
				req.Project = "proj-foo"

				list := func(b string) []string {
					req.PathBase = b

					l, err := svc.List(c, &req)
					So(err, ShouldBeRPCOK)
					return listPaths(l)
				}

				So(list(""), ShouldResemble, []string{"blargh/", "other/", "testing/"})
				So(list("blargh"), ShouldResemble, []string{"+/"})
				So(list("blargh/+"), ShouldResemble, []string{"foo$"})
				So(list("other"), ShouldResemble, []string{"+/"})
				So(list("other/+"), ShouldResemble, []string{"baz$", "foo/"})
				So(list("other/+/baz"), ShouldResemble, []string{})
				So(list("other/+/foo"), ShouldResemble, []string{"bar$"})
				So(list("testing"), ShouldResemble, []string{"+/"})
				So(list("testing/+"), ShouldResemble, []string{"foo$", "foo/"})
				So(list("testing/+/foo"), ShouldResemble, []string{"bar$"})
			})

			Convey(`Can use a cursor to fully enumerate a space.`, func() {
				req.MaxResults = 1

				all := []string{"blargh/", "other/", "testing/"}
				for len(all) > 0 {
					l, err := svc.List(c, &req)
					So(err, ShouldBeRPCOK)
					So(listPaths(l), ShouldResemble, all[:1])

					all = all[1:]
					req.Next = l.Next
				}
			})
		})

		Convey(`If the project does not exist, will return NotFound.`, func() {
			req.Project = "does-not-exist"

			_, err := svc.List(c, &req)
			So(err, ShouldBeRPCNotFound)
		})

	})
}
