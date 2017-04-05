package middleware_test

import (
  "testing"
  "encoding/base64"
  "github.com/mockup-broker/middleware"
)

func TestValidateTrue(t *testing.T) {
  username := "John"
  password := "test"

  if !validate(username, password) {
    t.Fatalf("Couldn't validate correct user")
  }
}

func TestValidateFalse(t *testing.T) {
  username := "Jhn"
  password := "test"

  if validate(username, password) {
    t.Fatalf("Validated wrong user")
  }
}

func TestHttpBaseAuthValid(t *testing.T) {
  auth := "John:test"
  auth64 := base64.StdEncoding.EncodeString(auth)
  httpAuth := "Basic " + auth64

  if validate(username, password) {
    t.Fatalf("Validated wrong user")
  }
}
