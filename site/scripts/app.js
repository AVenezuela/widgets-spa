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
    mounted: function(){
        this.disableActionButton = false
        this.disableActionButton = false
    },
    props: ['model', 'cancelText', 'actionText', 'modalSize', 'showActionButton', 'showCancelButton', 'disableActionButton', 'disableCancelButton']
})

Vue.component('pagination', {
    template: '#pagination-template',
    data: function(){
        return {
            currentPage: 1
        }
    },    
    props: {
        data: {
            required: true
        },        
        perPage: Number        
    },
    computed:{
        totalPages: function() {            
            return Math.ceil((this.data.length / this.perPage), -1);
        },
        showPagination: function() {
            return (this.totalPages > 1)
        }
    },
    methods: {                
        setPage: function(_page) {
            this.currentPage = _page
            this.$emit('page-change', _page)
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

Vue.use(VeeValidate)

new Vue({
    router,
    el: '#page-wrapper'
})