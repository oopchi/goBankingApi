package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func writeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	var err error

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(code)
		err = xml.NewEncoder(w).Encode(data)

	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		err = json.NewEncoder(w).Encode(data)
	}

	if err != nil {
		panic(err)
	}

}
