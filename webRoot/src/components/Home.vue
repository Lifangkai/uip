<template>
     <el-container width="100%" height="100%">
     <!-- <template @transferUser="getUser"></template> 
     <p>{{text}}</p> -->
    <el-aside v-bind:style="{width:(isCollapse?'200px':'64px')}" style="background-color: rgba(0, 21, 41, 1)" >
        <el-row style="height:64px">
          <el-col :span="6"><div class="grid-content bg-purple home-logo" :style="{backgroundImage: 'url(' + icon + ')' }"></div></el-col>
          <el-col :span="10" v-show="isCollapse"><div class="grid-content bg-purple home-text">统一接口</div></el-col>
        </el-row>
        <div class="nav-box">
          <template v-for="item in $router.options.routes[0].children" >
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
                    <!-- <i class="fl iconfont icon-daohangliebiao"  @click="changeLeftNav"></i> -->
              <p class="fl">{{$route.name}}</p>

                  </div>
                </el-col> 
          </el-row>
          <router-view class="right-box"></router-view>
    </el-main>
    
   </el-container>
</template>

<script>
import icon from '@/assets/icon.png'
export default {
  name: "HelloWorld",
  data() {
    return {
     icon:icon,
     isCollapse:true,
      navflag: ""
    };
  },
  methods: {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import "../../static/commonCss/home.css";
 
.el-container{
  width:100%;
  height:100%;
  box-sizing: border-box;
}
.el-main{
  padding:0 0 20px;
}
.el-aside{border-right:1px solid #D8D8D8;}
.home-logo{
    width: 48px;
    height: 48px;
    background-position:-155px -492px;
    background-repeat:  no-repeat;
    margin:10px 0 7px 8px;
}
.home-text{
  color: #14823B;
  height: 64px;
  line-height: 70px;
  font-size: 14px;
}
.el-main{
  height: 100%;
}
.main-top{
   height:64px;
   padding:0 20px;
   border-bottom: 1px solid rgba(226, 226, 226, 1);

}
.mt-left{
  padding-top:22px; 
}
.mt-left i{
  display: block;
  cursor: pointer;
  width:30px;
  height:30px;
  font-size: 28px;
  color: #14823B;
}
.mt-left p{
  padding:8px 0 0 22px;
  line-height: 20px;
  font-size: 14px;
  font-weight: 500;
}
.userinfo{
  float: right;
}
.main-top .userinfo .userinfo-inner{
    cursor: pointer;
    color: #354052;
    height: 100%;
    display:block;
    position: relative;
}
.main-top .userinfo .userinfo-inner span{
  line-height: 20px;
    font-size: 14px;
    padding-top: 20px;
    max-width: 170px;
    overflow: hidden;
    text-overflow:ellipsis;
    white-space: nowrap;
}
.main-top .userinfo .userinfo-inner i{
  padding-top: 23px;
  width:20px;
  height:20px;
  color: #14823B;
}
.main-top  .userinfo-inner img{
    display: block;
    width: 40px;
    height: 40px;
    border-radius: 20px;
    margin: 10px 10px 10px 0px;
}
.el-dropdown-menu__item{
  padding-left:56px;
   padding-right:36px;
  position: relative;
}
.el-dropdown-menu__item i{
  display: block;
  width:20px;
  height:20px;
  line-height: 20px;
  font-size: 20px;
  position: absolute;; 
  top:7px;
  left:22px;
  background-repeat:  no-repeat;
}
.el-dropdown-menu__item i.icon-tuichu{
  line-height: 33px;
}
.el-dropdown-menu__item:not(.is-disabled):hover{
  background: #ffffff;
  color: #14823B;
}
i.logout-item{
  margin: 0;
  top:14px;
  background-position:-168px -738px;
}
.el-dropdown-menu__item:not(.is-disabled):hover  i.logout-item{
  background-position:-168px -763px;
}
</style>
