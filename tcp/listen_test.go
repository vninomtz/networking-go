package main

import (
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ln.Close() }()

	t.Logf("bound to %q", ln.Addr().String())
}
