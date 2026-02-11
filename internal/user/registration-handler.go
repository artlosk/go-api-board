package user

import (
	"encoding/json"
	"io"
	"net/http"
)

func MakeRegHandler(registration *Registration) func(w http.ResponseWriter, r *http.Request) {
	if registration == nil {
		panic("nil registration")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var a Auth
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		err = json.Unmarshal(body, &a)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = registration.Register(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
