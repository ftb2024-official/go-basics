package howtotesthttphandlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExample1RR(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	Example1(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("got %d, want %d", rr.Result().StatusCode, http.StatusOK)
	}

	defer rr.Result().Body.Close()
	want := "FOO"
	got, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestExample1(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(Example1))
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d, want %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
	want := "FOO"
	got, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if string(got) != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
