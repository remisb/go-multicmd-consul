package health

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service health"))
}
