const WidgetsView = Vue.component('widgets-view', {
    template: '#widgets-view',
    data:function(){
        return  {
            showCreateModal:false
        }
    },
    computed: {        
        widgets() {
            return store.state.moduleWidget.list
        },
        widget(){
            return store.state.moduleWidget.model
        }
    },
    mounted:function(){
        store.dispatch('setWidgets')
    },
    methods:{
        saveWidget(){
            store.dispatch('saveWidget')
        }
    }
})

const ModalCreateWidget = Vue.component('create-widget', {
    template: '#modal-createwidget',
    props:['model']
}) 