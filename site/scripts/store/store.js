const HTTP = axios.create({
    baseURL: `http://localhost:666/api/`,
    headers: {
        "Content-Type": "application/json;charset=UTF-8"        
    }
})
HTTP.interceptors.request.use(config => {
    config.headers.Authorization = "Bearer " + localStorage.getItem(STR_TOKEN)
    return config
})
HTTP.interceptors.response.use(response => {    
    // …get the token from the header or response data if exists, and save it.
    const token = response.headers['Authorization'] || response.data['token']
    if (token) {
      //ls.set('jwt-token', token)
    }
    return response
  }, error => {    
    // Also, if we receive a Bad Request / Unauthorized error
    if (error.response.status === 400 || error.response.status === 401) {
      // and we're not trying to login
      if (!(error.config.method === 'post' && /\/api\/me\/?$/.test(error.config.url))) {
          alert('Needs login')
        // the token must have expired. Log out.
        //event.emit('logout')
      }
    }
    return Promise.reject(error)
})

const SET_USERS = 'SET_USERS'
const SET_WIDGETS = 'SET_WIDGETS'
const SAVE_WIDGET = 'SAVE_WIDGET'
const SAVE_USER = 'SAVE_USER'
const UPDATE_WIDGET = 'UPDATE_WIDGET'
const UPDATE_USER = 'UPDATE_USER'
const DEL_WIDGET = 'DEL_WIDGET'
const DEL_USER = 'DEL_USER'
const STR_TOKEN = 'STR_TOKEN'
const DO_LOGIN = 'DO_LOGIN'
const SORT_USERBY = 'SORT_USERBY'
const SORT_WIDGETBY = 'SORT_WIDGETBY'

const LoginModule ={
    state:{
        user:{
            username:'',
            password:'',
            token:''
        },
        isLogged:Boolean
    },
    mutations:{
        logAuth (state, user){
            localStorage.setItem(STR_TOKEN, user.token)            
        }
    },
    actions:{
        [DO_LOGIN](context, user){
            axios.post('http://localhost:666/login', user).then(function(response) {
                user.token = reponse.data.token
                context.commit('logAuth', user)
            }).catch(function(e) {
                alert("Ops! Something is wrong on login for user " + user.username +"\n"+ e.response.data.message)
            })            
        }
    },
    getters:{        
        isLogged (state){
            state.isLogged = (localStorage.getItem(STR_TOKEN) != null)
            return state.isLogged
        }
    }
}

const UserModule = {
    state: {
        list: [],
        model: {},
        sortKey: 'name'
    },
    mutations: {
        setUsers(state, users) {
            state.list = users
        },
        addUser(state, user) {
            state.list.push(user)
        },
        editUser(state, objUser) {
            state.list.splice(state.list.indexOf(objUser.oldUser), 1)
            state.list.push(objUser.newUser)
        },
        delUser(state, user) {
            state.list.splice(state.list.indexOf(user), 1)
        },
        [SORT_USERBY](state, sortKey){
            state.sortKey = sortKey
        }
    },
    actions: {
        [SET_USERS](context) {
            if (!this.getters.isUsersLoaded) {
                HTTP.get('users').then(function(response) {
                    context.commit('setUsers', response.data)
                }).catch(function(e) {
                    alert("Ops! Something is wrong loading users\n" + JSON.stringify(e))
                })
            }
        },
        [SAVE_USER](context, objUser) {
            return new Promise((resolve, reject) => {
                HTTP.post('user', objUser.newUser).then(function(response) {
                    context.commit('addUser', response.data)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [DEL_USER](context, user) {
            return new Promise((resolve, reject) => {
                HTTP.delete('user/' + user.id).then(function(response) {
                    context.commit('delUser', user)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [UPDATE_USER](context, objUser) {
            return new Promise((resolve, reject) => {
                HTTP.put('user', objUser.newUser).then(function(response) {
                    context.commit('editUser', objUser)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [SORT_USERBY](context, sortKey){
            context.dispatch(SORT_USERBY, sortKey)
        }
    },
    getters: {
        getSortedUsers(state) {
            var data = []
            if (state.sortKey) {
                data = state.list.slice().sort(function (a, b) {
                  a = a[state.sortKey]
                  b = b[state.sortKey]
                  return (a === b ? 0 : a > b ? 1 : -1) * 1
                })
              }
            return data
        },
        isUsersLoaded(state) {
            return (state.list.length > 0)
        }
    }
}

const WidgetModule = {
    state: {
        list: [],
        sortKey: 'name'
    },
    mutations: {
        setWidgets(state, widgets) {
            state.list = widgets
        },
        addWidget(state, widget) {
            state.list.push(widget)
        },
        editWidget(state, objWidget) {
            state.list.splice(state.list.indexOf(objWidget.oldWidget), 1)
            state.list.push(objWidget.newWidget)
        },
        delWidget(state, widget) {
            state.list.splice(state.list.indexOf(widget), 1)
        },
        [SORT_WIDGETBY](state, sortKey){
            state.sortKey = sortKey
        }
    },
    actions: {
        [SET_WIDGETS](context) {
            if (!this.getters.isWidgetsLoaded) {
                HTTP.get('widgets').then(function(response) {
                    context.commit('setWidgets', response.data)
                }).catch(function(e) {
                    //alert('Ops! Something is wrong loading widgets\n'+ JSON.stringify(e))
                })
            }
        },
        [SAVE_WIDGET](context, objWidget) {
            return new Promise((resolve, reject) => {
                HTTP.post('widget', objWidget.newWidget).then(function(response) {
                    context.commit('addWidget', response.data)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [DEL_WIDGET](context, widget) {
            return new Promise((resolve, reject) => {
                HTTP.delete('widget/' + widget.id).then(function(response) {
                    context.commit('delWidget', widget)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [UPDATE_WIDGET](context, objWidget) {
            return new Promise((resolve, reject) => {
                HTTP.put('widget', objWidget.newWidget).then(function(response) {
                    context.commit('editWidget', objWidget)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },        
        [SORT_WIDGETBY](context, sortKey){
            context.dispatch(SORT_WIDGETBY, sortKey)
        }
    },
    getters: {
        getSortedWidgets(state) {
            var data = []
            if (state.sortKey) {
                data = state.list.slice().sort(function (a, b) {
                  a = a[state.sortKey]
                  b = b[state.sortKey]
                  return (a === b ? 0 : a > b ? 1 : -1) * 1
                })
              }
            return data
        },
        isWidgetsLoaded(state) {
            return (state.list.length > 0)
        }
    }
}

const store = new Vuex.Store({
    modules: {
        moduleUser: UserModule,
        moduleWidget: WidgetModule,
        auth: LoginModule
    }
})