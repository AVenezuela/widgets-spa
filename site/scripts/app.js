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
    if (to.path != '/login') {
        store.dispatch(CHECK_LOGGED).then(function() {
            if (store.getters.isLogged) {
                next();
            } else {
                next('login');
            }
        })
    } else {
        next();
    }
})

Vue.config.productionTip = false

Vue.component('modal', {
    template: '#modal-template',
    props: ['model', 'cancelText', 'actionText', 'modalSize']
})

Vue.component('pagination', {
    template: '#pagination-template',
    data: function() {
        return {
            totalPages: 1,
            currentPage: 1
        }
    },
    props: {
        list: {
            required: true
        }
    },
    mounted: function() {
        if (this.list) this.calcPages()
    },
    methods: {
        calcPages: function() {
            this.totalPages = Math.ceil((this.list.length / this.perPage), -1);
        },
        showPagination: function() {
            return (this.totalPages > 1)
        },
        setPage: function(_page) {

        },
        nextPage: function() {
            this.setPage(this.currentPage + 1);
        },
        previousPage: function() {
            this.setPage(this.currentPage - 1);
        },
        lastPage: function() {
            this.setPage(this.totalPages);
        },
        showPrevious: function() {
            return (this.currentPage != 1)
        },
        showNext: function() {
            return (this.currentPage < this.totalPages)
        }
    }
})

new Vue({
    router,
    el: '#page-wrapper'
})