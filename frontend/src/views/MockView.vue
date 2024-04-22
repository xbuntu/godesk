<template>
  <div class="mock">
    <div class="mb-4">
      <el-button type="primary" size="mini" @click="getMockData">模拟mock请求</el-button>
      <el-button type="success" size="mini" @click="getApiData('get')">GET 请求</el-button>
      <el-button type="warning" size="mini" @click="getApiData('post')">POST 请求</el-button>
      <el-button type="danger" size="mini" @click="getGoData">调用Go代码</el-button>
      <el-button type="success" size="mini" :loading="loadingWallpaper" @click="setWallpaper('https://wallpaperswide.com/download/new_windows_11-wallpaper-1920x1080.jpg')">设置壁纸</el-button>
    </div>
    <div class="result">
      <json-viewer
          :value="jsonData"
          :expand-depth="5"
          boxed
          sort
          copyable
      />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import {GetConfig} from "../../wailsjs/go/controller/Config";
import {mockIndex , getApi, postApi} from "@/api/mock";
import {SetBackGround}  from "../../wailsjs/go/controller/Wallpaper";
import JsonViewer from 'vue-json-viewer'
export default {
  name: 'MockView',
  components:{
    JsonViewer
  },
  data(){
    return {
      loadingWallpaper:false,
      jsonData:{}
    }
  },
  created() {
    this.getMockData()
  },
  methods:{
    setWallpaper(url){
      this.loadingWallpaper = true
      SetBackGround(url).then(res => {
        this.loadingWallpaper = false
      })
    },
    getMockData(){
      mockIndex().then(res => {
        this.jsonData = res
      })
    },
    getGoData(){
      GetConfig().then(res => {
        this.jsonData = res
      })
    },
    getApiData(method){
      let res = method === 'get'?getApi():postApi();
      res.then(res => {
        this.jsonData = res
      }).finally(()=>{})
    }
  }
}
</script>

<style lang="scss" scoped>
  .mock{
    padding: 20px;
    .result{
      margin-top: 20px;
    }
  }
</style>
