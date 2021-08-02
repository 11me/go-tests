package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	// 	t.Run("compaering speeds of the servers returning the url of the fastest", func(t *testing.T) {
	//
	// 		serverFast := makeDelayedServer(0 * time.Millisecond)
	// 		serverSlow := makeDelayedServer(5 * time.Millisecond)
	//
	// 		defer serverFast.Close()
	// 		defer serverSlow.Close()
	//
	// 		want := serverFast.URL
	// 		got, err := Racer(serverFast.URL, serverSlow.URL)
	//
	// 		if err != nil {
	// 			t.Fatalf("did not expected an error %v", err)
	// 		}
	//
	// 		if got != want {
	// 			t.Errorf("got %s want %s", got, want)
	// 		}
	//
	// 	})

	t.Run("returns a error if a server does not reply within 10s", func(t *testing.T) {
		server := makeDelayedServer(12 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but did not get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

	return testServer
}
