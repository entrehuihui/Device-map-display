<template>
  <div id="account">
    <div>
      <div class="accountSearch">
        <div>
          <div class="accountSearchT">账号 :</div>
          <input type="text" class="accountSearchI" />
        </div>
        <div>
          <div class="accountSearchT">VIP :</div>
          <select class="accountSearchS">
            <option value="0">全部</option>
          </select>
        </div>
        <div class="accountSearchSearch">搜索</div>
      </div>
      <div class="accountList">
        <div class="accountListTitle">
          <!-- 选择框 -->
          <div class="accountList1" style="visibility:hidden">
            <!-- <div class="accountList1Insede"></div> -->
          </div>
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
          <div class="accountList6">创建时间</div>
          <!-- 状态 -->
          <div class="accountList7">状态</div>
          <!-- 操作 -->
          <div class="accountList8">操作</div>
        </div>
        <div class="accountListTitleD" v-for="(user) in users.data" :key="user.ID + 'user'">
          <!-- 选择框 -->
          <div class="accountList1" @click="user.Show=!user.Show;selectAll()">
            <div class="accountList1Insede" v-show="user.Show"></div>
          </div>
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
          <div class="accountList3">{{user.Status==2 ? '管理员':'普通用户'}}</div>
          <!-- VIP -->
          <div class="accountList4">VIP{{user.VIP}}</div>
          <!-- 子账户数 -->
          <div class="accountList5">
            <div v-show="user.List">{{user.List.all}}</div>
          </div>
          <!-- 创建时间 -->
          <div class="accountList6">{{new Date(user.Createtime*1000).toLocaleString()}}</div>
          <!-- 状态 -->
          <div class="accountList7">{{user.Status==1 ? '启用':'禁用'}}</div>
          <!-- 操作 -->
          <div class="accountList8">
            <div class="accountListDisable">
              <div v-show="user.Status==1">禁用</div>
              <div v-show="user.Status==2">启用</div>
            </div>
            <div class="accountListDisable">详情</div>
            <div class="accountListDisable">删除</div>
          </div>
        </div>
        <div class="accountListTitleB">
          <!-- 选择框 -->
          <div class="accountList1" @click="selectAll(true)">
            <div v-show="users.Show" class="accountList1Insede"></div>
          </div>
          <!-- 禁用 -->
          <div class="accountListDisable">禁用</div>
          <!-- 启用 -->
          <div class="accountListDisable">启用</div>
          <!-- 删除 -->
          <div class="accountListDisable">删除</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      users: {}
    };
  },
  methods: {
    delUser:function (id) {
      this.req.del("")
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
    getUserChildNum: function() {
      for (var user of this.users.data) {
        user.List = {
          all: 0,
          data: []
        };
        if (user.Permisson == 1) {
          continue;
        }
        this.req.get("/users?ownid=" + user.ID).then(response => {
          if (response.status != 200) {
            return;
          }
          user.List = response.data;
          this.$forceUpdate();
          return;
        });
      }
    },
    getUsers: function() {
      this.req.get("/users?limit=10&").then(response => {
        if (response.status != 200) {
          return;
        }
        this.users = response.data;
        this.getUserChildNum();
      });
    }
  },
  mounted() {
    this.getUsers();
  }
};
</script>

<style scoped>
@import url("./css/account.css");
</style>
