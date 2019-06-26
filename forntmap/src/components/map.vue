<template>
  <div class="lmap">
    <div id="top_div">
      <LMap
        :zoom="zoom"
        :maxZoom="18"
        :center="center"
        style="z-index: 0;"
        @click="addMarket"
        :options="{zoomControl: false}"
        ref="lmapmy"
      >
        <!-- 默认地图 如果不用地图选择器则不注释 -->
        <LTileLayer
          :key="tileProviders[mapindex].name"
          :name="tileProviders[mapindex].name"
          :visible="tileProviders[mapindex].visible"
          :url="tileProviders[mapindex].url"
          :attribution="tileProviders[mapindex].attribution"
          layer-type="base"
        ></LTileLayer>
        <!-- 画圆 -->
        <LCircle
          v-for="(v, i) in circle"
          :key="'circle'+i"
          :lat-lng="v.center"
          :radius="parseInt(v.radius)"
          :color="'red'"
        ></LCircle>
        <!-- 画标记 -->
        <LMarker v-for="(v, i) in markers" :key="'mark'+i" :lat-lng="v" :icon="eicon">
          <LTooltip :options="{ permanent: true, interactive: true}">
            <div>{{v.details["名称"]}}</div>
          </LTooltip>
          <LPopup>
            <div v-for="(v, j) in v.details" :key="i+j+'2'">
              <table>
                <tr>
                  <th>{{j}}</th>
                  <th>{{v}}</th>
                </tr>
              </table>
            </div>
          </LPopup>
        </LMarker>
        <!-- 画轨迹 -->
        <LPolyline
          v-for="(v, i) in polyline"
          :key="i+'guiji'"
          :lat-lngs="v"
          :color="global.color[i%9]"
        ></LPolyline>
        <!-- 画轨迹标记 -->
        <LMarker v-for="(v, i) in polyline" :key="'mark1'+i" :lat-lng="v[0]" :icon="eicon">
          <LTooltip :options="{ permanent: true, interactive: true}">
            <div>{{v[0].details["名称"]}}</div>
          </LTooltip>
          <LPopup>
            <div v-for="(v, j) in v[0].details" :key="i+j+'4'">
              <table>
                <tr>
                  <th>{{j}}</th>
                  <th>{{v}}</th>
                </tr>
              </table>
            </div>
          </LPopup>
        </LMarker>
        <!-- 画路线 -->
        <LPolyline
          v-for="(v, i) in routline"
          :key="i+'lx'"
          :lat-lngs="v"
          :color="global.color[i%9]"
        ></LPolyline>
        <!-- 画路线标记 起点-->
        <LMarker v-for="(v, i) in routline" :key="'mark2'+i" :lat-lng="v[0]" :icon="sicon">
          <LPopup>
            <table>
              <tr>
                <th>路线名:</th>
                <th>{{v.name}}--起点</th>
              </tr>
              <tr>
                <th>经度:</th>
                <th>{{v[0][1]}}</th>
              </tr>
              <tr>
                <th>纬度:</th>
                <th>{{v[0][0]}}</th>
              </tr>
            </table>
          </LPopup>
        </LMarker>
        <!-- 画路线标记 终点-->
        <LMarker v-for="(v, i) in routline" :key="'mark3'+i" :lat-lng="v[v.length-1]" :icon="eicon">
          <LPopup>
            <table>
              <tr>
                <th>路线名:</th>
                <th>{{v.name}}--终点</th>
              </tr>
              <tr>
                <th>经度:</th>
                <th>{{v[0][1]}}</th>
              </tr>
              <tr>
                <th>纬度:</th>
                <th>{{v[0][0]}}</th>
              </tr>
            </table>
          </LPopup>
        </LMarker>
        <!-- 比例尺 -->
        <LControlScale position="bottomright" :imperial="true" :metric="true"></LControlScale>
        <!-- 地图按钮 给用户图标挪位置 -->
        <!-- <LControlLayers position="topright"></LControlLayers> -->
        <!-- <LControl class="example-custom-control">
        </LControl>-->
        <!-- 地图选择按钮 -->
        <!-- <LControlLayers position="topright"></LControlLayers> -->
        <!-- 放大按钮 -->
        <!-- <LControlZoom></LControlZoom> -->
        <!-- <LTileLayer
          v-for="tileProvider in tileProviders"
          :key="tileProvider.name"
          :name="tileProvider.name"
          :visible="tileProvider.visible"
          :url="tileProvider.url"
          :attribution="tileProvider.attribution"
          layer-type="base"
        />-->
      </LMap>
    </div>
  </div>
