package goly

import (
	"errors"
	"net/url"
)

func (c *Client) CreateRequest(data url.Values) (*Request, error) {
	var response Request
	re := url.Values{}
	for k, v := range data {
		if len(v) > 0 {
			re.Set(k, v[0])
		} else {
			re.Set(k, "")
		}
	}
	re.Set("vendor_token", c.identity.PublicVendorToken)

	err := c.request("/request/do.json", re, &response)
	if err != nil {
		return nil, err
	}
	if response.Error != "0" {
		return nil, errors.New(response.Message)
	}
	return &response, nil
}

func (c *Client) GetRequestState(id string) (*RequestState, error) {
	var res RequestState
	var re = url.Values{}
	re.Set("request_id", id)
	re.Set("vendor_token", c.identity.PublicVendorToken)
	err := c.request("/request/state.json", re, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
