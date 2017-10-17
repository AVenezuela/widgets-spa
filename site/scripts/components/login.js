const LoginView = Vue.component('login-view', {
    template: '#login-view',
    data:function(){
        return  {
            showCreateModal:false,
            user:{}
        }
    },
    methods:{
        checkLogged (){
            this.showCreateModal = !store.getters.isLogged
        },
        actionLogin (){
            store.dispatch(DO_LOGIN, this.user).then(function(){
                router.go('/users')
            }).catch(function(){
                router.go('/login')
            })
        },
        closeModal (){
            return false
        }

    },
    mounted:function(){
        this.checkLogged()
        this.user = store.state.auth.user
    }
})