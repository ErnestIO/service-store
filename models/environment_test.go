/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"log"
	"strconv"
	"testing"

	"github.com/ernestio/service-store/tests"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

const TESTENVDB = "test_environments"

// EnvironmentTestSuite : Test suite for migration
type EnvironmentTestSuite struct {
	suite.Suite
}

// SetupTest : sets up test suite
func (suite *EnvironmentTestSuite) SetupTest() {
	err := tests.CreateTestDB(TESTENVDB)
	if err != nil {
		log.Fatal(err)
	}

	DB, err = gorm.Open("postgres", "user=postgres dbname="+TESTENVDB+" sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	//DB.LogMode(true)

	DB.AutoMigrate(Environment{})
	DB.Unscoped().Delete(Environment{})

	for i := 1; i <= 10; i++ {
		DB.Create(&Environment{
			Name:   "Test" + strconv.Itoa(i),
			Status: "in_progress",
			Options: map[string]interface{}{
				"sync":          true,
				"sync_type":     "hard",
				"sync_interval": 5,
			},
			Credentials: map[string]interface{}{},
		})
	}
}

func (suite *EnvironmentTestSuite) TestEnvironments() {
	suite.testFindEnvironments()
}

func (suite *EnvironmentTestSuite) testFindEnvironments() {
	environments, err := FindEnvironments(map[string]interface{}{
		"name":     "Test1",
		"group_id": 1,
	})

	suite.Nil(err)
	suite.Equal(len(environments), 1)
	suite.Equal(environments[0].ID, uint(1))
	suite.Equal(environments[0].Name, "Test1")
	suite.Equal(environments[0].Status, "in_progress")
	suite.Equal(environments[0].Options["sync"], true)
}

// TestEnvironmentTestSuite : Test suite for migration
func TestEnvironmentTestSuite(t *testing.T) {
	suite.Run(t, new(EnvironmentTestSuite))
}
