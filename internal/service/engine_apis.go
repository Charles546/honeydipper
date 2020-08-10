// Copyright 2019 Honey Science Corporation
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, you can obtain one at http://mozilla.org/MPL/2.0/.

package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/honeydipper/honeydipper/internal/api"
	"github.com/honeydipper/honeydipper/internal/config"
	"github.com/honeydipper/honeydipper/pkg/dipper"
)

func setupEngineAPIs() {
	engine.APIs["eventWait"] = handleEventWait
	engine.APIs["eventList"] = handleEventList
	engine.APIs["eventAdd"] = handleEventAdd
}

func handleEventWait(resp *api.Response) {
	resp.Request = dipper.DeserializePayload(resp.Request)
	eventID := dipper.MustGetMapDataStr(resp.Request.Payload, "eventID")
	sessions := sessionStore.ByEventID(eventID)
	if len(sessions) == 0 {
		return
	}

	resp.Ack()
	for _, session := range sessions {
		session.Watch()
	}
	for _, session := range sessions {
		<-session.Watch()
	}
	ret := make([]interface{}, len(sessions))
	for i, session := range sessions {
		status, reason := session.GetStatus()
		ret[i] = map[string]interface{}{
			"name":        session.GetName(),
			"description": session.GetDescription(),
			"exported":    session.GetExported(),
			"event":       session.GetEventName(),
			"status":      status,
			"reason":      reason,
		}
	}
	resp.Return(map[string]interface{}{
		"sessions": ret,
	})
}

func handleEventList(resp *api.Response) {
	resp.Request = dipper.DeserializePayload(resp.Request)
	sessions := sessionStore.GetEvents()
	ret := make([]interface{}, len(sessions))
	for i, session := range sessions {
		ret[i] = map[string]interface{}{
			"name":        session.GetName(),
			"description": session.GetDescription(),
			"exported":    session.GetExported(),
			"eventID":     session.GetEventID(),
			"event":       session.GetEventName(),
		}
	}
	resp.Return(map[string]interface{}{
		"sessions": ret,
	})
}

func handleEventAdd(resp *api.Response) {
	defer func() {
		if r := recover(); r != nil {
			resp.ReturnError(r.(error))
		}
	}()
	resp.Request = dipper.DeserializePayload(resp.Request)
	body := dipper.MustGetMapDataStr(resp.Request.Payload, "body")
	contentType := resp.Request.Labels["content-type"]
	if !strings.HasPrefix(contentType, "application/json") {
		panic(fmt.Errorf("%w: content-type: %s", http.ErrNotSupported, contentType))
	}

	type simulatedEvent struct {
		Do    config.Workflow
		Event map[string]interface{}
		With  map[string]interface{}
	}

	se := simulatedEvent{}
	dipper.Must(json.Unmarshal([]byte(body), &se))

	if se.With == nil {
		se.With = map[string]interface{}{}
	}
	se.With["_meta_event"] = "api:injected."

	msg := &dipper.Message{}
	msg.Payload = se.Event
	msg.Labels = map[string]string{}
	eventID := dipper.Must(uuid.NewRandom()).(uuid.UUID).String()
	msg.Labels["eventID"] = eventID

	go sessionStore.StartSession(&se.Do, msg, se.With)
	resp.Return(map[string]interface{}{
		"eventID": eventID,
	})
}
