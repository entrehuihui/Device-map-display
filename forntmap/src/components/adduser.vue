<template>
  <div id="adduser">
    <div class="adduserB"></div>
    <div class="adduserI">
      <div class="adduserI0">X</div>
      <div>
        <div class="adduserI1">
          <div class="adduserI2">用户名:</div>
          <div class="adduserI3">
            <input type="text" class="adduserI4" ref="name" v-model="name" />
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
        <div class="adduserI1" v-show="global.userinfo.permisson == 3">
          <div class="adduserI2">VIP:</div>
          <div class="adduserI3">
            <!-- <input type="text" class="adduserI4" ref="expired" /> -->
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
            <input type="text" class="adduserI4" ref="expired" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import times from "@/components/time.vue";
export default {
  components: { times },
  data() {
    return {
      name: "",
      password: "",
      password1: "",
      user: 1,
      vip: 1,
      details: ""
    };
  },
  methods: {
    setTime: function() {
      this.$refs.time.dayShortText = "永久";
    }
  }
};
</script>

<style scoped>
@import url("./css/adduser.css");
</style>
