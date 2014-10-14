package main

import (
  "testing"
  "encoding/json"
  "github.com/siliconvortex/golang-restexample/produce"
)

func TestGetJsonResponse(t *testing.T) {
  var j produce.Payload 
  resp, _ := getJsonResponse()
  err := json.Unmarshal(resp, &j)

  if err != nil {
    t.Errorf("Failed to unmarshal into produce.Payload object")
  }
}
