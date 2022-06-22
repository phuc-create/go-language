package golesson

import (
	"fmt"
	"net/http"
)

func HelloConnecter(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "server connected\n")

}

func HelloHeader(writer http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(writer, "%v: %v\n", name, h)
		}
	}
}
