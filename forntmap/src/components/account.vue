<template>
  <div id="account">
    <div>
      <!-- 搜索栏 -->
      <div class="accountSearch">
        <div class="accountSearchOne">
          <div>
            <div class="accountSearchT">账号 :</div>
            <input type="text" class="accountSearchI" v-model="name" />
          </div>
          <div>
            <div class="accountSearchT">VIP :</div>
            <select class="accountSearchS" v-model="vip">
              <option value="0">全部</option>
              <option value="1">VIP1</option>
              <option value="2">VIP2</option>
              <option value="3">VIP3</option>
              <option value="4">VIP4</option>
              <option value="5">VIP5</option>
            </select>
          </div>
          <div>
            <div class="accountSearchT">权限 :</div>
            <select class="accountSearchS" id="accountSearchPermisson" v-model="permisson">
              <option value="0">全部</option>
              <option value="2">管理员</option>
              <option value="1">普通用户</option>
            </select>
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
          <div v-show="global.userinfo.permisson == 3">
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
          <div class="accountSearchSearch" @click="clearQuery()">清除</div>
          <div class="accountSearchSearch" @click="querys()">搜索</div>
        </div>
      </div>
      <!-- 设备列表 -->
      <div class="accountList">
        <div class="accountListTitle">
          <!-- 选择框 -->
          <div class="accountList1" style="visibility:hidden">
            <!-- <div class="accountList1Insede"></div> -->
          </div>
          <!-- 序号 -->
          <div class="accountList0">序号</div>
          <!-- 头像 -->
          <div class="accountList9">头像</div>
          <!-- 账号名 -->
          <div class="accountList2">帐户名</div>
          <!-- 权限 -->
          <div class="accountList3">权限</div>
          <!-- VIP -->
          <div class="accountList4">VIP</div>
          <!-- 子账户数 -->
          <div class="accountList5">子账户数</div>
          <!-- 创建时间 -->
          <div class="accountList6">到期时间</div>
          <!-- 状态 -->
          <div class="accountList7">状态</div>
          <!-- 操作 -->
          <div class="accountList8">操作</div>
        </div>
        <div class="accountListTitleD" v-for="(user, index) in users.data" :key="user.ID + 'user'">
          <!-- 选择框 -->
          <div class="accountList1" @click="user.Show=!user.Show;selectAll()">
            <div class="accountList1Insede" v-show="user.Show"></div>
          </div>
          <!-- 序号 -->
          <div class="accountList0">{{index + 1 + offset}}</div>
          <!-- 图片 -->
          <div class="accountList9">
            <div class="accountList9P">
              <img :src="req.localhost + user.Photo" class="imgMax" />
            </div>
            <div class="accountList9PM">
              <img :src="req.localhost + user.Photo" class="imgMax" />
            </div>
          </div>
          <!-- 账号名 -->
          <div class="accountList2">{{user.Name}}</div>
          <!-- 权限 -->
          <div class="accountList3">{{user.Permisson==2 ? '管理员':'普通用户'}}</div>
          <!-- VIP -->
          <div class="accountList4">VIP{{user.VIP}}</div>
          <!-- 子账户数 -->
          <div class="accountList5">
            <div v-show="user.List">{{user.List.all}}</div>
          </div>
          <!-- 创建时间 -->
          <div class="accountList6">
            <div v-show="user.Expiredtime">{{new Date(user.Createtime*1000).toLocaleDateString()}}</div>
            <div v-show="!user.Expiredtime">永久</div>
          </div>
          <!-- 状态 -->
          <div class="accountList7">
            <div v-show="user.Expiredtime==0 || user.Expiredtime > nowTime">
              <div v-show="user.Status==1 ">启用</div>
              <div v-show="user.Status==2 ">禁用</div>
            </div>
            <div v-show="user.Status==3 || (nowTime >= user.Expiredtime && user.Expiredtime!=0)">到期</div>
          </div>
          <!-- 操作 -->
          <div class="accountList8">
            <div class="accountListDisable" v-show="user.Status!=3">
              <div v-show="user.Status==1" @click="disableUsers([user.ID], 2)">禁用</div>
              <div v-show="user.Status==2" @click="disableUsers([user.ID], 1)">启用</div>
              <!-- <div v-show="user.Status==3" @click="expireTime=user.ID">续期</div> -->
            </div>
            <div class="accountListDisable" @click="userDetails = user">详情</div>
            <div class="accountListDisable" @click="expireTimeID=user.ID">延期</div>
            <div class="accountListDisable" @click="delUser([user.ID])">删除</div>
          </div>
        </div>
      </div>
      <!-- 底部/页数 -->
      <div class="accountListTitleB">
        <!-- 选择框 -->
        <div class="accountList1" @click="selectAll(true)">
          <div v-show="users.Show" class="accountList1Insede"></div>
        </div>
        <!-- 禁用 -->
        <div class="accountListDisable" @click="changeAll(1)">禁用</div>
        <!-- 启用 -->
        <div class="accountListDisable" @click="changeAll(2)">启用</div>
        <!-- 删除 -->
        <div class="accountListDisable" @click="changeAll(3)">删除</div>
        <!-- 分页 -->
        <div class="accountListPages">
          <pages :all="allPage" v-on:page="page"></pages>
        </div>
      </div>
      <!-- 详情页 -->
      <div class="accountDetails" v-if="userDetails">
        <accountChild :users="userDetails" v-on:close="userDetails=null"></accountChild>
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
              <div class="accountListExpire4" @click="changeExpire()">确定</div>
              <div class="accountListExpire4" @click="clearTime3()">清除</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import times from "@/components/time.vue";
