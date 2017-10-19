const HTTP = axios.create({
    baseURL: 'http://localhost:666/api/', //`http://spa.tglrw.com:4000/`,
    headers: {
        "Content-Type": "application/json;charset=UTF-8"
    }
})
HTTP.interceptors.request.use(config => {
    config.headers.Authorization = "Bearer " + localStorage.getItem(STR_TOKEN)
    return config
})
HTTP.interceptors.response.use(response => {
    return response
}, error => {
    // Also, if we receive a Bad Request / Unauthorized error
    if (error.response.status === 400 || error.response.status === 401) {
        if (!(error.config.method === 'post' && /\/api\/me\/?$/.test(error.config.url))) {
            $.notify("Needs Login", { elementPosition: 'top center', className: "danger" });
            store.dispatch(DO_LOGOUT)
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
const DO_LOGOUT = 'DO_LOGOUT'
const CHECK_LOGGED = 'CHECK_LOGGED'
const SORT_USERBY = 'SORT_USERBY'
const SORT_WIDGETBY = 'SORT_WIDGETBY'
const SET_AUTH = 'SET_AUTH'
const SET_PAGE = 'SET_PAGE'

const Mixins = {
    updated: function() {
        $("[data-toggle='tooltip']").tooltip();
        $('[data-toggle="popover"]').popover();
    }
}