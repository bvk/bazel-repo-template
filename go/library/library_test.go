package library

import (
	"context"
	"testing"

	librarypb "github.com/bvk/bazel-repo-template/proto/library"
)

func TestHello(t *testing.T) {
	s, err := New()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := s.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	req := librarypb.CreateBookRequest{
		Name: "The Go Programming Language",
	}
	resp, err := s.CreateBook(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Book.Name != req.Name {
		t.Fatal(err)
	}
}
