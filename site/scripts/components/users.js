const UsersView = Vue.component('users-view', {
    template: '#users-view',
    computed: {        
        users() {
            return store.state.moduleUser.list
        }
    },
    mounted:function(){
        store.dispatch('setUsers')
    }
})