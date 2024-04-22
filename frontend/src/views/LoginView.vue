<template>
  <div class="login">
    <div class="left"></div>
    <div class="right">
      <div class="form">
        <el-form ref="form" :rules="rules" :model="form" label-width="80px">
          <el-form-item label="账号" prop="username">
            <el-input v-model="form.username" placeholder="用户名"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" placeholder="密码" type="password" show-password></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="onSubmit">登录</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
import {setToken} from "@/utils/token";
export default {
  name: 'LoginView',
  data() {
    return {
      form:{
        username:'admin',
        password:'admin'
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' }
        ]
      }
    }
  },
  methods:{
    onSubmit(){
      this.$refs['form'].validate((valid) => {
        if (valid) {
          if(this.form.username !== 'admin' || this.form.password !== 'admin'){
            this.$notify.error({
              title:'温馨提示',
              message:'用户名或密码错误'
            })
            return false
          }
          //1、登录请求获取token，这里直接存 用户名
          setToken('admin')
          this.$notify.success({
            title:'温馨提示',
            message:'登录成功'
          })
          this.$router.push({
            name:'home'
          })
        } else {
          return false;
        }
      });
    }
  }
}
</script>

<style lang="scss">
  .login{
    width: 100vw;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #ebf4ff;
    .left{
      width: 300px;
      height: 300px;
      background: url("../assets/loginBg.svg") 0 0 no-repeat scroll;
      background-size: 100% 100%;
    }
    .right{
      .form{
        width: 300px;
      }
    }
  }
</style>
