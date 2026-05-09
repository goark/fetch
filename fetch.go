package fetch

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// RequestOpts applies options to an HTTP request.
//
// The returned request value allows option functions to replace the request,
// for example when changing context with req.WithContext.
type RequestOpts func(*http.Request) *http.Request

// Client represents a fetch client.
type Client interface {
	// Get sends a GET request.
	// Deprecated: Use GetWithContext instead.
	Get(u *url.URL, opts ...RequestOpts) (Response, error)
	// GetWithContext sends a GET request with context.
	GetWithContext(ctx context.Context, u *url.URL, opts ...RequestOpts) (Response, error)
	// Post sends a POST request.
	// Deprecated: Use PostWithContext instead.
	Post(u *url.URL, payload io.Reader, opts ...RequestOpts) (Response, error)
	// PostWithContext sends a POST request with context.
	PostWithContext(ctx context.Context, u *url.URL, payload io.Reader, opts ...RequestOpts) (Response, error)
}

// Response wraps an HTTP response.
type Response interface {
	// Request returns the original request.
	Request() *http.Request
	// Header returns response headers.
	Header() http.Header
	// Body returns the response body reader.
	Body() io.ReadCloser
	// Close drains and closes the response body.
	Close() error
	// DumpBodyAndClose reads all body bytes and closes the response body.
	DumpBodyAndClose() ([]byte, error)
}

/* Copyright 2023-2026 Spiegel
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
