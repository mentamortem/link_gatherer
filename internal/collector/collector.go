package collector

import (
	"fmt"
	"net/http"
	// "log"
	"encoding/json"
)

type Links struct {
	links []string
}

func Run(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/postlinks", postHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var data Links
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

//func checkFormat() string{

// }
