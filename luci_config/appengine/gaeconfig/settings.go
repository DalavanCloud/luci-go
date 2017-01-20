// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package gaeconfig

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/luci-go/server/settings"
)

// DefaultExpire is a reasonable default expiration value.
const DefaultExpire = 10 * time.Minute

// DSCacheMode is the datastore cache mode.
type DSCacheMode string

const (
	// DSCacheDisabled means that the datastore cache is disabled.
	DSCacheDisabled DSCacheMode = ""
	// DSCacheEnabled means that the datastore cache is enabled.
	DSCacheEnabled DSCacheMode = "Enabled"
	// DSCacheStrict means that the datastore cache is enabled, and that a stale
	// datastore cache entry should return an error rather than fall through to
	// the real config service.
	DSCacheStrict DSCacheMode = "Strict"
)

// dsCacheDisabledSetting is the user-visible value for the "disabled" cache
// mode.
const dsCacheDisabledSetting = "Disabled"

// Settings are stored in the datastore via appengine/gaesettings package.
type Settings struct {
	// ConfigServiceURL is URL of luci-config service to fetch configs from.
	ConfigServiceURL string `json:"config_service_url"`

	// CacheExpirationSec is how long to hold configs in local cache.
	CacheExpirationSec int `json:"cache_expiration_sec"`

	// DatastoreCacheMode, is the datastore caching mode.
	DatastoreCacheMode DSCacheMode `json:"datastore_enabled"`
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
		return s, nil
	case settings.ErrNoSettings:
		return DefaultSettings(), nil
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
func DefaultSettings() Settings {
	return Settings{
		CacheExpirationSec: int(DefaultExpire.Seconds()),
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
			ID:    "ConfigServiceURL",
			Title: "Config service URL",
			Type:  settings.UIFieldText,
			Validator: func(v string) error {
				if v != "" {
					parsed, err := url.Parse(v)
					if err != nil {
						return fmt.Errorf("bad URL %q - %s", v, err)
					}
					if !parsed.IsAbs() || parsed.Path != "" {
						return fmt.Errorf("bad URL %q - must be host root URL", v)
					}
					if parsed.Scheme != "https" {
						return fmt.Errorf("bad URL %q - expecting https:// scheme", v)
					}
				}
				return nil
			},
			Help: `<p>The application may fetch configuration files stored centrally
in an instance of <a href="https://github.com/luci/luci-py/tree/master/appengine/config_service">luci-config</a>
service. This is an URL of such service. If you don't know what this is, you
probably don't use it and can keep this setting blank.</p>`,
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
				string(DSCacheStrict),
			},
			Help: `<p>For better performance and resilience against configuration
service outages, the local datastore can be used as a backing cache. When
enabled, an additional <strong>Strict</strong> mode is available where expired
datastore entries will error instead of fail-open to the config service.</p>`,
		},
	}, nil
}

func (settingsUIPage) ReadSettings(c context.Context) (map[string]string, error) {
	s := DefaultSettings()
	err := settings.GetUncached(c, settingsKey, &s)
	if err != nil && err != settings.ErrNoSettings {
		return nil, err
	}

	cacheMode := string(s.DatastoreCacheMode)
	if cacheMode == string(DSCacheDisabled) {
		cacheMode = dsCacheDisabledSetting
	}

	return map[string]string{
		"ConfigServiceURL":   s.ConfigServiceURL,
		"CacheExpirationSec": strconv.Itoa(s.CacheExpirationSec),
		"DatastoreCacheMode": cacheMode,
	}, nil
}

func (settingsUIPage) WriteSettings(c context.Context, values map[string]string, who, why string) error {
	dsMode := DSCacheMode(values["DatastoreCacheMode"])
	if dsMode == dsCacheDisabledSetting {
		dsMode = DSCacheDisabled
	}

	modified := Settings{
		ConfigServiceURL:   values["ConfigServiceURL"],
		DatastoreCacheMode: dsMode,
	}

	var err error
	modified.CacheExpirationSec, err = strconv.Atoi(values["CacheExpirationSec"])
	if err != nil {
		return err
	}

	return settings.SetIfChanged(c, settingsKey, &modified, who, why)
}

func init() {
	settings.RegisterUIPage(settingsKey, settingsUIPage{})
}