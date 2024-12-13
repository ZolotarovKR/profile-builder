document.addEventListener('alpine:init', () => {
    Alpine.store('id', new URLSearchParams(window.location.search).get('id'));
    Alpine.store('activeSection', 'personalDetails');
    Alpine.store('inEditMode', false);
    Alpine.store('deleteCv', async () => {
        if (!confirm('Are you sure?')) return;
        await fetch('/api/cv/' + Alpine.store('id'), {
            method: 'DELETE',
        });
        window.location.href = '/';
    });
    Alpine.data('personalDetails', () => ({
        isUpToDate: false,
        fullName: '',
        gender: '',
        age: '',
        citizenship: '',
        maritalStatus: '',
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/personal-details/' + Alpine.store('id'));
            Object.assign(this, await resp.json());
            this.isUpToDate = true;
        },
    }));
    Alpine.data('personalDetailsForm', () => ({
        fullName: '',
        gender: '',
        age: '',
        citizenship: '',
        maritalStatus: '',
        async save() {
            const data = {};
            data.fullName = this.fullName;
            data.gender = this.gender;
            data.age = Number(this.age);
            data.citizenship = this.citizenship;
            data.maritalStatus = this.maritalStatus;
            const resp = await fetch('/api/personal-details/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
    Alpine.data('contacts', () => ({
        isUpToDate: false,
        emails: [],
        phoneNumbers: [],
        location: {
            country: '',
            city: '',
        },
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/contacts/' + Alpine.store('id'));
            Object.assign(this, await resp.json());
            this.isUpToDate = true;
        },
    }));
    Alpine.data('contactsForm', () => ({
        emails: [],
        phoneNumbers: [],
        location: {
            country: '',
            city: '',
        },
        async save() {
            const data = {};
            data.emails = this.emails;
            data.phoneNumbers = this.phoneNumbers;
            data.location = this.location;
            const resp = await fetch('/api/contacts/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
    Alpine.data('education', () => ({
        isUpToDate: false,
        courses: [],
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/education/' + Alpine.store('id'));
            this.courses = await resp.json();
            this.isUpToDate = true;
        },
    }));
    Alpine.data('educationForm', () => ({
        courses: [],
        async save() {
            const data = this.courses;
            const resp = await fetch('/api/education/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
    Alpine.data('workExperience', () => ({
        isUpToDate: false,
        jobRecords: [],
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/work-experience/' + Alpine.store('id'));
            this.jobRecords = await resp.json();
            this.isUpToDate = true;
        },
    }));
    Alpine.data('workExperienceForm', () => ({
        jobRecords: [],
        async save() {
            const data = this.jobRecords;
            const resp = await fetch('/api/work-experience/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
    Alpine.data('languages', () => ({
        isUpToDate: false,
        languages: [],
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/languages/' + Alpine.store('id'));
            this.languages = await resp.json();
            this.isUpToDate = true;
        },
    }));
    Alpine.data('languagesForm', () => ({
        languages: [],
        async save() {
            const data = this.languages;
            const resp = await fetch('/api/languages/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
    Alpine.data('programmingLanguages', () => ({
        isUpToDate: false,
        programmingLanguages: [],
        async get() {
            if (this.isUpToDate) return;
            const resp = await fetch('/api/programming-languages/' + Alpine.store('id'));
            this.programmingLanguages = await resp.json();
            this.isUpToDate = true;
        },
    }));
    Alpine.data('programmingLanguagesForm', () => ({
        programmingLanguages: [],
        async save() {
            const data = this.programmingLanguages;
            const resp = await fetch('/api/programming-languages/' + Alpine.store('id'), {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status === 400) alert(await resp.text());
            return resp.status;
        },
    }));
});
