<template>
  <div id="selectUser">
    <!-- 用户选择 -->
    <div class="deivesUuser" @click="closed()"></div>
    <div class="deivesUuserI">
      <div class="deivesUuserIClose" @click="closed()">X</div>
      <!-- 用户列表 -->
      <div class="deivesUuserIU">
        <div class="deivesUuserIUS">
          <div>
            <input type="text" class="deivesUuserIUSI" placeholder="账户名称" v-model="name" />
            <div class="deivesUuserIUN" @click="querys()">搜索</div>
          </div>
          <div class="deivesUuserIUD">
            <div
              class="deivesUuserIUD1"
              v-for="(user) in users"
              :key="user.ID +'us'"
              @dblclick="seelctID(user.ID)"
            >{{user.Name}}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    allotDevices: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      users: [],
      name: "",
      query: ""
    };
  },
  methods: {
    seelctID: function(id) {
      this.$emit("seelctID", id);
    },
    querys: function() {
      this.query = "&name=" + this.name;
      this.getUser();
    },
    closed: function() {
      this.$emit("close");
    },
    getUser: function() {
      this.users = [
        {
          ID: this.allotDevices.Classid,
          Name: this.allotDevices.Name
        }
      ];
      this.req
        .get("/users?ownid=" + this.allotDevices.Classid + this.query)
        .then(response => {
          if (response.status != 200) {
            return;
          }
          this.users = response.data.data;
          this.users.unshift({
            ID: this.allotDevices.Classid,
            Name: this.allotDevices.Name
          });
          console.log(this.users);
        });
    }
  },
  mounted() {
    this.getUser();
  }
};
</script>

<style>
.deivesUuser {
  position: absolute;
  width: 1000px;
  height: 600px;
  background: rgb(168, 167, 167, 0.5);
  top: 0px;
  left: 0px;
}
.deivesUuserI {
  position: absolute;
  width: 350px;
  height: 400px;
  background: white;
  top: 100px;
  left: 300px;
  border-radius: 10px;
  overflow: hidden;
}
.deivesUuserIClose {
  position: relative;
  height: 30px;
  width: 30px;
  margin-left: 320px;
  line-height: 30px;
  text-align: center;
  font-size: 20px;
  cursor: pointer;
}
.deivesUuserIClose:hover {
  color: red;
}
.deivesUuserIU {
  position: relative;
  width: 250px;
  height: 300px;
  margin-left: 50px;
  margin-top: 30px;
  border: 1px solid rgb(168, 167, 167);
  border-radius: 10px;
  overflow: hidden;
}
.deivesUuserIUS {
  position: relative;
  float: left;
  height: 35px;
  width: 248px;
  border-bottom: 1px solid rgb(168, 167, 167);
}
.deivesUuserIUSI {
  position: relative;
  float: left;
  height: 25px;
  margin-top: 2px;
  margin-left: 10px;
  width: 160px;
  padding-left: 10px;
  padding-right: 10px;
}
.deivesUuserIUN {
  position: relative;
  float: left;
  border: 1px solid rgb(168, 167, 167);
  width: 40px;
  height: 29px;
  line-height: 29px;
  text-align: center;
  margin-top: 2px;
  margin-left: 5px;
  cursor: pointer;
}
.deivesUuserIUN:hover {
  color: rgb(5, 28, 235);
}
.deivesUuserIUD {
  position: relative;
  float: left;
  width: 250px;
  height: 255px;
  overflow: auto;
  margin-top: 5px;
}
.deivesUuserIUD1 {
  position: relative;
  float: left;
  height: 30px;
  width: 250px;
  /* background: #000; */
  margin-top: 5px;
  line-height: 30px;
  text-align: center;
  cursor: pointer;
  user-select: none;
}
.deivesUuserIUD1:hover {
  background: rgb(221, 221, 221);
}
</style>
