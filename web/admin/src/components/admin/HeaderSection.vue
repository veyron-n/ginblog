<template>
  <div @click='loginOut' class="logout-button">
    <el-icon><SwitchButton /></el-icon>&nbsp;<b>退出登录</b>
  </div>
</template>
<script>
import { SwitchButton } from '@element-plus/icons-vue'

export default {
  components: { SwitchButton },
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}
</script>
<style>
.logout-button {
  font-size: 15px;
  color: #545c64; /* 设置字体颜色为红色 */
  cursor: pointer; /* 添加鼠标指针样式，表示可点击 */
}
</style>