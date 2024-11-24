document.addEventListener('alpine:init', () => {
    Alpine.data('cv', () => ({
        id: 0,
        activeSection: 'personalDetails',
        relevantSections: [],
        inEditMode: false,

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

        pd: {
            fullName: '',
            gender: '',
            age: '',
            citizenship: '',
            maritalStatus: '',
        },
        c: {
            emails: [],
            phoneNumbers: [],
            location: {
                country: '',
                city: '',
            },
        },
        e: [],
        we: [],
        ls: [],
        pls: [],

        async init() {
            this.extractId();
            await this.fetchSection('personalDetails');
        },
        extractId() {
            let params = new URLSearchParams(window.location.search);
            this.id = Number(params.get('id'));
        },
        enterEditMode() {
            switch (this.activeSection) {
                case 'personalDetails':
                    this.pd = structuredClone(Alpine.raw(this.personalDetails));
                    break;
                case 'contacts':
                    this.c = structuredClone(Alpine.raw(this.contacts));
                    break;
                case 'education':
                    this.e = structuredClone(Alpine.raw(this.education));
                    break;
                case 'workExperience':
                    this.we = structuredClone(Alpine.raw(this.workExperience));
                    break;
                case 'languages':
                    this.ls = structuredClone(Alpine.raw(this.languages));
                    break;
                case 'programmingLanguages':
                    this.pls = structuredClone(Alpine.raw(this.programmingLanguages));
                    break;
                default:
                    return;
            }
            this.inEditMode = true;
        },
        async selectSection(s) {
            if (this.inEditMode) {
                return;
            }
            await this.fetchSection(s);
            if (this.relevantSections.includes(s)) {
                this.activeSection = s;
            }
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
        async saveActiveSection() {
            let url;
            let data;
            switch (this.activeSection) {
                case 'personalDetails':
                    url = '/api/personal-details/' + this.id;
                    data = structuredClone(Alpine.raw(this.pd));
                    data.age = Number(data.age);
                    break;
                case 'contacts':
                    url = '/api/contacts/' + this.id;
                    data = structuredClone(Alpine.raw(this.c));
                    break;
                case 'education':
                    url = '/api/education/' + this.id;
                    data = structuredClone(Alpine.raw(this.e));
                    break;
                case 'workExperience':
                    url = '/api/work-experience/' + this.id;
                    data = structuredClone(Alpine.raw(this.we));
                    break;
                case 'languages':
                    url = '/api/languages/' + this.id;
                    data = structuredClone(Alpine.raw(this.ls));
                    break;
                case 'programmingLanguages':
                    url = '/api/programming-languages/' + this.id;
                    data = structuredClone(Alpine.raw(this.pls));
                    break;
                default:
                    return;
            }
            let resp = await fetch(url, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data),
            });
            if (resp.status !== 200) {
                alert(await resp.text());
                return;
            }
            this.relevantSections.splice(this.relevantSections.indexOf(this.activeSection), 1);
            this.fetchSection(this.activeSection);
            this.inEditMode = false;
        },
    }));
});
