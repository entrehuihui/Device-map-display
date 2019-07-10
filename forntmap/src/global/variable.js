{/* <script>
export default {
    data:function () {
        return{
            // 警告图片
            alarm:9900,
            // 弹窗图层
            pop:8800,
            othen:7700,
        }
    }
}
</script> */}
import req from '@/global/req.js'

var code = []
var logo = {}
var userinfo = {
}
var state = []

var color = [
    "blueviolet",
    "aqua",
    "brown",
    "blueviolet",
    "cyan",
    "blue",
    "chartreuse",
    "yellow",
    "violet",
    "rgb(214, 7, 255)"
]
export default {
    logo,
    code,
    getcode,
    userinfo,
    color,
    state,
    sleep,
    getState,
    setCookie,
    getCookie,
    jwt
}
function jwt() {
    if (!this.userinfo.jwt && this.getCookie()) {
        getUserInfo(this);
        getlogo(this);
        setState(this);
        setCode(this);
    }
    return this.userinfo.jwt
}

function setState(event) {
    req.get("/configState").then(response => {
        if (response.status != 200) {
            return;
        }
        event.state = response.data;
    });
}
function setCode() {
    req.get("/login/code").then(response => {
        if (response.status != 200) {
            return;
        }
        event.code = response.data;
    });
}

function getlogo(event) {
    req.get("/login/logo").then(response => {
        if (response.status != 200) {
            return;
        }
        event.logo = response.data;
    });
}

function getState(index) {
    if (index > this.state.length) {
        return {
            States: "未知"
        }
    }
    return this.state[index - 1]
}

function getcode(id) {
    if (id == undefined) {
        return "失败!"
    }
    if (id > this.code.length) {
        return "失败!"
    }
    return this.code[id - 1].error
}

function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
}

var cookieName = "openMap"
function setCookie(value, expiredays = 30) {
    var exdate = new Date();
    exdate.setDate(exdate.getDate() + expiredays);
    document.cookie =
        cookieName +
        "=" +
        escape(value) +
        (expiredays == null ? "" : ";expires=" + exdate.toGMTString());
}
// 获取登陆cookie
function getCookie() {
    var key = cookieName;
    if (document.cookie.length == 0) {
        return false;
    }
    var start = document.cookie.indexOf(key + "=");
    if (start == -1) {
        return false;
    }
    start = start + key.length + 1;
    var end = document.cookie.indexOf(";", start);
    if (end === -1) end = document.cookie.length;
    this.userinfo.jwt = unescape(
        document.cookie.substring(start, end)
    );
    if (this.userinfo.jwt.length < 30) {
        this.setCookie("", -1);
        return false;
    }
    return true;
}

// 获取用户登陆信息
async function getUserInfo(event) {
    await req.get("/login/info").then(response => {
        if (response.status != 200) {
            event.setCookie(event.userinfo.jwt, -1);
            return;
        }
        event.userinfo = response.data;
    });
}