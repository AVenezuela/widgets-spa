const routes = [
    { path: '/', component: HomeView },
    { path: '/users', component: UsersView },
    { path: '/widgets', component: WidgetsView },
    { path: '/login', component: LoginView }
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, from, next) => {
    if(to.path != '/login') {
        if(store.getters.isLogged) {             
            next();
        } else {            
            next('login');
        }
    } else {        
        next();
    }
})

Vue.config.productionTip = false

Vue.component('modal', {
    template: '#modal-template',
    props: ['model']
})

new Vue({
    router,
    el: '#page-wrapper'
})