<template>
  <div class="app">
    <div class="appMenu">
      <Menu v-if="routeName !== 'login'" :routeName="routeName"/>
    </div>
    <div class="appMain">
      <router-view/>
    </div>
  </div>
</template>

<script>
import {GetOs} from "../wailsjs/go/controller/System";
import Menu from './components/Menu.vue'
export default {
  name: 'AppView',
  components: {
    Menu
  },
  computed: {
    routeName() {
      return this.$route.name;
    }
  },
  data(){
    return {
    }
  },
  created() {
    GetOs().then(os => {
      localStorage.setItem("os",os)
      // 根据操作系统类型进行判断
      let runOs = "未知操作系统"
      switch (os) {
        case "windows":runOs = "Windows系统";break;
        case "darwin":runOs = "Mac系统";break;
        case "linux":runOs = "Linux系统";break;
      }
      this.$notify.success({
        title: '温馨提示',
        message: `软件运行在${runOs}上，你可以根据系统版本做不同的界面效果`
      })
    })
  },
  methods:{

  }
}
</script>


<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
html,body{
  margin: 0;
  padding: 0;
}
body{
  width: 100vw;
  height: 100vh;
  background: #ffffff;
}
.app{
  display: flex;
  align-items: center;
  .appMenu{
    flex-shrink: 0;
  }
  .appMain{
    height: 100vh;
    overflow: auto;
    flex: 1;
  }
}
</style>
