package main

type request struct {
	Action     string `json:"action"`
	Payload    interface{}
	RawPayload []byte `json:"payload"`
}

type response struct {
	Action  string      `json:"action"`
	Success bool        `json:"success"`
	Payload interface{} `json:"payload,omitempty"`
	Error   string      `json:"error,omitempty"`
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
