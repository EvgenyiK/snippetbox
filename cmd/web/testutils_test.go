package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/EvgenyiK/snippetbox/pkg/models/mock"
	
	"github.com/golangcollege/sessions"
)

// Create a newTestApplication helper which returns an instance
func newTestApplication(t *testing.T)*application{
	// Create an instance of the template cache.
	templateCashe,err:=newTemplateCache("./../../ui/html/")
	if err != nil {
		t.Fatal(err)
	}

	// Create a session manager instance, with the same settings as production.
	session:=sessions.New([]byte("3dSm5MnygFHh7XidAtbskXrjbwfoJcbJ"))
	session.Lifetime = 12 * time.Hour
	session.Secure = true

	// Initialize the dependencies, using the mocks for the loggers and
    // database models.
	return &application{
		errorLog: log.New(ioutil.Discard,"",0),
		infoLog: log.New(ioutil.Discard,"",0),
		session: session,
		snippets: &mock.SnippetModel{},
		templateCache: templateCashe,
		users:    &mock.UserModel{},
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