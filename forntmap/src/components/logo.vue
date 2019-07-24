<template>
  <div id="logo">
    <div>
      <div class="logotitle" v-show="global.userinfo.permisson==3">
        <div class="logotitle1" @click="index=0" :id="index==0?'logotitle1':''">系统图片</div>
        <div class="logotitle1" @click="index=1" :id="index==1?'logotitle1':''">用户logo</div>
      </div>
      <div class="logoList" v-show="index==0" id="logoList">
        <!-- 背景 -->
        <div class="logoList1" v-if="global.userinfo.permisson ==3">
          <div class="logoList2" @dblclick="trigerUploadPhoto(1)">
            <img :src="req.localhost +logo[1]" class="logoimg" />
            <!-- 图片选择 -->
            <div v-if="select[1]" class="logoerror">
              <div class="logoerror1" @click="cancel(1)">取消</div>
              <div class="logoerror1" @click="postLogo(1)">确定</div>
            </div>
          </div>
          <div class="logoList3">
            <strong>登陆背景</strong>
          </div>
          <div class="logoList3">比例(192:1080)</div>
        </div>
        <!-- 登陆界面 -->
        <div class="logoList1" v-if="global.userinfo.permisson ==3">
          <div class="logoList2" @dblclick="trigerUploadPhoto(2)">
            <img :src="req.localhost +logo[2]" class="logoimg" />
            <!-- 图片选择 -->
            <div v-if="select[2]" class="logoerror">
              <div class="logoerror1" @click="cancel(2)">取消</div>
              <div class="logoerror1" @click="postLogo(2)">确定</div>
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
      <div v-show="index==1">
        <div class="logoList8">
          <div class="logoList10">用户:</div>
          <select class="logoList9" v-model="id">
            <option value="0">全部</option>
            <option
              :value="user.ID"
              v-for="(user, index) in userGroud"
              :key="index+'user'"
            >{{user.Name}}</option>
          </select>
          <div class="logoList11" @click="getUserLogo()">确定</div>
          <div class="logoList11" @click="id = 0;getUserLogo()">清除</div>
        </div>
        <div class="logoList">
          <div class="logoList1" v-for="(logos, index) in userLogo.data" :key="index+'logo'">
            <div class="logoList2">
              <img :src="req.localhost +logos.URL" class="logoimg" />
            </div>
            <div class="logoList4">
              <strong>LOGO</strong>
            </div>
            <div class="logoList4">{{logos.Uname}}</div>
            <div class="logoList5">
              <div class="logoList6" @click="delLogo(logos.ID)">删除</div>
            </div>
          </div>
        </div>
        <div class="logoList7">
          <pages :all="allPage" v-on:page="page"></pages>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import pages from "@/components/pages.vue";
export default {
  components: { pages },
  data() {
    return {
      logo: {},
      select: {},
      retphoto: "",
      index: 0,
      userLogo: {},
      offset: 0,
      allPage: 1,
      id: 0,
      userGroud: []
    };
  },
  methods: {
    delLogo: function(id) {
      this.req
        .del("/login/logo", {
          id: id
        })
        .then(response => {
          if (response.status != 200) {
            alert(response.msg);
          } else {
            alert(response.data);
          }
          this.getUserLogo();
        });
    },
    page: function(pages) {
      this.offset = 3 * (pages - 1);
      this.getUserLogo();
    },
    getUserGroud: function() {
      this.req.get("/users?permisson=2").then(response => {
        if (response.status == 200) {
          this.userGroud = response.data.data;
        }
      });
    },
    getUserLogo: function() {
      this.req
        .get(
          "/login/logo/user?limit=3&offset=" + this.offset + "&id=" + this.id
        )
        .then(response => {
          if (response.status == 200) {
            this.userLogo = response.data;
          }
        });
    },
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
        console.log(this.logo);
        this.$forceUpdate();
      });
      this.getLogo3();
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
    if (this.global.userinfo.permisson != 3) {
      return;
    }
    this.getUserLogo();
    this.getUserGroud();
  }
};
</script>

<style scoped>
@import url("./css/logo.css");
</style>
