package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}