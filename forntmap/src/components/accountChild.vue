<template>
  <div id="accountChild">
    <div class="accountChildO" @click="closed()"></div>
    <div class="accountChildI">
      <!-- 关闭 -->
      <div class="accountChildClose">
        <div class="accountChildCloseI" @click="closed()">X</div>
      </div>
      <!-- 头像 -->
      <div class="accountChildP">
        <div class="accountChildPI" v-show="!photo" @dblclick="trigerUploadPhoto">
          <img :src="req.localhost+ user.Photo" class="imgMax" />
          <div class="accountChildPRemind">双击修改头像</div>
          <input
            v-show="false"
            type="file"
            name="file"
            id="updatefile"
            accept="*"
            @change="changfile($event)"
          />
        </div>
        <div class="accountChildPI" v-show="photo" @dblclick="trigerUploadPhoto">
          <img :src="req.localhost+ photo" class="imgMax" />
          <div class="accountChildPRemind">双击修改头像</div>
        </div>
        <div class="accountChildPPost">
          <div v-show="!photo" class="accountChildPPost" id="updatefileerror" ref="updatefileerror"></div>
          <div v-show="photo">
            <div class="accountChildPPostD" @click="photo = ''">取消</div>
            <div class="accountChildPPostD" @click="putPhoto()">确定</div>
          </div>
        </div>
        <!-- 帐户名 -->
        <div class="accountChildPName">{{user.Name}}</div>
        <!-- 权限 -->
        <div class="accountChildPPermisson">
          <div v-show="user.Permisson == 2">管理员</div>
          <div v-show="user.Permisson == 1">普通用户</div>
        </div>
        <!-- vip -->
        <div class="accountChildPVip">
          <div class="accountChildPVipImg">
            <img src="/static/vip.png" class="imgMax" />
          </div>
          <div class="accountChildPVipN">VIP{{user.VIP}}</div>
        </div>
        <!-- <div class="accountChildPPassword" @click="updatePassword=true">修改密码</div> -->
      </div>
      <!-- 详情 -->
      <div class="accountChildD">
        <div class="accountChildDUpdate" @click="updatePassword= true">修改密码</div>
        <div class="accountChildDUpdate" @click="updateDetails= true">修改详情</div>
        <div>
          <div class="accountChildDD">
            <div class="accountChildDD1">电话 :</div>
            <div class="accountChildDD2">{{user.Mobile}}</div>
            <div class="accountChildDD3" v-show="user.Mobile">{{user.Mobile}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">邮箱 :</div>
            <div class="accountChildDD2">{{user.Email}}</div>
            <div class="accountChildDD3" v-show="user.Email">{{user.Email}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">地址 :</div>
            <div class="accountChildDD2">{{user.Address}}</div>
            <div class="accountChildDD3" v-show="user.Address">{{user.Address}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">备注 :</div>
            <div class="accountChildDD2">{{user.Details}}</div>
            <div class="accountChildDD3" v-show="user.Details">{{user.Details}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">状态 :</div>
            <div class="accountChildDD2">
              <div v-show="user.Status == 1">启用</div>
              <div v-show="user.Status == 2">禁用</div>
              <div v-show="user.Status == 3">到期</div>
            </div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">创建时间 :</div>
            <div class="accountChildDD2">{{new Date(user.Createtime*1000).toLocaleString()}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">到期时间 :</div>
            <div class="accountChildDD2">
              <div v-show="user.Expiredtime">{{new Date(user.Expiredtime*1000).toLocaleString()}}</div>
              <div v-show="!user.Expiredtime">永久</div>
            </div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">子账户数 :</div>
            <div class="accountChildDD2">{{user.List.all}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">设备数 :</div>
            <div class="accountChildDD2">
              <div v-show="user.Devices!=null">{{user.Devices}}</div>
              <div v-show="user.Devices==null">检索中...</div>
            </div>
          </div>
        </div>
      </div>
      <!-- 修改 -->
      <div class="accountChildU" v-show="updateDetails">
        <div>
          <div class="accountChildDD">
            <div class="accountChildDD1">电话 :</div>
            <input
              type="text"
              class="accountChildU2"
              placeholder="不修改不填"
              v-model="mobile"
              ref="mobile"
              @keyup="mobile.length > 11 ? mobile=mobile.substr(0, 11):''"
            />
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">邮箱 :</div>
            <input
              type="text"
              class="accountChildU2"
              placeholder="不修改不填"
              v-model="email"
              ref="email"
            />
            <div class="accountChildDD3" v-show="email">{{email}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">地址 :</div>
            <input
              type="text"
              class="accountChildU2"
              placeholder="不修改不填(最长200)"
              v-model="address"
              ref="address"
            />
          </div>
          <div class="accountChildDD">
            <div class="accountChildDD1">备注 :</div>
            <input
              type="text"
              class="accountChildU2"
              placeholder="不修改不填(最长200)"
              v-model="details"
              ref="details"
            />
            <div class="accountChildDD3" v-show="details">{{details}}</div>
          </div>
          <div class="accountChildDD">
            <div class="accountChildU3" @click="updateInfo()">确定</div>
            <div class="accountChildU3" @click="clearUpdare()">取消</div>
          </div>
        </div>
      </div>
      <!-- 修改密码 -->
      <div class="accountChildPassword" v-show="updatePassword">
        <div class="accountChildPassword1">
          <div class="accountChildPassword2" @click="clearUpdare()">x</div>
          <div>
            <div class="accountChildDD">
              <div class="accountChildDD1">密码 :</div>
              <input type="password" class="accountChildU2" v-model="password" ref="password" />
            </div>
            <div class="accountChildDD">
              <div class="accountChildDD1">重复密码 :</div>
              <input type="password" class="accountChildU2" v-model="password2" ref="password2" />
            </div>
            <div class="accountChildDD">
              <div class="accountChildU3" @click="postPassword()">确定</div>
              <div class="accountChildU3" @click="clearUpdare()">取消</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import check from "@/global/check.js";
import md5 from "js-md5";
export default {
  data() {
    return {
      photo: "",
      mobile: "",
      email: "",
      address: "",
      details: "",
      updateDetails: false,
      user: {},
      password: "",
      password2: "",
      updatePassword: false
    };
  },
  props: {
    users: {
      type: Object,
      default: () => {}
    }
  },
  methods: {
    postPassword: function() {
      if (this.password != this.password2) {
        this.$refs.password2.style = "border: 1px solid red";
        return;
      }
      this.$refs.password2.style = "border: 1px solid rgb(155, 154, 154)";
      this.req
        .put("/users/password", {
          id: this.user.ID,
          password: md5(this.password2)
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.clearUpdare();
          alert(response.data);
        });
    },
    updateInfo: function() {
      if (this.mobile != "" && !check.checkPhone(this.mobile)) {
        this.$refs.mobile.style = "border: 1px solid red";
        return;
      }
      this.$refs.mobile.style = "border: 1px solid rgb(155, 154, 154)";

      if (this.email != "" && !check.checkEmail(this.email)) {
        this.$refs.email.style = "border: 1px solid red";
        return;
      }
      this.$refs.email.style = "border: 1px solid rgb(155, 154, 154)";

      if (this.address != "" && this.address.length > 200) {
        this.$refs.address.style = "border: 1px solid red";
        return;
      }
      this.$refs.address.style = "border: 1px solid rgb(155, 154, 154)";

      if (this.details != "" && this.details.length > 200) {
        this.$refs.details.style = "border: 1px solid red";
        return;
      }
      this.$refs.details.style = "border: 1px solid rgb(155, 154, 154)";
      this.req
        .put("/users/info", {
          address: this.address,
          email: this.email,
          id: user.ID,
          mobile: this.mobile
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          alert(response.data);
          this.updateUser();
        });
    },
    updateUser: function() {
      if (this.mobile != "") {
        this.user.Mobile = this.mobile;
      }
      if (this.email != "") {
        this.user.Email = this.email;
      }
      if (this.address != "") {
        this.user.Address = this.address;
      }
      if (this.details != "") {
        this.user.Details = this.details;
      }
      this.clearUpdare();
    },
    clearUpdare: function() {
      this.mobile = "";
      this.email = "";
      this.address = "";
      this.details = "";
      this.updateDetails = false;
      this.password = "";
      this.password2 = "";
      this.updatePassword = false;
    },
    closed: function() {
      this.clearUpdare();
      this.$emit("close");
    },
    // 更改用户信息
    putPhoto: async function() {
      this.req
        .put("/users/info", {
          id: this.user.ID,
          photo: this.photo
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.photo = "";
          alert("Success");
        });
    },
    //检测图片
    changfile: function() {
      event.preventDefault();
      this.$refs.updatefileerror.innerHTML = "";
      var file = event.target.files;
      if (file.length != 1) {
        return;
      }
      file = file[0];
      if (
        file.type != "image/jpg" &&
        file.type != "image/png" &&
        file.type != "image/jpeg"
      ) {
        this.$refs.updatefileerror.innerHTML = "只能选择PNG/JPG";
      }
      if (file.size > 3 * 1024 * 1024) {
        this.$refs.updatefileerror.innerHTML = "图片尺寸过大(1M)";
      }
      let formData = new FormData();
      formData.append("file", file);
      event.target.value = "";
      this.req.postfile("/upload", formData, 1).then(response => {
        if (response.status != 200) {
          console.log(response);
          this.$refs.updatefileerror.innerHTML = response.msg;
          return;
        }
        this.photo = response.data;
        this.$refs.updatefileerror.innerHTML = "";
        this.$forceUpdate();
      });
    },
    // 触发上传图片
    trigerUploadPhoto: function() {
      document.getElementById("updatefile").click();
    },
    getDevices: function() {
      this.req.get("/devices?limit=1&uid=" + this.user.ID).then(response => {
        if (response.status != 200) {
          this.user.Devices = "未知";
          this.$forceUpdate();
          return;
        }
        this.user.Devices = response.data.all;
        this.$forceUpdate();
      });
    }
  },
  created() {
    this.user = this.users;
  },
  mounted() {
    this.getDevices();
  },
  watch: {}
};
</script>

<style scoped>
@import url("./css/accountChild.css");
</style>
