<template>
  <div class="wallpaper">
    <div class="cate-list">
      <div class="item" v-for="(item,index) in cate" @click="setCate(item,index)">
        <el-button size="small" :type="cid === Number(item['id'])?'success':'default'"  plain>{{item['title']}} {{item['total'] > 0 ?`(${item['total']}张)`:''}}</el-button>
      </div>
    </div>
    <div class="list" @scroll="handleScroll">
      <div class="item" v-for="(item,index) in list">
        <div class="toolbar">
          <div class="set" title="设置为系统壁纸" @click="setImg(item['url'])">设置</div>
          <div class="save" title="保存当前壁纸" @click="saveImg(item['url'])">保存</div>
        </div>
        <el-image class="image" title="点击预览" :preview-src-list="[item['url']]" lazy fit="cover" :src="item['url']">
          <div slot="placeholder" class="image-slot">
            加载中<span class="dot">...</span>
          </div>
        </el-image>
        <span class="tag" :title="item['tag']">{{item['tag']}}</span>
      </div>
    </div>
    <el-pagination
        style="margin-top: 20px"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[30,50,70,90,120]"
        :page-size="10"
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :total="total">
    </el-pagination>
  </div>
</template>

<script>
import {SetBackGround,GetCate,GetList,SaveWallpaper}  from "../../wailsjs/go/controller/Wallpaper";
export default {
  name: 'WallpaperView',
  data(){
    return {
      isLoading:false,
      cate:[],
      list:[],
      page:1,
      size:30,
      cid:0,
      cIndex:0,
      total:0,
      maxPage:0
    }
  },
  mounted() {
    this.getCate()
  },
  methods:{
    handleScroll(event){
      const { scrollTop, clientHeight, scrollHeight } = event.target;
      if (scrollTop + clientHeight >= scrollHeight) {
        // 滚动条到达底部，触发懒加载
        this.loadMore();
      }
    },
    // 每页显示数量变化
    handleSizeChange(val) {
      this.page = 1
      this.size = val
      this.getList(false)
    },

    //页码变化
    handleCurrentChange(val) {
      this.page = val
      this.getList(false)
    },
    setImg(url){
      if(url === undefined || url === ''){
        return
      }
      SetBackGround(url);
    },
    saveImg(url){
      if(url === undefined || url === ''){
        return
      }
      SaveWallpaper(url).then(res => {
        if(res['code'] === 200){
          this.$notify.success({
            title: '温馨提示',
            message: '保存成功'
          })
        }
      });
    },
    loadMore(){
      console.log(this.page)
      console.log(this.maxPage)
      if(this.page >= this.maxPage){
        return
      }
      this.page ++
      this.getList(true)
    },
    setCate(cate,index){
      this.cid = Number(cate['id'])
      this.cIndex = index
      this.page = 1
      this.getList(false)
    },
    getCate(){
      GetCate().then(res => {
        if(res['code'] === 200){
          this.cate = res["data"]
          this.cid = Number(this.cate[0]['id'])
          this.getList(false)
        }
        console.log(res)
      })
    },
    getList(loadMore){
      GetList().then(res => {
        if(res['code'] === 200){
          if(this.page === 1 || !loadMore){
            this.list = res["data"]["list"]
          }else{
            this.list.push(...res["data"]["list"])
          }
          this.cate[this.cIndex]['total'] = this.total = Number(res["data"]["total"])
          this.maxPage = Math.ceil(this.total/this.size)
        }
      })
    }
  }
}
</script>

<style lang="scss">
.wallpaper{
  padding-bottom: 10px;
  display: flex;
  flex-direction: column;
  height: calc(100vh - 10px);
  overflow: hidden;
  .cate-list{
    margin: 10px 0 0 10px;
    flex-shrink: 0;
    background: #ffffff;
    display: flex;
    align-items: center;
    flex-flow: wrap;
    .item{
      margin: 0 10px 10px 0;
    }
  }
  .list{
    flex: 1;
    overflow-x: hidden;
    overflow-y: auto;
    display: flex;
    flex-flow: wrap;
    padding: 0 0 20px 0;
    .item{
      width: calc((100vw - 200px)/3);
      margin: 0 0 10px 10px;
      display: flex;
      flex-direction: column;
      position: relative;
      .toolbar{
        position: absolute;
        right: 5px;
        top: 5px;
        z-index: 999;
        display: flex;
        .set,.save{
          cursor: pointer;
          margin-left: 5px;
          font-size: 12px;
          background: rgba(62, 211, 0, 0.6);
          color: #ffffff;
          width: 30px;
          height: 20px;
          text-align: center;
          line-height: 20px;
        }
        .set{
          background: rgba(0, 114, 255, 0.6);
          color: #ffffff;
        }
      }
      .tag{
        margin: 5px 0;
        height: 25px;
        line-height: 25px;
        width: 100%;
        white-space:nowrap;
        text-overflow:ellipsis;
        font-size: 14px;
        overflow: hidden;
      }
      .image{
        width: 100%;
        height: calc((100vw - 200px)/5);
        .image-slot{
          background: #eee;
          line-height: calc((100vw - 200px)/5);
          text-align: center;
        }
      }
    }
  }
}
</style>
