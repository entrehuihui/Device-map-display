<template>
  <div id="openmap" @keydown.esc="fullscreenShow(false)">
    <!-- 地图页 -->
    <div class="opeMapAllRight" ref="opeMapAllRight">
      <lmap
        ref="lmap"
        :changeZoom="changeZoom"
        :mapindex="mapIndex.index"
        :devicesMarks="devicesMarks"
        :devicePloy="orbitDevices"
        v-on:LatLng="LatLng"
      ></lmap>
    </div>
    <!-- 展开设备菜单 -->
    <div class="openMapDevicesList" v-show="devicesList">
      <!-- 收起 -->
      <div class="openMapDevicesListclose" @click="devicesList = false">
        <div class="openMapDevicesright">
          <img class="imgMax" src="/static/left.png" />
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
          >实时数据</div>
          <div
            v-show="titleselect == 2"
            class="openMapDevicesListTitlec"
            id="openMapDevicesListTitlecs"
            ref="devicetitle1"
          >实时数据</div>
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
        <div class="openMapdevicesinfo" v-show="titleselect==1">
          <!-- 搜索栏 -->
          <div class="openMapdevicessearch">
            <input
              type="text"
              class="openMapdevicessearchinside"
              placeholder="请输入关键字搜索"
              v-model="devicesQuery"
            />
            <div class="openMapdevicessearchinbotton" @click="getDevicesQuery()">搜索</div>
          </div>
          <!-- 设备分组 -->
          <div class="openMapdevicesGroup">
            <div class="openMapdevicesGroupName" v-for="(v, i) in groundUser" :key="i+v.Name">
              <div class="openMapdevicesGroupNameimg">
                <img src="/static/right.png" class="imgMax" v-show="!v.Show" />
                <img src="/static/bottom.png" class="imgMax" v-show="v.Show" />
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
                  <div
                    @click="changeLampZoom(v1.DeviceData.Latitude, v1.DeviceData.Longitude)"
                    @dblclick="; getDevicesOrbit(v1)"
                  >
                    <div class="openMapdevicesGroupDevicesName">{{v1.Name}}</div>
                    <!-- 显示状态 -->
                    <div class="openMapdevicesGroupDevicesState" v-if="v1.DeviceData">
                      <div
                        class="openMapdevicesGroupDevicesStateColor"
                        :id="'openMapdevicesGroupDevicesStateExplain'+v1.DeviceData.State"
                      ></div>
                      <div
                        class="openMapdevicesGroupDevicesStateExplain"
                      >{{global.getState(v1.DeviceData.State).States}}</div>
                      <div
                        class="openMapdevicesGroupDevicesStateTime"
                      >{{new Date(v1.DeviceData.Uptime * 1000).toLocaleString()}}</div>
                    </div>
                  </div>
                  <!-- 子菜单 -->
                  <div class="openMapdevicesGroupDevicesNameOption" v-if="v1.DeviceData">
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesGroupDevicesNameOptions"></div>
                    <div class="openMapdevicesDevicesOperat">
                      <div class="openMapdevicesDevicesOperatInside">
                        <div class="openMapdevicesDevicesOperatFrame">
                          <div class="openMapdevicesDevicesOperatFrameC">
                            <tr>
                              <th>EUI</th>
                              <th>&nbsp;:&nbsp;</th>
                              <th>{{v1.DevEUI}}</th>
                            </tr>
                          </div>
                          <div class="openMapdevicesDevicesOperatFrameC">
                            <tr>
                              <th>Latitude</th>
                              <th>&nbsp;:&nbsp;</th>
                              <th>{{v1.DeviceData.Latitude}}</th>
                            </tr>
                          </div>
                          <div class="openMapdevicesDevicesOperatFrameC">
                            <tr>
                              <th>Longitude</th>
                              <th>&nbsp;:&nbsp;</th>
                              <th>{{v1.DeviceData.Longitude}}</th>
                            </tr>
                          </div>
                          <div
                            class="openMapdevicesDevicesOperatFrameC"
                            v-for="(v2, v2index) in v1.DeviceData.Infos"
                            :key="v2 + i1"
                          >
                            <tr>
                              <th>{{v2index}}</th>
                              <th>&nbsp;:&nbsp;</th>
                              <th>{{v2}}</th>
                            </tr>
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
        <!-- 实时列表 -->
        <div class="openMapdevicesinfo" v-show="titleselect==2">
          <!-- 没有数据 -->
          <div class="openMapdeviceOrbitNo" v-show="!devicesState.length">暂无数据</div>
          <div class="openMapdevicesState" v-show="devicesState.length">
            <div>
              <div
                class="openMapdevicesGroupDevices"
                v-for="(v1, k) in devicesState"
                :key="k+'states'"
              >
                <div
                  @click="changeLampZoom(v1.DeviceData.Latitude, v1.DeviceData.Longitude, 'state', v1);"
                >
                  <div class="openMapdevicesGroupDevicesName">{{v1.Name}}</div>
                  <!-- 显示状态 -->
                  <div class="openMapdevicesGroupDevicesState">
                    <div
                      class="openMapdevicesGroupDevicesStateColor"
                      :id="'openMapdevicesGroupDevicesStateExplain'+v1.DeviceData.State"
                    ></div>
                    <div
                      class="openMapdevicesGroupDevicesStateExplain"
                    >{{global.getState(v1.DeviceData.State).States}}</div>
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
                      <div class="openMapdevicesDevicesOperatFrame">
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>EUI</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DevEUI}}</th>
                          </tr>
                        </div>
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>Latitude</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DeviceData.Latitude}}</th>
                          </tr>
                        </div>
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>Longitude</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DeviceData.Longitude}}</th>
                          </tr>
                        </div>
                        <div
                          class="openMapdevicesDevicesOperatFrameC"
                          v-for="(v2, v2index) in v1.DeviceData.Infos"
                          :key="v2 + 'state'"
                        >
                          <tr>
                            <th>{{v2index}}</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v2}}</th>
                          </tr>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- 轨迹列表 -->
        <div class="openMapdevicesinfo" v-show="titleselect==3">
          <!-- 没有选择设备 -->
          <div class="openMapdeviceOrbitNo" v-show="!orbitDevices.Name">双击设备列表选择设备</div>
          <!-- 选择设备后 -->
          <div v-show="orbitDevices.Name">
            <!-- 设备名称 -->
            <div class="openMapdeviceOrbitName">{{orbitDevices.Name}}</div>
            <!-- 检索条件 -->
            <div class="openMapdeviceOrbitTime">
              <div class="openMapdeviceOrbitTimeOne">
                <div class="openMapdeviceOrbitTimeTwo">开始时间:</div>
                <times ref="time1"></times>
              </div>
              <div class="openMapdeviceOrbitTimeOne">
                <div class="openMapdeviceOrbitTimeTwo">结束时间:</div>
                <times ref="time2"></times>
              </div>
            </div>
            <!-- 检索按钮 -->
            <div>
              <div class="openMapdeviceOrbitSearch" @click="getOrbit()">搜索</div>
              <div class="openMapdeviceOrbitClear" @click="clearTime()">清除</div>
            </div>
            <!-- 轨迹点列表 -->
            <div class="openMapdeviceOrbitList">
              <div
                class="openMapdevicesGroupDevices"
                v-for="(v1, k) in orbitDevices.Orbit"
                :key="k+'orbits'"
              >
                <div
                  @click="changeLampZoom(v1.DeviceData.Latitude, v1.DeviceData.Longitude, 'orbit', v1);"
                >
                  <div class="openMapdevicesGroupDevicesName">{{orbitDevices.Name}}</div>
                  <!-- 显示状态 -->
                  <div class="openMapdevicesGroupDevicesState">
                    <div
                      class="openMapdevicesGroupDevicesStateColor"
                      :id="'openMapdevicesGroupDevicesStateExplain'+v1.DeviceData.State"
                    ></div>
                    <div
                      class="openMapdevicesGroupDevicesStateExplain"
                    >{{global.getState(v1.DeviceData.State).States}}</div>
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
                      <div class="openMapdevicesDevicesOperatFrame">
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>EUI</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DevEUI}}</th>
                          </tr>
                        </div>
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>Latitude</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DeviceData.Latitude}}</th>
                          </tr>
                        </div>
                        <div class="openMapdevicesDevicesOperatFrameC">
                          <tr>
                            <th>Longitude</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v1.DeviceData.Longitude}}</th>
                          </tr>
                        </div>
                        <div
                          class="openMapdevicesDevicesOperatFrameC"
                          v-for="(v2, v2index) in v1.DeviceData.Infos"
                          :key=" v2  + v1.DeviceData.ID +'orbit'"
                        >
                          <tr>
                            <th>{{v2index}}</th>
                            <th>&nbsp;:&nbsp;</th>
                            <th>{{v2}}</th>
                          </tr>
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
    </div>
    <div v-show="!fullscreen.full">
      <!-- logo图标 -->
      <div class="openmaplogo">
        <img class="imgMax" :src="req.localhost + logo" />
      </div>
      <!-- 设备菜单 -->
      <div class="openMapDevices" v-show="!devicesList" @click="devicesList = true">
        <div class="openMapDevicesOpen">设备列表</div>
        <div class="openMapDevicesleft">
          <img class="imgMax" src="/static/right.png" />
        </div>
      </div>
      <!-- 用户头像信息 -->
      <div class="openMapUser" @mouseover="useropen = true" @mouseleave="useropen = false">
        <!-- 左侧展开 -->
        <div class="openMapUserOpen" v-show="useropen" @click="userInfo = false">
          <!-- 围栏增删 -->
          <div class="openMapUserPush" @mouseleave="orbit=false">
            <div @click="orbit=!orbit">
              <div class="openMapUserPushimg">
                <img class="imgMax" src="/static/fence.png" />
              </div>
              <div class="openMapUserPushExplain">围栏</div>
              <div class="openMapUserPushopenimg">
                <img v-show="!orbit" class="imgMax" src="/static/pushopen.png" />
                <img
                  v-show="orbit"
                  class="imgMax"
                  src="/static/pushopen.png"
                  id="openMapUserPushopenmg"
                />
              </div>
            </div>
            <div class="openMapUserPushbutton" v-show="orbit" @click="orbit=false">
              <div class="openMapUserPushbuttonExplain">
                <div class="openMapUserPushOrbit" @click="orbitList.Show = 1">添加围栏</div>
              </div>
              <div class="openMapUserPushbuttonExplain">
                <div class="openMapUserPushOrbit" @click="delFence(0)">删除</div>
              </div>
            </div>
          </div>
          <!-- 推送开关 -->
          <div class="openMapUserPush" @mouseleave="push=false">
            <div @click="push=!push">
              <div class="openMapUserPushimg">
                <img class="imgMax" src="/static/push.png" />
              </div>
              <div class="openMapUserPushExplain">推送</div>
              <div class="openMapUserPushopenimg">
                <img v-show="!push" class="imgMax" src="/static/pushopen.png" />
                <img
                  v-show="push"
                  class="imgMax"
                  src="/static/pushopen.png"
                  id="openMapUserPushopenmg"
                />
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
                  <div v-show="false">
                    <audio src="/static/sound.mp3" ref="pushSound" id="pushSound"></audio>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <!-- 语言选择 -->
          <div class="openMapUserPush" @mouseleave="language=false">
            <div @click="language=!language">
              <div class="openMapUserPushimg">
                <img class="imgMax" src="/static/language.png" />
              </div>
              <div class="openMapUserPushExplain">语言</div>
              <div class="openMapUserPushopenimg">
                <img v-show="!language" class="imgMax" src="/static/pushopen.png" />
                <img
                  v-show="language"
                  class="imgMax"
                  src="/static/pushopen.png"
                  id="openMapUserPushopenmg"
                />
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
          />
        </div>
      </div>
      <!-- 用户信息下部分张开 -->
      <div class="openMapUserInfo" v-show="userInfo">
        <div class="openMapUserInfoExplan" @click="jump(0)">
          <div class="openMapUserInfoExplanRight"></div>
          <div class="openMapUserInfoExplanLeft">设置中心</div>
        </div>
        <div class="openMapUserInfoExplan" @click="logout()">
          <div class="openMapUserInfoExplanLogout">
            <img src="/static/logout.png" class="imgMax" />
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
            <img src="/static/fullscreen.png" class="imgMax" />
          </div>
          <em class="openMapFullem"></em>
          <div class="openMapZoomOutLeft" v-show="fullscreen.explain">全屏显示</div>
        </div>
        <!-- 地图选择按钮 -->
        <div class="openMapMap" @mouseover="mapIndex.show=true" @mouseleave="mapIndex.show=false">
          <div class="openMapZoomOut">
            <img src="/static/map.png" class="imgMax" />
            <div class="openMapMapSelect" v-show="mapIndex.show">
              <div class="openMapMapSelectC">
                <div class="openMapMapSelectCc" @click="mapIndex.index = 0">
                  <div class="openMapMapSelectSe">
                    <div class="openMapMapSelectSeInside" v-show="mapIndex.index == 0"></div>
                  </div>
                  <div class="openMapMapSelectIndex">标准地图</div>
                </div>
                <div class="openMapMapSelectCc" @click="mapIndex.index = 1">
                  <div class="openMapMapSelectSe">
                    <div class="openMapMapSelectSeInside" v-show="mapIndex.index == 1"></div>
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
              <img src="/static/device.png" class="imgMax" />
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!deviceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/device.png" class="imgMax" />
            </div>
          </div>
          <em class="openMapFullem" @click="deviceShow.select=!deviceShow.select; changeMarkShow()"></em>
          <div
            class="openMapZoomOutLeft"
            v-show="deviceShow.explain"
          >{{deviceShow.select?'隐藏设备':'显示设备' }}</div>
        </div>
        <!-- 围栏显示 -->
        <div
          class="openMapFence"
          @mouseover="fenceShow.explain = true"
          @mouseleave="fenceShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="fenceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/fence.png" class="imgMax" />
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!fenceShow.select">
            <div class="openMapZoomOut">
              <img src="/static/fence.png" class="imgMax" />
            </div>
          </div>
          <em
            class="openMapFullem"
            @click="fenceShow.select=!fenceShow.select; showFence(fenceShow.select)"
          ></em>
          <div
            class="openMapZoomOutLeft"
            v-show="fenceShow.explain"
          >{{fenceShow.select?'隐藏围栏':'显示围栏' }}</div>
        </div>
        <!-- 轨迹显示 -->
        <div
          class="openMapFence"
          @mouseover="orbitShow.explain = true"
          @mouseleave="orbitShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="orbitShow.select">
            <div class="openMapZoomOut">
              <img src="/static/orbit.png" class="imgMax" />
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!orbitShow.select">
            <div class="openMapZoomOut">
              <img src="/static/orbit.png" class="imgMax" />
            </div>
          </div>
          <em class="openMapFullem" @click="orbitShow.select=!orbitShow.select; changeOrbit()"></em>
          <div
            class="openMapZoomOutLeft"
            v-show="orbitShow.explain"
          >{{orbitShow.select?'隐藏轨迹':'显示轨迹' }}</div>
        </div>
        <!-- 设备状态 -->
        <div
          class="openMapFence"
          @mouseover="statusShow.explain = true"
          @mouseleave="statusShow.explain = false"
        >
          <div class="openMapDeviceSelected" v-show="statusShow.select">
            <div class="openMapZoomOut">
              <img src="/static/status.png" class="imgMax" />
            </div>
          </div>
          <div class="openMapDeviceNotSelected" v-show="!statusShow.select">
            <div class="openMapZoomOut">
              <img src="/static/status.png" class="imgMax" />
            </div>
          </div>
          <em class="openMapFullem" @click="statusShow.select=!statusShow.select"></em>
          <div
            class="openMapZoomOutLeft"
            v-show="statusShow.explain"
          >{{statusShow.select?'隐藏数据':'显示数据' }}</div>
        </div>
      </div>
    </div>
    <div></div>
    <div class="openmapOrbit" id="openmapOrbitID" v-show="orbitList.Show">
      <div
        class="openmapOrbitTitle"
        @mousedown="mouseDownHandleelse($event,'openmapOrbitID')"
        @mouseup="mouseUpHandleelse($event,'openmapOrbitID')"
      >
        <strong>添加电子围栏</strong>
        <div class="openmapOrbitClose" @click="initFence()">X</div>
      </div>
      <div class="openmapOrbitTypes" v-show="orbitList.Show == 1">
        <div
          class="openmapOrbitTypesSelect"
          @click="orbitList.Types=1"
          @dblclick="orbitList.Types=1;orbitList.Show++"
        >
          <div class="openmapOrbitTypesName">添加圆形围栏</div>
          <div class="openmapOrbitTypesMark">
            <div class="openmapOrbitTypesMarkInside" v-show="orbitList.Types == 1"></div>
          </div>
        </div>
        <div
          class="openmapOrbitTypesSelect"
          @click="orbitList.Types=2"
          @dblclick="orbitList.Types=2; orbitList.Show++;setFence('exec', orbitList)"
        >
          <div class="openmapOrbitTypesName">添加多边形围栏</div>
          <div class="openmapOrbitTypesMark">
            <div class="openmapOrbitTypesMarkInside" v-show="orbitList.Types == 2"></div>
          </div>
        </div>
      </div>
      <!-- 画圆形围栏 -->
      <div class="openmapOrbitTypes" v-show="orbitList.Show==2 && orbitList.Types ==1">
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitCircleName">纬度:</div>
          <input
            type="number"
            class="openmapOrbitCircleInput"
            oninput="if(value>90)value=90;if(value<-90)value=-90"
            placeholder="点击地图获取"
            @keydown="LatLng(orbitList.Lat,orbitList.Lng, true)"
            v-model="orbitList.Lat"
          />
        </div>
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitCircleName">经度:</div>
          <input
            type="number"
            class="openmapOrbitCircleInput"
            oninput="if(value>180)value=180;if(value<-180)value=-180"
            @keydown="LatLng(orbitList.Lat,orbitList.Lng, true)"
            placeholder="或手动输入"
            v-model="orbitList.Lng"
          />
        </div>
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitCircleName">半径:</div>
          <input
            type="number"
            class="openmapOrbitCircleInput"
            placeholder="单位:米"
            v-model="orbitList.Radius"
          />
        </div>
        <div class="openmapOrbitCircle" v-show="orbitList.Radius&&orbitList.Lng&&orbitList.Lat">
          <div class="openmapOrbitCircleTrue" @click="setFence('exec', orbitList)">预览</div>
          <div class="openmapOrbitCircleTrue" @click="postFence('exec', orbitList)">确定</div>
        </div>
      </div>
      <!-- 画多边形围栏 -->
      <div class="openmapOrbitTypes" v-show="orbitList.Show==2 && orbitList.Types ==2">
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitCircleName">纬度:</div>
          <input
            type="number"
            class="openmapOrbitCircleInput"
            oninput="if(value>90)value=90;if(value<-90)value=-90"
            placeholder="点击地图获取"
            @keydown="LatLng(orbitList.Lat,orbitList.Lng, true)"
            v-model="orbitList.Lat"
          />
        </div>
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitCircleName">经度:</div>
          <input
            type="number"
            class="openmapOrbitCircleInput"
            oninput="if(value>180)value=180;if(value<-180)value=-180"
            @keydown="LatLng(orbitList.Lat,orbitList.Lng, true)"
            placeholder="或手动输入"
            v-model="orbitList.Lng"
          />
        </div>
        <div class="openmapOrbitCircle">
          <div class="openmapOrbitPolygon">第{{orbitList.Polygon.length}}个点</div>
          <div
            class="openmapOrbitPolygonNext"
            v-show="orbitList.Lng&&orbitList.Lat"
            @click="orbitList.Number++; orbitList.Lng='';orbitList.Lat=''"
          >Next</div>
          <div
            class="openmapOrbitPolygonNext"
            id="openmapOrbitPolygonNextRight"
            v-show="orbitList.Number > 0"
            @click="preFence()"
          >Pre</div>
        </div>
        <div
          class="openmapOrbitCircle"
          v-show="orbitList.Number > 1 &&orbitList.Lng&&orbitList.Lat"
        >
          <div class="openmapOrbitCircleTrue" @click="setFence('exec', orbitList)">取消</div>
          <div class="openmapOrbitCircleTrue" @click="postFence('exec', orbitList)">确定</div>
        </div>
      </div>
    </div>
    <div>
      <websocket v-on:retdata="socket" ref="websocket"></websocket>
    </div>
  </div>
