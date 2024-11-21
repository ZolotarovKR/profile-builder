document.addEventListener('alpine:init', () => {
    Alpine.data('cv', () => ({
        section: 'personalDetails',
        relevantSections: [],

        personalDetails: {
            fullName: '',
            gender: '',
            age: '',
            citizenship: '',
            maritalStatus: '',
        },
        contacts: {
            emails: [],
            phoneNumbers: [],
            location: {
                country: '',
                city: '',
            },
        },
        education: [],
        workExperience: [],
        languages: [],
        programmingLanguages: [],

        async fetchPersonalDetails() {
            if (this.relevantSections.includes('personalDetails')) {
                return;
            }
            let resp = await fetch('/api/personal-details/0');
            console.log(resp.status);
            this.personalDetails = await resp.json();
            this.relevantSections.push('personalDetails');
        },
        async fetchContacts() {
            if (this.relevantSections.includes('contacts')) {
                return;
            }
            let resp = await fetch('/api/contacts/0');
            console.log(resp.status);
            this.contacts = await resp.json();
            this.relevantSections.push('contacts');
        },
        async fetchEducation() {
            if (this.relevantSections.includes('education')) {
                return;
            }
            let resp = await fetch('/api/education/0');
            console.log(resp.status);
            this.education = await resp.json();
            this.relevantSections.push('education');
        },
        async fetchWorkExperience() {
            if (this.relevantSections.includes('workExperience')) {
                return;
            }
            let resp = await fetch('/api/work-experience/0');
            console.log(resp.status);
            this.workExperience = await resp.json();
            this.relevantSections.push('workExperience');
        },
        async fetchLanguages() {
            if (this.relevantSections.includes('languages')) {
                return;
            }
            let resp = await fetch('/api/languages/0');
            console.log(resp.status);
            this.languages = await resp.json();
            this.relevantSections.push('languages');
        },
        async fetchProgrammingLanguages() {
            if (this.relevantSections.includes('programmingLanguages')) {
                return;
            }
            let resp = await fetch('/api/programming-languages/0');
            console.log(resp.status);
            this.programmingLanguages = await resp.json();
            this.relevantSections.push('programmingLanguages');
        },
    }));
});
