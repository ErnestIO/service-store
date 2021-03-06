/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
)

// DB ...
var DB *gorm.DB

func structFields(x interface{}) []string {
	var sp []string

	rx := reflect.TypeOf(x)

	for i := 0; i < rx.NumField(); i++ {

		sp = append(sp, rx.Field(i).Tag.Get("json"))
	}

	return sp
}

func supported(f string, fields []string) bool {
	for _, field := range fields {
		sf := strings.Split(field, "->")[0]
		if f == sf {
			return true
		}
	}
	return false
}

func parse(f string, fields []string) string {
	for _, field := range fields {
		sf := strings.Split(field, "->")
		if sf[0] == f {
			return sf[1]
		}
	}
	return ""
}

func query(q map[string]interface{}, fields, qfields []string) *gorm.DB {
	qdb := DB

	for k, v := range q {
		var qs string

		if supported(k, fields) {
			qs = fmt.Sprintf("%s = ?", k)
		}

		if supported(k, qfields) {
			sf := parse(k, qfields)
			qs = fmt.Sprintf("%s in (?)", sf)
		}

		if qs != "" {
			qdb = qdb.Where(qs, v)
		}
	}

	return qdb
}
