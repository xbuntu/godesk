<template>
  <div class="photo">
    <div id="viewer"></div>
  </div>
</template>

<script>
import {Viewer} from "photo-sphere-viewer";
import "photo-sphere-viewer/dist/photo-sphere-viewer.css";
import "photo-sphere-viewer/dist/plugins/gallery.css";
import loader from '../assets/loader.gif'
import {GalleryPlugin} from "photo-sphere-viewer/dist/plugins/gallery";
import {GetList} from "../../wailsjs/go/controller/Photo";

export default {
  name: 'PhotoView',
  data() {
    return {}
  },
  created() {

  },
  mounted() {
    this.load()
  },
  methods: {
    load() {
      GetList().then(res => {
        try {
          const viewer = new Viewer({
            container: 'viewer',
            panorama: res[0]['panorama'],
            caption: res[0]['options']['caption'],
            loadingImg: loader,
            touchmoveTwoFingers: true,
            mousewheelCtrlKey: true,
            lang: {
              autorotate:"自动导览",
              zoom: "缩放",
              zoomOut: "缩小",
              zoomIn: "放大",
              move: "移动",
              moveUp: "向上移动",
              moveDown: "向下移动",
              moveLeft: "向左移动",
              moveRight: "向右移动",
              description: "描述",
              download: "下载",
              fullscreen: "全屏",
              loading: "加载中...",
              menu: "菜单",
              gallery: "图集",
              close: "关闭",
              twoFingers: "双指导航",
              ctrlZoom: "ctrl + 滚轮缩放",
              loadError: "图片加载失败"
            },
            plugins: [
              [GalleryPlugin, {
                visibleOnLoad: false,
                items: res
              }],
            ],
          });
        } catch (err) {
          this.$notify.error({
            title: '温馨提示',
            message: err
          })
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.photo {
  padding: 10px;

  #viewer {
    width: calc(100vw - 160px);
    height: calc(100vh - 20px);

    ::v-deep {
      .psv-canvas-container, .psv-canvas {
        width: 100% !important;
        height: 100% !important;
      }
    }
  }
}
</style>
