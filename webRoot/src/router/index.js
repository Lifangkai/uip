import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/Home'
import equip from '@/components/workItem/Equip'
import openinter from '@/components/workItem/Openinter'
import reginter from '@/components/workItem/Reginter'
import uipt from '@/components/workItem/Uipt'
import user from '@/components/workItem/User'
import equipAdd from '@/components/workItem/EquipAdd'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home,
      children: [
        { path: '/reginter', component: reginter, name: '接口登记管理', iconCls:"icon-dengjizongshu",hidden:true},
        { path: '/openinter', component: openinter, name: '接口开放管理',iconCls:"icon-fenxiang" ,hidden:true},
        { path: '/user', component: user, name: '用户管理',iconCls:"icon-yonghuguanli" , hidden:true},
        { path: '/equip', component: equip, name: '终端设备管理',iconCls:"icon-shebeiguanli", hidden:true,},
        { path: '/uipt', component:uipt, name: 'UIPT实例管理',iconCls:"icon-fuwushili", hidden:true },
        { path: '/equipAdd', component: equipAdd, name: '终端设备管理-新增', hidden:false},
    ]
    }
  ]
})
