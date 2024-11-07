package storage

import (
	"regexp"
	"slices"

	"github.com/ZolotarovKR/profile-builder/internal/models"
)

var dateRegexp = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}|\d{4}-\d{2}|\d{4}$`)
var fullNameRegexp = regexp.MustCompile(`^\w+( \w+){1,2}$`)
var emailRegexp = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
var phoneNumberRegexp = regexp.MustCompile(`^(\+380 \d{2}|0\d{2}) \d{3}[\- ]\d{2}[\- ]\d{2}$`)

var knownLocations = []models.Location{
	{Country: "Ukraine", City: "Kyiv"},
	{Country: "Ukraine", City: "Kharkiv"},
	{Country: "Ukraine", City: "Lviv"},
	{Country: "Ukraine", City: "Odesa"},
	{Country: "Germany", City: "Berlin"},
	{Country: "Germany", City: "Hamburg"},
}
var knownGenders = []string{
	"Male",
	"Female",
	"Other",
}
var knownCitizenships = []string{
	"Ukrainian",
	"German",
}
var knownMaritalStatuses = []string{
	"Single",
	"Married",
}
var knownLanguages = []string{
	"Ukrainian",
	"English",
	"German",
}
var knownLanguageLevels = []string{
	"Native",
	"A1", "A2",
	"B1", "B2",
	"C1", "C2",
}
var knownProgrammingLanguages = []string{
	"C++", "C#", "Java",
	"Python", "Ruby",
	"C", "Rust", "Golang",
}
var knownProgrammingLanguageLevels = []string{
	"Trainee",
	"Junior",
	"Middle",
	"Senior",
}

func checkPersonalDetails(pd models.PersonalDetails) error {
	if !fullNameRegexp.MatchString(pd.FullName) {
		return NewValidationError("Illegal name")
	}
	if !slices.Contains(knownGenders, pd.Gender) {
		return NewValidationError("Unknown gender")
	}
	if pd.Age < 16 || pd.Age > 99 {
		return NewValidationError("Illegal age")
	}
	if !slices.Contains(knownCitizenships, pd.Citizenship) {
		return NewValidationError("Unknown citizenship")
	}
	if !slices.Contains(knownMaritalStatuses, pd.MaritalStatus) {
		return NewValidationError("Unknown marital status")
	}
	return nil
}

func checkContacts(c models.Contacts) error {
	if len(c.Emails) < 1 {
		return NewValidationError("At least one email address must be supplied")
	}
	for _, e := range c.Emails {
		if !emailRegexp.MatchString(e) {
			return NewValidationError("Illegal email")
		}
	}
	if len(c.PhoneNumbers) < 1 {
		return NewValidationError("At least one phone number must be supplied")
	}
	for _, pn := range c.PhoneNumbers {
		if !phoneNumberRegexp.MatchString(pn) {
			return NewValidationError("Illegal phone number")
		}
	}
	if !slices.Contains(knownLocations, c.Location) {
		return NewValidationError("Unknown location")
	}
	return nil
}

func checkEducation(cs []models.Course) error {
	if len(cs) < 1 {
		return NewValidationError("At least one course must be supplied")
	}
	for _, c := range cs {
		if len(c.Name) == 0 {
			return NewValidationError("Empty course name")
		}
		if len(c.Institute) == 0 {
			return NewValidationError("Empty course institute")
		}
		if !dateRegexp.MatchString(c.Period.Start) {
			return NewValidationError("Illegal course start date")
		}
		if !dateRegexp.MatchString(c.Period.End) {
			return NewValidationError("Illegal course end date")
		}
	}
	return nil
}

func checkWorkExperience(jrs []models.JobRecord) error {
	if len(jrs) < 1 {
		return NewValidationError("At least one job record must be supplied")
	}
	for _, jr := range jrs {
		if len(jr.CompanyName) == 0 {
			return NewValidationError("Empty company name")
		}
		if !slices.Contains(knownLocations, jr.Location) {
			return NewValidationError("Unknown location")
		}
		if len(jr.Position) == 0 {
			return NewValidationError("Empty position")
		}
		if !dateRegexp.MatchString(jr.Period.Start) {
			return NewValidationError("Illegal job start date")
		}
		if !dateRegexp.MatchString(jr.Period.End) {
			return NewValidationError("Illegal job end date")
		}
		if len(jr.Achievements) == 0 {
			return NewValidationError("Empty achievements")
		}
	}
	return nil
}

func checkLanguages(ls []models.Language) error {
	if len(ls) < 1 {
		return NewValidationError("At least one language must be supplied")
	}
	for _, l := range ls {
		if !slices.Contains(knownLanguages, l.Name) {
			return NewValidationError("Unknown language")
		}
		if !slices.Contains(knownLanguageLevels, l.Level) {
			return NewValidationError("Unknown language level")
		}
	}
	return nil
}

func checkProgrammingLanguages(pls []models.ProgrammingLanguage) error {
	if len(pls) < 1 {
		return NewValidationError("At least one programming language must be supplied")
	}
	for _, pl := range pls {
		if !slices.Contains(knownProgrammingLanguages, pl.Name) {
			return NewValidationError("Unknown programming language")
		}
		if !slices.Contains(knownProgrammingLanguageLevels, pl.Level) {
			return NewValidationError("Unknown programming language level")
		}
	}
	return nil
}

func checkCv(cv models.Cv) error {
	var err error
	err = checkPersonalDetails(cv.PersonalDetails)
	if err != nil {
		return err
	}
	err = checkContacts(cv.Contacts)
	if err != nil {
		return err
	}
	err = checkEducation(cv.Education)
	if err != nil {
		return err
	}
	err = checkWorkExperience(cv.WorkExperience)
	if err != nil {
		return err
	}
	err = checkLanguages(cv.Languages)
	if err != nil {
		return err
	}
	err = checkProgrammingLanguages(cv.ProgrammingLanguages)
	if err != nil {
		return err
	}
	return nil
}
