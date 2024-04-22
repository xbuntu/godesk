<template>
  <div class="audio">
    <a-player
        v-if="show"
        :music="list[0]"
        :list="list"
        theme="#fc3d48"
        autoplay
        showLrc
    />
    <el-button title="点击选择音频文件所在目录" class="open-btn" type="primary" icon="el-icon-folder" @click="openDir" circle></el-button>
  </div>
</template>

<script>
import {GetList,GetFromJson} from "../../wailsjs/go/controller/Audio";
import APlayer  from 'vue-aplayer'
export default {
  name: 'AudioView',
  components: {
    APlayer
  },
  data() {
    return {
      show:false,
      list: []
    }
  },
  created() {
    this.loadJson()
  },
  mounted() {},
  methods: {
    //读取缓存文件
    loadJson(){
      //读取音乐列表缓存文件
      GetFromJson().then(res => {
        if(res == null || res.length < 1){
          //如果没有，提示打开目录选择框
          this.openDir()
          return
        }
        this.buildPlayList(res)
      })
    },
    //打开目录获取音频
    openDir(){
      GetList().then(res => {
        if(res['code'] === 200){
          this.buildPlayList(res['data'])
        }
      })
    },
    //构造APlayer数据
    buildPlayList(list){
      let data = []
      list.forEach((val,index) => {
        data.push({
          title: val['name'],
          artist: val['path'],
          src: val['src'],
          lrc: val['lrc'],
          pic: require('../assets/player.png')
        })
      })
      this.list = data
      this.show = true
    }
  }
}
</script>

<style lang="scss" scoped>
.audio {
  height: 100vh;
  ::v-deep{
    .aplayer {
      height: calc(100vh - 20px);
      margin: 0;
      display: flex;
      flex-direction: column-reverse;
      padding: 10px;
    }
    .aplayer-list-index {
      width: 20px;
      text-align: right;
      flex-shrink: 0;
    }
    .aplayer-list-title {
      flex: none;
      flex-grow: inherit;
      width: 300px;
    }
    .aplayer-info {
      border: 0!important;
    }
    .open-btn{
      position: fixed;
      right: 10px;
      bottom: 50px;
      z-index: 3;
      border-color: transparent;
      background: #e2e2e2;
      i{
        color: #8c8c8c !important;
      }
    }
    .aplayer-list{
      border-bottom: 1px solid #e9e9e9;
      height: calc(100vh - 105px)!important;
      overflow: auto;
    }
    .aplayer-icon-menu ,.aplayer-author{
      display: none!important;
    }
    .aplayer-lrc {
      text-indent: -100px;
      max-height: 30px;
      height: auto;
    }
    .aplayer-pic {
      height: 80px;
      width: 80px;
      border-radius: 80px;
      margin: 5px;
    }
  }
}
</style>
