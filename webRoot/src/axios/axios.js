import axios from 'axios';
import routes from '../router'
import QS from 'qs'

let base1 = 'http://172.16.0.13:7001';   //uip svc

//let base2 = 'http://172.16.0.14:31300';   //sns svc
let base2 = 'http://172.16.0.13:7501';   //sns svc


// 接口登记保存
export const sinteManage = params => { return instanceBase1.post(`/sinte/manage`, params).then(res => res.data); };
// 接口登记格式id的sns
export const sinteSns = params => { return instanceBase2.get(`/snList`, { params: params }).then(res => res.data); };
// 接口登记查询
export const sinteSearch = params => { return instanceBase1.get(`/sinte/manage/ftsearch`, { params: params }).then(res => res.data); };
// 接口登记弹框保存
export const funcManage = params => { return instanceBase1.post(`/func/manage`, params).then(res => res.data); };
// 接口登记弹框查询
export const funcManageAll = params => { return instanceBase1.get(`/func/manage`, { params: params }).then(res => res.data); };
// 接口登记数据格式保存
export const fmtManageSub = params => { return instanceBase1.post(`/fmt/manage/sub`, params).then(res => res.data); };
// 接口登记数据格式保存
export const fmtManageMain = params => { return instanceBase1.post(`/fmt/manage/main`, params).then(res => res.data); };
// 接口登记数据格式保存
export const fmtSubAll = params => { return instanceBase1.get(`/fmt/manage/main`, { params: params }).then(res => res.data); };
const instanceBase1 = axios.create({
    //  baseURL: config.BASE1_URL,
    baseURL: base1,
    // headers: {
    //     'Access-Control-Allow-Origin':'*',
    //     'Content-Type': 'application/json',
    //     'Access-Control-Allow-Headers':'Authorization,Origin, X-Requested-With, Content-Type, Accept',
    // }
    // transformRequest: [function (data) {
    //     data = QS.stringify(data);
    //     return data;
    // }]
});

instanceBase1.defaults.headers.post['Content-Type'] = 'application/json';
const instanceBase2 = axios.create({
    //  baseURL: config.BASE1_URL,
    baseURL: base2,
    // headers: {
    //     'Access-Control-Allow-Origin':'*',
    //     'Content-Type': 'application/json',
    //     'Access-Control-Allow-Headers':'Authorization,Origin, X-Requested-With, Content-Type, Accept',
    // }
    // transformRequest: [function (data) {
    //     data = QS.stringify(data);
    //     return data;
    // }]
});

instanceBase2.defaults.headers.post['Content-Type'] = 'application/json';
// instanceBase1.interceptors.request.use(
//        config => {
//         // if (request.getMethod().equals("OPTIONS")) {
//         //     HttpUtil.setResponse(response, HttpStatus.OK.value(), null);
//         //     return;
//         // }
//         //   config.headers.Authorization = store.state.loginInfo.access_token;
//           // var xtoken = getXtoken()
//           // if(xtoken != null){
//           //     config.headers['X-Token'] = xtoken
//           // }
//           // if(config.method=='post'){
//           //     config.data = {
//           //         ...config.data,
//           //         _t: Date.parse(new Date())/1000,
//           //     }
//           // }else if(config.method=='get'){
//           //     config.params = {
//           //         _t: Date.parse(new Date())/1000,
//           //         ...config.params
//           //     }
//           // }
//         return config;
//       },function(error){
//           store.dispatch("loginout");//登出删除状态
//           return Promise.reject(error)
//     });

