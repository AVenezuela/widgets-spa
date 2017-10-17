Vue.component('logout-view', {
    template: '#logout-view',
    data:function(){
        return  {
            showLogOut:false
        }
    },
    methods:{
        checkLogged (){
            this.showLogOut = store.getters.isLogged
        },
        actionLogout (){
            store.dispatch(DO_LOGOUT)
            router.go('/login');
        }

    },
    mounted:function(){
        this.checkLogged()
    }
})