document.addEventListener('alpine:init', () => {
    Alpine.data('form', () => ({
        personalDetails: {
            fullName: '',
            gender: 'Male',
            age: '',
            citizenship: 'Ukrainian',
            maritalStatus: 'Single',
        },
        contacts: {
            emails: [''],
            phoneNumbers: [''],
            location: {
                country: 'Ukraine',
                city: 'Kyiv',
            },
        },
        education: [{ name: '', institute: '', period: { start: '', end: '' } }],
        workExperience: [{ companyName: '', location: { country: 'Ukraine', city: 'Kyiv' }, position: '', period: { start: '', end: '' } }],
        languages: [{ name: 'Ukrainian', level: 'Native' }],
        programmingLanguages: [{ name: 'C++', level: 'Trainee' }],

        async submit() {
            const data = {};
            data.personalDetails = structuredClone(Alpine.raw(this.personalDetails));
            data.contacts = structuredClone(Alpine.raw(this.contacts));
            data.education = structuredClone(Alpine.raw(this.education));
            data.workExperience = structuredClone(Alpine.raw(this.workExperience));
            data.languages = structuredClone(Alpine.raw(this.languages));
            data.programmingLanguages = structuredClone(Alpine.raw(this.programmingLanguages));
            data.personalDetails.age = Number(data.personalDetails.age);
            let resp = await fetch('/api/cv', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            switch (resp.status) {
                case 201:
                    window.location.replace('/');
                    break;
                case 400:
                    alert(await resp.text());
                    break;
            }
        },
    }));
});
