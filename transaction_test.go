/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

/*
import (
	"testing"
	"time"

	"github.com/ernestio/service-store/models"
	"github.com/ernestio/service-store/tests"
	"github.com/nats-io/nats"
	. "github.com/smartystreets/goconvey/convey"
	graph "gopkg.in/r3labs/graph.v2"
)

func TestSetComponentHandler(t *testing.T) {
	_ = tests.CreateTestDB("test_transactions")

	setupNats()
	defer n.Close()

	_, _ = n.Subscribe("config.get.postgres", func(msg *nats.Msg) {
		_ = n.Publish(msg.Reply, []byte(`{"names":["services"],"password":"","url":"postgres://postgres@127.0.0.1","user":""}`))
	})

	setupPg("test_transactions")
	db.AutoMigrate(models.Environment{}, models.Build{})

	startHandler()

	db.Unscoped().Delete(models.Environment{}, models.Build{})
	CreateTestData(db, 20)

	Convey("Scenario: creating a service build", t, func() {
		Convey("When receiving two events that create a build using the same service name", func() {
			_ = n.Publish("service.set", []byte(`{"name": "Test1", "id": "uuid-98", "options":{"sync":false}}`))
			resp, err := n.Request("service.set", []byte(`{"name": "Test1", "id": "uuid-99", "options":{"sync":false}}`), time.Second)
			So(err, ShouldBeNil)
			So(string(resp.Data), ShouldEqual, `{"error": "could not create environment build: service in progress"}`)
		})
	})

	Convey("Scenario: setting multiple components on a service concurrently", t, func() {
		Convey("When receiving two events that update the same service mapping", func() {
			id := "uuid-1"

			_ = n.Publish("service.set.mapping.component", []byte(`{"_component_id":"network::test-1", "service":"`+id+`", "_state": "completed"}`))
			_, err := n.Request("service.set.mapping.component", []byte(`{"_component_id":"network::test-2", "service":"`+id+`", "_state": "completed"}`), time.Second)
			So(err, ShouldBeNil)

			Convey("It should update both the components", func() {
				var b models.Build
				db.Where("uuid = ?", id).First(&b)

				g := graph.New()
				So(g.Load(b.Mapping), ShouldBeNil)

				c1 := g.Component("network::test-1")
				So(c1, ShouldNotBeNil)
				So(c1.GetState(), ShouldEqual, "completed")

				c2 := g.Component("network::test-2")
				So(c2, ShouldNotBeNil)
				So(c2.GetState(), ShouldEqual, "completed")
			})
		})

	})

	Convey("Scenario: setting multiple changes on a service concurrently", t, func() {
		Convey("When receiving two events that update the same service mapping", func() {
			id := "uuid-3"

			_ = n.Publish("service.set.mapping.change", []byte(`{"_component_id":"network::test-3", "service":"`+id+`", "_state": "completed"}`))
			_, err := n.Request("service.set.mapping.change", []byte(`{"_component_id":"network::test-4", "service":"`+id+`", "_state": "completed"}`), time.Second)
			So(err, ShouldBeNil)

			Convey("It should update both the components", func() {
				var b models.Build
				db.Where("uuid = ?", id).First(&b)

				g := graph.New()
				So(g.Load(b.Mapping), ShouldBeNil)

				for _, change := range g.Changes {
					switch change.GetID() {
					case "network::test-3", "network::test-4":
						So(change.GetState(), ShouldEqual, "completed")
					}
				}
			})
		})

	})
}
*/
