package conroller

import (
	"Finale/result"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Server() {
	r := mux.NewRouter()
	r.Host("http://localhost:8080")
	r.HandleFunc("/api", handleConnection)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		return
	}
}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	a := result.ResultT{}
	b := result.GetResultData()
	if b.SMS != nil && b.MMS != nil && b.VoiceCall != nil && b.Email != nil && b.Incidents != nil && b.Support != nil {
		a.Status = true
		a.Data = b
	} else {
		a.Error = "Error on collect data"
	}

	res, err := json.Marshal(a)
	if err != nil {
		return
	}
	//fmt.Fprintf(w, "Category: %v\n", vars["category"])
	_, err = w.Write(res)
	if err != nil {
		return
	}
}
