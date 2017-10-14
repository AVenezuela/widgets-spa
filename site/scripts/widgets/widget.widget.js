Vue.component('widget-widget', {
    template: '#widget-widget',
    computed: {        
        widgets(){
            return store.state.moduleWidget.list        
        } 
    },    
    mounted:function(){
        store.dispatch('setWidgets')
    }
})