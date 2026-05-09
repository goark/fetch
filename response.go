package fetch

import (
	"io"
	"net/http"
	"os"

	"github.com/goark/errs"
)

// response wraps http.Response.
type response struct {
	*http.Response
}

// Request method returns Request element in http.Response.
func (resp *response) Request() *http.Request {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Request
}

// Header method returns Header element in http.Response.
func (resp *response) Header() http.Header {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Header
}

// Body method returns Body element in http.Response.
func (resp *response) Body() io.ReadCloser {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Body
}

// Close method safely drains and closes Response.Body.
func (resp *response) Close() (err error) {
	if resp == nil || resp.Response == nil {
		return nil
	}
	body := resp.Body()
	defer func() {
		if body == nil {
			return
		}
		if cerr := body.Close(); cerr != nil && !errs.Is(cerr, os.ErrClosed) {
			err = errs.Join(cerr, err)
		}
	}()
	_, err = io.Copy(io.Discard, body) // drain body to reuse connection
	err = errs.Wrap(err)
	return
}

// DumpBodyAndClose reads all response body bytes and closes Response.Body.
func (resp *response) DumpBodyAndClose() (b []byte, err error) {
	if resp == nil || resp.Response == nil {
		err = errs.Wrap(ErrNullPointer)
		return
	}
	body := resp.Body()
	defer func() {
		if body == nil {
			return
		}
		if cerr := body.Close(); cerr != nil && !errs.Is(cerr, os.ErrClosed) {
			err = errs.Join(cerr, err)
		}
	}()
	b, err = io.ReadAll(body)
	err = errs.Wrap(err)
	return
}

/* Copyright 2021-2026 Spiegel
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
