<template>
  <div id="userInfo">
    <!-- 头像 -->
    <div class="userInfoPhoto">
      <div class="userInfoPhotoInside" @dblclick="trigerUploadPhoto">
        <img v-show="!retphoto" :src="req.localhost + global.userinfo.photo" class="imgMax" />
        <img v-show="retphoto" :src="req.localhost + retphoto" class="imgMax" />
        <div class="userInfoPhotoPrompt">双击更改图片</div>
        <input
          v-show="false"
          type="file"
          name="file"
          id="updatefile"
          accept="*"
          @change="changfile($event)"
        />
      </div>
      <div v-show="!retphoto" class="updatefileerror" ref="updatefileerror"></div>
      <div v-show="retphoto" class="updatefilebutton">
        <div class="updatefileerrorSelect" @click="putPhoto()">确定</div>
        <div class="updatefileerrorSelect" @click="retphoto = ''">取消</div>
      </div>
      <div class="userInfoPhotoName">{{global.userinfo.name}}</div>
      <div class="userInfoPhotoPermisson">
        <div v-show="global.userinfo.permisson == 1">普通用户</div>
        <div v-show="global.userinfo.permisson == 2">管理员用户</div>
        <div v-show="global.userinfo.permisson == 3">超级管理员</div>
      </div>
      <div class="userInfoPhotoPermisson">
        <div class="userInfoPhotoVip">
          <img src="/static/vip.png" class="imgMax" />
        </div>
        <div class="userInfoPhotoVip">VIP{{global.userinfo.vip}}</div>
      </div>
    </div>
    <!-- 信息 -->
    <div class="userInfoDetails">
      <div class="userInfoDetailsTitle">修改</div>
      <div class="userInfoDetailsBody"></div>
    </div>
  </div>
</template>

<script>
export default {
  props: {},
  data() {
    return {
      retphoto: ""
    };
  },
  methods: {
    // 更改用户信息
    putPhoto: async function() {
      this.req
        .put("/users/info", {
          photo: this.retphoto
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.getUserinfo();
          this.retphoto = "";
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
      this.req.postfile("/upload", formData).then(response => {
        if (response.status != 200) {
          this.$refs.updatefileerror.innerHTML = response.msg;
          return;
        }
        this.retphoto = response.data;
        this.$refs.updatefileerror.innerHTML = "";
        this.$forceUpdate();
      });
    },
    // 触发上传图片
    trigerUploadPhoto: function() {
      document.getElementById("updatefile").click();
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
    }
  },
  async mounted() {
    await this.getUserinfo();
  }
};
</script>

<style scoped>
@import url("./css/user.css");
</style>
