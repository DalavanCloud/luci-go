// Copyright 2015 The LUCI Authors.
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

package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/common/gcloud/googleoauth"

	"golang.org/x/net/context"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

const testScope = "https://example.com/scopes/user.email"

type tokenInfo struct {
	Audience      string `json:"aud"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Error         string `json:"error_description"`
	ExpiresIn     string `json:"expires_in"`
	Scope         string `json:"scope"`
}

func TestGoogleOAuth2Method(t *testing.T) {
	t.Parallel()

	Convey("with mock backend", t, func(c C) {
		ctx := context.Background()

		info := tokenInfo{
			Audience:      "client_id",
			Email:         "abc@example.com",
			EmailVerified: "true",
			ExpiresIn:     "3600",
			Scope:         testScope + " other stuff",
		}
		status := http.StatusOK
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(status)
			c.So(json.NewEncoder(w).Encode(&info), ShouldBeNil)
		}))

		ctx = ModifyConfig(ctx, func(cfg Config) Config {
			cfg.Cache = &MemoryCache{LRU: lru.New(0)}
			cfg.AnonymousTransport = func(context.Context) http.RoundTripper {
				return http.DefaultTransport
			}
			return cfg
		})

		call := func(header string) (*User, error) {
			m := GoogleOAuth2Method{
				Scopes:            []string{testScope},
				tokenInfoEndpoint: ts.URL,
			}
			req, err := http.NewRequest("GET", "http://fake", nil)
			So(err, ShouldBeNil)
			req.Header.Set("Authorization", header)
			return m.Authenticate(ctx, req)
		}

		Convey("Works", func() {
			u, err := call("Bearer access_token")
			So(err, ShouldBeNil)
			So(u, ShouldResemble, &User{
				Identity: "user:abc@example.com",
				Email:    "abc@example.com",
				ClientID: "client_id",
			})
		})

		Convey("Bad header", func() {
			_, err := call("broken")
			So(err, ShouldErrLike, "oauth: bad Authorization header")
		})

		Convey("HTTP 500", func() {
			status = http.StatusInternalServerError
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "transient error")
		})

		Convey("Error response", func() {
			status = http.StatusBadRequest
			info.Error = "OMG, error"
			_, err := call("Bearer access_token")
			So(err, ShouldEqual, googleoauth.ErrBadToken)
		})

		Convey("No email", func() {
			info.Email = ""
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "not associated with an email")
		})

		Convey("Email not verified", func() {
			info.EmailVerified = "false"
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "not verified")
		})

		Convey("Bad expires_in", func() {
			info.ExpiresIn = "not a number"
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "json: invalid")
		})

		Convey("Zero expires_in", func() {
			info.ExpiresIn = "0"
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "not a positive integer")
		})

		Convey("Missing scope", func() {
			info.Scope = "some other scopes"
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "doesn't have scope")
		})

		Convey("Bad email", func() {
			info.Email = "@@@@"
			_, err := call("Bearer access_token")
			So(err, ShouldErrLike, "bad value")
		})
	})
}