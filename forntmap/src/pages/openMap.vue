<template>
  <div id="openmap" @keydown.esc="fullscreenShow(false)">
    <!-- 地图页 -->
    <lmap :changeZoom="changeZoom" :mapindex="mapIndex.index"></lmap>
    <div v-show="!fullscreen">
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
      <div class="openMapUser" @mouseover="useropen = true" @mouseleave="useropen = false">
        <!-- 左侧展开 -->
        <div class="openMapUserOpen" v-show="useropen" @click="userInfo = false">
          <!-- 推送开关 -->
          <div class="openMapUserPush" @mouseleave="push=false">
            <div @click="push=!push">
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
            <div class="openMapUserPushbutton" v-show="push">
              <div class="openMapUserPushbuttonExplain">
                <div class="openMapUserPushbuttonExplainleft">推送</div>
                <div
                  class="openMapUserPushbuttonExplainright"
                  @click="userConfig.Message =changeStatus(userConfig.Message, 'Message')"
                >
                  <div class="openMapUserPushbuttonExplainrightbg" v-show="userConfig.Message==1">
                    <div class="openMapUserPushbuttonExplainrightSwitch">开</div>
                    <div class="openMapUserPushbuttonExplainrightInside"></div>
                  </div>
                  <div
                    class="openMapUserPushbuttonExplainrightbg"
                    v-show="userConfig.Message==2"
                    id="openMapUserPushbuttonExplainrightbgclose"
                  >
                    <div class="openMapUserPushbuttonExplainrightInside"></div>
                    <div class="openMapUserPushbuttonExplainrightSwitch">关</div>
                  </div>
                </div>
              </div>
              <div class="openMapUserPushbuttonExplain">
                <div class="openMapUserPushbuttonExplainleft">声音</div>
                <div
                  class="openMapUserPushbuttonExplainright"
                  @click=" userConfig.Sound =changeStatus(userConfig.Sound, 'Sound')"
                >
                  <div class="openMapUserPushbuttonExplainrightbg" v-show="userConfig.Sound==1">
                    <div class="openMapUserPushbuttonExplainrightSwitch">开</div>
                    <div class="openMapUserPushbuttonExplainrightInside"></div>
                  </div>
                  <div
                    class="openMapUserPushbuttonExplainrightbg"
                    v-show="userConfig.Sound==2"
                    id="openMapUserPushbuttonExplainrightbgclose"
                  >
                    <div class="openMapUserPushbuttonExplainrightInside"></div>
                    <div class="openMapUserPushbuttonExplainrightSwitch">关</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- 语言选择 -->
          <div class="openMapUserPush" @mouseleave="language=false">
            <div @click="language=!language">
              <div class="openMapUserPushimg">
                <img class="imgMax" src="/static/language.png">
              </div>
              <div class="openMapUserPushExplain">语言</div>
              <div class="openMapUserPushopenimg">
                <img v-show="!language" class="imgMax" src="/static/pushopen.png">
                <img
                  v-show="language"
                  class="imgMax"
                  src="/static/pushopen.png"
                  id="openMapUserPushopenmg"
                >
              </div>
            </div>
            <div class="openMapUserPushbutton" id="openMapUserPushbuttonlanguage" v-show="language">
              <div
                class="openMapUserLanguage"
                @click="userConfig.Language = 1;updateConfig(1, 'Language')"
              >
                <div class="openMapUserLanguageName">简体中文</div>
                <div class="openMapUserLanguageSelect" v-show="userConfig.Language == 1">
                  <div class="openMapUserLanguageSelectInside"></div>
                </div>
              </div>
              <div
                class="openMapUserLanguage"
                @click="userConfig.Language = 2;updateConfig(2, 'Language')"
              >
                <div class="openMapUserLanguageName">繁体中文</div>
                <div class="openMapUserLanguageSelect" v-show="userConfig.Language == 2">
                  <div class="openMapUserLanguageSelectInside"></div>
                </div>
              </div>
              <div
                class="openMapUserLanguage"
                @click="userConfig.Language = 3;updateConfig(3, 'Language')"
              >
                <div class="openMapUserLanguageName">英文(US)</div>
                <div class="openMapUserLanguageSelect" v-show="userConfig.Language == 3">
                  <div class="openMapUserLanguageSelectInside"></div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="openMapUserOpenimg">
          <img
            class="imgMax"
            id="openMapUserphoto"
            :src="req.localhost + global.userinfo.photo"
            @click="userInfo = !userInfo"
          >
        </div>
      </div>
      <!-- 用户信息下部分张开 -->
      <div class="openMapUserInfo" v-show="userInfo">
        <div class="openMapUserInfoExplan">
          <div class="openMapUserInfoExplanRight"></div>
          <div class="openMapUserInfoExplanLeft">个人中心</div>
        </div>
        <div class="openMapUserInfoExplan">
          <div class="openMapUserInfoExplanRight"></div>
          <div class="openMapUserInfoExplanLeft">账户管理</div>
        </div>
        <div class="openMapUserInfoExplan" @click="logout()">
          <div class="openMapUserInfoExplanLogout">
            <img src="/static/logout.png" class="imgMax">
          </div>
          <div class="openMapUserInfoExplanLeft">注销登陆</div>
        </div>
      </div>
    </div>
    <div>
      <!-- 左侧功能按钮 -->
      <div class="openMapfunction">
        <!-- 全屏按钮 -->
        <div class="openMapFull" @click="fullscreenShow(true)" v-show="!fullscreen">
          <img src="/static/fullscreen.png" class="imgMax">
          <em class="openMapFullem"></em>
        </div>
        <!-- 地图选择按钮 -->
        <div class="openMapMap" @mouseover="mapIndex.show=true" @mouseleave="mapIndex.show=false">
          <img src="/static/map.png" class="imgMax">
          <em class="openMapFullem"></em>
          <div class="openMapMapSelect" v-show="mapIndex.show">
            <div class="openMapMapSelectC">
              <div class="openMapMapSelectCc" @click="mapIndex.index = 0">
                <div class="openMapMapSelectSe">
                  <div class="openMapMapSelectSeInside" v-show="mapIndex.index == 0"></div>
                </div>
                <div class="openMapMapSelectIndex">谷歌地图</div>
              </div>
              <div class="openMapMapSelectCc" @click="mapIndex.index = 1">
                <div class="openMapMapSelectSe">
                  <div class="openMapMapSelectSeInside" v-show="mapIndex.index == 1"></div>
                </div>
                <div class="openMapMapSelectIndex">谷歌卫星</div>
              </div>
              <div class="openMapMapSelectCc" @click="mapIndex.index = 2">
                <div class="openMapMapSelectSe">
                  <div class="openMapMapSelectSeInside" v-show="mapIndex.index == 2"></div>
                </div>
                <div class="openMapMapSelectIndex">离线地图</div>
              </div>
            </div>
          </div>
        </div>
        <!-- 放大缩小按钮 -->
        <div class="openMapZoon">
          <div class="openMapFull" @click="changeZoom+=1">
            <strong>+</strong>
            <em class="openMapFullem"></em>
          </div>
          <div class="openMapFull" @click="changeZoom-=1">
            <strong>-</strong>
            <em class="openMapFullem"></em>
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
      push: false,
      pushExpand: false,
      pushMessage: false,
      pushSound: false,
      language: false,
      useropen: false,
      userInfo: false,
      userConfig: { Message: 1, Sound: 1, Language: 1 },
      fullscreen: false,
      changeZoom: 0,
      mapIndex: {
        show: false,
        index: 0
      }
    };
  },
  methods: {
    fullscreenShow: function(metheds) {
      if (metheds == this.fullscreen) {
        return;
      }
      this.fullscreen = metheds;
      if (
        !(
          document.mozFullScreenEnabled ||
          document.fullscreenEnabled ||
          document.webkitFullscreenEnabled ||
          document.msFullscreenEnabled
        )
      ) {
        return;
      }

      if (this.fullscreen) {
        document.getElementById("openmap").requestFullscreen();
      } else {
        // if (!IsFull) {
        //   return;
        // }
        // if (document.exitFullscreen) {
        //   document.exitFullscreen();
        // } else if (document.webkitExitFullscreen) {
        //   document.webkitExitFullscreen();
        // } else if (document.msExitFullscreen) {
        //   document.msExitFullscreen();
        // } else if (document.mozCancelFullScreen) {
        //   document.mozCancelFullScreen();
        // }
      }
    },
    logout: function() {
      this.global.userinfo = {};
      this.$router.push({
        name: "index",
        params: {
          logout: true
        }
      });
    },
    changeStatus: function(status, key) {
      if (status == 1) {
        status = 2;
      } else {
        status = 1;
      }
      this.updateConfig(status, key);
      return status;
    },
    updateConfig: function(stauts, key) {
      var data = {};
      data[key] = stauts;
      this.req.put("/config", data);
    },
    getConfig: function() {
      this.req.get("/config").then(response => {
        if (response.status != 200) {
          return;
        }
        this.userConfig = response.data;
        this.$forceUpdate();
      });
    }
  },
  mounted() {
    this.getConfig();
  },
  watch: {
    changeZoom: function(v) {
      console.log(v, "++++++++");
    }
  }
};
</script>

<style scoped>
@import url("./css/openMap.css");
</style>
