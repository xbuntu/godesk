<template>
  <div class="config">
    <el-alert
        title="数据库保存在系统 我的文档-> GoDesk Files -> data.db"
        type="success">
    </el-alert>
    <el-container class="toolbar">
      <el-row>
        <el-button type="success" plain icon="el-icon-plus" round @click="handleAdd">添加配置项</el-button>
        <el-button type="primary" plain icon="el-icon-edit" round @click="handleSave">保存配置</el-button>
        <el-button type="info" plain icon="el-icon-edit" round @click="openDir">打开缓存目录</el-button>
      </el-row>
    </el-container>
    <el-form class="form-wrap" ref="form" label-width="80px">
      <el-form-item label="文件上传">
        <el-upload
            drag
            class="uploader"
            action="http://localhost:10086/upload"
            :show-file-list="false"
            :on-success="handleUpload">
          <el-image v-if="imageUrl" :src="imageUrl" class="avatar"/>
          <template v-else>
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">将文件拖到此处，或<em>点击上传</em></div>
          </template>
        </el-upload>
      </el-form-item>
      <el-form-item :label="item['title']" :key="index" v-for="(item,index) in list">
        <el-input v-model="item['value']">
          <span slot="prepend">{{ item['field'] }}</span>
          <el-button slot="append" icon="el-icon-delete" @click="del(item)"/>
        </el-input>
      </el-form-item>
    </el-form>

    <!-- 添加 -->
    <el-dialog title="添加配置项" :visible.sync="open" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="配置项" prop="title">
          <el-input v-model="form.title" placeholder="请输入配置项，如：软件名称"/>
        </el-form-item>
        <el-form-item label="字段名" prop="field">
          <el-input v-model="form.field" placeholder="请输入字段名，如：name"/>
        </el-form-item>
        <el-form-item label="字段值" prop="value">
          <el-input v-model="form.value" placeholder="请输入字段值，如：GoDesk"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {GetConfig, AddConfig, DelConfig, SaveConfig} from "../../wailsjs/go/controller/Config";
import {OpenConfigDir} from "../../wailsjs/go/controller/System";

export default {
  name: 'ConfigView',
  data() {
    return {
      open: false,
      imageUrl: null,
      form: {
        title: '',
        field: '',
        value: ''
      },
      rules: {
        title: [
          {required: true, message: "配置项不能为空", trigger: "blur"}
        ],
        field: [
          {required: true, message: "字段名不能为空", trigger: "blur"}
        ],
        value: [
          {required: true, message: "字段名不能为空", trigger: "blur"}
        ],
      },
      list: []
    }
  },
  created() {
  },
  mounted() {
    this.get()
  },
  methods: {
    handleUpload(res) {
      if(res.code === 200){
        this.imageUrl = res['data']['url']
      }
    },
    openDir() {
      OpenConfigDir().then(res => {
        console.log(res)
      })
    },
    get() {
      GetConfig().then(res => {
        this.list = res['data']
      })
    },

    //新增按钮操作
    handleAdd() {
      this.open = true;
    },

    //提交按钮
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          AddConfig(this.form).then(res => {
            this.cancel()
            this.get()
            this.$notify.success({
              title: '温馨提示',
              message: '添加成功'
            })
          })
        }
      });
    },
    /** 保存按钮 */
    handleSave() {
      if (this.list == null || this.list.length < 1) {
        this.$notify.error({
          title: '温馨提示',
          message: '请先添加配置项'
        })
        return
      }
      SaveConfig(this.list).then(res => {
        this.get()
        this.$notify.success({
          title: '温馨提示',
          message: '保存成功'
        })
      })
    },

    //删除
    del(config) {
      DelConfig(config).then(res => {
        this.get()
        this.$notify.success({
          title: '温馨提示',
          message: '删除成功'
        })
      })
    },
    // 取消按钮
    cancel() {
      this.open = false;
    },
  }
}
</script>

<style lang="scss" scoped>
.config {
  margin: 10px;

  .toolbar {
    margin-top: 10px;
  }

  .form-wrap {
    margin: 20px 20px 10px 0;
  }
}
</style>
