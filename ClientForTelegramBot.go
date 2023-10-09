package ClientForTelegramBot

import "net/http"

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func NewClient(host string, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}
