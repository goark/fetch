package fetch

import (
	"io"
	"net/http"
	"os"

	"github.com/goark/errs"
)

// response is wrapper class of http.Response.
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

// Header method returns Body element in http.Response.
func (resp *response) Body() io.ReadCloser {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Body
}

// Close method closes Response.Body safety.
func (resp *response) Close() (err error) {
	if resp == nil || resp.Response == nil {
		return nil
	}
	defer func() {
		if resp == nil {
			return
		}
		if cerr := resp.Body().Close(); cerr != nil {
			err = errs.Join(cerr, err)
		}
	}()
	_, err = io.Copy(io.Discard, resp.Body())
	err = errs.Wrap(err)
	return
}

func (resp *response) DumpBodyAndClose() (b []byte, err error) {
	if resp == nil || resp.Response == nil {
		err = errs.Wrap(ErrNullPointer)
		return
	}
	defer func() {
		if resp == nil {
			return
		}
		if cerr := resp.Body().Close(); cerr != nil && !errs.Is(err, os.ErrClosed) {
			err = errs.Join(cerr, err)
		}
	}()
	b, err = io.ReadAll(resp.Body())
	err = errs.Wrap(err)
	return
}

/* Copyright 2021-2025 Spiegel
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
