<template>
  <div id="status">
    <div>
      <div class="statusList">
        <div class="statusList1">
          <div class="statusList3">序号</div>
          <div class="statusList3">Types</div>
          <div class="statusList3">图标</div>
          <div class="statusList4">状态说明</div>
          <div class="statusList4">更新时间</div>
          <div class="statusList5">操作</div>
        </div>
        <div class="statusList2" v-for="(State, index) in configState" :key="index+'State'">
          <div class="statusList3">{{index +1}}</div>
          <div class="statusList3">{{State.Types}}</div>
          <div class="statusList3">
            <div class="statusList6">
              <img :src="'/static/mark' + State.Types+ '.png'" class="imgMax" />
            </div>
          </div>
          <div class="statusList4">
            <div v-show="!State.Update">{{State.States}}</div>
            <div v-show="State.Update">
              <input type="text" class="statusList9" v-model="State.States" />
            </div>
          </div>
          <div class="statusList4">
            <div v-show="!State.Updatetime">无更新</div>
            <div v-show="State.Updatetime">{{new Date(State.Updatetime*1000).toLocaleDateString()}}</div>
          </div>
          <div class="statusList5">
            <div v-show="!configState.Update">
              <div class="statusList7" @click="changeUpdate(State, true)">修改</div>
            </div>
            <div v-show="configState.Update">
              <div class="statusList7" @click="getConfigState()">取消</div>
              <div class="statusList7" @click="putConfigState(State)">确定</div>
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
      configState: []
    };
  },
  methods: {
    putConfigState: async function(State) {
      var response = await this.req.put("/configState", {
        id: State.ID,
        states: State.States
      });
      if (response.status != 200) {
        alert(response.msg);
      } else {
        alert(response.data);
      }
      this.getConfigState();
    },
    changeUpdate: function(State, methods = false) {
      State.Update = methods;
      this.configState.Update = methods;
      this.$forceUpdate();
    },
    getConfigState: function() {
      this.req.get("/configState").then(response => {
        if (response.status == 200) {
          this.configState = response.data;
        }
      });
    }
  },
  mounted() {
    this.getConfigState();
  }
};
</script>

<style scoped>
@import url("./css/status.css");
@import url("./css/status.css");
</style>
