// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package dipper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// ExtractWebRequestExceptBody put needed information except body from a request in a map.
func ExtractWebRequestExceptBody(r *http.Request) map[string]interface{} {
	Must(r.ParseForm())

	req := map[string]interface{}{
		"url":        r.URL.Path,
		"method":     r.Method,
		"form":       r.Form,
		"headers":    r.Header,
		"host":       r.Host,
		"remoteAddr": r.RemoteAddr,
	}

	return req
}

// ExtractWebRequest put needed information from a request in a map.
func ExtractWebRequest(r *http.Request) map[string]interface{} {
	req := ExtractWebRequestExceptBody(r)

	if r.Method == http.MethodPost {
		req["body"] = Must(ioutil.ReadAll(r.Body))
		if strings.HasPrefix(r.Header.Get("content-type"), "application/json") {
			bodyObj := map[string]interface{}{}
			Must(json.Unmarshal(req["body"].([]byte), &bodyObj))
			req["json"] = bodyObj
		}
	}

	return req
}
