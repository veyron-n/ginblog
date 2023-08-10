<template>
  <el-aside class='nav'>
    <el-row>
      <el-col>
        <div class='log'>
          <span>GinBlog</span>
        </div>
        <el-menu
          active-text-color="#ffd04b"
          background-color="#545c64"
          class="el-menu-vertical-demo"
          default-active="2"
          text-color="#fff"
          @open="handleOpen"
          @close="handleClose"
          router
        >
          <el-menu-item index="index">
            <template #title>
              <el-icon><Odometer /></el-icon>
              <span>仪表盘</span>
            </template>
          </el-menu-item>
          <el-sub-menu index="2">
            <template #title>
              <el-icon><MessageBox /></el-icon>
              <span>文章管理</span>
            </template>
            <el-menu-item index="addart">
              <el-icon><EditPen /></el-icon>
              <span>写文章</span>
            </el-menu-item>
            <el-menu-item index="artlist">
              <el-icon><Document /></el-icon>
              <span>文章列表</span>
            </el-menu-item>
          </el-sub-menu>

          <el-menu-item index="catelist">
            <el-icon><Guide /></el-icon>
            <span>分类列表</span>
          </el-menu-item>
          <el-menu-item index="userlist">
            <el-icon><User /></el-icon>
            <span>用户列表</span>
          </el-menu-item>
        </el-menu>
      </el-col>
    </el-row>
  </el-aside>
</template>

<script>
import { Odometer, MessageBox, EditPen, Document, Guide, User } from '@element-plus/icons-vue'

export default {
  components: { Odometer, MessageBox, EditPen, Document, Guide, User },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },
  }
}

</script>

<style>
.nav{
  background-color: #545c64;
  text-align: center;
}
.nav span{
  font-size: 18px;
}
.log {
  height: 32px;
  margin: 16px;
  color: #ffd04b;
  background-color: #545c64;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 17px;
}
</style>

