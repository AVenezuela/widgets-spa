const HTTP = axios.create({
    baseURL: `http://localhost:666/`
    ,headers: {
        "Content-Type":"application/json;charset=UTF-8"
    }
})

const SAVE_SUCCESS = 'SAVE_SUCCESS'
const SAVE_FAILURE = 'SAVE_FAILURE'
const SAVE_RESET = 'SAVE_RESET'

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
        saveStatus:null
    },
    mutations:{
        setWidgets (state, widgets){
            state.list = widgets
        },
        [SAVE_SUCCESS](state){
            state.saveStatus = 'sucess'
        },
        [SAVE_FAILURE](state){
            state.saveStatus = 'failed'
        },
        [SAVE_RESET](state){
            state.saveStatus = null
        }
    },
    actions:{
        setWidgets(context){
            if(!this.getters.isWidgetsLoaded){
                HTTP.get('widgets').then(function(response){
                    context.commit('setWidgets', response.data)                    
                }
                ).catch(function(e){                    
                    //alert('Ops! Something is wrong loading widgets\n'+ JSON.stringify(e))
                })
            }
        },
        saveWidget(context, widget){
            return new Promise((resolve, reject) =>{
                console.log(JSON.stringify(widget))
                HTTP.post('widget', JSON.stringify(widget)).then(function(response){
                    context.commit(SAVE_SUCCESS)
                    resolve(response);
                }).catch(function(error){
                    context.commit(SAVE_FAILURE)
                    reject(error);
                })  
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