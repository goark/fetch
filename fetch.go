package fetch

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type RequestOpts func(*http.Request) *http.Request

// Client is inteface class for HTTP client.
type Client interface {
	// Deprecated: Should use GetWithContext() method instead of Get() method.
	Get(u *url.URL, opts ...RequestOpts) (Response, error)
	GetWithContext(ctx context.Context, u *url.URL, opts ...RequestOpts) (Response, error)
	// Deprecated: Should use PostWithContext() method instead of Post() method.
	Post(u *url.URL, payload io.Reader, opts ...RequestOpts) (Response, error)
	PostWithContext(ctx context.Context, u *url.URL, payload io.Reader, opts ...RequestOpts) (Response, error)
}

// Response is inteface class for HTTP response.
type Response interface {
	Request() *http.Request
	Header() http.Header
	Body() io.ReadCloser
	Close()
	DumpBodyAndClose() ([]byte, error)
}

/* Copyright 2023 Spiegel
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
