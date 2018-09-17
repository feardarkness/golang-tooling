package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"ariel@golang.org", "gopher ariel"},
		{"alvarado@hotmail.org", "dear alvarado"},
	}

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://127.0.0.1:8087/"+c.in,
			nil)

		if err != nil {
			t.Fatalf("Could not create  request %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rec.Code)
		}

		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("Unexpected body in response %q", rec.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://127.0.0.1:8087/alvarado@golang.org",
			nil)

		if err != nil {
			b.Fatalf("Could not create  request %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected status 200, got %d", rec.Code)
		}

		if !strings.Contains(rec.Body.String(), "gopher alvarado") {
			b.Errorf("Unexpected body in response %q", rec.Body.String())
		}

	}
}
