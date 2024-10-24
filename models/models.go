package models

type Period struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Location struct {
	Country, City string
}

type PersonalDetails struct {
	Fullname      string `json:"fullname"`
	Gender        string `json:"gender"`
	Age           uint8  `json:"age"`
	Citizenship   string `json:"citizenship"`
	MaritalStatus string `json:"maritalStatus"`
}

type Contacts struct {
	Emails       []string `json:"emails"`
	PhoneNumbers []string `json:"phoneNumbers"`
	Location     Location `json:"location"`
}

type Course struct {
	Name      string `json:"name"`
	Institute string `json:"institute"`
	Period    Period `json:"period"`
}

type Education struct {
	Courses []Course `json:"courses"`
}

type JobRecord struct {
	CompanyName  string   `json:"companyName"`
	Location     Location `json:"location"`
	Position     string   `json:"position"`
	Period       Period   `json:"period"`
	Achievements string   `json:"achievements"`
}

type Language struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type ProgrammingLanguage struct {
	Name  string `json:"name"`
	Level string `json:"level"`
}

type Cv struct {
	PersonalDetails      PersonalDetails       `json:"personalDetails"`
	Contacts             Contacts              `json:"contacts"`
	Education            Education             `json:"education"`
	WorkExperience       []JobRecord           `json:"workExperience"`
	Languages            []Language            `json:"languages"`
	ProgrammingLanguages []ProgrammingLanguage `json:"programmingLanguages"`
}

type Preview struct {
	Id                   int                   `json:"id"`
	PersonalDetails      PersonalDetails       `json:"personalDetails"`
	Languages            []Language            `json:"languages"`
	ProgrammingLanguages []ProgrammingLanguage `json:"programmingLanguages"`
}
