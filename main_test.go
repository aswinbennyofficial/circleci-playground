package main

import (
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
)

func TestHandler(t *testing.T) {
    // Set up a test request
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Create a mock environment variable for testing
    os.Setenv("SECRET_KEY", "test_secret_key")

    // Call the handler function directly, passing in the ResponseRecorder and Request
    handler(rr, req)

    // Check the status code
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body
    expected := "Hello World, The secret env is test_secret_key"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
