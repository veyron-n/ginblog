<template>
  <div class='container'>
    <div class='loginBox'>
      <h3>登录</h3>
      <el-form class='loginForm' :model="formData" :rules="rules" ref='loginFormRef'>

        <el-form-item prop="username">
          <el-input v-model="formData.username" placeholder="请输入账号">
            <template #prefix>
              <el-icon class="el-icon--left">
                <User />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>

        <el-form-item prop="password">
          <el-input v-model="formData.password"
                    type="password" placeholder="请输入密码" show-password
                    v-on:keyup.enter='login'>
            <template #prefix>
              <el-icon class="el-icon--left">
                <Lock />
              </el-icon>
            </template>
          </el-input>
        </el-form-item>

        <el-button type="success" plain @click="login">登录</el-button>
        <el-button type="info" plain @click="resetForm">取消</el-button>
      </el-form>
    </div>
  </div>
</template>

<script>
import { User, Lock } from '@element-plus/icons-vue'

export default {
  name: 'LoginView',
  computed: {},
  components: { User, Lock },
  data() {
    return {
      formData: {
        username: '',
        password: '',
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          { min: 4, max: 20, message: '长度在 4 - 20 个字符', trigger: 'blur'}
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {min: 6, max: 20, message: '长度在 6 - 20 个字符', trigger: 'blur'}
        ]
      }
    }
  },
  methods: {
    login() {
      this.$refs.loginFormRef.validate(async (valid)=> {
        if (!valid) return this.$message.error("输入非法数据请重新输入")
        const { data: res } = await this.$http.post("login", this.formData)
        if (res.status !== 200) return this.$message.error(res.message)
        window.localStorage.setItem('token', res.token)
        this.$router.push('admin/index')
        this.$message.success('登录成功')
      })
    },
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    }
  }
}
</script>

<style scoped>
  .container{
    height: 100%;
    background-color: #46CEDA;
  }
  .loginBox{
    width: 450px;
    height: 300px;
    background-color: #fff;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 9px;
  }
  .loginForm{
    width: 100%;
    position: absolute;
    bottom: 10%;
    padding: 0 20px;
    box-sizing: border-box;
  }
</style>
