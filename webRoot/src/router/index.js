import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/Home'
import equip from '@/components/workItem/Equip'
import openinter from '@/components/workItem/Openinter'
import uipt from '@/components/workItem/Uipt'
import user from '@/components/workItem/User'
import equipAdd from '@/components/workItem/EquipAdd'
import uiptAdd from '@/components/workItem/UiptAdd'
import uiptEditor from '@/components/workItem/UiptEditor'
// 接口登记管理
import reginter from '@/components/workItem/Reginter'
import reginterAdd from '@/components/workItem/reginterAdd'
import reginterEditor from '@/components/workItem/reginterEditor'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home,
      children: [
        { path: '/reginter', component: reginter, name: '接口登记管理', iconCls: "icon-queshudengji1", hidden: true },
        { path: '/openinter', component: openinter, name: '接口开放管理', iconCls: "icon-fenxiang", hidden: true },
        { path: '/user', component: user, name: '用户管理', iconCls: "icon-yonghubangding", hidden: true },
        { path: '/equip', component: equip, name: '终端设备管理', iconCls: "icon-shebei", hidden: true, },
        { path: '/uipt', component: uipt, name: 'UIPT实例管理', iconCls: "icon-fuwushili", hidden: true },
        { path: '/equipAdd', component: equipAdd, name: '终端设备管理-新增', hidden: false },
        { path: '/uiptAdd', component: uiptAdd, name: '新增实例', hidden: false },
        { path: '/uiptEditor', component: uiptEditor, name: '编辑实例', hidden: false },
        { path: '/reginterAdd', component: reginterAdd, name: '接口登记管理-新增', hidden: false },
        { path: '/reginterEditor', component: reginterEditor, name: '接口登记管理-编辑', hidden: false },
      ]
    }
  ]
})
