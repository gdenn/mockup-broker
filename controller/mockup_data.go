package controller

import (
  "encoding/json"
  "log"
)

type MockupData struct {
}

var mockupData MockupData

func ParseMockupData() {
	err := json.Unmarshal([]byte("./mockup_data.json"), &mockupData)
  if err != nil {
    log.Fatal(err)
  }
}
