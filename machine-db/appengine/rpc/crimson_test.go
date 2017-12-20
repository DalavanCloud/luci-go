// Copyright 2017 The LUCI Authors.
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

package rpc

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"

	"github.com/DATA-DOG/go-sqlmock"

	"go.chromium.org/luci/common/data/stringset"

	"go.chromium.org/luci/machine-db/api/crimson/v1"
	"go.chromium.org/luci/machine-db/appengine/database"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestGetDatacenters(t *testing.T) {
	Convey("getDatacenters", t, func() {
		db, m, _ := sqlmock.New()
		defer db.Close()
		c := database.With(context.Background(), db)
		selectStmt := `
			^SELECT d.name, d.description
			FROM datacenters d$
		`
		columns := []string{"d.name", "d.description"}
		rows := sqlmock.NewRows(columns)

		Convey("query failed", func() {
			names := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnError(fmt.Errorf("error"))
			dcs, err := getDatacenters(c, names)
			So(err, ShouldErrLike, "failed to fetch datacenters")
			So(dcs, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("empty", func() {
			names := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			dcs, err := getDatacenters(c, names)
			So(err, ShouldBeNil)
			So(dcs, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("no matches", func() {
			names := stringset.NewFromSlice("datacenter")
			rows.AddRow("datacenter 1", "description 1")
			rows.AddRow("datacenter 2", "description 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			dcs, err := getDatacenters(c, names)
			So(err, ShouldBeNil)
			So(dcs, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("matches", func() {
			names := stringset.NewFromSlice("datacenter 2", "datacenter 3")
			rows.AddRow("datacenter 1", "description 1")
			rows.AddRow("datacenter 2", "description 2")
			rows.AddRow("datacenter 3", "description 3")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			dcs, err := getDatacenters(c, names)
			So(err, ShouldBeNil)
			So(dcs, ShouldResemble, []*crimson.Datacenter{
				{
					Name:        "datacenter 2",
					Description: "description 2",
				},
				{
					Name:        "datacenter 3",
					Description: "description 3",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("ok", func() {
			names := stringset.New(0)
			rows.AddRow("datacenter 1", "description 1")
			rows.AddRow("datacenter 2", "description 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			dcs, err := getDatacenters(c, names)
			So(err, ShouldBeNil)
			So(dcs, ShouldResemble, []*crimson.Datacenter{
				{
					Name:        "datacenter 1",
					Description: "description 1",
				},
				{
					Name:        "datacenter 2",
					Description: "description 2",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func TestGetRacks(t *testing.T) {
	Convey("getRacks", t, func() {
		db, m, _ := sqlmock.New()
		defer db.Close()
		c := database.With(context.Background(), db)
		selectStmt := `
			^SELECT r.name, r.description, d.name
			FROM racks r, datacenters d
			WHERE r.datacenter_id = d.id$
		`
		columns := []string{"r.name", "r.description", "d.name"}
		rows := sqlmock.NewRows(columns)

		Convey("query failed", func() {
			names := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnError(fmt.Errorf("error"))
			racks, err := getRacks(c, names, dcs)
			So(err, ShouldErrLike, "failed to fetch racks")
			So(racks, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("empty", func() {
			names := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			racks, err := getRacks(c, names, dcs)
			So(err, ShouldBeNil)
			So(racks, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("no matches", func() {
			names := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			rows.AddRow("rack 1", "description 1", "datacenter 1")
			rows.AddRow("rack 2", "description 2", "datacenter 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			racks, err := getRacks(c, names, dcs)
			So(err, ShouldBeNil)
			So(racks, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("matches", func() {
			names := stringset.NewFromSlice("rack 2", "rack 3")
			dcs := stringset.NewFromSlice("datacenter 2", "datacenter 3")
			rows.AddRow("rack 1", "description 1", "datacenter 1")
			rows.AddRow("rack 2", "description 2", "datacenter 2")
			rows.AddRow("rack 3", "description 3", "datacenter 3")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			racks, err := getRacks(c, names, dcs)
			So(err, ShouldBeNil)
			So(racks, ShouldResemble, []*crimson.Rack{
				{
					Name:        "rack 2",
					Description: "description 2",
					Datacenter:  "datacenter 2",
				},
				{
					Name:        "rack 3",
					Description: "description 3",
					Datacenter:  "datacenter 3",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("ok", func() {
			names := stringset.New(0)
			dcs := stringset.New(0)
			rows.AddRow("rack 1", "description 1", "datacenter 1")
			rows.AddRow("rack 2", "description 2", "datacenter 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			racks, err := getRacks(c, names, dcs)
			So(err, ShouldBeNil)
			So(racks, ShouldResemble, []*crimson.Rack{
				{
					Name:        "rack 1",
					Description: "description 1",
					Datacenter:  "datacenter 1",
				},
				{
					Name:        "rack 2",
					Description: "description 2",
					Datacenter:  "datacenter 2",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}

func TestGetSwitches(t *testing.T) {
	Convey("getSwitches", t, func() {
		db, m, _ := sqlmock.New()
		defer db.Close()
		c := database.With(context.Background(), db)
		selectStmt := `
			^SELECT s.name, s.description, s.ports, r.name, d.name
			FROM switches s, racks r, datacenters d
			WHERE s.rack_id = r.id AND r.datacenter_id = d.id$
		`
		columns := []string{"s.name", "s.description", "s.ports", "r.name", "d.name"}
		rows := sqlmock.NewRows(columns)

		Convey("query failed", func() {
			names := stringset.NewFromSlice("switch")
			racks := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnError(fmt.Errorf("error"))
			switches, err := getSwitches(c, names, racks, dcs)
			So(err, ShouldErrLike, "failed to fetch switches")
			So(switches, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("empty", func() {
			names := stringset.NewFromSlice("switch")
			racks := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			switches, err := getSwitches(c, names, racks, dcs)
			So(err, ShouldBeNil)
			So(switches, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("no matches", func() {
			names := stringset.NewFromSlice("switch")
			racks := stringset.NewFromSlice("rack")
			dcs := stringset.NewFromSlice("datacenter")
			rows.AddRow("switch 1", "description 1", 10, "rack 1", "datacenter 1")
			rows.AddRow("switch 2", "description 2", 20, "rack 2", "datacenter 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			switches, err := getSwitches(c, names, racks, dcs)
			So(err, ShouldBeNil)
			So(switches, ShouldBeEmpty)
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("matches", func() {
			names := stringset.NewFromSlice("switch 3", "switch 4")
			racks := stringset.NewFromSlice("rack 3", "rack 4")
			dcs := stringset.NewFromSlice("datacenter 2", "datacenter 3", "datacenter 4")
			rows.AddRow("switch 1", "description 1", 10, "rack 1", "datacenter 1")
			rows.AddRow("switch 2", "description 2", 20, "rack 2", "datacenter 2")
			rows.AddRow("switch 3", "description 3", 30, "rack 3", "datacenter 3")
			rows.AddRow("switch 4", "description 4", 40, "rack 4", "datacenter 4")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			switches, err := getSwitches(c, names, racks, dcs)
			So(err, ShouldBeNil)
			So(switches, ShouldResemble, []*crimson.Switch{
				{
					Name:        "switch 3",
					Description: "description 3",
					Ports:       30,
					Rack:        "rack 3",
					Datacenter:  "datacenter 3",
				},
				{
					Name:        "switch 4",
					Description: "description 4",
					Ports:       40,
					Rack:        "rack 4",
					Datacenter:  "datacenter 4",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})

		Convey("ok", func() {
			names := stringset.New(0)
			racks := stringset.New(0)
			dcs := stringset.New(0)
			rows.AddRow("switch 1", "description 1", 10, "rack 1", "datacenter 1")
			rows.AddRow("switch 2", "description 2", 20, "rack 2", "datacenter 2")
			m.ExpectQuery(selectStmt).WillReturnRows(rows)
			switches, err := getSwitches(c, names, racks, dcs)
			So(err, ShouldBeNil)
			So(switches, ShouldResemble, []*crimson.Switch{
				{
					Name:        "switch 1",
					Description: "description 1",
					Ports:       10,
					Rack:        "rack 1",
					Datacenter:  "datacenter 1",
				},
				{
					Name:        "switch 2",
					Description: "description 2",
					Ports:       20,
					Rack:        "rack 2",
					Datacenter:  "datacenter 2",
				},
			})
			So(m.ExpectationsWereMet(), ShouldBeNil)
		})
	})
}