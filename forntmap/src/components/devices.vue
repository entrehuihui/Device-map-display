<template>
  <div id="devices">
    <div>
      <!-- 搜索栏 -->
      <div class="accountSearch">
        <div class="accountSearchOne">
          <div>
            <div class="accountSearchT">设备 :</div>
            <input type="text" class="accountSearchI" v-model="name" />
          </div>
          <div>
            <div class="accountSearchT">状态 :</div>
            <select class="accountSearchS" v-model="status">
              <option value="0">全部</option>
              <option value="1">启用</option>
              <option value="2">禁用</option>
              <option value="3">到期</option>
            </select>
          </div>
        </div>
        <div class="accountSearchOne">
          <div v-show="global.userinfo.permisson != 1">
            <div class="accountSearchT">分组 :</div>
            <select class="accountSearchS" id="accountSearchGround" v-model="ground">
              <option value="0">全部</option>
              <option
                :value="user.ID"
                v-for="(user) in grounds"
                :key="user.ID +'suser'"
              >{{user.Name}}</option>
            </select>
          </div>
          <div>
            <div class="accountSearchT" id="accountSearchTime">到期时间 :</div>
            <div class="accountSearchSTimes">
              <times ref="time1"></times>
            </div>
            <div class="accountSearchSTimes" id="accountSearchSTimes">
              <times ref="time2"></times>
            </div>
          </div>
          <div class="accountSearchSearch">清除</div>
          <div class="accountSearchSearch" @click="querys()">搜索</div>
        </div>
      </div>
      <!-- 设备列表 -->
      <div class="accountList">
        <div class="accountListTitle">
          <!-- 选择框 -->
          <div class="accountList1" style="visibility:hidden">
            <div class="accountList1Insede"></div>
          </div>
          <!-- 序号 -->
          <div class="accountList0">序号</div>
          <!-- 设备名 -->
          <div class="accountList6">设备名</div>
          <!-- 群组名 -->
          <div class="accountList6">群组名</div>
          <!-- 用户名 -->
          <div class="accountList6">用户名</div>
          <!-- 创建时间 -->
          <div class="accountList6">创建时间</div>
          <!-- 到期时间 -->
          <div class="accountList6">到期时间</div>
          <!-- 状态 -->
          <div class="accountList7">状态</div>
          <!-- 操作 -->
          <div class="accountList8">操作</div>
        </div>
        <div
          class="accountListTitleD"
          v-for="(device, index) in devices.data"
          :key="device.ID+'device'"
        >
          <!-- 选择框 -->
          <div class="accountList1" @click="device.Show=!device.Show;selectAll()">
            <div class="accountList1Insede" v-show="device.Show"></div>
          </div>
          <!-- 序号 -->
          <div class="accountList0">{{index + 1 + offset}}</div>
          <!-- 设备名 -->
          <div class="accountList6">{{device.Name}}</div>
          <!-- 群组名 -->
          <div class="accountList6">{{device.Classname}}</div>
          <!-- 用户名 -->
          <div class="accountList6">{{device.Uname}}</div>
          <!-- 创建时间 -->
          <div class="accountList6">{{new Date(device.Createtime*1000).toLocaleDateString()}}</div>
          <!-- 到期时间 -->
          <div class="accountList6">
            <div
              v-show="device.Expiredtime"
            >{{new Date(device.Createtime*1000).toLocaleDateString()}}</div>
            <div v-show="!device.Expiredtime">永久</div>
          </div>
          <!-- 状态 -->
          <div class="accountList7">
            <div v-show="device.Expiredtime==0 || device.Expiredtime > nowTime">
              <div v-show="device.Status==1 ">启用</div>
              <div v-show="device.Status==2 ">禁用</div>
            </div>
            <div
              v-show="device.Status==3 || (nowTime >= device.Expiredtime && device.Expiredtime!=0)"
            >到期</div>
          </div>
          <!-- 操作 -->
          <div class="accountList8">
            <div v-show="global.userinfo.permisson != 1">
              <div class="accountListDisable" v-show="device.Status!=3">
                <div v-show="device.Status==1" @click="updateStatus([device.ID],2)">禁用</div>
                <div v-show="device.Status==2" @click="updateStatus([device.ID],1)">启用</div>
              </div>
              <div
                class="accountListDisable"
                @click="allotDevices.Classid=device.Classid; allotDevices.IDs=[device.ID]; allotDevices.Name=device.Classname"
              >分配</div>
              <div class="accountListDisable" @click="expireTimeID = [device.ID]">延期</div>
              <div class="accountListDisable" @click="delDevices([device.ID])">删除</div>
            </div>
          </div>
        </div>
      </div>
      <!-- 底部/页数 -->
      <div class="accountListTitleB">
        <!-- 选择框 -->
        <div class="accountList1" @click="selectAll(true)">
          <div class="accountList1Insede" v-show="devices.Show"></div>
        </div>
        <div v-show="global.userinfo.permisson != 1">
          <!-- 禁用 -->
          <div class="accountListDisable" @click="updateAlls(1)">禁用</div>
          <!-- 启用 -->
          <div class="accountListDisable" @click="updateAlls(2)">启用</div>
          <!-- 删除 -->
          <div class="accountListDisable" @click="updateAlls(3)">删除</div>
          <!-- 分配 -->
          <div class="accountListDisable" @click="updateAlls(4)">分配</div>
          <!-- 延期 -->
          <div class="accountListDisable" @click="updateAlls(5)">延期</div>
        </div>
        <!-- 分页 -->
        <div class="accountListPages">
          <pages :all="allPage" v-on:page="page"></pages>
        </div>
      </div>
      <!-- 续期 -->
      <div v-show="expireTimeID">
        <div class="accountListExpire" @click="expireTimeID=null"></div>
        <div class="accountListExpire1">
          <div class="accountListExpire6">时间为空表示永远不过期</div>
          <div class="accountListExpire5" @click="expireTimeID=null">X</div>
          <div class="accountListExpire2">
            <div class="accountListExpire3">到期时间 :</div>
            <div class="accountListExpire3">
              <times :type="'day'" ref="time3"></times>
            </div>
            <div class="accountListExpire3">
              <div class="accountListExpire4" @click="updateExoired()">确定</div>
              <div class="accountListExpire4" @click="clearTime3()">清除</div>
            </div>
          </div>
        </div>
      </div>
      <!-- 分配 -->
      <div v-if="allotDevices.Classid">
        <selectUser
          v-on:close="allotDevices = {Classid:0, IDs:[], Name:''}"
          :allotDevices="allotDevices"
          v-on:seelctID="seelctID"
        ></selectUser>
      </div>
    </div>
  </div>
