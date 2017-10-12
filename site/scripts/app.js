const routes = [
    { path: '/', component: HomeView },
    { path: '/users', component: UsersView },
    { path: '/widgets', component: WidgetsView }
]

const router = new VueRouter({
    routes
})

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
    router,
    el: '#page-wrapper'
})
