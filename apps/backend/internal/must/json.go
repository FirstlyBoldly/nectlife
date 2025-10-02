package must

import (
	"encoding/json"
	"net/http"
)

func MarshalEnforce(w http.ResponseWriter, r *http.Request, v ...any) ([]byte, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return []byte{}, err
	}

	return jsonBytes, nil
}

func UnmarshalEnforce(w http.ResponseWriter, r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return err
	}

	return nil
}
