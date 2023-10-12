package ClientForTelegramBot

import "net/url"

type UpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}

// testing types

type testCasesNewClient struct {
	name             string
	host             string
	token            string
	expectedHost     string
	expectedBasePath string
}

type testCasesNewBasePath struct {
	name             string
	token            string
	expectedBasePath string
}

type testCasesSendMessage struct {
	name    string
	chatID  int
	text    string
	wantErr bool
}

type testCasesMakeQuery struct {
	name         string
	args         []interface{}
	expectedData []interface{}
}

type testCasesDoRequest struct {
	name           string
	client         Client
	query          url.Values
	method         string
	expectedOutput []byte
	expectedError  error
}

type testCasesUpdates struct {
	name            string
	offset          int
	limit           int
	expectedUpdates []Update
	expectedError   error
}
