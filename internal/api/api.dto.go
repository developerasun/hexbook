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

type QRCodeDataDto struct {
	Wallet  string `form:"wallet" binding:"required"`
	Amount  string `form:"amount" binding:"required"`
	AppType string `form:"apptype" binding:"required"`
}
