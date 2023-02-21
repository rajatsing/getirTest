package testing

import (
	"bytes"
	"encoding/json"
	"getir/inmemory"
	"getir/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
	TestLocalMemory_GetInMemoryHandler
	- Test the POST request
	- Test the GET request
	Matches the expected response with the actual response
*/

func TestLocalMemory_GetInMemoryHandler(t *testing.T) {
	payload := models.Payload{
		Key:   "test",
		Value: "some value",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// making a POST request
	req, err := http.NewRequest("POST", "/in-memory", bytes.NewReader(jsonPayload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	localMemory := &inmemory.LocalMemory{Data: make(map[string]string)}
	handler := http.HandlerFunc(localMemory.GetInMemoryHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"key":"test","value":"some value"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// making a GET request
	req, err = http.NewRequest("GET", "/in-memory?key=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(localMemory.GetInMemoryHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected = `{"key":"test","value":"some value"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

}
