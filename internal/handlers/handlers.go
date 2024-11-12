package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ZolotarovKR/profile-builder/internal/models"
	"github.com/ZolotarovKR/profile-builder/internal/storage"
	"github.com/gorilla/mux"
)

func GetPreviews(w http.ResponseWriter, r *http.Request) {
	ps := storage.ReadPreviews()
	err := json.NewEncoder(w).Encode(ps)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PostCv(w http.ResponseWriter, r *http.Request) {
	var cv models.Cv
	err := json.NewDecoder(r.Body).Decode(&cv)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	_, err = storage.SaveCv(cv)
	if err != nil {
		switch err.(type) {
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func DeleteCv(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
	}

	err = storage.DeleteCv(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetPersonalDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pd, err := storage.ReadPersonalDetails(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(pd)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutPersonalDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var pd models.PersonalDetails
	err = json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdatePersonalDetails(id, pd)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	c, err := storage.ReadContacts(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutContacts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var c models.Contacts
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdateContacts(id, c)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	cs, err := storage.ReadEducation(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(cs)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutEducation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var cs []models.Course
	err = json.NewDecoder(r.Body).Decode(&cs)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdateEducation(id, cs)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetWorkExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	jrs, err := storage.ReadWorkExperiece(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(jrs)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutWorkExperience(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var jrs []models.JobRecord
	err = json.NewDecoder(r.Body).Decode(&jrs)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdateWorkExperience(id, jrs)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetLanguages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	ls, err := storage.ReadLanguages(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(ls)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutLanguages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var ls []models.Language
	err = json.NewDecoder(r.Body).Decode(&ls)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdateLanguages(id, ls)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}

func GetProgrammingLanguages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	pls, err := storage.ReadProgrammingLanguages(id)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(pls)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func PutProgrammingLanguages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var pls []models.ProgrammingLanguage
	err = json.NewDecoder(r.Body).Decode(&pls)
	if err != nil {
		http.Error(w, "Request body cannot be parsed", http.StatusBadRequest)
		return
	}

	err = storage.UpdateProgrammingLanguages(id, pls)
	if err != nil {
		switch err.(type) {
		case storage.NotFoundError:
			http.NotFound(w, r)
			return
		case storage.ValidationError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		default:
			log.Print(err.Error())
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}
	}
}
