package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

func TestServerRun(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatalf("failed to listen port: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	mux := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	})
	group.Go(func() error {
		svr := NewServer(l, mux)
		return svr.Run(ctx)
	})

	input := "message"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), input)
	t.Logf("try request to %q", url)

	response, err := http.Get(url)
	if err != nil {
		t.Fatalf("failed to get: %v", err)
	}
	defer response.Body.Close()

	got, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}

	cancel()

	if err := group.Wait(); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, fmt.Sprintf("Hello, %s!", input), string(got))
}
