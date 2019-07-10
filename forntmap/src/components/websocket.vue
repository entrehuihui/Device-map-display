<template>
  <div id="mywebsocket"></div>
</template>
<script>
export default {
  data() {
    return {
      websock: null,
      oldwebsocket: null,
      url: "",
      wsuri: ""
    };
  },
  props: {},
  methods: {
    retdata: function(data) {
      this.$emit("retdata", data);
    },
    initWebpack() {
      // 正式地址;
      // this.wsuri = this.url + this.global.jwt();
      // // 初始化websocket 调试用
      // this.wsuri =
      //   "ws://120.78.76.139:8999/ws?Authorization=" +
      //   this.global.jwt()
      this.wsuri = "ws://127.0.0.1:8800/ws?Authorization=" + this.global.jwt();
      this.oldwebsocket = this.websock;
      this.websock = new WebSocket(this.wsuri); //这里面的this都指向vue
      this.websock.onopen = this.websocketopen;
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onclose = this.websocketclose;
      this.websock.onerror = this.websocketerror;
      if (this.oldwebsocket) {
        this.oldwebsocket.close();
      }
    },
    websocketopen() {
      //打开
      console.log("WebSocket连接成功");
    },
    websocketonmessage(e) {
      //数据接收
      this.retdata(e.data);
    },
    async websocketclose(e) {
      //关闭 --重新连接
      console.log("WebSocket关闭");
      // 如果当前连接断开--重新连接
      if (e.target.url == this.wsuri) {
        // await this.global.sleep(2000);
        // this.initWebpack();
      }
    },
    websocketerror() {}
  },
  watch: {},
  mounted() {
    var a = window.location.href;
    a = a.replace(/https{0,1}/, "ws");
    a = a.split("#")[0] + "ws?Authorization=";
    this.url = a;
    this.initWebpack();
  }
};
</script>