import pages from "@/components/pages.vue";
import accountChild from "@/components/accountChild.vue";
export default {
  components: { times, pages, accountChild },
  data() {
    return {
      users: {},
      query: "",
      name: "",
      vip: 0,
      permisson: 0,
      status: 0,
      offset: 0,
      allPage: 1,
      userDetails: null,
      expireTimeID: null,
      nowTime: parseInt(new Date().getTime() / 1000),
      ground: 0,
      grounds: []
    };
  },
  methods: {
    getFrounds: function() {
      this.req.get("/users?permisson=2").then(response => {
        if (response.status != 200) {
          return;
        }
        this.grounds = response.data.data;
      });
    },
    changeExpire: function() {
      var now = this.$refs.time3.dayShortText;
      if (now) {
        now = parseInt(new Date(now).getTime() / 1000);
      } else now = 0;
      this.req
        .put("/users/expire", {
          expiredtime: now,
          id: this.expireTimeID
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.getUsers();
          alert(response.data);
        });
      this.expireTimeID = null;
    },
    clearTime3: function() {
      this.$refs.time3.dayShortText = "";
    },
    page: function(pages) {
      var offset = (pages - 1) * 10;
      if (this.offset == offset) {
        return;
      }
      this.offset = offset;
      this.getUsers();
    },
    changeAll: function(methods = 0) {
      var ids = new Array();
      for (const user of this.users.data) {
        if (user.Show) {
          ids.push(user.ID);
        }
      }
      switch (methods) {
        case 1:
          this.disableUsers(ids, 2);
          break;
        case 2:
          this.disableUsers(ids, 1);
          break;
        case 3:
          this.delUser(ids);
          break;
      }
    },
    disableUsers: function(ids, status) {
      this.req
        .put("/users/status", {
          ids: ids,
          status: status
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.getUsers();
          alert(response.data);
        });
    },
    clearQuery: function() {
      this.name = "";
      this.vip = 0;
      this.permisson = 0;
      this.status = 0;
      this.query = "";
      this.$refs.time1.minuteShortText = "";
      this.$refs.time2.minuteShortText = "";
      this.ground = 0;
    },
    querys: function() {
      this.query = "&name=" + this.name;
      this.query += "&vip=" + this.vip;
      this.query += "&permisson=" + this.permisson;
      this.query += "&status=" + this.status;
      this.query += "&ownid=" + this.ground;
      var time1 = new Date(this.$refs.time1.minuteShortText).getTime() / 1000;
      var time2 = new Date(this.$refs.time2.minuteShortText).getTime() / 1000;
      if (time1) {
        this.query += "&starttime=" + time1;
      }
      if (time2) {
        this.query += "&endtime=" + time2;
      }
      this.allPage = 1;
      this.getUsers();
    },
    delUser: function(ids) {
      this.req.del("/users", { id: ids }).then(response => {
        if (response.status != 200) {
          alert(response.msg);
          return;
        }
        this.getUsers();
        alert("Success");
      });
    },
    selectAll: function(methods = false) {
      if (methods) {
        this.users.Show = !this.users.Show;
        methods = this.users.Show;
        for (var user of this.users.data) {
          user.Show = methods;
        }
        this.$forceUpdate();
        return;
      }
      // methods.Show = !methods.Show;
      for (var user of this.users.data) {
        if (!user.Show) {
          this.users.Show = false;
          this.$forceUpdate();
          return;
        }
      }
      this.users.Show = true;
      this.$forceUpdate();
    },
    getusersChildNums: function(user) {
      this.req.get("/users?ownid=" + user.ID).then(response => {
        if (response.status != 200) {
          return;
        }
        user.List = response.data;
        this.$forceUpdate();
        return;
      });
    },
    getUserChildNum: function() {
      for (var user of this.users.data) {
        user.List = {
          all: 0,
          data: []
        };
        if (user.Permisson == 1) {
          continue;
        }
        this.getusersChildNums(user);
      }
    },
    getUsers: function() {
      this.req
        .get("/users?limit=10&offset=" + this.offset + this.query)
        .then(response => {
          if (response.status != 200) {
            return;
          }
          this.users = response.data;
          this.allPage = parseInt(this.users.all / 10);
          if (this.users.all % 10) {
            this.allPage++;
          }
          if (!this.allPage) {
            this.allPage = 1;
          }
          this.getUserChildNum();
        });
    }
  },
  mounted() {
    this.getFrounds();
    this.getUsers();
  }
};
</script>

<style scoped>
@import url("./css/account.css");
</style>
