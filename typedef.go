package main

import "encoding/json"

type request struct {
	Action     string `json:"action"`
	Payload    interface{}
	RawPayload json.RawMessage `json:"payload"`
}

type response struct {
	Action  string `json:"action"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type push struct {
	Action  string      `json:"action"`
	Payload interface{} `json:"payload"`
}

type requestLogin struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type requestBroadcast struct {
	Text string `json:"text"`
}

type pushMessage struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type pushUsers struct {
	List []*userDetails `json:"list"`
}

type userDetails struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
