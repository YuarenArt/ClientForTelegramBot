package ClientForTelegramBot

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	host       string
	basePath   string
	httpClient http.Client
}

const (
	getUpdates  = "getUpdates"
	sendMessage = "sendMessage"
)

func NewClient(host string, token string) Client {
	return Client{
		host:       host,
		basePath:   newBasePath(token),
		httpClient: http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (client *Client) SendMessage(chatId int, text string) error {
	query := makeQueryWithUrlValues(chatId, text)

	_, err := client.doRequest(query, sendMessage)
	if err != nil {
		return Wrap("can't send message: ", err)
	}
	return nil
}

func makeQueryWithUrlValues(args ...interface{}) url.Values {
	query := url.Values{}

	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			query.Add(v, v)
		case int:
			query.Add(fmt.Sprintf("%d", v), "")
		case bool:
			query.Add(fmt.Sprintf("%t", v), "")
		default:
			log.Fatal(fmt.Errorf("unsupported type %T", arg))
		}
	}

	return query
}

func (client *Client) doRequest(query url.Values, method string) ([]byte, error) {
	const errorMessageDoRequest = "can't do request"

	url := url.URL{
		Scheme: "https",
		Host:   client.host,
		Path:   path.Join(client.basePath, method),
	}

	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, Wrap(errorMessageDoRequest, err)
	}

	request.URL.RawQuery = query.Encode()

	response, err := client.httpClient.Do(request)
	if err != nil {
		return nil, Wrap(errorMessageDoRequest, err)
	}

	defer func() { _ = response.Body.Close() }()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, Wrap(errorMessageDoRequest, err)
	}
	return body, nil
}

func (client *Client) Updates(offset int, limit int) ([]Update, error) {
	query := makeQueryWithUrlValues(offset, limit)

	data, err := client.doRequest(query, getUpdates)
	if err != nil {
		return nil, err
	}

	var result UpdatesResponse
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result.Result, err
}
