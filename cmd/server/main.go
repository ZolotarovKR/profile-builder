package main

import (
	"log"
	"net/http"

	"github.com/ZolotarovKR/profile-builder/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/previews", handlers.GetPreviews).Methods("GET")
	r.HandleFunc("/api/cv", handlers.PostCv).Methods("POST")
	r.HandleFunc("/api/cv/{id:[0-9]+}", handlers.DeleteCv).Methods("DELETE")

	r.HandleFunc("/api/personal-details/{id:[0-9]+}", handlers.GetPersonalDetails).Methods("GET")
	r.HandleFunc("/api/personal-details/{id:[0-9]+}", handlers.PutPersonalDetails).Methods("PUT")

	r.HandleFunc("/api/contacts/{id:[0-9]+}", handlers.GetContacts).Methods("GET")
	r.HandleFunc("/api/contacts/{id:[0-9]+}", handlers.PutContacts).Methods("PUT")

	r.HandleFunc("/api/education/{id:[0-9]+}", handlers.GetEducation).Methods("GET")
	r.HandleFunc("/api/education/{id:[0-9]+}", handlers.PutEducation).Methods("PUT")

	r.HandleFunc("/api/work-experience/{id:[0-9]+}", handlers.GetWorkExperience).Methods("GET")
	r.HandleFunc("/api/work-experience/{id:[0-9]+}", handlers.PutWorkExperience).Methods("PUT")

	r.HandleFunc("/api/languages/{id:[0-9]+}", handlers.GetLanguages).Methods("GET")
	r.HandleFunc("/api/languages/{id:[0-9]+}", handlers.PutLanguages).Methods("PUT")

	r.HandleFunc("/api/programming-languages/{id:[0-9]+}", handlers.GetProgrammingLanguages).Methods("GET")
	r.HandleFunc("/api/programming-languages/{id:[0-9]+}", handlers.PutProgrammingLanguages).Methods("PUT")

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
