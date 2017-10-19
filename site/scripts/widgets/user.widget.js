Vue.component('user-widget', {
    template: '#user-widget',
    data: function() {
        return {
            gridColumns: ['id', 'name']
        }
    },
    computed: {
        users() {
            return store.state.moduleUser.list
        },
        totalUsers() {
            return this.users.length
        }
    },
    methods: {
        setPage(page) {
            this.page = page
        }
    },
    mounted: function() {
        store.dispatch(SET_USERS)
    }
})