const HTTP = axios.create({
    baseURL: `http://localhost:666/api/`,
    headers: {
        "Content-Type": "application/json;charset=UTF-8",
        "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDgxOTAwMTMsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUwODE4NjQxM30.EthCZPXmFtJVNQygO-0jWBFEmSGAe25c28sv7RIKXmI"
    }
})

const SET_USERS = 'SET_USERS'
const SET_WIDGETS = 'SET_WIDGETS'
const SAVE_WIDGET = 'SAVE_WIDGET'
const SAVE_USER = 'SAVE_USER'
const UPDATE_WIDGET = 'UPDATE_WIDGET'
const UPDATE_USER = 'UPDATE_USER'
const DEL_WIDGET = 'DEL_WIDGET'
const DEL_USER = 'DEL_USER'

const UserModule = {
    state: {
        list: [],
        model: {}
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
        }
    },
    getters: {
        getSortedUsers(state, sort) {
            return state.list
        },
        isUsersLoaded(state) {
            return (state.list.length > 0)
        }
    }
}

const WidgetModule = {
    state: {
        list: []
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
        }
    },
    getters: {
        getSortedWidgets(state, sort) {
            return state.list
        },
        isWidgetsLoaded(state) {
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