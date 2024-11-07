package storage

import (
	"sync"

	"github.com/ZolotarovKR/profile-builder/internal/models"
)

var (
	cvs    map[int]models.Cv = make(map[int]models.Cv)
	nextId int
	m      sync.Mutex
)

func ReadPreviews() []models.Preview {
	m.Lock()
	defer m.Unlock()

	r := make([]models.Preview, 0, len(cvs))
	for id, cv := range cvs {
		p := models.Preview{
			Id:                   id,
			PersonalDetails:      cv.PersonalDetails,
			Languages:            cv.Languages,
			ProgrammingLanguages: cv.ProgrammingLanguages,
		}
		r = append(r, p)
	}
	return r
}

func SaveCv(cv models.Cv) (int, error) {
	err := checkCv(cv)
	if err != nil {
		return -1, err
	}

	m.Lock()
	defer m.Unlock()

	cvs[nextId] = cv
	nextId += 1
	return nextId - 1, nil
}

func DeleteCv(id int) error {
	m.Lock()
	defer m.Unlock()

	_, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	delete(cvs, id)
	return nil
}

func ReadPersonalDetails(id int) (models.PersonalDetails, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.PersonalDetails, NewNotFoundError()
	}

	return cv.PersonalDetails, nil
}

func UpdatePersonalDetails(id int, pd models.PersonalDetails) error {
	err := checkPersonalDetails(pd)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.PersonalDetails = pd
	cvs[id] = cv
	return nil
}

func ReadContacts(id int) (models.Contacts, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.Contacts, NewNotFoundError()
	}

	return cv.Contacts, nil
}

func UpdateContacts(id int, c models.Contacts) error {
	err := checkContacts(c)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.Contacts = c
	cvs[id] = cv
	return nil
}

func ReadEducation(id int) ([]models.Course, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.Education, NewNotFoundError()
	}

	return cv.Education, nil
}

func UpdateEducation(id int, cs []models.Course) error {
	err := checkEducation(cs)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.Education = cs
	cvs[id] = cv
	return nil
}

func ReadWorkExperiece(id int) ([]models.JobRecord, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.WorkExperience, NewNotFoundError()
	}

	return cv.WorkExperience, nil
}

func UpdateWorkExperience(id int, jrs []models.JobRecord) error {
	err := checkWorkExperience(jrs)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.WorkExperience = jrs
	cvs[id] = cv
	return nil
}

func ReadLanguages(id int) ([]models.Language, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.Languages, NewNotFoundError()
	}

	return cv.Languages, nil
}

func UpdateLanguages(id int, ls []models.Language) error {
	err := checkLanguages(ls)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.Languages = ls
	cvs[id] = cv
	return nil
}

func ReadProgrammingLanguages(id int) ([]models.ProgrammingLanguage, error) {
	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return cv.ProgrammingLanguages, NewNotFoundError()
	}

	return cv.ProgrammingLanguages, nil
}

func UpdateProgrammingLanguages(id int, pls []models.ProgrammingLanguage) error {
	err := checkProgrammingLanguages(pls)
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	cv, ok := cvs[id]
	if !ok {
		return NewNotFoundError()
	}

	cv.ProgrammingLanguages = pls
	cvs[id] = cv
	return nil
}
