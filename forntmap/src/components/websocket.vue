<template>
  <div id="mywebsocket"></div>
</template>
<script>
export default {
  data() {
    return {
      timer: null,
      websock: null,
      oldwebsocket: null,
      url: "",
      wsuri: "",
      close: false
    };
  },
  props: {},
  methods: {
    retdata: function(data) {
      console.log(data);
      if (data == "Heartbeat") {
        return;
      }
      try {
        data = JSON.parse(data);
      } catch (error) {
        console.log(error);
        return;
      }
      this.$emit("retdata", data);
    },
    setTimer() {
      // 测试例子-- 定时器
      if (this.timer == null) {
        this.timer = setInterval(() => {
          console.log("开始定时...每过60秒发送一次心跳");
          this.websock.send("Heartbeat");
        }, 60000);
      }
    },
    initWebpack() {
      // 正式地址;
      this.wsuri = this.url + this.global.jwt();
      // // 初始化websocket 调试用
      // this.wsuri =
      //   "ws://120.78.76.139:8999/ws?Authorization=" +
      //   this.global.jwt()
      // this.wsuri = "ws://127.0.0.1:8800/ws?Authorization=" + this.global.jwt();
      if (this.websock) {
        this.websock.close;
      }
      this.websock = new WebSocket(this.wsuri); //这里面的this都指向vue
      this.websock.onopen = this.websocketopen;
      this.websock.onmessage = this.websocketonmessage;
      this.websock.onclose = this.websocketclose;
      this.websock.onerror = this.websocketerror;
    },
    websocketopen() {
      clearInterval(this.timer);
      this.setTimer();
      //打开
      console.log("WebSocket连接成功");
    },
    websocketonmessage(e) {
      //数据接收
      this.retdata(e.data);
    },
    async websocketclose(e) {
      clearInterval(this.timer);
      //关闭 --重新连接
      console.log("WebSocket关闭");
      // 如果当前连接断开--重新连接
      if (e.target.url == this.wsuri && !this.close) {
        console.log("websocket意外断开, 2秒后重连");
        await this.global.sleep(2000);
        this.initWebpack();
      } else {
        console.log("主动关闭websocket");
        this.close = false;
      }
    },
    async websocketerror(e) {}
  },
  watch: {
    close: function(v) {
      if (v) {
        this.wsuri.close();
      }
    }
  },
  mounted() {
    var a = window.location.href;
    if (this.req.localhost != "") {
      a = this.req.localhost + "/"
    }
    a = a.replace(/https{0,1}/, "ws");
    a = a.split("#")[0] + "ws?Authorization=";
    this.url = a;
    this.initWebpack();
  },
  beforeDestroy() {
    this.websock.close();
  }
};
</script>