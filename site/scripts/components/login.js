const LoginView = Vue.component('login-view', {
    template: '#login-view',
    data: function() {
        return {
            showCreateModal: false,
            user: {}
        }
    },
    methods: {
        checkLogged() {
            this.showCreateModal = !store.getters.isLogged
        },
        actionLogin() {
            var meRouter = router
            store.dispatch(DO_LOGIN, this.user).then(function() {
                meRouter.go('/')
            }).catch(function(error) {
                $('#btnActionModal').notify(error.response.data.message, { elementPosition: 'top center', className: "danger" });
            })
        },
        closeModal() {
            return false
        }

    },
    mounted: function() {
        this.checkLogged()
        this.user = store.state.auth.user
    }
})