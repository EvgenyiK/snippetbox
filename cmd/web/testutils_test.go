package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

// Create a newTestApplication helper which returns an instance
func newTestApplication(t *testing.T)*application{
	return &application{
		errorLog: log.New(ioutil.Discard,"",0),
		infoLog: log.New(ioutil.Discard,"",0),
	}
}

// Define a custom testServer type which anonymously embeds a httptest.
type testServer struct{
	*httptest.Server
}

// Create a newTestServer helper which initalizes and returns a new instance
func newTestServer(t*testing.T,h http.Handler)*testServer{
	ts:=httptest.NewTLSServer(h)
	
	// Initialize a new cookie jar.
	jar,err:=cookiejar.New(nil)
	if err != nil{
		t.Fatal(err)
	}
	// Add the cookie jar to the client, so that response cookies are stored
	ts.Client().Jar = jar

	// Disable redirect-following for the client.
	ts.Client().CheckRedirect = func (req *http.Request,via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

// Implement a get method on our custom testServer type.
func(ts *testServer)get(t *testing.T, urlPath string)(int,http.Header, []byte){
	rs,err:=ts.Client().Get(ts.URL + urlPath)
	if err != nil{
		t.Fatal(err)
	}

	defer rs.Body.Close()
	body,err:=ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	return rs.StatusCode, rs.Header, body
}