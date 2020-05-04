package middleware

import "net/http"

func method(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Allow", method)
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	}
}

func Get(handler http.HandlerFunc) http.HandlerFunc {
	return method("GET", handler)
}

func Delete(handler http.HandlerFunc) http.HandlerFunc {
	return method("DELETE", handler)
}

func Patch(handler http.HandlerFunc) http.HandlerFunc {
	return method("PATCH", handler)
}

func Post(handler http.HandlerFunc) http.HandlerFunc {
	return method("POST", handler)
}

func Put(handler http.HandlerFunc) http.HandlerFunc {
	return method("PUT", handler)
}
