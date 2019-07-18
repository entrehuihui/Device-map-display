<template>
  <div id="pages">
    <!-- 显示总条数和页数 -->
    <div class="pagesAllNum">{{start+index}}/{{all}}</div>
    <!-- 最后一页-->
    <div class="pagesAll" @click="selectEnd()">
      <div class="pagesimg2">
        <img src="/static/right2.png" class="imgMax" />
      </div>
    </div>
    <!-- 向后一页 -->
    <div class="pagesAll" @click="selectIndex(index+1)">
      <div class="pagesimg">
        <img src="/static/right.png" class="imgMax" />
      </div>
    </div>
    <!-- 选择页 -->
    <div class="pagesAllOut">
      <div
        class="pagesAll pagesAllOuts"
        v-for="item in maxPage"
        :key="item +'page'"
        @click="selectIndex(item)"
      >
        <div class="pagesAllInside" v-show="item == index">{{item+start}}</div>
        <div v-show="item != index">{{item+start}}</div>
      </div>
    </div>
    <!-- 向前一页 -->
    <div class="pagesAll" @click="selectIndex(index-1)">
      <div class="pagesimg">
        <img src="/static/left.png" class="imgMax" />
      </div>
    </div>
    <!-- 第一页 -->
    <div class="pagesAll" @click="selectStart()">
      <div class="pagesimg2">
        <img src="/static/left2.png" class="imgMax" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      maxPage: 1,
      index: 1,
      start: 0
    };
  },
  props: {
    all: {
      default: 1
    }
  },
  methods: {
    selectEnd: function() {
      if (this.all < 6) {
        this.index = this.all;
      } else {
        this.index = 5;
        this.start = this.all - 5;
      }
      this.selectIndex(this.index);
    },
    selectStart: function() {
      this.start = 0;
      this.index = 1;
      this.selectIndex(this.index);
    },
    selectIndex: function(index) {
      if (index > 0 && index <= this.maxPage) {
        this.countIndex(index);
      }
      this.$emit("page", this.index + this.start);
    },
    countIndex: function(index) {
      if (index == 3 || this.index == index) {
        this.index = index;
        return;
      }
      if (this.start + index < 3) {
        this.index = index;
        return;
      }
      if (this.all < 6) {
        this.index = index;
        return;
      }
      if (this.start + index + 2 <= this.all) {
        this.index = 3;
        this.start += index - 3;
        return;
      }
      this.index = index;
    }
  },
  watch: {
    all: function(v) {
      if (v > 5) {
        this.maxPage = 5;
      } else {
        this.maxPage = v;
      }
      this.start = 0;
      this.index = 1;
    }
  },
  mounted() {
    if (this.all > 5) {
      this.maxPage = 5;
    } else {
      this.maxPage = this.all;
    }
  }
};
</script>

<style scoped>
@import url("./css/pages.css");
</style>
