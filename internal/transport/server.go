package transport

import "net/http"

func InitializeServer() {
	http.ListenAndServe(":8000", nil)
}
