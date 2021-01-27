package fetch

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
)

// client is client class for fetching (internal).
type client struct {
	client *http.Client
}

type ClientOpts func(*client)

// New function returns Client instance.
func New(opts ...ClientOpts) Client {
	cli := &client{client: &http.Client{}}
	for _, opt := range opts {
		opt(cli)
	}
	return cli
}

// WithProtocol returns function for setting http.Client.
func WithHTTPClient(cli *http.Client) ClientOpts {
	return func(c *client) {
		c.client = cli
	}
}

// Get method returns respons data from URL by GET method.
func (c *client) Get(u *url.URL, opts ...RequestOpts) (Response, error) {
	req, err := request(http.MethodGet, u, nil, opts...)
	if err != nil {
		return nil, errs.Wrap(ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", u.String()))
	}
	resp, err := c.fetch(req)
	if err != nil {
		return nil, errs.Wrap(ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", u.String()))
	}
	return resp, nil
}

// Post method returns respons data from URL by POST method.
func (c *client) Post(u *url.URL, payload io.Reader, opts ...RequestOpts) (Response, error) {
	req, err := request(http.MethodPost, u, payload, opts...)
	if err != nil {
		return nil, errs.Wrap(ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", u.String()))
	}
	resp, err := c.fetch(req)
	if err != nil {
		return nil, errs.Wrap(ErrInvalidRequest, errs.WithCause(err), errs.WithContext("url", u.String()))
	}
	return resp, nil
}

// WithProtocol returns function for setting context.Context.
func WithContext(ctx context.Context) RequestOpts {
	return func(req *http.Request) *http.Request {
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		return req
	}
}

// WithRequestHeaderAdd returns function for adding request header in http.Request.
func WithRequestHeaderAdd(name, value string) RequestOpts {
	return func(req *http.Request) *http.Request {
		req.Header.Add(name, value)
		return req
	}
}

// WithRequestHeaderSet returns function for setting request header in http.Request.
func WithRequestHeaderSet(name, value string) RequestOpts {
	return func(req *http.Request) *http.Request {
		req.Header.Set(name, value)
		return req
	}
}

func request(method string, u *url.URL, payload io.Reader, opts ...RequestOpts) (*http.Request, error) {
	req, err := http.NewRequest(method, u.String(), payload)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	for _, opt := range opts {
		req = opt(req)
	}
	return req, nil
}

func (c *client) fetch(request *http.Request) (Response, error) {
	if c == nil {
		c = New().(*client)
	}
	r, err := c.client.Do(request)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	resp := &response{r}
	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		resp.Close()
		return nil, errs.Wrap(ErrHTTPStatus, errs.WithContext("status", resp.StatusCode))
	}
	return resp, nil
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
