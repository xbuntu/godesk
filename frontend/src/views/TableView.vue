<template>
  <div class="table">
    <el-form class="query-form" :model="queryParams" ref="queryForm" size="small" :inline="true" label-width="68px">
      <el-form-item label="姓名" prop="postName">
        <el-input
            v-model="queryParams.postName"
            placeholder="请输入姓名"
            clearable
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
      </el-form-item>
      <el-form-item>
        <el-button
            type="primary"
            plain
            icon="el-icon-plus"
            size="mini"
            @click="handleAdd">新增
        </el-button>
      </el-form-item>
      <el-form-item>
        <el-button
            type="warning"
            plain
            icon="el-icon-download"
            size="mini"
            @click="handleExport">导出
        </el-button>
      </el-form-item>
    </el-form>

    <el-table :height="tableHeight" v-loading="loading" :data="list" @selection-change="handleSelectionChange">
      <el-table-column
          prop="date"
          label="日期"
          width="150">
      </el-table-column>
      <el-table-column
          prop="name"
          label="姓名"
          width="120">
      </el-table-column>
      <el-table-column
          prop="province"
          label="省份"
          width="120">
      </el-table-column>
      <el-table-column
          prop="city"
          label="市区"
          width="120">
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          width="100">
        <template slot-scope="scope">
          <el-link @click="handleUpdate(scope.row)" type="success" :underline="false">编辑</el-link>
          <el-link @click="handleDelete(scope.row)" type="danger" :underline="false">删除</el-link>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
        style="margin-top: 20px"
        background
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="1"
        :page-sizes="[50, 100]"
        :page-size="10"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total">
    </el-pagination>

    <!-- 添加或修改用户对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="500px" append-to-body>
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
        <el-form-item label="姓名" prop="postName">
          <el-input v-model="form.postName" placeholder="请输入姓名"/>
        </el-form-item>
        <el-form-item label="备注" prop="remark">
          <el-input v-model="form.remark" type="textarea" placeholder="请输入内容"/>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" size="mini" @click="submitForm">确 定</el-button>
        <el-button @click="cancel" size="mini">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'TableView',
  data() {
    return {
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 50,
        postName: undefined
      },
      // 选中数组
      ids: [],
      total: 0,
      // 遮罩层
      loading: true,
      // 用户表格数据
      list: [],
      tableHeight:null,
      title: "添加",
      open: false,
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        postName: [
          {required: true, message: "姓名不能为空", trigger: "blur"}
        ],
      }
    }
  },
  created() {
    this.getList()
  },
  mounted() {
    this.setTableHeight()
    window.onresize = () => {
      return (() => {
        this.setTableHeight()
      })()
    }
  },
  methods: {
    //设置表格高度
    setTableHeight(){
      try{
        let formHeight = document.querySelector('.el-form').clientHeight;
        let pagerHeight = document.querySelector('.el-pagination').clientHeight
        this.tableHeight = `calc(100vh - ${formHeight + pagerHeight + 20}px)`
      }catch (e) {
        console.log(e)
      }
    },
    /** 查询用户列表 */
    getList() {
      this.loading = true;
      this.list = [];
      for (let i = 0; i < 100; i++) {
        this.list.push({
          date: '2016-05-02',
          name: '王小虎',
          province: '上海',
          city: '普陀区'
        })
      }
      this.total = 100;
      this.loading = false;
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1;
      this.getList();
    },

    /** 新增按钮操作 */
    handleAdd() {
      this.open = true;
      this.title = "添加用户";
    },
    /** 导出按钮操作 */
    handleExport() {

    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.open = true;
      this.title = "修改用户";
      const postId = row.postId || this.ids
      // 修改用户
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const postIds = row.postId || this.ids;
      this.$confirm('是否确认删除用户编号为"' + postIds + '"的数据项？').then(function () {
        // 删除成功
      }).then(() => {

      }).catch(() => {
      });
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.postId)
      this.single = selection.length != 1
      this.multiple = !selection.length
    },
    //提交按钮
    submitForm() {
      this.$refs["form"].validate(valid => {
        if (valid) {

        }
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
    },
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
    }
  }
}
</script>

<style lang="scss" scoped>
.table {
  padding: 10px;
  display: flex;
  flex-direction: column;
  height: calc(100vh - 20px);

  .query-form{
    margin-top: 20px;
  }

  ::v-deep {
    flex: 1;
    .el-link+.el-link {
      margin-left: 10px;
    }
    .el-table .cell {
      font-size: 14px;
      color: rgb(36, 47, 87);
    }
    .el-table thead {
      color: #000;
      font-weight: 500;
    }
  }
}
</style>
