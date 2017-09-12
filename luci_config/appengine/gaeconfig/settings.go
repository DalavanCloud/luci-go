// Copyright 2016 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gaeconfig

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/info"
	"go.chromium.org/luci/server/settings"
)

// DefaultExpire is a reasonable default expiration value to use on prod GAE.
const DefaultExpire = 10 * time.Minute

// DSCacheMode is the datastore cache mode.
type DSCacheMode string

const (
	// DSCacheDisabled means that the datastore cache is disabled.
	DSCacheDisabled DSCacheMode = ""
	// DSCacheEnabled means that the datastore cache is enabled.
	//
	// When enabled, all requests that have cached entries will hit the cache,
	// regardless of whether the cache is stale or not. If the cache is stale,
	// a warning will be printed during fetch.
	DSCacheEnabled DSCacheMode = "Enabled"
)

// dsCacheDisabledSetting is the user-visible value for the "disabled" cache
// mode.
const dsCacheDisabledSetting = "Disabled"

// Settings are stored in the datastore via appengine/gaesettings package.
type Settings struct {
	// ConfigServiceHost is host name (and port) of the luci-config service to
	// fetch configs from.
	//
	// For legacy reasons, the JSON value is "config_service_url".
	ConfigServiceHost string `json:"config_service_url"`

	// CacheExpirationSec is how long to hold configs in local cache.
	CacheExpirationSec int `json:"cache_expiration_sec"`

	// DatastoreCacheMode, is the datastore caching mode.
	DatastoreCacheMode DSCacheMode `json:"datastore_enabled"`
}

// SetIfChanged sets "s" to be the new Settings if it differs from the current
// settings value.
func (s *Settings) SetIfChanged(c context.Context, who, why string) error {
	return settings.SetIfChanged(c, settingsKey, s, who, why)
}

// CacheExpiration returns a Duration representing the configured cache
// expiration. If the cache expiration is not configured, CacheExpiration
// will return 0.
func (s *Settings) CacheExpiration() time.Duration {
	if s.CacheExpirationSec > 0 {
		return time.Second * time.Duration(s.CacheExpirationSec)
	}
	return 0
}

// FetchCachedSettings fetches Settings from the settings store.
//
// Uses in-process global cache to avoid hitting datastore often. The cache
// expiration time is 1 min (see gaesettings.expirationTime), meaning
// the instance will refetch settings once a minute (blocking only one unlucky
// request to do so).
//
// Returns errors only if there's no cached value (i.e. it is the first call
// to this function in this process ever) and datastore operation fails.
func FetchCachedSettings(c context.Context) (Settings, error) {
	s := Settings{}
	switch err := settings.Get(c, settingsKey, &s); err {
	case nil:
		// Backwards-compatibility with full URL: translate to host.
		s.ConfigServiceHost = translateConfigURLToHost(s.ConfigServiceHost)
		return s, nil
	case settings.ErrNoSettings:
		return DefaultSettings(c), nil
	default:
		return Settings{}, err
	}
}

func mustFetchCachedSettings(c context.Context) *Settings {
	settings, err := FetchCachedSettings(c)
	if err != nil {
		panic(err)
	}
	return &settings
}

// DefaultSettings returns Settings to use if setting store is empty.
func DefaultSettings(c context.Context) Settings {
	// Disable local cache on devserver by default to allows changes to local
	// configs to propagate instantly. This is usually preferred when developing
	// locally.
	exp := 0
	if !info.IsDevAppServer(c) {
		exp = int(DefaultExpire.Seconds())
	}
	return Settings{
		CacheExpirationSec: exp,
		DatastoreCacheMode: DSCacheDisabled,
	}
}

////////////////////////////////////////////////////////////////////////////////
// UI for settings.

// settingsKey is used internally to identify gaeconfig settings in settings
// store.
const settingsKey = "gaeconfig"

type settingsUIPage struct {
	settings.BaseUIPage
}

func (settingsUIPage) Title(c context.Context) (string, error) {
	return "Configuration service settings", nil
}

