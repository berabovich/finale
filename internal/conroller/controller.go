package conroller

import (
	"Finale/internal/result"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func Server() {
	r := mux.NewRouter()
	r.Host("http://localhost:8080")
	r.HandleFunc("/api", handleConnection)
	//r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.Dir("internal/web"))))
	http.Handle("/web/", http.FileServer(http.Dir("./internal/web")))
	r.HandleFunc("/", homePage)

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
	_, err = w.Write(res)
	if err != nil {
		return
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/status_page.html")
	t.Execute(w, nil)
}
