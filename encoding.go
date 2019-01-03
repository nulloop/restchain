package restchain

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON(r io.Reader, ptr interface{}) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(ptr)
}

func WriteJSON(ptr interface{}, w io.Writer) error {
	switch v := ptr.(type) {
	case Error:
		return json.NewEncoder(w).Encode(v)
	case error:
		return json.NewEncoder(w).Encode(NewError(500, v.Error(), nil))
	default:
		return json.NewEncoder(w).Encode(v)
	}
}

func WriteResponse(ptr interface{}, w http.ResponseWriter) error {
	switch v := ptr.(type) {
	case Error:
		w.WriteHeader(v.Code)
		return WriteJSON(v, w)
	case error:
		w.WriteHeader(500)
		return WriteJSON(v, w)
	default:
		w.WriteHeader(200)
		return WriteJSON(v, w)
	}
}
