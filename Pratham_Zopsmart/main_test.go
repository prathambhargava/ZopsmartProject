package main_test

import (
  "bytes"
  "net/http"
  "testing"
  "time"

  "gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
  time.Sleep(5 * time.Second) // Adjust based on server startup time

  tests := []struct {
    desc      string
    method    string
    endpoint  string
    statusCode int
    body      []byte
  }{
    // Example tests for a library management API:
    {"get all books", http.MethodGet, "/books", http.StatusOK, nil},
    {"create book", http.MethodPost, "/books", http.StatusCreated, []byte(`{"name": "The Name of the Wind", "author": "Patrick Rothfuss", "student_id": 1234, "status": "Available"}`)},
    {"get specific book", http.MethodGet, "/books/1", http.StatusOK, nil},
    {"update book", http.MethodPut, "/books/1", http.StatusOK, []byte(`{"status": "Borrowed"}`)},
    {"delete book", http.MethodDelete, "/books/1", http.StatusNoContent, nil},
    {"get non-existent book", http.MethodGet, "/books/999", http.StatusNotFound, nil},
    {"bad request for create", http.MethodPost, "/books", http.StatusBadRequest, []byte(`{"invalid_field": "invalid_value"}`)},
  }

  for i, tc := range tests {
    req, _ := request.NewMock(tc.method, "http://localhost:8080"+tc.endpoint, bytes.NewBuffer(tc.body))
    c := http.Client{}

    resp, err := c.Do(req)
    if err != nil {
      t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
      continue 
    }

    if resp.StatusCode != tc.statusCode {
      t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
    }

    _ = resp.Body.Close()
  }
}
