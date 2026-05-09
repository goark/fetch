package fetch

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDumpBodyAndClose(t *testing.T) {
	r := &response{&http.Response{Body: io.NopCloser(strings.NewReader("payload"))}}
	b, err := r.DumpBodyAndClose()
	if err != nil {
		t.Fatalf("DumpBodyAndClose() error = %v", err)
	}
	if string(b) != "payload" {
		t.Fatalf("DumpBodyAndClose() body = %q, want %q", string(b), "payload")
	}
}

func TestDumpBodyAndCloseWithNullPointer(t *testing.T) {
	var r *response
	b, err := r.DumpBodyAndClose()
	if err == nil {
		t.Fatal("DumpBodyAndClose() error is nil, want ErrNullPointer")
	}
	if !errors.Is(err, ErrNullPointer) {
		t.Fatalf("DumpBodyAndClose() error = %v, want ErrNullPointer", err)
	}
	if b != nil {
		t.Fatalf("DumpBodyAndClose() body = %v, want nil", b)
	}
}

/* Copyright 2026 Spiegel
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
