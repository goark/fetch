package fetch

import (
	"net/http"
	"testing"
)

func TestWithHTTP11Client(t *testing.T) {
	cli, ok := New(WithHTTP11Client()).(*client)
	if !ok {
		t.Fatal("New(WithHTTP11Client()) did not return *client")
	}
	if cli.client == nil {
		t.Fatal("http client is nil")
	}

	tr, ok := cli.client.Transport.(*http.Transport)
	if !ok {
		t.Fatalf("transport type = %T, want *http.Transport", cli.client.Transport)
	}
	if tr.ForceAttemptHTTP2 {
		t.Fatal("ForceAttemptHTTP2 = true, want false")
	}
	if tr.TLSNextProto == nil {
		t.Fatal("TLSNextProto is nil, want non-nil map")
	}

	if _, exists := tr.TLSNextProto["h2"]; exists {
		t.Fatal("TLSNextProto unexpectedly has h2 entry")
	}

	if _, exists := tr.TLSNextProto["http/1.1"]; exists {
		t.Fatal("TLSNextProto has unexpected http/1.1 entry")
	}

	if len(tr.TLSNextProto) != 0 {
		t.Fatalf("TLSNextProto length = %d, want 0", len(tr.TLSNextProto))
	}
}
