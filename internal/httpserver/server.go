package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jwoglom/macos-findmy-mqtt/internal/fmipcore"
)

func index(w http.ResponseWriter, req *http.Request) {
	items, err := fmipcore.ReadItems()
	if err != nil {
		fmt.Fprintf(w, "error: %q", err)
		return
	}

	out, err := json.Marshal(items)
	if err != nil {
		fmt.Fprintf(w, "error: %q", err)
		return
	}

	fmt.Fprintf(w, "%s", string(out))
}

func Run(port int) {
	if port == 0 {
		return
	}

	fmt.Printf("webserver listening on port :%d\n", port)

	http.HandleFunc("/", index)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
