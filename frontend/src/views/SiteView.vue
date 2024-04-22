<template>
  <div class="site">
    <div class="nav" v-for="(item,i) in list">
      <div class="cate">{{item['title']}}</div>
      <div class="site-list">
        <div class="item" v-for="(site,j) in item['list']" @click="openUrl(site['url'],false)">
          <div class="image">
            <el-image style="width: 36px; height: 36px" :src="site['icon']" fit="cover" />
          </div>
          <div class="desc" :title="site['remark']">
            <span class="title">{{site['title']}}</span>
            <span class="remark">{{site['remark']}}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

import {getSiteList} from "@/api/site";
import {BrowserOpenURL} from "../../wailsjs/runtime";

export default {
  name: 'SiteView',
  data(){
    return {
      list:getSiteList()
    }
  },
  created() {
  },
  methods:{
    openUrl(url,self){
      //是否在当前窗口打开
      if(self){
        location.href = url
      }
      BrowserOpenURL(url)
    }
  }
}
</script>

<style lang="scss" scoped>
  .site{
    margin: 10px 0 10px 10px;
    .nav{
      .cate{
        margin: 10px 0;
        font-size: 16px;
        color: #000000;
        font-weight: 600;
      }
      .site-list{
        display: flex;
        flex-flow: wrap;
        .item{
          padding: 10px;
          background: #f5fafd;
          margin: 10px 10px 0 0;
          width: calc((100vw - 255px)/3);
          height: 60px;
          display: flex;
          align-items: center;
          cursor: pointer;
          &:hover{
            box-shadow: 0 5px 12px #e8e8e8
          }
          .desc{
            margin-left: 10px;
            display: flex;
            flex-direction: column;
            font-size: 14px;
            overflow: hidden;
            text-overflow: ellipsis;
            width: 100%;
            .title{
              white-space: nowrap;
              height: 22px;
              color: #242f40;
            }
            .remark{
              height:34px;
              display: -webkit-box;
              -webkit-line-clamp: 2;
              -webkit-box-orient: vertical;
              color: #b2b0c0;
              font-size: 12px;
            }
          }
        }
      }
    }
  }
</style>
