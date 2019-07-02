<template>
  <div id="openmap" @keydown.esc="fullscreenShow(false)">
    <!-- 地图页 -->
    <div class="opeMapAllRight" ref="opeMapAllRight">
      <lmap
        ref="lmap"
        :changeZoom="changeZoom"
        :mapindex="mapIndex.index"
        :devicesMarks="devicesMarks"
      ></lmap>
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
            <input
              type="text"
              class="openMapdevicessearchinside"
              placeholder="请输入关键字搜索"
              v-model="devicesQuery"
            >
            <div class="openMapdevicessearchinbotton" @click="getDevicesQuery()">搜索</div>
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
                @click="v.Show = !v.Show; getDevicesState(v)"
              >{{v.Name}} ( {{v.Devices.all}} )</div>
              <div v-show="v.Show">
                <div
                  class="openMapdevicesGroupDevices"
                  v-for="(v1, i1) in v.Devices.data"
                  :key="i1+'-'+v1.UID"
                >
                  <div @click="changeLampZoom(v1.DeviceData.Latitude, v1.DeviceData.Longitude)">
                    <div class="openMapdevicesGroupDevicesName">{{v1.Name}}</div>
                    <!-- 显示状态 -->
                    <div class="openMapdevicesGroupDevicesState" v-if="v1.DeviceData">
                      <div
                        class="openMapdevicesGroupDevicesStateColor"
                        :id="'openMapdevicesGroupDevicesStateExplain'+v1.DeviceData.State"
                      ></div>
                      <div
                        class="openMapdevicesGroupDevicesStateExplain"
                      >{{global.state[v1.DeviceData.State-1].States}}</div>
                      <div
                        class="openMapdevicesGroupDevicesStateTime"
                      >{{new Date(v1.DeviceData.Uptime * 1000).toLocaleString()}}</div>
                    </div>
                  </div>
                  <!-- 子菜单 -->
                  <div class="openMapdevicesGroupDevicesNameOption">
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesDevicesOperat">
                      <div class="openMapdevicesDevicesOperatInside">
                        <div class="openMapdevicesDevicesOperatFrame"></div>
                      </div>
                    </div>
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
          <em
            class="openMapFullem"
            @click="deviceShow.select=!deviceShow.select; deviceShow.select ? devicesMarks={}: getDevicesQuery()"
          ></em>
          <div
            class="openMapZoomOutLeft"
            v-show="deviceShow.explain"
          >{{deviceShow.select?'显示设备':'隐藏设备' }}</div>
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
          <div
            class="openMapZoomOutLeft"
            v-show="fenceShow.explain"
          >{{fenceShow.select?'显示围栏':'隐藏围栏' }}</div>
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
          <div
            class="openMapZoomOutLeft"
            v-show="orbitShow.explain"
          >{{orbitShow.select?'显示轨迹':'隐藏轨迹' }}</div>
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
          <div
            class="openMapZoomOutLeft"
            v-show="statusShow.explain"
          >{{statusShow.select?'显示数据':'隐藏数据' }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import lmap from "@/components/lmap.vue";
export default {
  components: {
    lmap
  },
  data() {
    return {
      maxMinCenter: [22.593262, 113.925971, 22.593262, 113.925971],
      // ---//
      devicesMarks: {},
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
      },
      // ---//
      screenWidth: ""
    };
  },
  methods: {
    // 计算放大等级
    countZoom: function() {
      var y = getMaxArea(this.maxMinCenter[0], this.maxMinCenter[1], 0);
      var x = getMaxArea(this.maxMinCenter[2], this.maxMinCenter[3], 1);
      this.setZoom(x > y ? y : x);
    },
    // 设置放大等级
    setZoom: function(v) {
      this.$refs.lmap.zoom = v;
    },
    // 计算经纬度最大最小值
    countMaxMinCenter: function(lat, long) {
      if (lat > this.maxMinCenter[0]) {
        this.maxMinCenter[0] = lat;
      } else if (lat < this.maxMinCenter[1]) {
        this.maxMinCenter[1] = lat;
      }
      if (long > this.maxMinCenter[2]) {
        this.maxMinCenter[2] = long;
      } else if (long < this.maxMinCenter[3]) {
        this.maxMinCenter[3] = long;
      }
    },
    // 改变中心点
    setCenter: function() {
      if (this.maxMinCenter.length != 4) {
        return;
      }
      this.$refs.lmap.center = [
        (this.maxMinCenter[0] + this.maxMinCenter[1]) / 2,
        (this.maxMinCenter[2] + this.maxMinCenter[3]) / 2,
        1
      ];
      this.countZoom();
    },
    // 改变视野到单个标记
    changeLampZoom: function(lat, long) {
      this.$refs.lmap.center = [lat, long];
      // this.$set(this.$refs.lmap.center, "time", new Date().getTime());
      this.$refs.lmap.zoom = 16;
    },
    // 计算中心点
    countCenter: function() {
      this.maxMinCenter = new Array();
      for (const userid in this.devicesMarks) {
        for (const data of this.devicesMarks[userid]) {
          if (this.maxMinCenter.length) {
            this.countMaxMinCenter(
              data.DeviceData.Latitude,
              data.DeviceData.Longitude
            );
          } else {
            this.maxMinCenter = [
              data.DeviceData.Latitude,
              data.DeviceData.Latitude,
              data.DeviceData.Longitude,
              data.DeviceData.Longitude
            ];
          }
        }
      }
      this.setCenter();
    },
    // 添加标记
    pushDevicesMark: function(v, id) {
      // 判断是否隐藏标记
      if (this.deviceShow.select) {
        return;
      }
      var data = this.devicesMarks[id];
      if (!data) {
        data = new Array();
      }
      v.DeviceData.Infos = JSON.parse(v.DeviceData.Infos);
      data.push(v);
      this.$set(this.devicesMarks, id, data);
      this.countCenter();
    },
    pullDevicesMark: function(v) {
      // 删除标记
      this.$delete(this.devicesMarks, v.ID);
      this.countCenter();
    },
    getDevicesQuery: async function() {
      this.devicesMarks = {};
      for (var user of this.groundUser) {
        user.Devices = await this.getDevices(user.ID, this.devicesQuery);
        this.getDevicesState(user);
      }
      this.$forceUpdate();
    },
    getDevicesState: function(v) {
      if (v.Show) {
        for (const v1 of v.Devices.data) {
          this.getDeviceData(v1, v.ID);
        }
      } else {
        this.pullDevicesMark(v);
      }
    },
    getDeviceData: function(v, id) {
      this.req.get("/devicedata?limit=1&id=" + v.ID).then(response => {
        if (response.status != 200) {
          return;
        }
        if (response.data.all == 0) {
          return;
        }
        v.DeviceData = response.data.data[0];
        this.pushDevicesMark(v, id);
        this.$forceUpdate();
      });
    },
    getDevices: async function(id, query = "") {
      var methed = "uid=" + id;
      if (this.global.userinfo.permisson == 3) {
        methed = "classid" + id;
      }
      var response = await this.req.get(
        "/devices?" + methed + "&name=" + query
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
      if (this.global.userinfo.permisson != 3) {
        this.groundUser.push({
          ID: this.global.userinfo.id,
          Name: this.global.userinfo.name,
          Devices: await this.getDevices(this.global.userinfo.id),
          Show: false
        });
      }
      if (this.global.userinfo.permisson != 1) {
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
            Devices: await this.getDevices(v.ID),
            Show: false
          });
        }
      }
    },
    // ----------//
    fullscreenShow: function(metheds) {
      if (metheds == this.fullscreen.full) {
        return;
      }
      this.fullscreen.full = metheds;
      if (!metheds) {
        return;
      }
      this.devicesList = false;
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
      document.getElementById("openmap").requestFullscreen();
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
    },
    getState: function() {
      if (this.global.userinfo.permisson == 3) {
        return;
      }
      this.req.get("/configState").then(response => {
        if (response.status != 200) {
          return;
        }
        this.global.state = response.data;
      });
    },
    updateSreen: function() {
      var width = document.body.clientWidth;
      var v = this.devicesList;
      if (v) {
        width -= 360;
        this.$refs.opeMapAllRight.style.width = width + "px";
      } else {
        this.$refs.opeMapAllRight.style.width = "100%";
      }
      this.$forceUpdate();
    }
  },
  mounted() {
    this.getUserGround();
    this.getConfig();
    this.getState();
    this.screenWidth = document.body.clientWidth;
    window.onresize = () => {
      return (() => {
        this.screenWidth = document.body.clientWidth;
      })();
    };
  },
  watch: {
    // 浏览器尺寸改变
    screenWidth: function() {
      this.updateSreen();
    },
    devicesList: async function() {
      this.updateSreen();
    }
  }
};

function getMaxArea(max, min, refer = 0) {
  var y = 90;
  var ym = 0.02;
  if (refer) {
    y = 180;
    ym = 0.009;
  }
  var ly = (max - min).toFixed(4);
  var zy = 1;
  if (ly != 0) {
    while (1) {
      if (ly > y) {
        break;
      }
      if (y < ym) {
        zy = 15;
        break;
      }
      y = y / 2;
      zy++;
    }
  } else {
    zy = 16;
  }
  return zy;
}
</script>

<style scoped>
@import url("./css/openMap.css");
</style>
