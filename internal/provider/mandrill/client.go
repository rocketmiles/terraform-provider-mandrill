package mandrill

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultApiBaseUrl = "https://mandrillapp.com/api/1.0"
	contentType       = "application/json"
)

func NewMandrillClient(apiKey string, apiBaseUrl string) (*mandrillClient, error) {

	if apiKey == "" {
		return nil, errors.New("mandrillClient: missing api key")
	}

	if apiBaseUrl == "" {
		return nil, errors.New("mandrillClient: missing api base url")
	}

	return &mandrillClient{
		apiKey:     apiKey,
		apiBaseUrl: apiBaseUrl,
	}, nil
}

type mandrillClient struct {
	apiKey     string
	apiBaseUrl string
}

func (client mandrillClient) post(endpoint string, request interface{}, responseModel interface{}) (*http.Response, error) {
	requestBuffer := new(bytes.Buffer)
	json.NewEncoder(requestBuffer).Encode(request)

	response, error := http.Post(client.apiBaseUrl+endpoint, contentType, requestBuffer)

	if error != nil {
		log.WithError(error).Error("sendersCheckDomain")
		return response, error
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBody := getResponseBodyAsString(response)
		return response, errors.New(responseBody)
	}

	decodeError := json.NewDecoder(response.Body).Decode(responseModel)

	return response, decodeError
}

func getResponseBodyAsString(resp *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithField("status code", resp.StatusCode).WithError(err).Error("getResponseBodyAsString")
		return "Could not convert response body to string"
	}
	return string(bodyBytes)
}

type MandrillClient interface {
	SendersCheckDomain(SendersCheckDomainRequest) (SendersCheckDomainResponse, error)
}

func (client mandrillClient) SendersCheckDomain(request SendersCheckDomainRequest) (SendersCheckDomainResponse, error) {

	request.Key = client.apiKey

	responseModel := SendersCheckDomainResponse{}

	_, error := client.post("/senders/check-domain", request, &responseModel)

	if error != nil {
		log.WithError(error).Error("sendersCheckDomain")
		return SendersCheckDomainResponse{}, error
	}

	return responseModel, nil
}
