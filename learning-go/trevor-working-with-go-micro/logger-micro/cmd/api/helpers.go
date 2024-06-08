package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// readJSON is a method on the Config type that reads and decodes JSON from the request body
// into the provided data structure. It also ensures the request body does not exceed a specified size
// and contains only a single JSON object.
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	// Set the maximum number of bytes to be read from the request body to 1 MB
	maxBytes := 1048576 // one megabyte

	// Limit the size of the request body to prevent large payloads from being processed
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Create a new JSON decoder for the request body
	dec := json.NewDecoder(r.Body)

	// Decode the JSON from the request body into the provided data structure
	err := dec.Decode(data)
	if err != nil {
		// Return an error if decoding fails
		return err
	}

	// Try to decode one more JSON object to ensure there is only one JSON value in the body
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		// If the second decode does not return an EOF, it means there is more than one JSON value
		return errors.New("body must have only one single json value")
	}

	// Return nil if everything is successful
	return nil
}

// writeJSON is a method on the Config type that sends a JSON response to the client.
// It takes an HTTP response writer, a status code, the data to be encoded to JSON, and optional headers.
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	// Marshal the provided data into a JSON format.
	out, err := json.Marshal(data)
	if err != nil {
		// Return an error if the data cannot be marshaled to JSON.
		return err
	}

	// If any additional headers are provided, add them to the response.
	if len(headers) > 0 {
		// Loop through the first set of headers (since headers is a variadic parameter).
		for key, value := range headers[0] {
			// Set each header key-value pair on the response writer.
			w.Header()[key] = value
		}
	}

	// Set the Content-Type header to indicate that the response is in JSON format.
	w.Header().Set("Content-Type", "application/json")

	// Write the provided status code to the response.
	w.WriteHeader(status)

	// Write the JSON output to the response body.
	_, err = w.Write(out)
	if err != nil {
		// Return an error if writing the response fails.
		return err
	}

	// Return nil if everything is successful, indicating no errors.
	return nil
}

// errorJSON is a method on the Config type that sends a JSON-formatted error response to the client.
// It takes an HTTP response writer, an error, and an optional status code.
func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	// Default the status code to http.StatusBadRequest (400) if not provided.
	statusCode := http.StatusBadRequest

	// If a status code is provided, use it instead of the default.
	if len(status) > 0 {
		statusCode = status[0]
	}

	// Create a jsonResponse struct to hold the error details.
	var payload jsonResponse
	payload.Error = true          // Indicate that this response is an error.
	payload.Message = err.Error() // Set the error message from the provided error.

	// Use the writeJSON method to send the JSON response with the appropriate status code.
	return app.writeJSON(w, statusCode, payload)
}
