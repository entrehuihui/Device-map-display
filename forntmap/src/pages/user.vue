<template>
  <div id="user">
    <div class="userleft">
      <!-- logo -->
      <div class="openmaplogo">
        <img class="imgMax" :src="req.localhost + global.logo[3]" />
      </div>
      <!-- 列表 -->
      <div>
        <div class="userleftTitle">
          <div class="userleftTitle1" @click="index=0" :id="index==0 ? 'userleftTitle1':''">
            <div class="userleftTitleName">个人中心</div>
            <div class="userleftTitlefold">
              <img src="/static/right.png" class="imgMax" v-show="index == 0" />
              <img src="/static/top.png" class="imgMax" v-show="index != 0" />
            </div>
          </div>
        </div>
        <div class="userleftTitle" v-show="global.userinfo.permisson == 3">
          <div class="userleftTitle1" @click="index=1" :id="index==1 ? 'userleftTitle1':''">
            <div class="userleftTitleName">权限管理</div>
            <div class="userleftTitlefold">
              <img src="/static/right.png" class="imgMax" v-show="index == 1" />
              <img src="/static/top.png" class="imgMax" v-show="index != 1" />
            </div>
          </div>
        </div>
        <div class="userleftTitle" v-show="global.userinfo.permisson != 1">
          <div class="userleftTitle1" @click="index=2" :id="index==2 ? 'userleftTitle1':''">
            <div class="userleftTitleName">账户管理</div>
            <div class="userleftTitlefold">
              <img src="/static/right.png" class="imgMax" v-show="index == 2" />
              <img src="/static/top.png" class="imgMax" v-show="index != 2" />
            </div>
          </div>
        </div>
        <div class="userleftTitle">
          <div class="userleftTitle1" @click="index=3" :id="index==3 ? 'userleftTitle1':''">
            <div class="userleftTitleName">设备管理</div>
            <div class="userleftTitlefold">
              <img src="/static/right.png" class="imgMax" v-show="index == 3" />
              <img src="/static/top.png" class="imgMax" v-show="index != 3" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- 头部 -->
    <div class="userHeard">
      <div class="userOut" @click="ligout()">注销</div>
      <!-- 用户头像 -->
      <div class="userPhoto">
        <div class="userPhotoImg">
          <img :src="req.localhost + global.userinfo.photo" class="imgMax" />
        </div>
        <div class="userPhotoName">{{global.userinfo.name}}</div>
      </div>
      <!-- retMap -->
      <div class="userReturn">
        <a href="/openMap#/openMap">返回地图监控</a>
      </div>
    </div>
    <!-- 中部 -->
    <div class="userBody">
      <div class="userBodyInside">
        <!-- 个人中心 -->
        <user v-if="index==0"></user>
        <!-- 权限管理 -->
        <permisson v-if="index==1"></permisson>
        <!-- 用户管理 -->
        <account v-if="index==2"></account>
        <!-- 设备管理 -->
        <deives v-if="index==3"></deives>
      </div>
    </div>
    <!-- 尾部 -->
    <div class="userBottom" :userinfo="global.userinfo"></div>
  </div>
</template>

<script>
import user from "@/components/user.vue";
import account from "@/components/account.vue";
import deives from "@/components/devices.vue";
import permisson from "@/components/permisson.vue";
export default {
  components: {
    user,
    account,
    deives,
    permisson
  },
  data() {
    return {
      index: 0
    };
  },
  methods: {
    // 注销
    ligout: function() {
      this.global.userinfo = {};
      this.$router.push({
        name: "index",
        params: {
          logout: true
        }
      });
    },
    // 获取用户信息
    getUserinfo: async function() {
      await this.req.get("/login/info").then(response => {
        if (response.status != 200) {
          this.global.setCookie("", -1);
          return;
        }
        this.global.userinfo = response.data;
        // this.$set(this.global, "userinfo", response.data);
        this.$forceUpdate();
      });
    },
    // 获取状态
    getState: function() {
      this.req.get("/configState").then(response => {
        if (response.status != 200) {
          return;
        }
        this.global.state = response.data;
        this.$forceUpdate();
      });
    }
  },
  async mounted() {
    await this.getUserinfo();
    this.getState();
  }
};
</script>

<style>
@import url("./css/user.css");
</style>
