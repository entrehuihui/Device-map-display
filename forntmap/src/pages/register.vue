<template>
  <div id="register">
    <div class="registerInside">
      <div class="registerTitle">
        <h1>欢迎注册地图服务</h1>
      </div>
      <div class="registerInput">
        <!-- name -->
        <div class="registerInputt">
          <input
            type="text"
            :class="namecheck?'registerInput1red':'registerInput1'"
            placeholder="设置登陆名 (4-20个字符)"
            v-model="name"
            @blur="checkresult(1)"
          />
          <div class="registerInput2" ref="nametext" v-show="namecheck">
            <div v-show="!namecheck1">会员名为5-20字符</div>
            <div v-show="namecheck1">会员名已经存在</div>
          </div>
        </div>
        <!-- password -->
        <div class="registerInputt">
          <input
            type="password"
            :class="passwordcheck?'registerInput1red':'registerInput1'"
            placeholder="设置你的登陆密码 (6-32个字符)"
            v-model="password"
            @blur="checkresult(2)"
          />
          <div class="registerInput2" ref="passwordtext" v-show="passwordcheck">密码设置不符合要求</div>
        </div>
        <!-- password1 -->
        <div class="registerInputt">
          <input
            type="password"
            :class="password1check?'registerInput1red':'registerInput1'"
            placeholder="再次输入你的密码"
            v-model="password1"
            @blur="checkresult(3)"
          />
          <div class="registerInput2" ref="password1text" v-show="password1check">两次输入的密码不一致</div>
        </div>
        <!-- mobile -->
        <div class="registerInputt">
          <input
            type="text"
            :class="mobilecheck?'registerInput1red':'registerInput1'"
            placeholder="输入手机号码"
            v-model="mobile"
            @keyup="mobile.length > 11 ? mobile = mobile.substr(0,11):''"
            @blur="checkresult(4)"
          />
          <div class="registerInput2" ref="mobiletext" v-show="mobilecheck">手机号码格式不正确</div>
        </div>
        <!-- email -->
        <div class="registerInputt">
          <input
            type="text"
            :class="emailcheck?'registerInput1red':'registerInput1'"
            placeholder="请输入密保邮箱"
            v-model="email"
            @blur="checkresult(5)"
          />
          <div class="registerInput2" ref="mobiletext" v-show="emailcheck">邮箱格式不正确</div>
        </div>
        <!-- able -->
        <div class="registerInputt">
          <div class="registerInput1" id="registerable" @click="registerUser()">同意条款并注册</div>
        </div>
        <!-- service -->
        <div class="registerInputt">
          <div class="container" @click="pact=!pact">
            <input type="checkbox" id="cbtest" />
            <label for="cbtest" class="check-box"></label>
          </div>
          <div class="registerInputservice">《服务条款》|《法律声明和隐私权政策》</div>
          <div v-show="pact" class="registerInputpact">请勾选同意协议</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import check from "@/global/check.js";
import md5 from "js-md5";
export default {
  data() {
    return {
      name: "",
      namecheck: false,
      password: "",
      passwordcheck: false,
      password1: "",
      password1check: false,
      mobile: "",
      mobilecheck: false,
      email: "",
      emailcheck: false,
      namecheck1: false,
      pact: true
    };
  },
  methods: {
    checkresult: function(methods = 0) {
      switch (methods) {
        case 1:
          if (
            this.name.length < 4 ||
            this.name.length > 20 ||
            !check.checkName(this.name)
          ) {
            this.namecheck = true;
          } else {
            this.namecheck = false;
            this.checkUserName();
          }
          break;
        case 2:
          if (this.password.length < 6 || this.password.length > 32) {
            this.passwordcheck = true;
          } else {
            this.passwordcheck = false;
          }
          break;
        case 3:
          if (this.password != this.password1) {
            this.password1check = true;
          } else {
            this.password1check = false;
          }
          break;
        case 4:
          if (this.mobile.length != 11 || !check.checkPhone(this.mobile)) {
            this.mobilecheck = true;
          } else {
            this.mobilecheck = false;
          }
          break;
        case 5:
          if (!check.checkEmail(this.email)) {
            this.emailcheck = true;
          } else {
            this.emailcheck = false;
          }
          break;
      }
      this.$forceUpdate();
    },
    checkUserName: function() {
      this.req.get("/login/register/name?name=" + this.name).then(response => {
        if (response.status != 200) {
          return;
        }
        if (response.data == true) {
          this.namecheck = true;
          this.namecheck1 = true;
        } else {
          this.namecheck = false;
          this.namecheck1 = false;
        }
        this.$forceUpdate();
      });
    },
    registerUser: async function() {
      if (
        this.namecheck ||
        this.name == "" ||
        this.passwordcheck ||
        this.password == "" ||
        this.password1check ||
        this.mobilecheck ||
        this.mobile == "" ||
        this.emailcheck ||
        this.email == "" ||
        this.pact
      ) {
        return;
      }
      var response = await this.req.post("/login/register", {
        code: "string",
        codeJWT: "string",
        email: this.email,
        mobile: this.mobile,
        name: this.name,
        password: md5(this.password)
      });
      if (response.status != 200) {
        alert(response.msg);
        return;
      }
      alert("注册成功,跳转至登陆界面");
      this.$router.push({
        name: "index",
        params: {
          name: this.name,
          password: this.password
        }
      });
    }
  }
};
</script>

<style scoped>
@import url("./css/register.css");
</style>
