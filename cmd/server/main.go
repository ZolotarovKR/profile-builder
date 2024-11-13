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
	r.HandleFunc("/api/cv/{id}", handlers.DeleteCv).Methods("DELETE")

	r.HandleFunc("/api/personal-details/{id}", handlers.GetPersonalDetails).Methods("GET")
	r.HandleFunc("/api/personal-details/{id}", handlers.PutPersonalDetails).Methods("PUT")

	r.HandleFunc("/api/contacts/{id}", handlers.GetContacts).Methods("GET")
	r.HandleFunc("/api/contacts/{id}", handlers.PutContacts).Methods("PUT")

	r.HandleFunc("/api/education/{id}", handlers.GetEducation).Methods("GET")
	r.HandleFunc("/api/education/{id}", handlers.PutEducation).Methods("PUT")

	r.HandleFunc("/api/work-experience/{id}", handlers.GetWorkExperience).Methods("GET")
	r.HandleFunc("/api/work-experience/{id}", handlers.PutWorkExperience).Methods("PUT")

	r.HandleFunc("/api/languages/{id}", handlers.GetLanguages).Methods("GET")
	r.HandleFunc("/api/languages/{id}", handlers.PutLanguages).Methods("PUT")

	r.HandleFunc("/api/programming-languages/{id}", handlers.GetProgrammingLanguages).Methods("GET")
	r.HandleFunc("/api/programming-languages/{id}", handlers.PutProgrammingLanguages).Methods("PUT")

	fs := http.FileServer(http.Dir("./web/static"))
	r.PathPrefix("/").Handler(fs)

	log.Print("Starting server at :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}
