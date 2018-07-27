package httpserver

import (
	"io"
	"net/http"
)

func ItemServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "20")
	io.WriteString(w, "20")
}
