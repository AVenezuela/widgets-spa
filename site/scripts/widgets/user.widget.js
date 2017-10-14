Vue.component('user-widget', {
    template: '#user-widget',
    computed: {        
        users() {
            return store.state.moduleUser.list
        }
    },
    mounted:function(){
        store.dispatch('setUsers')
    }
})