<template>
  <el-button type='danger' @click='loginOut' plain>退出</el-button>

</template>
<script>
export default {
  methods: {
    loginOut() {
      window.sessionStorage.clear('token')
      this.$router.push('/login')
    }
  }
}

</script>