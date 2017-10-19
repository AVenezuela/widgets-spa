const LoginModule = {
    state: {
        user: {
            username: '',
            password: ''
        },
        isLogged: Boolean
    },
    mutations: {
        [SET_AUTH](state, user) {
            state.user = user
        },
        [DO_LOGOUT](state) {
            localStorage.removeItem(STR_TOKEN)
            state.user = {
                username: '',
                password: ''
            }
            state.isLogged = false
        }
    },
    actions: {
        [DO_LOGIN](context, user) {
            return new Promise((resolve, reject) => {
                axios.post('http://localhost:666/login', user).then(function(response) {
                    context.commit(SET_AUTH, user)
                    localStorage.setItem(STR_TOKEN, response.data.token)
                    resolve(response);
                }).catch(function(e) {
                    reject(e);
                })
            })
        },
        [DO_LOGOUT](context) {
            context.commit(DO_LOGOUT)
        },
        [CHECK_LOGGED]() {
            return this.getters.isLogged
        }
    },
    getters: {
        isLogged(state) {
            state.isLogged = (localStorage.getItem(STR_TOKEN) != null)
            return state.isLogged
        }
    }
}

const UserModule = {
    state: {
        list: [],
        model: {},
        filteredList: []
    },
    mutations: {
        [SET_USERS](state, users) {
            state.list = users
        },
        [SAVE_USER](state, user) {
            state.list.push(user)
        },
        [UPDATE_USER](state, objUser) {
            state.list.splice(state.list.indexOf(objUser.oldUser), 1)
            state.list.push(objUser.newUser)
        },
        [DEL_USER](state, user) {
            state.list.splice(state.list.indexOf(user), 1)
        },
        [SORT_USERBY](state, sortKey) {
            state.sortKey = sortKey
        }
    },
    actions: {
        [SET_USERS](context) {
            if (!this.getters.isUsersLoaded) {
                HTTP.get('users').then(function(response) {
                    context.commit(SET_USERS, response.data)
                }).catch(function(e) {
                    $.notify("Ops! Something is wrong loading users.", { elementPosition: 'top center', className: "danger" });
                })
            }
        },
        [SAVE_USER](context, objUser) {
            return new Promise((resolve, reject) => {
                HTTP.post('user', objUser.newUser).then(function(response) {
                    context.commit(SAVE_USER, response.data)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [DEL_USER](context, user) {
            return new Promise((resolve, reject) => {
                HTTP.delete('user/' + user.id).then(function(response) {
                    context.commit(DEL_USER, user)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [UPDATE_USER](context, objUser) {
            return new Promise((resolve, reject) => {
                HTTP.put('user', objUser.newUser).then(function(response) {
                    context.commit(UPDATE_USER, objUser)
                    resolve(response);
                }).catch(function(error) {
                    reject(error);
                })
            })
        },
        [SORT_USERBY](context, param) {
            context.state.filteredList = context.state.list

        }
    },
    getters: {
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
        [SORT_WIDGETBY](state, sortKey) {
            state.sortKey = sortKey
        }
    },
    actions: {
        [SET_WIDGETS](context) {
            if (!this.getters.isWidgetsLoaded) {
                HTTP.get('widgets').then(function(response) {
                    context.commit('setWidgets', response.data)
                }).catch(function(e) {
                    $.notify("Ops! Something is wrong loading widgets.", { elementPosition: 'top center', className: "danger" });
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
        [SORT_WIDGETBY](context, sortKey) {
            context.dispatch(SORT_WIDGETBY, sortKey)
        }
    },
    getters: {
        getSortedWidgets(state) {
            var data = []
            if (state.sortKey) {
                data = state.list.slice().sort(function(a, b) {
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