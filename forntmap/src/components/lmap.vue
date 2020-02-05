<template>
  <div id="lmap">
    <div id="top_div">
      <LMap
        :zoom="zoom"
        :maxZoom="24"
        :center="center"
        style="z-index: 0;"
        @click="addMarket"
        @mouseleave="mapMove()"
        :options="{zoomControl: false}"
        :noBlockingAnimations="false"
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
        <!-- 围栏点 -->
        <LMarker
          v-for="(mark, marki) in execMark"
          :key="marki + 'exec'"
          :lat-lng="mark"
          :icon="execicon"
        >
          <LPopup>
            <table>
              <tr>
                <th>纬度:</th>
                <th>{{mark[0]}}</th>
              </tr>
              <tr>
                <th>经度:</th>
                <th>{{mark[1]}}</th>
              </tr>
            </table>
          </LPopup>
        </LMarker>
        <!-- 画点 -->
        <div>
          <div v-for="(user, userindex) in devicesMarks" :key="userindex+'devicesMarks'">
            <div v-if="devicesMarks.Show || userindex =='orbit' || userindex == 'state'">
              <LMarker
                v-for="(device, deviceindex) in user"
                :key="device.Name+deviceindex"
                :lat-lng="[device.DeviceData.Latitude, device.DeviceData.Longitude]"
                :icon="icons[device.DeviceData.State-1]"
                @click="showMark([device.DeviceData.Latitude, device.DeviceData.Longitude])"
              >
                <LTooltip :options="{ permanent: true, interactive: true}">
                  <div>{{device.Name}}</div>
                </LTooltip>
                <LPopup>
                  <tr>
                    <th>EUI:</th>
                    <th>{{device.DevEUI}}</th>
                  </tr>
                  <div
                    v-for="(infokey, infovalue) in  device.DeviceData.Infos"
                    :key="infokey+deviceindex+'explain'"
                  >
                    <tr>
                      <th>{{infokey}}:</th>
                      <th>{{infovalue}}</th>
                    </tr>
                  </div>
                  <tr>
                    <th>{{new Date(device.DeviceData.Uptime*1000).toLocaleString()}}</th>
                  </tr>
                </LPopup>
              </LMarker>
            </div>
          </div>
        </div>

        <!-- 画围栏 -->
        <div v-for="(orbit, orbiti) in orbitLists" :key="orbiti+'mark'">
          <!-- 画圆 -->
          <LCircle
            v-if="orbit.Types==1 && parseInt(orbit.Radius)"
            :lat-lng="[orbit.Lat, orbit.Lng]"
            :radius="parseInt(orbit.Radius)"
            :color="global.color[9]"
            :fillColor="global.color[9]"
            :fillOpacity="0.05"
          ></LCircle>
          <!-- 画矩形 -->
          <LPolygon
            v-if="orbit.Types==2 && orbit.Polygon.length > 2"
            :lat-lngs="orbit.Polygon"
            :color="global.color[9]"
            :fillColor="global.color[9]"
            :fillOpacity="0.05"
          ></LPolygon>
        </div>
        <!-- 画线 轨迹 -->
        <div v-if="devicePloy.OrbitList && devicePloy.OrbitList.length > 0 && devicePloy.Show">
          <LPolyline :lat-lngs="devicePloy.OrbitList" :color="global.color[4]"></LPolyline>
          <!-- 标记起始点 -->
          <LMarker :lat-lng="devicePloy.OrbitList[0]" :icon="icons[9]"></LMarker>
          <!-- 标记终点 -->
          <LMarker :lat-lng="devicePloy.OrbitList[devicePloy.OrbitList.length-1]" :icon="icons[10]"></LMarker>
        </div>
        <!-- 画线 回放 -->
        <div v-if="devicePloyReturn.OrbitList && devicePloyReturn.OrbitList.length > 0 && devicePloyReturn.Show">
          <LPolyline :lat-lngs="devicePloyReturn.OrbitList" :color="global.color[1]"></LPolyline>
        </div>
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
  LPolygon,
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
  components: {
    LMap,
    LTileLayer,
    LMarker,
    LPolyline,
    LPolygon,
    LPopup,
    LTooltip,
    LCircle,
    LControlScale,
    LControlZoom,
    LControlLayers,
    LControl
  },
  methods: {
    showMark: function(v) {
      this.center = v;
      // this.zoom = 18;
    },
    mapMove: function() {
      var lang = this.$refs.lmapmy.mapObject.getCenter();
      this.center = [lang.lat, lang.lng];
    },
    addMarket: function(event) {
      console.log(event.latlng);
      this.$emit("LatLng", event.latlng.lat, event.latlng.lng);
    }
  },
  props: {
    devicePloy: {
      type: Object,
      default: () => {}
    },
    devicePloyReturn: {
      type: Object,
      default: () => {}
    },
    changeZoom: {
      type: Number,
      default: 0
    },
    mapindex: {
      type: Number,
      default: 0
    },
    devicesMarks: {
      type: Object,
      default: () => {}
    }
  },
  data() {
    return {
      rang: 0,
      orbitLists: {},
      execMark: [],
      execicon: icon({
        iconUrl: "/static/execmark.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      zoom: 15,
      center: [22.593262, 113.925971],
      url: "http://{s}.tile.osm.org/{z}/{x}/{y}.png",
      attribution:
        '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      route: false,
      icons: [],
      tileProviders: [
        {
          name: "街道地图",
          visible: true,
          attribution:
            '&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
          url: "http://{s}.tile.osm.org/{z}/{x}/{y}.png"
        },
        {
          name: "离线地图",
          visible: true,
          attribution:
            '&copy; <a target="_blank" href="http://osm.org/copyright">OpenStreetMap</a> contributors',
          url: "/map/{z}/{x}/{y}.jpg"
        }
      ]
    };
  },
  mounted() {
    this.icons = [
      icon({
        iconUrl: "/static/mark1.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark2.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -32]
      }),
      icon({
        iconUrl: "/static/mark3.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark4.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark5.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark6.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark7.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark8.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/mark9.png",
        iconSize: [24, 34],
        shadowSize: [24, 34],
        iconAnchor: [12, 34],
        shadowAnchor: [24, 34],
        popupAnchor: [0, -34]
      }),
      icon({
        iconUrl: "/static/start.png",
        iconSize: [15, 15],
        shadowSize: [15, 15],
        iconAnchor: [8, 8],
        shadowAnchor: [8, 8]
      }),
      icon({
        iconUrl: "/static/end.png",
        iconSize: [20, 20],
        shadowSize: [15, 15],
        iconAnchor: [10, 20],
        shadowAnchor: [8, 8]
      })
    ];
  },
  watch: {
    center: function() {
      // console.log(this.center, "+++");
    },
    changeZoom: function(v, o) {
      if (v > o) {
        this.zoom = this.$refs.lmapmy.mapObject.getZoom() + 1;
      } else {
        this.zoom = this.$refs.lmapmy.mapObject.getZoom() - 1;
      }
    }
  }
};
</script>
<style scoped>
#lmap {
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
