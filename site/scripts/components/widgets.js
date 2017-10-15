const WidgetsView = Vue.component('widgets-view', {
    template: '#widgets-view',
    data:function(){
        return  {
            showCreateModal:false,
            widget:{}
        }
    },
    computed: {        
        widgets() {
            return store.state.moduleWidget.list
        },
        saveStatus(){
            return store.state.moduleWidget.saveStatus
        }
    },
    mounted:function(){
        store.dispatch('setWidgets')
        this.widget = this.getNewWidget()        
    },
    methods:{
        saveWidget(){
            var me = this            
            store.dispatch('saveWidget', me.widget).then(function(response){
                alert(response)
                me.widget = me.getNewWidget()
                this.showCreateModal = false
            }).catch(function(error){
                alert(error)
            })
        },
        closeModal(){
            this.showCreateModal = false
            this.widget = this.getNewWidget()
            store.commit(SAVE_RESET)
        },
        getNewWidget(){
            return JSON.parse(JSON.stringify({
                id:null,
                name:'',
                color:'',
                price:'',
                inventory:null,
                melts:false
            }))
        }
    }
})