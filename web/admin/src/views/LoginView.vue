<template>
  <div class='container'>
    <div class="loginBox">
      <h1>登录页面</h1>
      <el-button type='primary'>登录</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LoginView',
  components: {}
}
</script>

<style scoped>
  .container{
    height: 100%;
    background-color: pink;
  }
  .loginBox{
    width: 450px;
    height: 300px;
    background-color: #fff;
    position: absolute;
  }
</style>
