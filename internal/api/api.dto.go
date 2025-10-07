package api

import (
	"encoding/json"
)

type HealthResponse struct {
	Message string `json:"message"`
}

type FetchDummyDataResponse struct {
	Data json.RawMessage `json:"data"`
}
