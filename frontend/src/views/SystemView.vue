<template>
  <div class="system">
    <div class="nav">
      <div class="cate">系统 API</div>
      <div class="site-list">
        <div class="item" v-for="(cmd,index) in cmdList" @click="runCmd(cmd.cmd,cmd.param)">
          <div class="image">
            <el-image style="width: 36px; height: 36px" :src="cmd.icon" fit="cover"/>
          </div>
          <div class="desc" :title="cmd.desc">
            <span class="title">{{ cmd.name }}</span>
            <span class="remark">{{ cmd.desc }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {ShellCMD} from "../../wailsjs/go/controller/System";
import {getCMDList} from "@/api/system";

export default {
  name: 'SystemView',
  data() {
    return {
      cmdList: getCMDList()
    }
  },
  created() {
  },
  methods: {
    //运行cmd命令
    runCmd(cmd, param) {
      ShellCMD(cmd, param)
    }
  }
}
</script>

<style lang="scss" scoped>
.system {
  margin: 10px 0 10px 10px;

  .nav {
    .cate {
      margin: 10px 0;
      font-size: 16px;
      color: #000000;
      font-weight: 600;
    }

    .site-list {
      display: flex;
      flex-flow: wrap;

      .item {
        padding: 10px;
        background: #f5fafd;
        margin: 10px 10px 0 0;
        width: calc((100vw - 255px)/3);
        height: 60px;
        display: flex;
        align-items: center;
        cursor: pointer;

        &:hover {
          box-shadow: 0 5px 12px #e8e8e8
        }

        .desc {
          margin-left: 10px;
          display: flex;
          flex-direction: column;
          font-size: 14px;
          overflow: hidden;
          text-overflow: ellipsis;
          width: 140px;

          .title {
            white-space: nowrap;
            height: 22px;
            color: #242f40;
          }

          .remark {
            height: 34px;
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
