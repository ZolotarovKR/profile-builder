document.addEventListener('alpine:init', () => {
    Alpine.data('cv', () => ({
        id: 0,
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

        async init() {
            this.extractId();
            await this.fetchSection('personalDetails');
        },
        extractId() {
            let params = new URLSearchParams(window.location.search);
            this.id = Number(params.get('id'));
        },
        async fetchSection(s) {
            if (this.relevantSections.includes(s)) {
                return;
            }
            switch (s) {
                case 'personalDetails':
                    this.personalDetails = await (await fetch('/api/personal-details/' + this.id)).json();
                    break;
                case 'contacts':
                    this.contacts = await (await fetch('/api/contacts/' + this.id)).json();
                    break;
                case 'education':
                    this.education = await (await fetch('/api/education/' + this.id)).json();
                    break;
                case 'workExperience':
                    this.workExperience = await (await fetch('/api/work-experience/' + this.id)).json();
                    break;
                case 'languages':
                    this.languages = await (await fetch('/api/languages/' + this.id)).json();
                    break;
                case 'programmingLanguages':
                    this.programmingLanguages = await (await fetch('/api/programming-languages/' + this.id)).json();
                    break;
                default:
                    return;
            }
            this.relevantSections.push(s);
        },
    }));
});