func (settingsUIPage) Fields(c context.Context) ([]settings.UIField, error) {
	return []settings.UIField{
		{
			ID:    "ConfigServiceHost",
			Title: `Config service host`,
			Type:  settings.UIFieldText,
			Validator: func(v string) error {
				if strings.ContainsRune(v, '/') {
					return fmt.Errorf("host must be a host name, not a URL")
				}
				return nil
			},
			Help: `<p>The application may fetch configuration files stored centrally ` +
				`in an instance of <a href="https://chromium.googlesource.com/infra/luci/luci-py/+/master/appengine/config_service">luci-config</a> ` +
				`service. This is the host name (e.g., "example.com") of such service. For legacy purposes, this may be an ` +
				`URL, in which case the host component will be used. If you don't know what this is, you probably don't ` +
				`use it and can keep this setting blank.</p>`,
		},
		{
			ID:    "CacheExpirationSec",
			Title: "Cache expiration, sec",
			Type:  settings.UIFieldText,
			Validator: func(v string) error {
				if i, err := strconv.Atoi(v); err != nil || i < 0 {
					return errors.New("expecting a non-negative integer")
				}
				return nil
			},
			Help: `<p>For better performance configuration files fetched from remote
service are cached in memcache for specified amount of time. Set it to 0 to
disable local cache.</p>`,
		},
		{
			ID:    "DatastoreCacheMode",
			Title: "Enable datastore-backed config caching",
			Type:  settings.UIFieldChoice,
			ChoiceVariants: []string{
				dsCacheDisabledSetting,
				string(DSCacheEnabled),
			},
			Help: `<p>For better performance and resilience against configuration
service outages, the local datastore can be used as a backing cache. When
enabled, all configuration requests will be made against a cached configuration
in the datastore. This configuration will be updated periodically by an
independent cron job out of band with any user requests. See
<a href="https://godoc.org/go.chromium.org/luci/appengine/gaemiddleware/#hdr-Cron_setup">gaemiddleware</a>
package doc for instructions how to setup this cron job.</p>`,
		},
	}, nil
}

func (settingsUIPage) ReadSettings(c context.Context) (map[string]string, error) {
	s := DefaultSettings(c)
	err := settings.GetUncached(c, settingsKey, &s)
	if err != nil && err != settings.ErrNoSettings {
		return nil, err
	}

	// Translate the DSCacheMode into a user-readable string.
	var cacheModeString string
	switch s.DatastoreCacheMode {
	// Recognized modes.
	case DSCacheEnabled:
		cacheModeString = string(s.DatastoreCacheMode)

		// Any unrecognized mode translates to "disabled".
	case DSCacheDisabled:
		fallthrough
	default:
		cacheModeString = dsCacheDisabledSetting
	}

	return map[string]string{
		"ConfigServiceHost":  s.ConfigServiceHost,
		"CacheExpirationSec": strconv.Itoa(s.CacheExpirationSec),
		"DatastoreCacheMode": cacheModeString,
	}, nil
}

func (settingsUIPage) WriteSettings(c context.Context, values map[string]string, who, why string) error {
	dsMode := DSCacheMode(values["DatastoreCacheMode"])
	switch dsMode {
	case DSCacheEnabled: // Valid.

	// Any unrecognized mode translates to disabled.
	case dsCacheDisabledSetting:
		fallthrough
	default:
		dsMode = DSCacheDisabled
	}

	modified := Settings{
		ConfigServiceHost:  values["ConfigServiceHost"],
		DatastoreCacheMode: dsMode,
	}

	var err error
	modified.CacheExpirationSec, err = strconv.Atoi(values["CacheExpirationSec"])
	if err != nil {
		return err
	}

	// Backwards-compatibility with full URL: translate to host.
	modified.ConfigServiceHost = translateConfigURLToHost(modified.ConfigServiceHost)

	return modified.SetIfChanged(c, who, why)
}

func translateConfigURLToHost(v string) string {
	// If the host is a full URL, extract just the host component.
	switch u, err := url.Parse(v); {
	case err != nil:
		return v
	case u.Host != "":
		// If we have a host (e.g., "example.com"), this will parse into the "Path"
		// field with an empty host value. Therefore, if we have a "Host" value,
		// we will use it directly (e.g., "http://example.com")
		return u.Host
	case u.Path != "":
		// If this was just an empty (correct) host, it will have parsed into the
		// Path field with an empty Host value.
		return u.Path
	default:
		return v
	}
}

func init() {
	settings.RegisterUIPage(settingsKey, settingsUIPage{})
}
