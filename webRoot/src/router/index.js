import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/Home'
import equip from '@/components/workItem/Equip'
import openinter from '@/components/workItem/Openinter'
import reginter from '@/components/workItem/Reginter'
import uipt from '@/components/workItem/Uipt'
import user from '@/components/workItem/User'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home,
      children: [
        { path: '/reginter', component: reginter, name: '接口登记管理', iconCls:"icon-dengjizongshu"},
        { path: '/openinter', component: openinter, name: '接口开放管理',iconCls:"icon-fenxiang"},
        { path: '/user', component: user, name: '用户管理',iconCls:"icon-yonghuguanli" },
        { path: '/equip', component: equip, name: '终端设备管理',iconCls:"icon-shebeiguanli" },
        { path: '/uipt', component:uipt, name: 'UIPT实例管理',iconCls:"icon-fuwushili" },
    ]
    }
  ]
})
