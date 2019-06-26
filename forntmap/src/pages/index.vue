<template>
  <div id="index">
    <div class="indexbackgroud">
      <img class="indexbackgroud" :src="req.localhost + global.logo[1]">
    </div>
    <div class="indexlogin">
      <img class="indexbackgroud" :src="req.localhost + global.logo[2]">
      <div class="indexloginone">
        <div class="indexloginonetwo">
          <div class="indexloginprompt" ref="prompt"></div>
          <input
            class="indexlogininput"
            placeholder="账号"
            type="text"
            v-model="account"
            ref="account"
            @keyup.enter="checkAccount()"
          >
        </div>
        <div class="indexloginonetwo">
          <div class="indexloginprompt"></div>
          <input
            class="indexlogininput"
            placeholder="密码"
            type="password"
            v-model="password"
            ref="password"
            @keyup.enter="checkPassword()"
          >
        </div>
        <div class="indexloginonethree">
          <div class="indexloginremember" @click="selectRemember()">
            <div class="indexloginselect">
              <div class="indexloginselectinside" v-show="remember"></div>
            </div>
            <div class="indexloginexplain">自动登陆</div>
          </div>
          <div class="indexloginforget">
            <a href="/login/forget">忘记密码</a>
          </div>
        </div>
        <div class="indexloginonethree" id="indexloginonelogin" @click="checkPassword()">登陆</div>
        <div class="indexloginonethree">
          <div class="indexregister">
            <a href="/login/register">立即注册</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import md5 from "js-md5";
export default {
  data() {
    return {
      remember: true,
      account: "",
      password: "",
      cookieName: "openMap"
    };
  },
  methods: {
    jump: function() {
      this.$router.push("/openMap");
    },
    login: function() {
      this.req
        .post("/login", {
          code: "110",
          name: this.account,
          password: md5(this.password)
        })
        .then(response => {
          if (response.status != 200) {
            this.$refs.prompt.innerHTML = response.msg;
          }
          this.global.userinfo = response.data;
          // 判断是否记住登陆
          if (this.remember) {
            this.setCookie(response.data.jwt);
          } else {
            this.setCookie(response.data.jwt, -1);
          }
          // 跳转地图页面
          this.jump();
        });
    },
    checkPassword: function() {
      if (!this.checkAccount()) {
        this.$refs.account.focus();
        return;
      }
      if (this.password == "") {
        this.$refs.password.style = " border: 1px solid red;";
        this.$refs.prompt.innerHTML = "密码不能为空";
        return;
      }
      this.$refs.password.style = "border: 1px solid rgb(163, 163, 163);";
      this.$refs.prompt.innerHTML = "";
      this.login();
    },
    checkAccount: function() {
      if (this.account == "") {
        this.$refs.account.style = " border: 1px solid red;";
        this.$refs.prompt.innerHTML = "账号不能为空";
        return false;
      }
      this.$refs.account.style = "border: 1px solid rgb(163, 163, 163);";
      this.$refs.prompt.innerHTML = "";
      this.$refs.password.focus();
      return true;
    },
    selectRemember: function() {
      this.remember = !this.remember;
    },
    // 获取错误代码
    getCode: async function() {
      await this.req.get("/login/code").then(response => {
        if (response.status != 200) {
          return;
        }
        this.global.code = response.data;
      });
    },
    // 获取logo图标
    getLogo: function() {
      this.req.get("/login/logo").then(response => {
        if (response.status != 200) {
          return;
        }
        this.global.logo = response.data;
        this.$forceUpdate();
      });
    },
    // 获取用户登陆信息
    getUserInfo: async function() {
      await this.req.get("/login/info").then(response => {
        if (response.status != 200) {
          // alert(response.msg);
          // 删除cookie
          this.setCookie(this.global.userinfo.jwt, -1);
          return;
        }
        this.global.userinfo = response.data;
        this.jump();
      });
    },
    // 获取登陆cookie
    getCookie: function() {
      var key = this.cookieName;
      if (document.cookie.length == 0) {
        return;
      }
      var start = document.cookie.indexOf(key + "=");
      if (start == -1) {
        return;
      }
      start = start + key.length + 1;
      var end = document.cookie.indexOf(";", start);
      if (end === -1) end = document.cookie.length;
      this.global.userinfo.jwt = unescape(
        document.cookie.substring(start, end)
      );
      this.getUserInfo();
    },
    // expiredays 正数为设置, 负数为删除
    setCookie: function(value, expiredays = 30) {
      var exdate = new Date();
      exdate.setDate(exdate.getDate() + expiredays);
      document.cookie =
        this.cookieName +
        "=" +
        escape(value) +
        (expiredays == null ? "" : ";expires=" + exdate.toGMTString());
    }
  },
  watch: {},
  async mounted() {
    await this.getCode();
    if (this.$route.params.logout) {
      this.setCookie(this.global.userinfo.jwt, -1);
    } else {
      this.getCookie();
    }
    this.getLogo();
  }
};
</script>

<style scoped>
@import url("./css/index.css");
</style>
