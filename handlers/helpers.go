/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package handlers

import (
	"encoding/json"
	"log"
)

// Error : default error message
type Error struct {
	Error string `json:"_error"`
}

// Message ...
type Message struct {
	ID         string                 `json:"id"`
	Definition string                 `json:"definition"`
	Mapping    map[string]interface{} `json:"mapping"`
	Validation map[string]interface{} `json:"validation"`
}

func response(reply string, data *[]byte, err *error) {
	var rdata []byte
	if data != nil {
		rdata = *data
	}

	if *err != nil {
		log.Println("[ ERROR ] " + (*err).Error())
		rdata, _ = json.Marshal(Error{Error: (*err).Error()})
	}

	if reply != "" {
		NC.Publish(reply, rdata)
	}
}

func pub(subject string, data []byte) {
	if err := NC.Publish(subject, data); err != nil {
		log.Println("[ERROR] : " + err.Error())
	}
}
