package fetch

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spiegel-im-spiegel/errs"
)

//response is wrapper class of http.Response.
type response struct {
	*http.Response
}

//Header method returns Header element in http.Response.
func (resp *response) Header() http.Header {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Header
}

//Header method returns Body element in http.Response.
func (resp *response) Body() io.ReadCloser {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Body
}

//Close method closes Response.Body safety.
func (resp *response) Close() {
	if resp == nil || resp.Response == nil {
		return
	}
	_, _ = io.Copy(ioutil.Discard, resp.Body())
	resp.Body().Close()
}

func (resp *response) DumpBodyAndClose() ([]byte, error) {
	if resp == nil || resp.Response == nil {
		return nil, errs.Wrap(ErrNullPointer)
	}
	defer resp.Body().Close()
	b, err := ioutil.ReadAll(resp.Body())
	return b, errs.Wrap(err)
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