</template>

<script>
import selectUser from "@/components/selectUser.vue";
import times from "@/components/time.vue";
import pages from "@/components/pages.vue";
export default {
  components: { times, pages, selectUser },
  data() {
    return {
      name: "",
      status: 0,
      ground: 0,
      grounds: [],
      allPage: 1,
      offset: 0,
      devices: {},
      query: "",
      nowTime: parseInt(new Date().getTime() / 1000),
      expireTimeID: null,
      allotDevices: {
        IDs: [],
        Classid: 0,
        Name: ""
      }
    };
  },
  methods: {
    seelctID: function(id) {
      this.allot(this.allotDevices.IDs, id);
      this.allotDevices = {
        ID: [],
        Classid: 0,
        Name: ""
      };
      this.$forceUpdate();
    },
    updateAlls: function(methods, uid = 0) {
      // 分配
      if (methods == 4) {
        var ids = new Array();
        var name = "";
        var classid = 0;
        for (const device of this.devices.data) {
          if (device.Show) {
            ids.push(device.ID);
            if (classid == 0) {
              classid = device.Classid;
              name = device.Classname;
            } else if (classid != device.Classid) {
              alert("只能分配同一群组下设备");
              return;
            }
          }
        }
        this.allotDevices = {
          IDs: ids,
          Classid: classid,
          Name: name
        };
        this.$forceUpdate();
        return;
      }
      var ids = new Array();
      for (const device of this.devices.data) {
        if (device.Show) {
          ids.push(device.ID);
        }
      }
      if (ids.length == 0) {
        return;
      }
      switch (methods) {
        case 1: //禁用
          this.updateStatus(ids, 2);
          break;
        case 2: //启用
          this.updateStatus(ids, 1);
          break;
        case 3: //删除
          this.delDevices(ids);
          break;
        case 5: //延期
          this.expireTimeID = ids;
          this.$forceUpdate();
          break;
      }
    },
    allot: function(ids, uid) {
      this.req
        .put("/devices/user", {
          ids: ids,
          uid: uid
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
          } else {
            alert(response.data);
            this.getDevices();
          }
        });
    },
    clearTime3: function() {
      this.$refs.time3.dayShortText = "";
    },
    updateExoired: function() {
      var now = this.$refs.time3.dayShortText;
      if (now) {
        now = parseInt(new Date(now).getTime() / 1000);
      } else {
        now = 0;
      }
      this.req
        .put("/devices/expire", {
          expiredtime: now,
          ids: this.expireTimeID
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
          } else {
            alert(response.data);
            this.getDevices();
          }
          this.expireTimeID = null;
        });
    },
    delDevices: function(ids) {
      this.req
        .del("/devices", {
          ids: ids
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
          } else {
            alert(response.data);
            this.getDevices();
          }
        });
    },
    updateStatus: function(ids, status) {
      this.req
        .put("/devices/status", {
          ids: ids,
          status: status
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
          } else {
            alert(response.data);
            this.getDevices();
          }
        });
    },
    selectAll: function(methods = false) {
      if (methods) {
        this.devices.Show = !this.devices.Show;
        methods = this.devices.Show;
        for (var device of this.devices.data) {
          device.Show = methods;
        }
        this.$forceUpdate();
        return;
      }
      // methods.Show = !methods.Show;
      for (var device of this.devices.data) {
        if (!device.Show) {
          this.devices.Show = false;
          this.$forceUpdate();
          return;
        }
      }
      this.devices.Show = true;
      this.$forceUpdate();
    },
    page: function(pages) {
      var offset = (pages - 1) * 10;
      if (this.offset == offset) {
        return;
      }
      this.offset = offset;
      this.getDevices();
    },
    querys: function() {
      this.query = "&name=" + this.name;
      this.query += "&status=" + this.status;
      if (this.global.userinfo.permisson == 3) {
        this.query += "&classid=" + this.ground;
      } else {
        this.query += "&uid=" + this.ground;
      }
      var time1 = new Date(this.$refs.time1.minuteShortText).getTime() / 1000;
      var time2 = new Date(this.$refs.time2.minuteShortText).getTime() / 1000;
      if (time1) {
        this.query += "&starttime=" + time1;
      }
      if (time2) {
        this.query += "&endtime=" + time2;
      }
      this.allPage = 1;
      this.offset = 0;
      this.getDevices();
    },
    clearQuery: function() {
      this.name = "";
      this.status = 0;
      this.query = "";
      this.$refs.time1.minuteShortText = "";
      this.$refs.time2.minuteShortText = "";
      this.ground = 0;
    },
    getFrounds: function() {
      if (this.global.userinfo.permisson != 2) {
        return;
      }
      this.grounds = [
        {
          ID: this.global.userinfo.id,
          Name: this.global.userinfo.name
        }
      ];
      var query = "";
      if (this.global.userinfo.permisson == 3) {
        query = "permisson=2";
      }
      this.req.get("/users?" + query).then(response => {
        if (response.status != 200) {
          this.$forceUpdate();
          return;
        }
        this.grounds = response.data.data;
        this.grounds.unshift({
          ID: this.global.userinfo.id,
          Name: this.global.userinfo.name
        });
        this.$forceUpdate();
      });
    },
    // 获取设备
    getDevices: function() {
      this.req
        .get("/devices?limit=10&offset=" + this.offset + this.query)
        .then(response => {
          if (response.status != 200) {
            return;
          }
          this.devices = response.data;
          this.allPage = parseInt(this.devices.all / 10);
          if (this.devices.all % 10) {
            this.allPage++;
          }
          if (!this.allPage) {
            this.allPage = 1;
          }
        });
    }
  },
  mounted() {
    this.getFrounds();
    this.getDevices();
  }
};
</script>

<style scoped>
@import url("./css/devices.css");
@import url("./css/account.css");
</style>
