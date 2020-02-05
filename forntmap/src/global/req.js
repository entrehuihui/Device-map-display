import router from '@/router/index'
import Vue from "vue";
import axios from "axios";
import global from '@/global/variable.js'
Vue.prototype.$axios = axios;

// --------------- //
var localhost = ""; //正式
// localhost = "http://127.0.0.1:8800"; //调试地址
localhost = "http://120.78.76.139:8800"; //调试地址/

function getAuthorization() {
    return {
        headers: {
            'Authorization': global.jwt(),
        }
    }
}
function ret(response) {
    if (response == undefined) {
        return {
            status: 500,
            msg: "net::ERR_CONNECTION_REFUSED"
        }
    }
    if (response.status == 200) {
    } else if (response.status == 301) {
        console.log(response, "+++++++++++++++++");
        router.push('/');
        response.msg = global.getcode(response.data)
    } else {
        response.msg = global.getcode(response.data)
    }
    return response
}

export default {
    localhost,
    post: async function (url, postData) {
        var response = await Vue.prototype.$axios
            .post(localhost + url, postData, getAuthorization()).catch(error => {
                return error.response
            })
        return ret(response)
    },
    get: async function (url) {
        var response = await Vue.prototype.$axios
            .get(localhost + url, getAuthorization()).then(response => {
                return response
            }).catch(error => {
                return error.response
            })
        return ret(response)
    },
    put: async function (url, putData) {
        var response = await Vue.prototype.$axios
            .put(localhost + url, putData, getAuthorization()).catch(error => {
                return error.response
            })
        return ret(response)
    },
    del: async function (url, delData) {
        var response = await Vue.prototype.$axios
            .delete(localhost + url, {
                data: delData,
                headers: getAuthorization().headers
            }).catch(error => {
                return error.response
            })
        return ret(response)
    },
    postfile: async function (url, postData, config = 0) {
        var response = await Vue.prototype.$axios
            .post(localhost + url + "?Authorization=" + global.userinfo.jwt, postData, {
                headers: {
                    "Content-Type": "multipart/form-data",
                    "Authorization": global.jwt(),
                    "Config": config,
                }
            }).catch(error => {
                return error.response
            })
        return ret(response)
    },
};
