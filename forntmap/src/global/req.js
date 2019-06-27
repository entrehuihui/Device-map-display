import router from '@/router/index'
import Vue from "vue";
import axios from "axios";
import global from '@/global/variable.js'
Vue.prototype.$axios = axios;

// --------------- //
// const localhost = "http://120.78.76.139:8800"; //调试地址/
const localhost = "http://localhost:8800"; //调试地址
// const localhost = ""; //正式

function getAuthorization() {
    return {
        headers: {
            'Authorization': global.userinfo.jwt,
        }
    }
}


export default {
    localhost,
    post: async function (url, postData) {
        var response = await Vue.prototype.$axios
            .post(localhost + url, postData, getAuthorization())
            .then(response => {
                return response;
            })
            .catch(error => {
                if (error.response == undefined) {
                    return {
                        status: 500
                    }
                } else {
                    return error.response
                }
            });
        if (response == undefined) {
            return {
                status: 500
            }
        }
        if (response.status == 200) {
        } else if (response.status == 301) {
            router.push('/');
            response.msg = global.getcode(response.data)
        } else {
            response.msg = global.getcode(response.data)
        }
        return response
    },
    get: async function (url) {
        var response = await Vue.prototype.$axios
            .get(localhost + url, getAuthorization())
            .then(response => {
                return response;
            })
            .catch(error => {
                if (error.response == undefined) {
                    return {
                        status: 500
                    }
                } else {
                    return error.response
                }
            });
        if (response == undefined) {
            return {
                status: 500
            }
        }
        if (response.status == 200) {
        } else if (response.status == 301) {
            router.push('/');
            response.msg = global.getcode(response.data)
        } else {
            response.msg = global.getcode(response.data)
        }
        return response
    },
    put: async function (url, putData) {
        var response = await Vue.prototype.$axios
            .put(localhost + url, putData, getAuthorization())
            .then(response => {
                return response;
            })
            .catch(error => {
                if (error.response == undefined) {
                    return {
                        status: 500
                    }
                } else {
                    return error.response
                }
            });
        if (response == undefined) {
            return {
                status: 500
            }
        }
        if (response.status == 200) {
        } else if (response.status == 301) {
            router.push('/');
            response.msg = global.getcode(response.data)
        } else {
            response.msg = global.getcode(response.data)
        }
        return response
    },
    del: async function (url, delData) {
        var response = await Vue.prototype.$axios
            .delete(localhost + url, {
                data: delData,
                headers: getAuthorization().headers
            })
            .then(response => {
                return response;
            })
            .catch(error => {
                if (error.response == undefined) {
                    return {
                        status: 500
                    }
                } else {
                    return error.response
                }
            });
        if (response == undefined) {
            return {
                status: 500
            }
        }
        if (response.status == 200) {
        } else if (response.status == 301) {
            router.push('/');
            response.msg = global.getcode(response.data)
        } else {
            response.msg = global.getcode(response.data)
        }
        return response
    },
    postfile: async function (url, postData, config = {
        headers: {
            "Content-Type": "multipart/form-data"
        }
    }) {
        var response = await Vue.prototype.$axios
            .post(localhost + url + "?Authorization=" + global.userinfo.jwt, postData, config)
            .then(response => {
                return response;
            })
            .catch(error => {
                if (error.response == undefined) {
                    return {
                        status: 500
                    }
                } else {
                    return error.response
                }
            });
        if (response == undefined) {
            return {
                status: 500
            }
        }
        if (response.status == 200) {
        } else if (response.status == 301) {
            router.push('/');
            response.msg = global.getcode(response.data)
        } else {
            response.msg = global.getcode(response.data)
        }
        return response
    },
};
