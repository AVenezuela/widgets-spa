Vue.component('widget-widget', {
    template: '#widget-widget',
    data: function() {
        return {
            gridColumns: ['id', 'name']
        }
    },
    computed: {
        widgets() {
            return store.state.moduleWidget.list
        }
    },
    mounted: function() {
        store.dispatch(SET_WIDGETS)
    }
})