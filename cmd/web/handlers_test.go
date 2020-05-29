package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	)


func TestPing(t *testing.T){
	rr:=httptest.NewRecorder()

	r,err:=http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the ping handler function,
	ping(rr,r)

	// Call the Result() method on the http.ResponseRecorder
	rs:=rr.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d;got %d",http.StatusOK, rs.StatusCode)
	}

	// And we can check that the response body written by the ping handler
	defer rs.Body.Close()
	body, err:=ioutil.ReadAll(rs.Body)
	if err !=nil {
		t.Fatal(err)
	}
	if string(body)!= "OK" {
		t.Errorf("want body to equal %q","OK")
	}
}