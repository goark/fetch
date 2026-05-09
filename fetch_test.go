package fetch_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goark/fetch"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestURL(t *testing.T) {
	u, err := fetch.URL("foo\nbar")
	if err == nil {
		t.Fatal("fetch.URL() error is nil, want ErrInvalidURL")
	}
	if !errors.Is(err, fetch.ErrInvalidURL) {
		t.Fatalf("fetch.URL() = %v, want ErrInvalidURL", err)
	}
	if u != nil {
		t.Fatal("fetch.URL() returned non-nil URL for invalid input")
	}
}

func TestGetWithTransportError(t *testing.T) {
	u, err := fetch.URL("http://example.test/")
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	cli := fetch.New(fetch.WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("transport failure")
	})}))

	resp, err := cli.GetWithContext(context.Background(), u)
	if err == nil {
		t.Fatal("GetWithContext() error is nil, want ErrInvalidRequest")
	}
	if !errors.Is(err, fetch.ErrInvalidRequest) {
		t.Fatalf("GetWithContext() = %v, want ErrInvalidRequest", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestGetWithHTTPStatusError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, "not found")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().GetWithContext(context.Background(), u)
	if err == nil {
		t.Fatal("GetWithContext() error is nil, want ErrInvalidRequest")
	}
	if !errors.Is(err, fetch.ErrInvalidRequest) {
		t.Fatalf("GetWithContext() = %v, want ErrInvalidRequest", err)
	}
	if !errors.Is(err, fetch.ErrHTTPStatus) {
		t.Fatalf("GetWithContext() = %v, want ErrHTTPStatus in cause chain", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestGetWithSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().GetWithContext(context.Background(), u)
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

func TestPostWithTransportError(t *testing.T) {
	u, err := fetch.URL("http://example.test/")
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	cli := fetch.New(fetch.WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, errors.New("transport failure")
	})}))

	resp, err := cli.PostWithContext(context.Background(), u, strings.NewReader("a=1"))
	if err == nil {
		t.Fatal("PostWithContext() error is nil, want ErrInvalidRequest")
	}
	if !errors.Is(err, fetch.ErrInvalidRequest) {
		t.Fatalf("PostWithContext() = %v, want ErrInvalidRequest", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestPostWithHTTPStatusError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "bad request")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().PostWithContext(context.Background(), u, strings.NewReader("a=1"))
	if err == nil {
		t.Fatal("PostWithContext() error is nil, want ErrInvalidRequest")
	}
	if !errors.Is(err, fetch.ErrInvalidRequest) {
		t.Fatalf("PostWithContext() = %v, want ErrInvalidRequest", err)
	}
	if !errors.Is(err, fetch.ErrHTTPStatus) {
		t.Fatalf("PostWithContext() = %v, want ErrHTTPStatus in cause chain", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestPostWithSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		body, _ := io.ReadAll(r.Body)
		_ = r.Body.Close()
		if string(body) != "a=1" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().PostWithContext(context.Background(), u, strings.NewReader("a=1"))
	if err != nil {
		t.Fatalf("PostWithContext() error = %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if cerr := resp.Close(); cerr != nil {
		t.Fatalf("resp.Close() error = %v", cerr)
	}
}

func TestPostWithNilURL(t *testing.T) {
	resp, err := fetch.New().PostWithContext(context.Background(), nil, strings.NewReader("a=1"))
	if err == nil {
		t.Fatal("error is nil, want ErrInvalidURL")
	}
	if !errors.Is(err, fetch.ErrInvalidURL) {
		t.Fatalf("PostWithContext(nil) is %v, want ErrInvalidURL", err)
	}
	if resp != nil {
		t.Fatal("response is not nil, want nil")
	}
}

func TestRequestHeaderOptions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xVals := r.Header.Values("X-Test")
		if len(xVals) != 2 || xVals[0] != "first" || xVals[1] != "second" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if got := r.Header.Get("X-Mode"); got != "final" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().GetWithContext(context.Background(), u,
		fetch.WithRequestHeaderAdd("X-Test", "first"),
		fetch.WithRequestHeaderAdd("X-Test", "second"),
		fetch.WithRequestHeaderSet("X-Mode", "initial"),
		fetch.WithRequestHeaderSet("X-Mode", "final"),
	)
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

func TestURLWithTrimmedInput(t *testing.T) {
	u, err := fetch.URL("  https://example.test/path?q=1  ")
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}
	if got, want := u.String(), "https://example.test/path?q=1"; got != want {
		t.Fatalf("fetch.URL() string = %q, want %q", got, want)
	}
}

func TestDeprecatedGetDelegatesToGetWithContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().Get(u)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if cerr := resp.Close(); cerr != nil {
		t.Fatalf("resp.Close() error = %v", cerr)
	}
}

func TestDeprecatedPostDelegatesToPostWithContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok")
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().Post(u, strings.NewReader("a=1"))
	if err != nil {
		t.Fatalf("Post() error = %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if cerr := resp.Close(); cerr != nil {
		t.Fatalf("resp.Close() error = %v", cerr)
	}
}

func TestGetWithHTTPStatus3xx(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotModified)
	}))
	defer ts.Close()

	u, err := fetch.URL(ts.URL)
	if err != nil {
		t.Fatalf("fetch.URL() error = %v", err)
	}

	resp, err := fetch.New().GetWithContext(context.Background(), u)
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
