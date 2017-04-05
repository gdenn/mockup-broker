package controller

type MockupData struct {
  
}

var mockupData MockupData

func ParseMockupData() &MockupData {
  err := json.Unmarshal("./mockup_data.json", &mockupData)
}