</template>

<script>
import Vue2Leaflet from "vue2-leaflet";
import "leaflet/dist/leaflet.css";
import { latLng, icon } from "leaflet";
import {
  LMap,
  LTileLayer,
  LMarker,
  LPolyline,
  LPopup,
  LTooltip,
  LCircle,
  LControlScale,
  LControlZoom,
  LControlLayers,
  LControl
} from "vue2-leaflet";
/* leaflet icon */
delete L.Icon.Default.prototype._getIconUrl;
L.Icon.Default.mergeOptions({
  iconRetinaUrl: require("leaflet/dist/images/marker-icon-2x.png"),
  iconUrl: require("leaflet/dist/images/marker-icon.png"),
  shadowUrl: require("leaflet/dist/images/marker-shadow.png")
});
export default {
  name: "example",
  components: {
    LMap,
    LTileLayer,
    LMarker,
    LPolyline,
    LPopup,
    LTooltip,
    LCircle,
    LControlScale,
    LControlZoom,
    LControlLayers,
    LControl
  },
  props: {
    polyline: {
      default: ""
    },
    markers: {
      default: ""
    },
    changeZoom: {
      default: 0
    },
    mapindex: {
      default: 0
    }
  },
  data() {
    return {
      // polyline: [],
      routline: [],
      // markers: [],
      circle: [],
      zoom: 15,
      zoomp: 15,
      zoomm: 15,
      center: [22.593262, 113.925971],
      centerp: [],
      centerm: [],
      url: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",
      attribution:
        '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      route: false,
      sicon: icon({
        iconUrl: "static/image/start.png",
        iconSize: [32, 37],
        iconAnchor: [16, 37]
      }),
      eicon: icon({
        iconUrl: "static/image/end.png",
        iconSize: [32, 37],
        iconAnchor: [16, 37]
      }),
      tileProviders: [
        {
          name: "谷歌街道图",
          visible: true,
          attribution:
            '&copy; <a target="_blank" href="http://www.google.cn/maps">谷歌地图</a> contributors',
          url: "http://www.google.cn/maps/vt?lyrs=m@189&gl=cn&x={x}&y={y}&z={z}"
        },
        {
          name: "谷歌影像图",
          visible: true,
          attribution:
            '&copy; <a target="_blank" href="http://www.google.cn/maps">谷歌地图</a> contributors',
          url: "http://www.google.cn/maps/vt?lyrs=s@189&gl=cn&x={x}&y={y}&z={z}"
        },
        {
          name: "离线地图",
          visible: true,
          attribution:
            '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors',
          url: "http://120.78.76.139:8998/1818940751/{z}/{x}/{y}"
        }
      ]
    };
  },
  methods: {
    addorbits: function(v) {
      this.polyline = v;
      this.$forceUpdate();
    },
    addmark: function(mark) {
      this.markers = mark;
    },
    addMarket: function(event) {
      console.log(event.latlng);
    },
    setTimer() {
      // 测试例子-- 定时器
      if (this.timer == null) {
        this.timer = setInterval(() => {
          console.log("开始定时...每过一秒执行一次");
          // this.marker.push([this.l, (this.g += 0.002)]);
        }, 5000);
      }
    },
    updatecenter: function() {
      if (this.centerp.length == 0 && this.centerm.length == 0) {
        this.center = [22.593262, 113.925971];
        return;
      }
      if (this.centerp.length == 0) {
        this.center = this.centerm;
        return;
      }
      if (this.centerm.length == 0) {
        this.center = this.centerp;
        return;
      }
      var l = (this.centerp[0] + this.centerm[0]) / 2;
      var m = (this.centerp[1] + this.centerm[1]) / 2;
      this.center = [l, m];
    }
  },
  mounted() {},
  watch: {
    mapindex: function(v) {
      console.log(v, "=====");
    },
    changeZoom: function(v, o) {
      // console.log(this.$refs.lmapmy.mapObject.getCenter());
      // console.log(this.$refs.lmapmy.mapObject.getZoom(), "+++++22");
      // console.log(this.$refs.lmapmy.mapObject.getBounds());
      if (v > o) {
        this.zoom = this.$refs.lmapmy.mapObject.getZoom() + 1;
      } else {
        this.zoom = this.$refs.lmapmy.mapObject.getZoom() - 1;
      }
    },
    routline: function(v) {
      var maxx = 0;
      var maxy = 0;
      var minx = 0;
      var miny = 0;
      var m = true;
      for (const k in v) {
        if (k == 0) {
          return;
        }
        for (const z of v[k]) {
          if (m) {
            maxy = z[1];
            maxx = z[0];
            miny = z[1];
            minx = z[0];
            m = false;
            continue;
          }
          if (z[0] > maxx) {
            maxx = z[0];
          } else if (z[0] < minx) {
            minx = z[0];
          }
          if (z[1] > maxy) {
            maxy = z[1];
          } else if (z[1] < miny) {
            miny = z[1];
          }
        }
      }
      if (m) {
        this.updatecenter();
        this.$forceUpdate();
        return;
      }
      var middlex = (maxx + minx) / 2;
      var middley = (maxy + miny) / 2;
      // this.center = new Array();
      this.center = [middlex, middley];
      var it = true;
      var ly = (maxy - miny).toFixed(4);
      //计算Y方向放大系数
      var y = 90;
      var zy = 1;
      if (ly != 0) {
        while (it) {
          if (ly > y) {
            break;
          }
          if (y < 0.02) {
            zy = 15;
            break;
          }
          y = y / 2;
          zy++;
        }
      } else {
        zy = 17;
      }

      //技术X方向放大系数
      var x = 180;
      var zx = 1;
      var lx = (maxx - minx).toFixed(4);
      if (lx != 0) {
        while (it) {
          if (lx > x) {
            break;
          }
          if (lx < 0.009) {
            zy = 15;
            break;
          }
          x = x / 2;
          zx++;
          if (zx >= zy) {
            break;
          }
        }
      } else {
        zx = 17;
      }
      var zoom = zy > zx ? zx : zy;
      this.zoomm = this.zoomm < zoom ? this.zoomm : zoom;
      this.zoom = this.zoomm < this.zoomp ? this.zoomm : this.zoomp;
      // this.updatecenter();
      this.$forceUpdate();
    },
    circle: function(v) {
      if (v.length == 0) {
        return;
      }
      var y = 90;
      var zy = 2;
      var ly = v[0].radius / 111000;
      while (true) {
        if (ly > y) {
          break;
        }
        if (y < 0.002) {
          zy = 15;
          break;
        }
        y = y / 2;
        zy++;
      }
      this.center = new Array();
      this.center = v[0].center;
      this.zoom = zy;
      this.$forceUpdate();
    },
    polyline: function(v) {
      this.zoomp = 15;
      var maxx = 0;
      var maxy = 0;
      var minx = 0;
      var miny = 0;
      if (v.length < 1) {
        this.centerp = [];
        this.updatecenter();
        return;
      }
      if (v[0].length < 1) {
        this.centerp = [];
        this.updatecenter();
        return;
      }
      maxx = v[0][0][0];
      maxy = v[0][0][1];
      minx = v[0][0][0];
      miny = v[0][0][1];
      // 求最大最小坐标点
      for (const k of v) {
        for (const z of k) {
          if (z[0] > maxx) {
            maxx = z[0];
          } else if (z[0] < minx) {
            minx = z[0];
          }
          if (z[1] > maxy) {
            maxy = z[1];
          } else if (z[1] < miny) {
            miny = z[1];
          }
        }
      }
      var middlex = (maxx + minx) / 2;
      var middley = (maxy + miny) / 2;

      // this.center = new Array();
      // this.centerp = [middlex, middley];
      this.center = [middlex, middley];
      var it = true;
      var ly = (maxy - miny).toFixed(4);
      //计算Y方向放大系数
      var y = 90;
      var zy = 1;
      if (ly != 0) {
        while (it) {
          if (ly > y) {
            break;
          }
          if (y < 0.002) {
            zy = 15;
            break;
          }
          y = y / 2;
          zy++;
        }
      } else {
        zy = 17;
      }

      //技术X方向放大系数
      var x = 180;
      var zx = 1;
      var lx = (maxx - minx).toFixed(4);
      if (lx != 0) {
        while (it) {
          if (lx > x) {
            break;
          }
          if (lx < 0.0009) {
            zy = 15;
            break;
          }
          x = x / 2;
          zx++;
          if (zx >= zy) {
            break;
          }
        }
      } else {
        zx = 17;
      }
      // this.zoom = zy > zx ? zx : zy;
      var zoom = zy > zx ? zx : zy;
      this.zoomp = this.zoomp < zoom ? this.zoomp : zoom;
      this.zoom = this.zoomm < this.zoomp ? this.zoomm : this.zoomp;
      // this.updatecenter();
      this.$forceUpdate();
    },
    markers: function(v) {
      this.zoomm = 15;
      var maxx = 0;
      var maxy = 0;
      var minx = 0;
      var miny = 0;
      if (v.length < 1) {
        this.centerm = [];
        this.updatecenter();
        return;
      }
      if (v[0].length < 1) {
        return;
      }
      maxx = v[0][0];
      maxy = v[0][1];
      minx = v[0][0];
      miny = v[0][1];
      // 求最大最小坐标点
      for (const z of v) {
        if (z[0] > maxx) {
          maxx = z[0];
        } else if (z[0] < minx) {
          minx = z[0];
        }
        if (z[1] > maxy) {
          maxy = z[1];
        } else if (z[1] < miny) {
          miny = z[1];
        }
      }
      var middlex = (maxx + minx) / 2;
      var middley = (maxy + miny) / 2;
      // this.center = new Array();
      // this.centerm = [middlex, middley];
      this.center = [middlex, middley];
      var it = true;
      var ly = (maxy - miny).toFixed(4);
      //计算Y方向放大系数
      var y = 90;
      var zy = 1;
      if (ly != 0) {
        while (it) {
          if (ly > y) {
            break;
          }
          if (y < 0.02) {
            zy = 15;
            break;
          }
          y = y / 2;
          zy++;
        }
      } else {
        zy = 17;
      }

      //技术X方向放大系数
      var x = 180;
      var zx = 1;
      var lx = (maxx - minx).toFixed(4);
      if (lx != 0) {
        while (it) {
          if (lx > x) {
            break;
          }
          if (lx < 0.009) {
            zy = 15;
            break;
          }
          x = x / 2;
          zx++;
          if (zx >= zy) {
            break;
          }
        }
      } else {
        zx = 17;
      }
      var zoom = zy > zx ? zx : zy;
      this.zoomm = this.zoomm < zoom ? this.zoomm : zoom;
      this.zoom = this.zoomm < this.zoomp ? this.zoomm : this.zoomp;
      // this.updatecenter();
      this.$forceUpdate();
    }
  }
};
</script>

<style>
.lmap {
  width: 100%;
  height: 100%;
}

#top_div {
  overflow-x: auto;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  height: 100%;
  width: 100%;
}
.example-custom-control {
  background: #fff;
  padding: 0 0.5em;
  border: 1px solid #aaa;
  border-radius: 0.1em;
  /* width: 1px; */
  height: 40px;
  border-radius: 10px;
}
.leaflet-fake-icon-image-2x {
  background-image: url(../../node_modules/leaflet/dist/images/marker-icon-2x.png);
}
.leaflet-fake-icon-shadow {
  background-image: url(../../node_modules/leaflet/dist/images/marker-shadow.png);
}
</style>
