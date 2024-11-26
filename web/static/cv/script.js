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
            let params = new URLSearchParams(window.location.search);
            this.id = Number(params.get('id'));
            await this.fetchActiveSection();
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
            if (this.inEditMode) return;

            this.activeSection = s;
            await this.fetchActiveSection();
        },
        async fetchActiveSection() {
            if (this.relevantSections.includes(this.activeSection)) return;

            let url = this.getActiveSectionUrl();

            let resp = await fetch(url);
            if (resp.status !== 200) {
                alert(await resp.text());
                return;
            }
            let data = await resp.json();
            switch (this.activeSection) {
                case 'personalDetails':
                    this.personalDetails = data;
                    break;
                case 'contacts':
                    this.contacts = data;
                    break;
                case 'education':
                    this.education = data;
                    break;
                case 'workExperience':
                    this.workExperience = data;
                    break;
                case 'languages':
                    this.languages = data;
                    break;
                case 'programmingLanguages':
                    this.programmingLanguages = data;
                    break;
                default:
                    return;
            }
            this.relevantSections.push(this.activeSection);
        },
        async saveActiveSection() {
            let url = this.getActiveSectionUrl();
            let data;
            switch (this.activeSection) {
                case 'personalDetails':
                    data = structuredClone(Alpine.raw(this.pd));
                    data.age = Number(data.age);
                    break;
                case 'contacts':
                    data = structuredClone(Alpine.raw(this.c));
                    break;
                case 'education':
                    data = structuredClone(Alpine.raw(this.e));
                    break;
                case 'workExperience':
                    data = structuredClone(Alpine.raw(this.we));
                    break;
                case 'languages':
                    data = structuredClone(Alpine.raw(this.ls));
                    break;
                case 'programmingLanguages':
                    data = structuredClone(Alpine.raw(this.pls));
                    break;
                default:
                    console.error('Unknown section: ' + s);
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
            await this.fetchActiveSection();
            this.inEditMode = false;
        },
        async deleteThisCV() {
            if (!confirm('Are you sure?')) return;

            let resp = await fetch('/api/cv/' + this.id, {
                method: 'DELETE',
            });
            if (resp.status !== 200) {
                alert(await resp.text());
                return;
            }
            window.location.href = '/';
        },
        getActiveSectionUrl() {
            switch (this.activeSection) {
                case 'personalDetails':
                    return '/api/personal-details/' + this.id;
                case 'contacts':
                    return '/api/contacts/' + this.id;
                case 'education':
                    return '/api/education/' + this.id;
                case 'workExperience':
                    return '/api/work-experience/' + this.id;
                case 'languages':
                    return '/api/languages/' + this.id;
                case 'programmingLanguages':
                    return '/api/programming-languages/' + this.id;
                default:
                    console.error('Unknown section: ' + s);
                    return null;
            }
        },
    }));
});
