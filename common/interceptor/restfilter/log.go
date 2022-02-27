package restfilter

import "net/http"

func HttpLogFilter(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)

	}
}
