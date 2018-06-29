package gpkg

import (
	"net/http"
)

func param(r *http.Request, params map[string]string, name string, fallback string) string {
	value := r.URL.Query().Get(name)
	if len(value) == 0 {
		value, ok := params[name]
		if !ok {
			return fallback
		} else {
			return value
		}
	} else {
		return value
	}
}
