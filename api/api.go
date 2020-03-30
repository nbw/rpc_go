package api

import (
	"encoding/json"
	"net/http"
)

type MoveResponse struct {
	Weapon string `json:"weapon"`
}

func DecodeMoveResponse(res *http.Response, decoded *MoveResponse) error {
	err := json.NewDecoder(res.Body).Decode(&decoded)
	return err
}
