package bitbucket_v2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/DaoCloud/go-bitbucket/oauth1"
)

const (
	APIURL = "https://api.bitbucket.org/2.0"
)

var (
	// Returned if the specified resource does not exist.
	ErrNotFound = errors.New("Not Found")

	// Returned if the caller attempts to make a call or modify a resource
	// for which the caller is not authorized.
	//
	// The request was a valid request, the caller's authentication credentials
	// succeeded but those credentials do not grant the caller permission to
	// access the resource.
	ErrForbidden = errors.New("Forbidden")

	// Returned if the call requires authentication and either the credentials
	// provided failed or no credentials were provided.
	ErrNotAuthorized = errors.New("Unauthorized")

	// Returned if the caller submits a badly formed request. For example,
	// the caller can receive this return if you forget a required parameter.
	ErrBadRequest = errors.New("Bad Request")

	ErrWrongParamType = errors.New("Wrong Param Type")
)

// DefaultClient uses DefaultTransport, and is used internall to execute
// all http.Requests. This may be overriden for unit testing purposes.
//
// IMPORTANT: this is not thread safe and should not be touched with
// the exception overriding for mock unit testing.
var DefaultClient = http.DefaultClient

func (c *Client) do(method string, path string, params url.Values, values url.Values, jsonbody string, v interface{}) error {

	// if this is the guest client then we don't need
	// to sign the request ... we will execute just
	// a simple http request.
	if c == Guest {
		return c.guest(method, path, params, values, jsonbody, v)
	}

	// create the client
	var client = oauth1.Consumer{
		ConsumerKey:    c.ConsumerKey,
		ConsumerSecret: c.ConsumerSecret,
	}

	// create the URI
	uri, err := url.Parse(APIURL + path)
	if err != nil {
		return err
	}

	if params != nil && len(params) > 0 {
		uri.RawQuery = params.Encode()
	}

	// create the access token
	token := oauth1.NewAccessToken(c.AccessToken, c.TokenSecret, nil)

	// create the request
	req, err := http.NewRequest(method, uri.String(), strings.NewReader(jsonbody))

	if jsonbody != "" {
		req.Header.Add("Content-Type", "application/json")
	}

	if values != nil {
		req.Form = values
	}

	// sign the request
	if err := client.Sign(req, token); err != nil {
		return err
	}

	// make the request using the default http client
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read the bytes from the body (make sure we defer close the body)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for an http error status (ie not 200 StatusOK)
	switch resp.StatusCode {
	case 404:
		return ErrNotFound
	case 403:
		return ErrForbidden
	case 401:
		return ErrNotAuthorized
	case 400:
		return ErrBadRequest
	}

	// Unmarshall the JSON response
	if v != nil {
		return json.Unmarshal(body, v)
	}

	return nil
}

func (c *Client) guest(method string, path string, params url.Values, values url.Values, jsonbody string, v interface{}) error {

	// create the URI
	uri, err := url.Parse(APIURL + path)
	if err != nil {
		return err
	}

	if params != nil && len(params) > 0 {
		uri.RawQuery = params.Encode()
	}

	// create the request
	req, err := http.NewRequest(method, uri.String(), strings.NewReader(jsonbody))

	if jsonbody != "" {
		req.Header.Add("Content-Type", "application/json")
	}

	if values != nil {
		req.Form = values
	}

	// make the request using the default http client
	resp, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Read the bytes from the body (make sure we defer close the body)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for an http error status (ie not 200 StatusOK)
	switch resp.StatusCode {
	case 404:
		return ErrNotFound
	case 403:
		return ErrForbidden
	case 401:
		return ErrNotAuthorized
	case 400:
		return ErrBadRequest
	}

	// Unmarshall the JSON response
	if v != nil {
		return json.Unmarshal(body, v)
	}

	return nil
}
