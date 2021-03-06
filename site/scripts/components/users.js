const UsersView = Vue.component('users-view', {
    template: '#users-view',
    mixins: [Mixins],
    data: function() {
        return {
            showCreateModal: false,
            user: {},
            editedUser: null,
            searchFor: '',
            totalPerPage: 10
        }
    },
    computed: {
        users() {
            var filteredData = store.state.moduleUser.list
            var searchKey = this.searchFor && this.searchFor.toLowerCase()
            if (searchKey) {
                filteredData = filteredData.filter(function(row) {
                    return Object.keys(row).some(function(key) {
                        return String(row[key]).toLowerCase().indexOf(searchKey) > -1
                    })
                })
            }
            return filteredData
        },
        searchData() {
            var list = this.users

        }
    },
    mounted: function() {
        store.dispatch(SET_USERS)
        this.user = this.getNewUser()
    },
    methods: {
        actionUser() {
            var me = this
            var ACTION = (me.user.id === null) ? SAVE_USER : UPDATE_USER
            store.dispatch(ACTION, { newUser: me.user, oldUser: me.editedUser }).then(function(response) {
                this.showCreateModal = false
                me.user = me.getNewUser()
                this.editedUser = null
                $('#btnActionModal').notify("Success!", { elementPosition: 'top center', className: "success" });
            }).catch(function(error) {
                $.notify(error.response.data.message, { elementPosition: 'top center', className: "danger" });
            })
        },
        delUser(user) {
            store.dispatch(DEL_USER, user)
        },
        editUser(user) {
            this.editedUser = user
            this.user = JSON.parse(JSON.stringify(user));
            this.showCreateModal = true
        },
        closeModal() {
            this.showCreateModal = false
            this.user = this.getNewUser()
        },
        getNewUser() {
            return JSON.parse(JSON.stringify({
                id: null,
                name: '',
                gravatar: ''
            }))
        },
        closeModal() {
            this.showCreateModal = false
            this.user = this.getNewUser()
        }
    }
})