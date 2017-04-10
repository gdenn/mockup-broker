package controller

import (
  "net/http"
)

type Catalog struct {
}

func GetCatalog(w http.ResponseWriter, r *http.Request) {
	// return demo catalog
}
