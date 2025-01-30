package basic-go-web-app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Providing useful helpers for encoding and decoding allows for a quick and easy way to always
// return from handlers. At the basic level, just providing a basic `encode` and `decode`.
// Some prefer a way to encode certain responses (think OK, InternalError, etc) and that can be
// useful if you prefer that. Provided here is the basics.
//
// NOTE: This is assuming and basing off of typical web app request/response in JSON format. Also a
// world where you need different formats. If so, consider specific format helpers or some kind of flag param.

// encode takes the given value and encodes it to JSON with a status code. Compiler can infer the type to be encoded here.
func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// decode takes the request and decodes the body from JSON to the given type. You will need to specify type when
// calling decode. I find this to be helpful as you know what you are decoding to:
//
// Example: struct, err := decode[SomeRequest](request)
func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
