<template>
  <div id="openmap" @keydown.esc="fullscreenShow(false)">
    <!-- 地图页 -->
    <div class="opeMapAllRight" ref="opeMapAllRight">
      <lmap :changeZoom="changeZoom" :mapindex="mapIndex.index"></lmap>
    </div>
    <!-- 展开设备菜单 -->
    <div class="openMapDevicesList" v-show="devicesList">
      <!-- 收起 -->
      <div class="openMapDevicesListclose" @click="devicesList = false">
        <div class="openMapDevicesright">
          <img class="imgMax" src="/static/left.png">
        </div>
      </div>
      <!-- 设备列表 -->
      <div class="openMapDevicesListInfo">
        <div class="openMapDevicesListTitle">
          <div
            v-show="titleselect != 1"
            class="openMapDevicesListTitlec"
            ref="devicetitle1"
            @click="titleselect=1"
          >设备</div>
          <div
            v-show="titleselect == 1"
            class="openMapDevicesListTitlec"
            id="openMapDevicesListTitlecs"
            ref="devicetitle1"
          >设备</div>
          <div
            v-show="titleselect != 2"
            class="openMapDevicesListTitlec"
            ref="devicetitle1"
            @click="titleselect=2"
          >状态</div>
          <div
            v-show="titleselect == 2"
            class="openMapDevicesListTitlec"
            id="openMapDevicesListTitlecs"
            ref="devicetitle1"
          >状态</div>
          <div
            v-show="titleselect != 3"
            class="openMapDevicesListTitlec"
            ref="devicetitle1"
            @click="titleselect=3"
          >轨迹</div>
          <div
            v-show="titleselect == 3"
            class="openMapDevicesListTitlec"
            id="openMapDevicesListTitlecs"
            ref="devicetitle1"
          >轨迹</div>
        </div>
        <!-- 设备列表 -->
        <div class="openMapdevicesinfo">
          <!-- 搜索栏 -->
          <div class="openMapdevicessearch">
            <input type="text" class="openMapdevicessearchinside" placeholder="请输入关键字搜索">
            <div class="openMapdevicessearchinbotton">搜索</div>
          </div>
          <!-- 设备分组 -->
          <div class="openMapdevicesGroup">
            <div class="openMapdevicesGroupName" v-for="(v, i) in groundUser" :key="i+v.Name">
              <div class="openMapdevicesGroupNameimg">
                <img src="/static/right.png" class="imgMax" v-show="!v.Show">
                <img src="/static/bottom.png" class="imgMax" v-show="v.Show">
              </div>
              <div
                class="openMapdevicesGroupNameExplain"
                @click="v.Show = !v.Show"
              >{{v.Name}} ( {{v.Devices.all}} )</div>
              <div v-show="v.Show">
                <div
                  class="openMapdevicesGroupDevices"
                  v-for="(v1, i1) in v.Devices.data"
                  :key="i1+'-'+v1.UID"
                >
                  <div class="openMapdevicesGroupDevicesName">{{v1.Name}}</div>
                  <div class="openMapdevicesGroupDevicesState">
                    <div class="openMapdevicesGroupDevicesStateColor"></div>
                    <div class="openMapdevicesGroupDevicesStateExplain">离线</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-show="!fullscreen.full">
      <!-- logo图标 -->
      <div class="openmaplogo">
        <img class="imgMax" :src="req.localhost + global.logo[3]">
      </div>
      <!-- 设备菜单 -->
      <div class="openMapDevices" v-show="!devicesList" @click="devicesList = true">
        <div class="openMapDevicesOpen">设备列表</div>
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
        <div
          class="openMapFull"
          @click="fullscreenShow(true)"
          v-show="!fullscreen.full"
          @mouseover="fullscreen.explain = true"
          @mouseleave="fullscreen.explain = false"
        >
          <div class="openMapZoomOut">
            <img src="/static/fullscreen.png" class="imgMax">
          </div>
          <em class="openMapFullem"></em>
          <div class="openMapZoomOutLeft" v-show="fullscreen.explain">全屏显示</div>
        </div>
        <!-- 地图选择按钮 -->
        <div class="openMapMap" @mouseover="mapIndex.show=true" @mouseleave="mapIndex.show=false">
          <div class="openMapZoomOut">
            <img src="/static/map.png" class="imgMax">
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
          <em class="openMapFullem"></em>
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
        <!-- 设备显示 -->
        <div
          class="openMapDeviceShow"
          @mouseover="deviceShow.explain = true"
          @mouseleave="deviceShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="deviceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/device.png" class="imgMax">
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!deviceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/device.png" class="imgMax">
            </div>
          </div>
          <em class="openMapFullem" @click="deviceShow.select=!deviceShow.select"></em>
          <div class="openMapZoomOutLeft" v-show="deviceShow.explain">设备</div>
        </div>
        <!-- 围栏显示 -->
        <div
          class="openMapFence"
          @mouseover="fenceShow.explain = true"
          @mouseleave="fenceShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="fenceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/fence.png" class="imgMax">
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!fenceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/fence.png" class="imgMax">
            </div>
          </div>
          <em class="openMapFullem" @click="fenceShow.select=!fenceShow.select"></em>
          <div class="openMapZoomOutLeft" v-show="fenceShow.explain">围栏</div>
        </div>
        <!-- 轨迹显示 -->
        <div
          class="openMapFence"
          @mouseover="orbitShow.explain = true"
          @mouseleave="orbitShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="orbitShow.select">
            <div class="openMapZoomOut">
              <img src="/static/orbit.png" class="imgMax">
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!orbitShow.select">
            <div class="openMapZoomOut">
              <img src="/static/orbit.png" class="imgMax">
            </div>
          </div>
          <em class="openMapFullem" @click="orbitShow.select=!orbitShow.select"></em>
          <div class="openMapZoomOutLeft" v-show="orbitShow.explain">轨迹</div>
        </div>
        <!-- 设备状态 -->
        <div
          class="openMapFence"
          @mouseover="statusShow.explain = true"
          @mouseleave="statusShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="statusShow.select">
            <div class="openMapZoomOut">
              <img src="/static/status.png" class="imgMax">
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!statusShow.select">
            <div class="openMapZoomOut">
              <img src="/static/status.png" class="imgMax">
            </div>
          </div>
          <em class="openMapFullem" @click="statusShow.select=!statusShow.select"></em>
          <div class="openMapZoomOutLeft" v-show="statusShow.explain">设备状态</div>
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
      devicesQuery: "",
      groundUser: [],
      titleselect: 1,
      // ------------//
      devicesList: false,
      push: false,
      pushExpand: false,
      pushMessage: false,
      pushSound: false,
      language: false,
      useropen: false,
      userInfo: false,
      userConfig: { Message: 1, Sound: 1, Language: 1 },
      fullscreen: {
        full: false,
        explain: false
      },
      changeZoom: 0,
      mapIndex: {
        show: false,
        index: 0
      },
      deviceShow: {
        select: false,
        explain: false
      },
      fenceShow: {
        select: false,
        explain: false
      },
      orbitShow: {
        select: false,
        explain: false
      },
      statusShow: {
        select: false,
        explain: false
      }
    };
  },
  methods: {
    getDevicesNumber: async function(id) {
      var response = await this.req.get(
        "/devices?uid=" + id + this.devicesQuery
      );
      if (response.status != 200) {
        return {
          all: 0,
          data: []
        };
      }
      return response.data;
    },
    getUserGround: async function() {
      this.groundUser.push({
        ID: this.global.userinfo.id,
        Name: this.global.userinfo.name,
        Devices: await this.getDevicesNumber(this.global.userinfo.id),
        Show: false
      });
      var permisson = this.global.userinfo.permisson - 1;
      var url = "/users?permisson=" + permisson;
      var response = await this.req.get(url);
      if (response.status != 200) {
        return;
      }
      for (const v of response.data.data) {
        this.groundUser.push({
          ID: v.ID,
          Name: v.Name,
          Devices: await this.getDevicesNumber(v.ID),
          Show: false
        });
      }
    },
    // ----------//
    fullscreenShow: function(metheds) {
      if (metheds == this.fullscreen.full) {
        return;
      }
      this.fullscreen.full = metheds;
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

      if (this.fullscreen.full) {
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
    this.getUserGround();
    this.getConfig();
  },
  watch: {
    devicesList: function(v) {
      var width = document.body.clientWidth;
      if (v) {
        width -= 360;
        this.$refs.opeMapAllRight.style = "width:" + width + "px";
      } else {
        this.$refs.opeMapAllRight.style = "width:" + width + "px";
      }
      this.$forceUpdate();
    }
  }
};
</script>

<style scoped>
@import url("./css/openMap.css");
</style>
