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
        }
    }
})

const ModalCreateWidget = Vue.component('create-widget', {
    template: '#modal-createwidget'
  })
  