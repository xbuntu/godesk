import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/SiteView.vue'
import MockView from '../views/MockView.vue'
import PhotoView from '../views/PhotoView.vue'
import {getToken} from "@/utils/token";
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
import SystemView from "@/views/SystemView";
import TableView from "@/views/TableView";
import ConfigView from "@/views/ConfigView";
import AudioView from "@/views/AudioView";
import WallpaperView from "@/views/WallpaperView";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },{
    path: '/wallpaper',
    name: 'wallpaper',
    component: WallpaperView
  },{
    path: '/photo',
    name: 'photo',
    component: PhotoView
  },{
    path: '/table',
    name: 'table',
    component: TableView
  },
  {
    path: '/mock',
    name: 'mock',
    component: MockView
  }, {
    path: '/system',
    name: 'system',
    component: SystemView
  },{
    path: '/config',
    name: 'config',
    component: ConfigView
  },{
    path: '/audio',
    name: 'audio',
    component: AudioView
  },
  {
    path: '/login',
    name: 'login',
    component: function () {
      return import(/* webpackChunkName: "about" */ '../views/LoginView.vue')
    }
  }
]

const router = new VueRouter({
  routes
})

NProgress.configure({
  easing: 'ease',  // 动画方式，和css动画属性一样（默认：ease）
  speed: 500,  // 递增进度条的速度，单位ms（默认： 200）
})


router.beforeEach((to, from , next) => {
  //每次切换页面时，调用进度条
  //NProgress.start();
  var token = getToken();
  if(!token && to.name !== 'login'){
    next({
      name:'login'
    })
    return
  }
  // 这个一定要加，没有next()页面不会跳转的。
  next();
});
router.afterEach(() => {
  // 在即将进入新的页面组件前，关闭掉进度条
  //NProgress.done()
})

export default router