</template>

<script>
import lmap from "@/components/lmap.vue";
import times from "@/components/time.vue";
import websocket from "@/components/websocket.vue";
export default {
  components: {
    lmap,
    websocket,
    times
  },
  data() {
    return {
      logo: "",
      time1: "",
      time2: "",
      devicesState: [],
      // -----//
      websock: null,
      // -------.//
      orbit: false,
      orbitList: {
        Show: 0,
        Types: 0,
        Lat: "",
        Lng: "",
        Radius: 100,
        Number: 0,
        Polygon: []
      },
      moveDataelse: {
        x: null,
        y: null
      },
      // --//
      maxMinCenter: [22.593262, 113.925971, 22.593262, 113.925971],
      // ---//
      devicesMarks: {
        Show: true
      },
      devicesQuery: "",
      groundUser: [],
      titleselect: 1, //设备菜单选择
      // ------------//
      orbitDevices: {}, //轨迹
      devicesList: true,
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
        select: true,
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
    jump: function(index = 0) {
      this.$router.push({
        name: "user",
        params: {
          index: index
        }
      });
    },
    changeMarkShow: function() {
      this.$set(this.devicesMarks, "Show", this.deviceShow.select);
    },
    // 显示隐藏轨迹
    changeOrbit: function(motheds = false) {
      this.$set(this.orbitDevices, "Show", this.orbitShow.select);
    },
    // 清除时间
    clearTime: function() {
      this.$refs.time1.minuteShortText = "";
      this.$refs.time2.minuteShortText = "";
      this.time1 = "";
      this.time2 = "";
    },
    getOrbit: function() {
      this.time1 = new Date(this.$refs.time1.minuteShortText).getTime() / 1000;
      this.time2 = new Date(this.$refs.time2.minuteShortText).getTime() / 1000;
      var query = "";
      if (this.time1) {
        query += "&starttime=" + this.time1;
      }
      if (this.time2) {
        query += "&endtime=" + this.time2;
      }
      this.req
        .get("/devicedata?limit=200&id=" + this.orbitDevices.ID + query)
        .then(response => {
          if (response.status != 200) {
            return;
          }
          var data = response.data.data;
          var devices = [];
          var OrbitList = [];
          var methods = true;
          for (const v of data) {
            v.Infos = JSON.parse(v.Infos);
            devices.push({
              Name: this.orbitDevices.Name,
              DevEUI: this.orbitDevices.DevEUI,
              DeviceData: v
            });
            OrbitList.unshift([v.Latitude, v.Longitude]);
            this.countMaxMinCenter(v.Latitude, v.Longitude);
            if (methods) {
              this.maxMinCenter = [
                v.Latitude,
                v.Latitude,
                v.Longitude,
                v.Longitude
              ];
              methods = false;
            } else {
              this.countMaxMinCenter(v.Latitude, v.Longitude);
            }
          }
          this.setCenter();
          this.$set(this.orbitDevices, "Orbit", devices);
          this.$set(this.orbitDevices, "OrbitList", OrbitList);
        });
    },
    // 获取设备轨迹
    getDevicesOrbit: function(v) {
      this.orbitDevices = v;
      this.titleselect = 3;
      this.orbitShow.select = true;
      this.changeOrbit();
      this.getOrbit();
    },
    // websocket数据
    socket: async function(msg) {
      // 是否接收实时推送
      if (this.userConfig.Message == 2) {
        return;
      }
      // 是否开启声音提示
      if (this.userConfig.Sound == 1) {
        this.$refs.pushSound.play();
      }
      // 检测是否存在设备
      var uid = 0;
      if (this.global.userinfo.permisson == 3) {
        uid = msg.Classid;
      } else {
        uid = msg.UID;
      }
      for (const user of this.groundUser) {
        if (user.ID != uid) {
          continue;
        }
        var check = true;
        for (const device of user.Devices.data) {
          if (msg.Did == device.ID) {
            check = false;
            msg.Infos = JSON.parse(msg.Infos);
            // 设备标记点
            if (user.Show) {
              device.DeviceData = msg;
              this.pushDevicesMark(device, user.ID);
            }
            // 实时数据点
            var devices = {
              Name: device.Name,
              DevEUI: device.DevEUI,
              DeviceData: msg
            };
            this.devicesState.unshift(devices);
            if (this.devicesState.length > 200) {
              this.$delete(this.devicesState, 200);
            }
            this.$forceUpdate();
            break;
          }
        }
        //设备不存在 -- 添加重新拉取设备列表
        if (check) {
          user.Devices = await this.getDevices(uid);
        }
        this.$forceUpdate();
        break;
      }

      // 轨迹点
      if (this.orbitDevices.ID == msg.Did) {
        if (this.time1 && msg.Uptime < this.time1) {
          return;
        }
        if (this.time2 && msg.Uptime > this.time2) {
          return;
        }
        this.orbitDevices.Orbit.unshift({
          Name: this.orbitDevices.Name,
          DevEUI: this.orbitDevices.DevEUI,
          DeviceData: msg
        });
        this.$set(
          this.orbitDevices.OrbitList,
          this.orbitDevices.OrbitList.length,
          [msg.Latitude, msg.Longitude]
        );
        if (this.orbitDevices.Orbit.length > 200) {
          this.$delete(this.orbitDevices.Orbit, 200);
          this.$delete(this.orbitDevices.OrbitList, 0);
        }
        // 如果轨迹是显示的
        if (this.orbitShow.select) {
          this.countMaxMinCenter(msg.Latitude, msg.Longitude);
          this.setCenter();
        }
      }
    },
    // 返回地图坐标点
    LatLng: function(lat, lng, motheds = false) {
      if (lat == "" || lng == "") {
        return;
      }
      if (this.orbitList.Show == 2 && this.orbitList.Types) {
        this.orbitList.Lat = lat;
        this.orbitList.Lng = lng;
        this.$set(this.orbitList.Polygon, this.orbitList.Number, [lat, lng]);
        this.$refs.lmap.execMark = this.orbitList.Polygon;
      }
      if (motheds) {
        this.$refs.lmap.center = [lat, lng];
      }
    },
    // 上一个围栏点
    preFence: function() {
      this.$delete(this.orbitList.Polygon, this.orbitList.Number, 1);
      this.orbitList.Number--;
    },
    // 初始化围栏点
    initFence: function(name = "exec") {
      this.orbitList = {
        Show: 0,
        Types: 0,
        Lat: "",
        Lng: "",
        Radius: 100,
        Number: 0,
        Polygon: []
      };
      this.$delete(this.$refs.lmap.orbitLists, name);
      this.$refs.lmap.execMark = new Array();
    },
    //显示隐藏围栏
    showFence: function(methods = false) {
      for (const key in this.groundUser) {
        if (this.groundUser[key].Show) {
          this.setFence(this.groundUser[key].ID, this.groundUser[key].Orbit);
        }
      }
    },
    // 设置围栏
    setFence: function(name, orbitList) {
      if (name == "exec") {
        this.$set(this.$refs.lmap.orbitLists, name, orbitList);
        return;
      }
      if (orbitList && this.fenceShow.select) {
        this.$set(this.$refs.lmap.orbitLists, name, orbitList);
      } else {
        this.$delete(this.$refs.lmap.orbitLists, name);
      }
    },
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
      this.countZoom();
      this.$refs.lmap.center = [
        (this.maxMinCenter[0] + this.maxMinCenter[1]) / 2,
        (this.maxMinCenter[2] + this.maxMinCenter[3]) / 2,
        1
      ];
    },
    // 改变视野到单个标记
    changeLampZoom: function(lat, long, metheds = null, v) {
      this.$refs.lmap.center = [lat, long];
      this.$refs.lmap.zoom = 16;
      // 如果是实时状态点击--则显示当前点击点
      if (metheds) {
        this.$set(this.devicesMarks, metheds, [v]);
      }
    },
    // 计算中心点
    countCenter: function() {
      this.maxMinCenter = new Array();
      for (const userid in this.devicesMarks) {
        var devicesMarks = this.devicesMarks[userid];
        for (const deviceid in devicesMarks) {
          var data = devicesMarks[deviceid];
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
    pushDevicesMark: function(v, uid) {
      var data = this.devicesMarks[uid];
      if (!data) {
        data = {};
      }
      data[v.ID] = v;
      this.$set(this.devicesMarks, uid, data);
      if (this.orbitShow.select && this.orbitDevices.Orbit) {
        return;
      }
      this.countCenter();
    },
    // 取消标记
    pullDevicesMark: function(v) {
      // 删除标记
      this.$delete(this.devicesMarks, v.ID);
      this.countCenter();
    },
    // 检索设备
    getDevicesQuery: async function() {
      this.devicesMarks = {};
      for (var user of this.groundUser) {
        user.Devices = await this.getDevices(user.ID, this.devicesQuery);
        this.getDevicesState(user);
      }
      this.$forceUpdate();
    },
    // 获取设备状态
    getDevicesState: async function(v) {
      if (v.Show) {
        for (const v1 of v.Devices.data) {
          this.getDeviceData(v1, v.ID);
          await this.global.sleep(100);
          if (!v.Show) {
            return;
          }
        }
        // 展示围栏
        // 获取围栏
        this.getFence(v);
      } else {
        this.pullDevicesMark(v);
        // 取消围栏展示
        this.setFence(v.ID);
      }
    },
    // 获取设备数据
    getDeviceData: function(v, id) {
      this.req.get("/devicedata?limit=1&id=" + v.ID).then(response => {
        if (response.status != 200) {
          return;
        }
        if (response.data.all == 0) {
          return;
        }
        var DeviceData = response.data.data[0];
        DeviceData.Infos = JSON.parse(DeviceData.Infos);
        // v.DeviceData.Infos = JSON.parse(v.DeviceData.Infos);
        this.$set(v, "DeviceData", DeviceData);
        this.pushDevicesMark(v, id);
      });
    },
    // 获取设备
    getDevices: async function(id, query = "") {
      var methed = "uid=" + id;
      if (this.global.userinfo.permisson == 3) {
        methed = "classid=" + id;
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
    // 获取分组
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
    // 全屏
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
    // 注销
    logout: function() {
      this.global.userinfo = {};
      this.$router.push({
        name: "index",
        params: {
          logout: true
        }
      });
    },
    // 状态改变
    changeStatus: function(status, key) {
      if (status == 1) {
        status = 2;
      } else {
        status = 1;
      }
      this.updateConfig(status, key);
      return status;
    },
    // 更新配置
    updateConfig: function(stauts, key) {
      var data = {};
      data[key] = stauts;
      this.req.put("/config", data);
    },
    // 获取配置
    getConfig: function() {
      this.req.get("/config").then(response => {
        if (response.status != 200) {
          return;
        }
        this.userConfig = response.data;
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
      });
    },
    // 改变窗口大小
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
    },
    // 获取围栏
    getFence: function(v) {
      this.req.get("/fence?id=" + v.ID).then(response => {
        if (response.status != 200) {
          return;
        }
        if (response.data.ID == 0) {
          return;
        }
        if (response.data.Types == 1) {
          var data = JSON.parse(response.data.Info);
          v.Orbit = {
            Types: 1,
            Lat: data.Lat,
            Lng: data.Lng,
            Radius: data.Radius
          };
        } else if (response.data.Types == 2) {
          v.Orbit = { Types: 2 };
          var data = JSON.parse(response.data.Info);
          var Polygon = new Array();
          for (const v of data) {
            Polygon.push([v.Lat, v.Lng]);
          }
          v.Orbit.Polygon = Polygon;
        }
        this.setFence(v.ID, v.Orbit);
      });
    },
    // 添加围栏
    postFence: function(name, v) {
      var data = {};
      // 圆形
      if (v.Types == 1) {
        data = {
          Lat: parseFloat(v.Lat),
          Lng: parseFloat(v.Lng),
          Radius: parseInt(v.Radius)
        };
      } else if (v.Types == 2) {
        if (v.Polygon.length < 3) {
          return;
        }
        data = new Array();
        for (const lang of v.Polygon) {
          data.push({
            Lat: parseFloat(lang[0]),
            Lng: parseFloat(lang[1])
          });
        }
      } else {
        return;
      }
      this.req
        .post("/fence", {
          Types: v.Types,
          Info: data
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            this.initFence();
            return;
          }
          alert("success");
          this.initFence();
        });
    },
    // 删除围栏
    delFence: function(id = 0) {
      this.req
        .del("/fence", {
          id: id
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          alert("Success");
          if (id == 0) {
            this.setFence(this.groundUser[0].ID, false);
            this.groundUser[0].Orbit = null;
          }
        });
    },
    // 监控鼠标
    mouseDownHandleelse(event, ID) {
      this.moveDataelse.x =
        event.pageX - document.getElementById(ID).offsetLeft;
      this.moveDataelse.y = event.pageY - document.getElementById(ID).offsetTop;
      window.onmousemove = event => {
        let moveLeft = event.pageX - this.moveDataelse.x + "px";
        let moveTop = event.pageY - this.moveDataelse.y + "px";
        document.getElementById(ID).style.left = moveLeft;
        document.getElementById(ID).style.top = moveTop;
      };
    },
    mouseUpHandleelse(event) {
      window.onmousemove = null;
    },
    // 获取用户信息
    getUserinfo: async function() {
      await this.req.get("/login/info").then(response => {
        if (response.status != 200) {
          this.global.setCookie("", -1);
          return;
        }
        this.global.userinfo = response.data;
      });
    },
    // getLogo
    getLogo: function() {
      // this.logo
      this.req.get("/login/logo/3").then(response => {
        if (response.status == 200) {
          this.logo = response.data;
        }
        this.$forceUpdate();
      });
    }
  },
  async mounted() {
    await this.getUserinfo();
    this.getLogo();
    this.getState();
    this.getUserGround();
    this.getConfig();
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
    devicesList: async function(v) {
      this.updateSreen();
    },
    titleselect: function(v) {
      if (v != 2) {
        this.$delete(this.devicesMarks, "state");
      }
      if (v != 3) {
        this.$delete(this.devicesMarks, "orbit");
      }
    }
  },
  beforeDestroy() {
    this.$refs.websocket.close = true;
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
