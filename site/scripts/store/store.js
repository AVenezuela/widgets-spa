const HTTP = axios.create({
    baseURL: `http://localhost:666/`
    /*,headers: {
      Authorization: 'Bearer {token}'
    }*/
})

const UserModule = {
    state:{
        list:[],
        model:{}
    },
    mutations: {      
        setUsers (state, users){
            state.list = users
        }
    },
    actions:{
        setUsers(context){
            if(!this.getters.isUsersLoaded){
                HTTP.get('users').then(function(response){
                    context.commit('setUsers', response.data)                    
                }
                ).catch(function(e){
                    alert("Ops! Something is wrong loading users\n" + JSON.stringify(e))
                })
            }
        }
    },
    getters:{
        getSortedUsers(state, sort){
            return state.list
        },
        isUsersLoaded(state){
            return (state.list.length > 0)
        }
    }
}

const WidgetModule ={
    state:{
        list:[],
        model:{
            name:'',
            color:'',
            price:'',
            inventory:0,
            melts:false
        }
    },
    mutations:{
        setWidgets (state, widgets){
            state.list = widgets
        }
    },
    actions:{
        setWidgets(context){
            if(!this.getters.isWidgetsLoaded){
                HTTP.get('widgets').then(function(response){
                    context.commit('setWidgets', response.data)                    
                }
                ).catch(function(e){
                    alert('Ops! Something is wrong loading widgets\n'+ JSON.stringify(e))
                })
            }
        },
        saveWidget(context){
            HTTP.post('widgets/', context.state.model).then(function(response){
                alert('Done!')
            }).catch(function(e){
                console.log('Ops! Something is wrong saving widget\n'+ JSON.stringify(e))
            })            
        }
    },
    getters:{
        getSortedWidgets(state, sort){
            return state.list
        },
        isWidgetsLoaded(state){
            return (state.list.length > 0)
        }
    }
}

const store = new Vuex.Store({
    modules: {
        moduleUser: UserModule,
        moduleWidget: WidgetModule
      }
})