package handler

import "encoding/json"

func jsonError(mgs string) []byte {
	e := struct {
		Message string `json:"message"`
	}{
		mgs,
	}
	r, err := json.Marshal(e)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
