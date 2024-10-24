package main

import (
	"log"
	"net/http"

	"github.com/ZolotarovKR/profile-builder/routers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/cv", routers.PostCv).Methods("POST")
	r.HandleFunc("/api/cv", routers.DeleteCv).Methods("DELETE")
	r.HandleFunc("/api/previews", routers.GetPreviews).Methods("GET")

	r.HandleFunc("/api/personal-details/{id:[0-9]+}", routers.GetPersonalDetails).Methods("GET")
	r.HandleFunc("/api/personal-details/{id:[0-9]+}", routers.PutPersonalDetails).Methods("PUT")

	r.HandleFunc("/api/contacts/{id:[0-9]+}", routers.GetContacts).Methods("GET")
	r.HandleFunc("/api/contacts/{id:[0-9]+}", routers.PutContacts).Methods("PUT")

	r.HandleFunc("/api/education/{id:[0-9]+}", routers.GetEducation).Methods("GET")
	r.HandleFunc("/api/education/{id:[0-9]+}", routers.PutEducation).Methods("PUT")

	r.HandleFunc("/api/work-experience/{id:[0-9]+}", routers.GetWorkExperience).Methods("GET")
	r.HandleFunc("/api/work-experience/{id:[0-9]+}", routers.PutWorkExperience).Methods("PUT")

	r.HandleFunc("/api/languages/{id:[0-9]+}", routers.GetLanguages).Methods("GET")
	r.HandleFunc("/api/languages/{id:[0-9]+}", routers.PutLanguages).Methods("PUT")

	r.HandleFunc("/api/programming-languages/{id:[0-9]+}", routers.GetProgrammingLanguages).Methods("GET")
	r.HandleFunc("/api/programming-languages/{id:[0-9]+}", routers.PutProgrammingLanguages).Methods("PUT")

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
