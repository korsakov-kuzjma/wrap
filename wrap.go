package wrap

import (
	"net/http"
)

func Wrap(fw ...func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, ffw := range fw {
			ffw(w, r)
		}
	}
}
