package conroller

import (
	"Finale/result"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Server() {
	r := mux.NewRouter()
	r.Host("http://localhost:8282")
	r.HandleFunc("/", handleConnection)
	http.ListenAndServe("localhost:8282", r)
}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
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
	w.Write(res)
}
