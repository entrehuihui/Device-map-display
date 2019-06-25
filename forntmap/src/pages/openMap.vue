<template>
  <div id="openmap">
    <!-- 地图页 -->
    <lmap></lmap>
    <div>
      <!-- logo图标 -->
      <div class="openmaplogo">
        <img class="imgMax" :src="req.localhost + global.logo[3]">
      </div>
      <!-- 设备菜单 -->
      <div class="openMapDevices">
        <div class="openMapDevicesOpen">展开菜单</div>
        <div class="openMapDevicesleft">
          <img class="imgMax" src="/static/right.png">
        </div>
      </div>
      <!-- 用户头像信息 -->
      <div class="openMapUser">
        <img class="imgMax" id="openMapUserphoto" :src="req.localhost + global.userinfo.photo">
        <div class="openMapUserOpen">
          <!-- 推送开关 -->
          <div class="openMapUserPush" @click="push=!push">
            <div class="openMapUserPushimg">
              <img class="imgMax" src="/static/push.png">
            </div>
            <div class="openMapUserPushExplain">推送</div>
            <div class="openMapUserPushopenimg">
              <img v-show="!push" class="imgMax" src="/static/pushopen.png">
              <img
                v-show="push"
                class="imgMax"
                src="/static/pushopen.png"
                id="openMapUserPushopenmg"
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import lmap from "@/components/map.vue";
export default {
  components: {
    lmap
  },
  data() {
    return {
      push: false
    };
  }
};
</script>

<style scoped>
@import url("./css/openMap.css");
</style>
