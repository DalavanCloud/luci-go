// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package settings implements storage for infrequently changing global
// settings.
//
// Settings are represented as (key, value) pairs, where value is JSON
// serializable struct. Settings are cached internally in the process memory to
// avoid hitting the storage all the time.
package settings

import (
	"encoding/json"
	"errors"
	"reflect"
	"sync"
	"time"

	"golang.org/x/net/context"

	"github.com/luci/luci-go/common/lazyslot"
)

var (
	// ErrNoSettings can be returned by Get and Set on fatal errors.
	ErrNoSettings = errors.New("settings are not available")

	// ErrBadType is returned if Get(...) receives unexpected type.
	ErrBadType = errors.New("bad type")
)

// Bundle contains all latest settings along with the timestamp when they need
// to be refetched. Stored in lazyslot.Slot.
type Bundle struct {
	Values map[string]*json.RawMessage // immutable
	Exp    time.Time

	lock     sync.Mutex             // protects 'unpacked'
	unpacked map[string]interface{} // deserialized RawMessages
}

// get deserializes value for given key.
func (b *Bundle) get(key string, value interface{}) error {
	raw, ok := b.Values[key]
	if !ok || raw == nil || len(*raw) == 0 {
		return ErrNoSettings
	}

	typ := reflect.TypeOf(value)

	b.lock.Lock()
	defer b.lock.Unlock()

	cached, ok := b.unpacked[key]
	if !ok {
		// 'value' must be a pointer to a struct.
		if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
			return ErrBadType
		}
		// 'cached' is &Struct{}.
		cached = reflect.New(typ.Elem()).Interface()
		if err := json.Unmarshal([]byte(*raw), cached); err != nil {
			return err
		}
		if b.unpacked == nil {
			b.unpacked = make(map[string]interface{}, 1)
		}
		b.unpacked[key] = cached
	}

	// All calls to 'get' must use same type consistently.
	if reflect.TypeOf(cached) != typ {
		return ErrBadType
	}

	// 'value' and 'cached' are &Struct{}, Do *value = *cached.
	reflect.ValueOf(value).Elem().Set(reflect.ValueOf(cached).Elem())
	return nil
}

// Storage knows how to fetch settings from permanent storage and mutate
// them there. Methods of Storage can be called concurrently.
type Storage interface {
	// FetchAllSettings fetches all latest settings at once.
	FetchAllSettings(c context.Context) (*Bundle, error)

	// UpdateSetting updates a setting at the given key.
	UpdateSetting(c context.Context, key string, value json.RawMessage, who, why string) error
}

// Settings represent process global cache of all settings. Exact same instance
// of Settings should be injected into the context used by request handlers.
type Settings struct {
	storage Storage       // used to load and save settings
	values  lazyslot.Slot // cached settings
}

// New creates new Settings object that uses given Storage to fetch and save
// settings.
func New(storage Storage) *Settings {
	return &Settings{
		storage: storage,
		values: lazyslot.Slot{
			Fetcher: func(c context.Context, _ lazyslot.Value) (lazyslot.Value, error) {
				bundle, err := storage.FetchAllSettings(c)
				if err != nil {
					return lazyslot.Value{}, err
				}
				return lazyslot.Value{
					Value:      bundle,
					Expiration: bundle.Exp,
				}, nil
			},
		},
	}
}

// Get returns setting value for the given key. It will be deserialized into
// the supplied value. Caller is responsible to pass correct type and pass same
// type to all calls. If the setting is not set returns ErrNoSettings.
func (s *Settings) Get(c context.Context, key string, value interface{}) error {
	lazyValue, err := s.values.Get(c)
	if err != nil {
		return err
	}
	return lazyValue.Value.(*Bundle).get(key, value)
}

// Set changes a setting value for the given key. New settings will apply only
// when existing in-memory cache expires. In particular, Get() right after Set()
// may still return old value.
func (s *Settings) Set(c context.Context, key string, value interface{}, who, why string) error {
	blob, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.storage.UpdateSetting(c, key, json.RawMessage(blob), who, why)
}
