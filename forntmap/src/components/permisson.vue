<template>
  <div id="permisson">
    <div>
      <div class="permissonList">
        <div class="permissonList1">
          <div class="permissonList3">VIP</div>
          <div class="permissonList3">子用户数</div>
          <div class="permissonList3">设备数</div>
          <div class="permissonList3">轨迹</div>
          <div class="permissonList3">围栏</div>
          <div class="permissonList3">围栏报警</div>
          <div class="permissonList3">实时数据</div>
          <div class="permissonList3">LOGO</div>
          <div class="permissonList3">状态</div>
          <div class="permissonList4">操作</div>
        </div>
        <div class="permissonList2" v-for="(vip) in vips" :key="vip.ID +'vip'">
          <div class="permissonList3">VIP{{vip.ID}}</div>
          <div class="permissonList3">
            <div v-show="!vip.Update">{{vip.Users}}</div>
            <div v-show="vip.Update">
              <input type="number" class="permissonList8" min="1" max="1000" v-model="vip.Users" />
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">{{vip.Devices}}</div>
            <div v-show="vip.Update">
              <input type="number" class="permissonList8" min="1" max="1000" v-model="vip.Devices" />
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.Orbit==1" class="permissonList6">允许</div>
              <div v-show="vip.Orbit==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.Orbit">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.Fence==1" class="permissonList6">允许</div>
              <div v-show="vip.Fence==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.Fence">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.FenceAlarm==1" class="permissonList6">允许</div>
              <div v-show="vip.FenceAlarm==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.FenceAlarm">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.Real==1" class="permissonList6">允许</div>
              <div v-show="vip.Real==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.Real">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.Logo==1" class="permissonList6">允许</div>
              <div v-show="vip.Logo==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.Logo">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList3">
            <div v-show="!vip.Update">
              <div v-show="vip.State==1" class="permissonList6">允许</div>
              <div v-show="vip.State==2" class="permissonList5">禁止</div>
            </div>
            <div v-show="vip.Update">
              <select class="permissonList9" v-model="vip.State">
                <option value="1">允许</option>
                <option value="2">禁止</option>
              </select>
            </div>
          </div>
          <div class="permissonList4">
            <div v-show="!vip.Update &&!vips.Update">
              <div class="permissonList7" @click="changeShow(vip, true)">修改</div>
            </div>
            <div v-show="vip.Update">
              <div class="permissonList7" @click="changeShow(vip, false)">取消</div>
              <div class="permissonList7" @click="updateVip(vip)">确定</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      vips: []
    };
  },
  methods: {
    updateVip: async function(vip) {
      var response = await this.req.put("/vip", {
        ID: parseInt(vip.ID),
        Users: parseInt(vip.Users),
        Devices: parseInt(vip.Devices),
        Orbit: parseInt(vip.Orbit),
        Fence: parseInt(vip.Fence),
        FenceAlarm: parseInt(vip.FenceAlarm),
        Real: parseInt(vip.Real),
        Logo: parseInt(vip.Logo),
        State: parseInt(vip.State)
      });
      if (response.status == 200) {
        alert(response.data);
      } else {
        alert(response.msg);
      }
      this.changeShow(vip, false);
    },
    changeShow: function(vip, methods = false) {
      vip.Update = methods;
      this.vips.Update = methods;
      if (!methods) {
        this.getVip();
      }
      this.$forceUpdate();
    },
    getVip: function() {
      this.req.get("/vip").then(response => {
        if (response.status == 200) {
          this.vips = response.data;
          console.log(this.vips);
        }
      });
    }
  },
  mounted() {
    this.getVip();
  }
};
</script>

<style scoped>
@import url("./css/permisson.css");
</style>
