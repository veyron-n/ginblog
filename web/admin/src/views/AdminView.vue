<template>
  <el-container class='container'>
    <NavSection/>
    <el-container>
      <el-header class='headerBtn'>
        <HeaderSection/>
      </el-header>
      <el-main>
        <router-view></router-view>
      </el-main>
      <el-footer>
        <FooterSection/>
      </el-footer>
    </el-container>
  </el-container>
</template>

<script>
import NavSection from '../components/admin/NavSection.vue'
import FooterSection from '../components/admin/FooterSection.vue'
import HeaderSection from '../components/admin/HeaderSection.vue'
export default {
  components: { NavSection, FooterSection, HeaderSection },
  methods: {

  }
}
</script>

<style scoped>

.container {
  height: 100%
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}

</style>