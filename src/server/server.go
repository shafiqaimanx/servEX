package server

import (
	"fmt"
	"net/http"
)

type statusCodeResponse struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusCodeResponse) WriteHeader(code int) {
	if w.statusCode == 0 {
		w.statusCode = code
	}
	w.ResponseWriter.WriteHeader(code)
}

func StartTheServer(defaultPort int) {
	fileServer := http.FileServer(http.Dir("./"))

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			statusWriter := &statusCodeResponse{ResponseWriter: w}
			fileServer.ServeHTTP(statusWriter, r)
			loggingFormat := fmt.Sprintf("[%s %s] %s %d -\n", DateFormat(), TimeFormat(), ProtoMethod(r), statusWriter.statusCode)
			fmt.Printf("%s - - %s", GetClientIP(r), loggingFormat)
		} //else {
			//http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		//}
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", defaultPort), nil)
	if err != nil {
		fmt.Println(err)
	}
}