<template>
  <div id="adduser">
    <div class="adduserB" @click="closed()"></div>
    <div class="adduserI">
      <div class="adduserI0" @click="closed()">X</div>
      <div>
        <div class="adduserI1">
          <div class="adduserI2">用户名:</div>
          <div class="adduserI3">
            <input
              type="text"
              class="adduserI4"
              ref="name"
              v-model="name"
              placeholder="长度4- 20"
              @blur="checkName()"
            />
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">密码:</div>
          <div class="adduserI3">
            <input type="password" class="adduserI4" ref="password" v-model="password" />
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">重复密码:</div>
          <div class="adduserI3">
            <input type="password" class="adduserI4" ref="password1" v-model="password1" />
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">过期时间:</div>
          <div class="adduserI3">
            <!-- <input type="text" class="adduserI4" ref="expired" /> -->
            <div class="adduserI5">
              <times class="adduserI6" :type="'day'" ref="time"></times>
            </div>
          </div>
          <div class="adduserI8" @click="setTime()">永久</div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">权限:</div>
          <div class="adduserI3">
            <input
              type="text"
              class="adduserI4"
              readonly
              :value="global.userinfo.permisson == 3? '管理员':'普通用户'"
            />
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">状态:</div>
          <div class="adduserI3">
            <select class="adduserI7" v-model="status">
              <option value="1">启用</option>
              <option value="2">禁用</option>
            </select>
          </div>
        </div>
        <div class="adduserI1" v-show="global.userinfo.permisson == 3">
          <div class="adduserI2">VIP:</div>
          <div class="adduserI3">
            <select class="adduserI7" v-model="vip">
              <option value="1">VIP1</option>
              <option value="2">VIP2</option>
              <option value="3">VIP3</option>
              <option value="4">VIP4</option>
              <option value="5">VIP5</option>
              <option value="6">VIP6</option>
            </select>
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI2">备注:</div>
          <div class="adduserI3">
            <input type="text" class="adduserI4" ref="details" placeholder="长度0- 200" />
          </div>
        </div>
        <div class="adduserI1">
          <div class="adduserI8" id="adduserdisable" @click="closed()">取消</div>
          <div class="adduserI8" @click="check()">确定</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import times from "@/components/time.vue";
import md5 from "js-md5";
export default {
  components: { times },
  data() {
    return {
      name: "",
      password: "",
      password1: "",
      user: 1,
      status: 1,
      vip: 1,
      details: "",
      samename: false
    };
  },
  methods: {
    postUser: function() {
      if (!this.samename) {
        this.$refs.name.style = "border: 1px solid red;";
        return;
      }
      var times = 0;
      var time1 = this.$refs.time.dayShortText;
      if (time1 != "" && time1 != "永久") {
        times = parseInt(new Date(time1).getTime() / 1000);
      }
      var data = {
        details: this.details,
        expired: times,
        name: this.name,
        password: md5(this.password),
        status: parseInt(this.status),
        vip: parseInt(this.vip)
      };
      this.req.put("/users", data).then(response => {
        if (response.status != 200) {
          alert(response.msg);
          return;
        }
        this.closed()
        alert("Success");
      });
    },
    checkName: async function() {
      this.samename = false;
      var response = await this.req.get(
        "/login/register/name?name=" + this.name
      );
      if (response.status != 200) {
        return;
      }
      if (response.data == false) {
        this.samename = true;
        this.$refs.name.style = "border: 1px solid rgb(184, 182, 182);";
        return;
      }
      this.$refs.name.style = "border: 1px solid red;";
    },
    check: function() {
      if (this.name == "" || this.name.length > 20 || this.name.length < 4) {
        this.$refs.name.style = "border: 1px solid red;";
        return;
      }
      this.$refs.name.style = "border: 1px solid rgb(184, 182, 182);";
      if (this.password == "") {
        this.$refs.password.style = "border: 1px solid red;";
        return;
      }
      this.$refs.password.style = "border: 1px solid rgb(184, 182, 182);";
      if (this.password != this.password1) {
        this.$refs.password1.style = "border: 1px solid red;";
        return;
      }
      this.$refs.password1.style = "border: 1px solid rgb(184, 182, 182);";
      if (this.details.length > 200) {
        this.$refs.details.style = "border: 1px solid red;";
        return;
      }
      this.$refs.details.style = "border: 1px solid rgb(184, 182, 182);";
      this.postUser();
    },
    closed: function() {
      this.$emit("close");
    },
    setTime: function() {
      this.$refs.time.dayShortText = "永久";
    }
  }
};
</script>

<style scoped>
@import url("./css/adduser.css");
</style>
