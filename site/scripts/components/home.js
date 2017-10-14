const HomeView = Vue.component('home-view', {
    template: '#home-view',
    computed: {
        usersCounter () {
            return store.state.moduleUser.list.length
        },
        widgetsCounter (){
            return store.state.moduleWidget.list.length
        }
    }
})