Vue.component('widget', {
    template: '#widget-template',
    data: function() {
        var sortOrders = {}
        this.columns.forEach(function(key) {
            sortOrders[key] = 1
        })
        return {
            sortKey: '',
            sortOrders: sortOrders,
            searchFor: '',
            totalPerPage: 5,
            currentPage: 1
        }
    },
    props: {
        list: {
            required: true
        },
        title: {
            type: String,
            required: true
        },
        columns: Array
    },
    mounted: function() {
        //store.dispatch(SET_WIDGETS) 
    },
    methods: {
        sortBy (key) {
            this.sortKey = key
            this.sortOrders[key] = this.sortOrders[key] * -1
        },
        doPageAction(page){
            this.currentPage = page
        }
    },
    filters: {
        capitalize: function(str) {
            return str.charAt(0).toUpperCase() + str.slice(1)
        }
    },
    computed: {
        filteredData() {
            var sortKey = this.sortKey
            var filterKey = this.searchFor && this.searchFor.toLowerCase()
            var order = this.sortOrders[sortKey] || 1
            var data = this.list
            var me = this
            if (filterKey) {
                data = data.filter(function(row) {                    
                    return Object.keys(row).some(function(key) {
                        return ((String(row[key]).toLowerCase().indexOf(filterKey) > -1) && (me.columns.indexOf(key) > -1 ))
                    })
                })
            }
            if (sortKey) {
                data = data.slice().sort(function(a, b) {
                    a = a[sortKey]
                    b = b[sortKey]
                    return (a === b ? 0 : a > b ? 1 : -1) * order
                })
            }
            return data
        },
        paginateData(){
            var data = this.filteredData
            var from = ((this.currentPage - 1) * this.totalPerPage)
            var to = (this.currentPage * this.totalPerPage)

            return data.slice(from, to)
        }
    }
})