<template>
  <div id="jump"></div>
</template>
<script>
export default {
  methods: {
    // 跳转
    jump: function(router, params) {
      this.$router.push({
        name: router,
        params: params
      });
    }
  },
  mounted() {
    var params = this.$route.params.params;
    var router = this.$route.params.router;
    this.jump(router, params);
  }
};
</script>
