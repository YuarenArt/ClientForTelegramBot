package ClientForTelegramBot

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	testCases := makeTestCasesNewClient()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			clients := NewClient(tc.host, tc.token)

			if clients.host != tc.expectedHost {
				t.Errorf("Expected host to be %s, but got %s", tc.expectedHost, clients.host)
			}

			if clients.basePath != tc.expectedBasePath {
				t.Errorf("Expected basePath to be %s, but got %s", tc.expectedBasePath, clients.basePath)
			}
			// Добавьте другие проверки по мере необходимости.
		})
	}
}

func makeTestCasesNewClient() []testCasesNewClient {
	return []testCasesNewClient{
		{
			name:             "Test Case 1",
			host:             "example.com",
			token:            "yourtoken",
			expectedHost:     "example.com",
			expectedBasePath: "botyourtoken",
		},
		{
			name:             "Test Case 2",
			host:             "testhost.com",
			token:            "anothertoken",
			expectedHost:     "testhost.com",
			expectedBasePath: "botanothertoken",
		},
		// Добавьте другие тестовые случаи по мере необходимости.
	}
}

func TestNewBasePath(t *testing.T) {

	testCases := makeTestCasesNewBasePath()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			basePath := newBasePath(tc.token)

			if basePath != tc.expectedBasePath {
				t.Errorf("Expected basePath to be %s, but got %s", tc.expectedBasePath, basePath)
			}
			// Добавьте другие проверки по мере необходимости.
		})
	}
}

func makeTestCasesNewBasePath() []testCasesNewBasePath {

	return []testCasesNewBasePath{
		{
			name:             "Test Case 1",
			token:            "yourtoken",
			expectedBasePath: "botyourtoken",
		},
		{
			name:             "Test Case 2",
			token:            "anothertoken",
			expectedBasePath: "botanothertoken",
		},
		// Добавьте другие тестовые случаи по мере необходимости.
	}
}

func TestSendMessage(t *testing.T) {
	client := NewClient("example.com", "yourtoken")

	testCases := makeTestCasesSendMessage()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := client.SendMessage(tc.chatID, tc.text)

			if err != nil {
				t.Errorf("Expected no error, but got an error: %v", err)
			}
			// Добавьте другие проверки по мере необходимости.
		})
	}
}

func makeTestCasesSendMessage() []testCasesSendMessage {

	return []testCasesSendMessage{
		{
			name:    "Test Case 1",
			chatID:  12345,
			text:    "Hello, world!",
			wantErr: false,
		},
		{
			name:    "Test Case 2",
			chatID:  54321,
			text:    "Test message",
			wantErr: false,
		},
		// Добавьте другие тестовые случаи по мере необходимости.
	}
}

// not complete
func TestMakeQueryWithUrlValues(t *testing.T) {
	t.Errorf("Not complete function")
}

// not complete
func makeTestCasesMakeQuery() []testCasesMakeQuery {
	return []testCasesMakeQuery{
		{
			name: "Test Case 1",
			args: []interface{}{
				"key1", "value1",
				"key2", "value2",
			},
			expectedData: []interface{}{
				"key1", "value1",
				"key2", "value2",
			},
		},
		{
			name: "Test Case 2",
			args: []interface{}{
				"key", "value",
				"intKey", 123,
			},
			expectedData: []interface{}{
				"key", "value",
				"intKey", 123,
			},
		},
		// Добавьте другие тестовые случаи по мере необходимости.
	}
}

// not complete
func TestDoRequest(t *testing.T) {
	t.Errorf("Not complete function")
}

func TestUpdates(t *testing.T) {

}
