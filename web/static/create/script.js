document.addEventListener('alpine:init', () => {
    Alpine.data('form', () => ({
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

        submit() {
            console.log(this.workExperience);
        },
    }));
});
