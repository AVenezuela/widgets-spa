Vue.component('widget-widget', {
    template: '#widget-widget',
    data:function(){
        return  {
            searchFor:''
        }
    },
    computed: {        
        widgets(){            
            var data = store.state.moduleWidget.list
            var searchKey = this.searchFor && this.searchFor.toLowerCase()
            var searchKey = this.searchFor && this.searchFor.toLowerCase()
            if(searchKey){
                data = data.filter(function (row) {
                        return Object.keys(row).some(function (key) {
                            return String(row[key]).toLowerCase().indexOf(searchKey) > -1
                        })
                })
            }
            return data     
        } 
    },    
    mounted:function(){
        store.dispatch(SET_WIDGETS)
    }
})