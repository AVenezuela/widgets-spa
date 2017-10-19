const WidgetsView = Vue.component('widgets-view', {
    template: '#widgets-view',
    mixins: [Mixins],
    data: function() {
        return {
            showCreateModal: false,
            widget: {},
            editedWidget: null,
            searchFor: ''
        }
    },
    computed: {
        widgets() {
            var filteredData = store.getters.getSortedWidgets
            var searchKey = this.searchFor && this.searchFor.toLowerCase()
            if (searchKey) {
                filteredData = filteredData.filter(function(row) {
                    return Object.keys(row).some(function(key) {
                        return String(row[key]).toLowerCase().indexOf(searchKey) > -1
                    })
                })
            }
            return filteredData
        }
    },
    mounted: function() {
        store.dispatch(SET_WIDGETS)
        this.widget = this.getNewWidget()
    },
    methods: {
        actionWidget() {
            var me = this
            var ACTION = (me.widget.id === null) ? SAVE_WIDGET : UPDATE_WIDGET
            store.dispatch(ACTION, { newWidget: me.widget, oldWidget: me.editedWidget }).then(function(response) {
                this.showCreateModal = false
                me.widget = me.getNewWidget()
                this.editedWidget = null
                $('#btnActionModal').notify("Success!", { elementPosition: 'top center', className: "success" });
            }).catch(function(error) {
                $.notify(error.response.data.message, { elementPosition: 'top center', className: "danger" });
            })
        },
        delWidget(widget) {
            store.dispatch(DEL_WIDGET, widget)
        },
        editWidget(widget) {
            this.editedWidget = widget
            this.widget = JSON.parse(JSON.stringify(widget));
            this.showCreateModal = true
        },
        closeModal() {
            this.showCreateModal = false
            this.widget = this.getNewWidget()
        },
        getNewWidget() {
            return JSON.parse(JSON.stringify({
                id: null,
                name: '',
                color: '#ffffff',
                price: '',
                inventory: null,
                melts: false
            }))
        }
    }
})