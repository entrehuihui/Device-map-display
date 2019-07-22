<template>
  <div id="logo">
    <div>
      <div class="logoList">
        <!-- 背景 -->
        <div class="logoList1" v-if="global.permisson ==3">
          <div class="logoList2" @dblclick="trigerUploadPhoto(1)">
            <img :src="req.localhost +logo[1]" class="logoimg" />
            <!-- 图片选择 -->
            <div v-if="select[1]" class="logoerror">
              <div class="logoerror1" @click="postLogo(1)">确定</div>
              <div class="logoerror1" @click="cancel(1)">取消</div>
            </div>
          </div>
          <div class="logoList3">
            <strong>登陆背景</strong>
          </div>
          <div class="logoList3">比例(192:1080)</div>
        </div>
        <!-- 登陆界面 -->
        <div class="logoList1" v-if="global.permisson ==3">
          <div class="logoList2" @dblclick="trigerUploadPhoto(2)">
            <img :src="req.localhost +logo[2]" class="logoimg" />
            <!-- 图片选择 -->
            <div v-if="select[2]" class="logoerror">
              <div class="logoerror1" @click="postLogo(2)">确定</div>
              <div class="logoerror1" @click="cancel(2)">取消</div>
            </div>
          </div>
          <div class="logoList3">
            <strong>登陆界面</strong>
          </div>
          <div class="logoList3">比例(600:400)</div>
        </div>
        <!-- LOGO -->
        <div class="logoList1">
          <div class="logoList2" @dblclick="trigerUploadPhoto(3)">
            <img :src="req.localhost +logo[3]" class="logoimg" />
            <!-- 图片选择 -->
            <div v-if="select[3]" class="logoerror">
              <div class="logoerror1" @click="cancel(3)">取消</div>
              <div class="logoerror1" @click="postLogo(3)">确定</div>
            </div>
          </div>
          <div class="logoList3">
            <strong>LOGO</strong>
          </div>
          <div class="logoList3">比例(300:600)</div>
        </div>
        <!-- 图片选自框 -->
        <input
          v-show="false"
          type="file"
          name="file"
          id="logoupdatefile"
          accept="*"
          @change="changfile($event)"
        />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      logo: {},
      select: {},
      retphoto: ""
    };
  },
  methods: {
    postLogo: async function(photo) {
      var response = await this.req.put("/login/logo", {
        photo: this.logo[photo],
        types: photo
      });
      if (response.status != 200) {
        alert(response.msg);
      } else {
        alert("Success");
      }
      if (photo == 3) {
        this.jump();
      } else {
        this.getLogo();
      }
      this.select = {};
    },
    cancel: function(photo) {
      this.select = {};
      this.getLogo();
      this.$forceUpdate();
    },
    // 获取logo图标
    getLogo: async function() {
      await this.req.get("/login/logo").then(response => {
        if (response.status != 200) {
          return;
        }
        this.logo = response.data;
        this.getLogo3();
        this.$forceUpdate();
      });
    },
    // getLogo
    getLogo3: function() {
      // this.logo
      this.req.get("/login/logo/3").then(response => {
        if (response.status == 200) {
          this.logo[3] = response.data;
          this.$forceUpdate();
        }
      });
    },
    //检测图片
    changfile: function() {
      event.preventDefault();
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
        alert("只能选择PNG/JPG");
        return;
      }
      if (file.size > 3 * 1024 * 1024) {
        alert("图片尺寸过大(1M)");
        return;
      }
      let formData = new FormData();
      formData.append("file", file);
      event.target.value = "";
      this.req
        .postfile("/upload", formData, this.retphoto + 1)
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
            return;
          }
          this.logo[this.retphoto] = response.data;
          this.select[this.retphoto] = true;
          this.$forceUpdate();
        });
    },
    // 触发上传图片
    trigerUploadPhoto: function(photo = 0) {
      this.retphoto = photo;
      document.getElementById("logoupdatefile").click();
    },
    // 跳转
    jump: function(index = 5) {
      this.$router.push({
        name: "jump",
        params: {
          params: {
            index: index
          },
          router: "user"
        }
      });
    }
  },
  mounted() {
    this.getLogo();
  }
};
</script>

<style scoped>
@import url("./css/logo.css");
</style>
