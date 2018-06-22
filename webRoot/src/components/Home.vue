<template>
    <el-container width="100%" height="100%">
     <el-aside v-bind:style="{width:(isCollapse?'200px':'64px')}" style="background-color: rgba(0, 21, 41, 1)" >
          <el-row style="height:64px">
            <el-col :span="6"><div class="grid-content bg-purple home-logo">
               <i class="fl icon-clolor iconfont icon-hisgai"></i>
              </div></el-col>
            <el-col :span="14" v-show="isCollapse"><div class="grid-content bg-purple home-text">统一接口平台</div></el-col>
          </el-row>
          <div class="nav-box">
            <template v-for="item in $router.options.routes[0].children"  v-if="item.hidden" !=false >
              <router-link  v-bind:to="item.path">
                    <i class="iconfont" v-bind:class="item.iconCls"  v-bind:title="isCollapse?'':item.name"></i>
                    <span v-show="isCollapse" class="">{{item.name}}</span>
              </router-link>  
            </template>
          </div>
      </el-aside>
    <el-main>
      <el-row class="main-top">
        <el-col :span="6">
          <div class="grid-content bg-purple mt-left">
            <i class="fl icon-clolor iconfont icon-daohangliebiao" @click="changeLeftNav"></i>
            <template><p class="fr">{{$route.name}}</p></template>
          </div>
        </el-col>

        <el-col :span="6" :offset="12">
          <div class="grid-content bg-purple clearf">
            <el-dropdown trigger="click">
              <span class="el-dropdown-link userinfo-inner">
                <img :src="this.sysUserAvatar" class='fr'/>管理员<i class="el-icon-arrow-down el-icon--right"></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item>我的消息</el-dropdown-item>
                    <el-dropdown-item>设置</el-dropdown-item>
                    <el-dropdown-item divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
        </el-col>
      </el-row>
      <transition name="fade" mode="out-in">
        <router-view class="right-box"></router-view>
      </transition>
    </el-main>
    
   </el-container>
</template>

<script>
import icon from "@/assets/icon.png";
export default {
  name: "HelloWorld",
  data() {
    return {
      icon: icon,
      isCollapse: true,
      sysUserAvatar:
        "https://raw.githubusercontent.com/taylorchen709/markdown-images/master/vueadmin/user.png",
      navflag: ""
    };
  },
  methods: {
    changeLeftNav() {
      this.isCollapse = !this.isCollapse;
    }
  }
};
</script>

<style scoped>
@import "../../static/commonCss/home.css";
.el-container {
  width: 100%;
  height: 100%;
  box-sizing: border-box;
}
.el-main {
  padding: 0 0 20px;
}
.el-aside {
  border-right: 1px solid #d8d8d8;
}
.icon-clolor {
  color: #1c93ef !important;
}
.home-logo {
  width: 42px;
  height: 42px;
  margin: 16px 5px 8px 10px;
}
.home-logo i {
  font-size: 42px;
}
.home-text {
  background: linear-gradient(to right, #1782e5, #67a0e4, #c5d2e5);
  -webkit-background-clip: text;
  color: transparent;
  height: 64px;
  line-height: 74px;
  font-size: 18px;
  margin-left: 2px;
}
.el-main {
  height: 100%;
}
.main-top {
  height: 64px;
  padding: 0 20px;
  border-bottom: 1px solid rgba(226, 226, 226, 1);
}
.mt-left {
  padding-top: 22px;
}
.mt-left i {
  display: block;
  cursor: pointer;
  width: 30px;
  height: 30px;
  font-size: 28px;
  color: #14823b;
  float: left;
}
.mt-left p {
  padding: 8px 0 0 20px;
  line-height: 10px;
  font-size: 16px;
  font-weight: 500;
}
.userinfo {
  float: right;
  line-height: 64px;
}
.main-top .userinfo .userinfo-inner {
  cursor: pointer;
  color: #354052;
  height: 100%;
  display: block;
  position: relative;
}
.main-top .userinfo .userinfo-inner span {
  line-height: 20px;
  font-size: 14px;
  padding-top: 20px;
  max-width: 170px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.main-top .userinfo .userinfo-inner i {
  padding-top: 23px;
  width: 20px;
  height: 20px;
  color: #14823b;
}
.main-top .userinfo-inner img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 20px;
  margin: 10px 10px 10px 0px;
}
.fr {
  float: left;
  font-size: 18px;
}
.tools {
  line-height: 64px;
  text-align: left;
}
.right-box {
  background: #eff3f6;
  min-height: 700px;
}
.clearf {
  text-align: right;
  margin-right: 20px;
  line-height: 64px;
}
.clearf i {
  font-size: 10px;
  margin-left: 6px;
}
</style>
