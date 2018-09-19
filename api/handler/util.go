package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yasushiasahi/karabiner-members/api/db"
)

func decodeBody(r *http.Request) (db.User, error) {
	d := json.NewDecoder(r.Body)
	var u db.User
	for {
		err := d.Decode(&u)
		if err == io.EOF {
			err = nil
			break
		}
		if err != nil {
			return u, err
		}
	}
	return u, nil
}
