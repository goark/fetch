package fetch_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goark/fetch"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		s    string
		err1 error
		err2 error
	}{
		{s: "foo\nbar", err1: fetch.ErrInvalidURL, err2: fetch.ErrInvalidRequest},
		{s: "http://foo.bar", err1: nil, err2: fetch.ErrInvalidRequest},
		{s: "https://text.baldanders.info/not-exist/", err1: nil, err2: fetch.ErrInvalidRequest},
		{s: "https://github.com/spiegel-im-spiegel.gpg", err1: nil, err2: nil},
	}
	for _, tc := range testCases {
		u, err := fetch.URL(tc.s)
		if err != nil {
			if !errors.Is(err, tc.err1) {
				t.Errorf("fetch.Client.URL(%s) is \"%v\", want \"%+v\"", tc.s, err, tc.err1)
			}
			fmt.Printf("Info: %+v\n", err)
		} else {
			resp, err := fetch.New(
				fetch.WithHTTPClient(&http.Client{}),
			).GetWithContext(context.Background(), u)
			if err != nil {
				if !errors.Is(err, tc.err2) {
					t.Errorf("fetch.Client.Get() is \"%v\", want \"%+v\"", err, tc.err2)
				}
				fmt.Printf("Info: %+v\n", err)
			} else if cerr := resp.Close(); cerr != nil {
				t.Errorf("resp.Close() is \"%v\", want nil", cerr)
			}
		}
	}
}

func TestGetWithNilURL(t *testing.T) {
	resp, err := fetch.New().GetWithContext(context.Background(), nil)
	if err == nil {
		t.Fatal("error is nil, want ErrInvalidURL")
	}
	if !errors.Is(err, fetch.ErrInvalidURL) {
		t.Fatalf("GetWithContext(nil) is %v, want ErrInvalidURL", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestWithHTTPClientNilFallback(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New(fetch.WithHTTPClient(nil)).GetWithContext(context.Background(), u)
	if err != nil {
		t.Fatalf("GetWithContext() error = %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if cerr := resp.Close(); cerr != nil {
		t.Fatalf("resp.Close() error = %v", cerr)
	}
}

/* Copyright 2023-2025 Spiegel
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
