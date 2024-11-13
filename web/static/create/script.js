document.addEventListener('alpine:init', () => {
    Alpine.data('form', () => ({
        personalDetails: {
            fullName: '',
            gender: '',
            age: 18,
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

        async submit() {
            const data = {};
            data.personalDetails = Alpine.raw(this.personalDetails);
            data.contacts = Alpine.raw(this.contacts);
            data.education = Alpine.raw(this.education);
            data.workExperience = Alpine.raw(this.workExperience);
            data.languages = Alpine.raw(this.languages);
            data.programmingLanguages = Alpine.raw(this.programmingLanguages);
            let resp = await fetch('/api/cv', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            switch (resp.status) {
                case 201:
                    alert('Created');
                    break;
                case 400:
                    alert(await resp.text());
                    break;
            }
        },
    }));
});
