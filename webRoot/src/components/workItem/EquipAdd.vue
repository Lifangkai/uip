<template>
    <div>
        <router-view></router-view>
        <div class="main_block block-height1">
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端SN：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入" v-model="termSn"></el-input>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端分类：</label>
                    <div class="add_form_item_con">
                        <el-select size="small" v-model="termType" placeholder="请选择">
                            <el-option
                            v-for="item in termType_options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端名称：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入" v-model="termName"></el-input>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">说明：</label>
                    <div class="add_form_item_con">
                        <el-input
                        type="textarea"
                        :rows="2"
                        placeholder="请输入内容"
                        v-model="cmt">
                        </el-input>
                    </div>
                </div>   
            </el-col>
        </div>
        <div class="main_block block-height3">
            <div>
                <div class="block-title"><p>宿主信息</p></div>
                <hr style="height:1px;border:none;border-top:1px solid #EFF3F6;" />
                <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">实例名称：</label>
                        <div class="add_form_item_con">
                            <!-- <el-input size="small" placeholder="请输入"></el-input> -->
                            <el-autocomplete
                                size="small"
                                v-model="queryString"
                                :fetch-suggestions="querySearchAsync"
                                placeholder="请输入内容"
                                @select="handleSelect"></el-autocomplete>
                        </div>
                    </div>   
                </el-col>
                <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24"><div class="grid-content"></div></el-col>

                <!-- <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div></el-col> -->
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">IP地址：</label>
                        <div class="add_form_item_con">
                            <el-input size="small" placeholder="请输入"></el-input>
                        </div>
                    </div>   
                </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">服务端口：</label>
                        <div class="add_form_item_con">
                            <el-input size="small" placeholder="请输入"></el-input>
                        </div>
                    </div>   
                </el-col>
                <!-- <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div></el-col> -->
            </div>
        </div>
        <div class="main_block block-height2">
            <div>
                <div class="block-title"><p>驱动信息</p></div>
                <hr style="height:1px;border:none;border-top:1px solid #EFF3F6;" />
                <div class="block-header">
                <el-col :xs="12" :sm="12" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">驱动类型：</label>
                        <div class="add_form_item_con">
                            <el-select size="small" v-model="drivType" placeholder="请选择">
                                    <el-option
                                    v-for="item in drivType_options"
                                    :key="item.value"
                                    :label="item.label"
                                    :value="item.value">
                                    </el-option>
                                </el-select>
                        </div>
                    </div>
                </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row">
                        <el-button type="primary" size="small" @click="drivParamAdd">新增</el-button></div>
                </el-col>
                </div>
                <div class="block-table">
                    <el-table :data="tableData" height="200" border style="width: 90%">
                    <el-table-column prop="" label="参数名称" style="width: 20%">
                        <template slot-scope="scope">
                            <span v-if="!scope.row.editFlag"><el-input v-model="paramName" placeholder="请输入内容"></el-input></span>
                            <span v-if="scope.row.editFlag">{{ scope.row.paramName }}</span>
                            <!-- <span v-if="scope.row.editFlag" class="cell-edit-input"><el-input v-model="inputColumn1" placeholder="请输入内容"></el-input></span> -->
                        </template>
                    </el-table-column>
                    <el-table-column prop="termName" label="参数编码" style="width: 25%">
                        <template slot-scope="scope">
                        <span v-if="!scope.row.editFlag"><el-input v-model="paramCode" placeholder="请输入内容"></el-input></span>
                            <span v-if="scope.row.editFlag">{{ scope.row.paramCode }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column prop="cmt" label="参数值" style="width: 25%">	
                        <template slot-scope="scope">
                            <span v-if="!scope.row.editFlag"><el-input v-model="paramValue" placeholder="请输入内容"></el-input></span>	
                            <span v-if="scope.row.editFlag">{{ scope.row.paramValue }}</span>
                        </template>
                    </el-table-column>
                    <el-table-column label="操作" style="width: 25%">
                        <template slot-scope="scope">
                            <el-button
                                size="mini"
                                type="danger"
                                @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                            <el-button
                                size="mini"
                                @click="handleSave(scope.$index, scope.row)">保存</el-button>
                        </template>
                    </el-table-column>
                </el-table>
                </div>
                <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16"><p>&nbsp;</p></el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row">
                    <el-button type="primary" size="small" @click="equipAdd">保存</el-button></div>
                </el-col>
            </div>
        </div>
    </div>
</template>
<script>
import axios from "axios";

export default {
//<<<<<<< HEAD
//   data() {
//     return {
//       termType_options: [
//         {
//           value: "选项1",
//           label: "黄金糕"
//         },
//         {
//           value: "选项2",
//           label: "双皮奶"
//         },
//         {
//           value: "选项3",
//           label: "蚵仔煎"
// =======
    data() {
        return {
            termType_options:[{
                value: '选项1',
                label: '黄金糕'
                }, {
                value: '选项2',
                label: '双皮奶'
                }, {
                value: '选项3',
                label: '蚵仔煎'
                }, {
                value: '选项4',
                label: '龙须面'
                }, {
                value: '选项5',
                label: '北京烤鸭'
                }],
            drivType_options:[{
                value: '驱动1',
                label: '驱动1'
            }],
            drivType:'',
            termType:'',
            instances: [],
            // state4: '',
            timeout:  null,
            value:'',
            cmt:'',
            tableData:[],
            queryString:'',
            instCode:'',//实例编码
            termName: '',
            termSn:'',
            paramName:'',
            paramCode:'',
            paramValue:''
        };
    },
    methods: {
        //新增终端
        equipAdd() {
            axios.defaults.baseURL = "http://172.16.0.13:31425"
            axios.post('/terminal/manage', {
                com:'POST',
                data: {
                    "groupId":"12345678",
                    "termName":"终端名4",
                    "termType": "终端类型4",
                    "instCode": "sd3434002", 
                    "drivType": "驱动4", 
                    "drivParam":'',
                    "cmt":"说明4"
                }          
            })
            .then(function (response) {
                console.log(response);
                if (response.data.code) {
                    alert(response.data.msg);
                }
            })
            .catch(function (error) {
                console.log(error);
                if (response.data.code) {
                    alert(response.data.msg);
                }
            });
        },
        //添加驱动参数的方法
        drivParamAdd(){
            this.tableData.push({ 
                paramName:'',
                paramCode:'',              
                paramValue:'',                
                editeFlage:true
            })
        },
        handleSave(index, row){
            this.tableData[index].editeFlage = false;
            this.tableData[index].paramName = this.paramName;
            this.tableData[index].paramCode = this.paramCode;
            this.tableData[index].paramValue = this.paramValue;
            console.log(this.tableData);
            editeFlage:false;
        },
        handleDelete(index, row){
            this.tableData.splice(index, 1);
        },
        loadAll() {
            return []
// >>>>>>> f7f4c84af31bc98f92f7a01c0a74f0b97d2dab34
//         },
//         {
//           value: "选项4",
//           label: "龙须面"
//         },
//         {
//           value: "选项5",
//           label: "北京烤鸭"
//         }
//       ],
//       drivType_options: [
//         {
//           value: "驱动1",
//           label: "驱动1"
//         }
//       ],
//       drivType: "",
//       termType: "",
//       instances: [],
//       // state4: '',
//       timeout: null,
//       value: "",
//       textarea: "",
//       tableData: [],
//       queryString: ""
//     };
//   },
//   methods: {
//     //新增终端
//     equipAdd() {
//       axios.defaults.baseURL = "http://172.16.0.13:31425";
//       axios
//         .post("/terminal/manage", {
//           com: "POST",
//           data: {
//             groupId: "12345678",
//             termName: "终端名4",
//             termType: "终端类型4",
//             instCode: "sd3434002",
//             drivType: "驱动4",
//             drivParam: "",
//             cmt: "说明4"
//           }
//         })
//         .then(function(response) {
//           console.log(response);
//           if (response.data.code) {
//             alert(response.data.msg);
//           }
//         })
//         .catch(function(error) {
//           console.log(error);
//           if (response.data.code) {
//             alert(response.data.msg);
//           }
//         });
//     },
//     loadAll() {
//       return [];
    },
    querySearchAsync(queryString, cb) {
      console.log(this.queryString);
      let _this = this;
      axios.defaults.baseURL = "http://172.16.0.13:31425";
      axios
        .get("/instance/manage/ftsearch", {
          params: {
            groupId: "12345678",
            fttext: _this.queryString
          }
        })
        .then(function(response) {
          console.log(response);
          if (response.data.code === "200000") {
            // _this.tableData = response.data.data;
          } else {
          }
        })
        .catch(function(error) {
          console.log(error);
          // _this.tableData = [];
        });
    },
    handleSelect(item) {
      console.log(item);
    }
  },
  mounted() {
    this.instances = this.loadAll();
  }
};
</script>

<style scoped>
.main_row {
  margin-top: 20px;
}
.block-height1 {
  height: 256px;
}
.block-height3 {
  height: 184px;
}
.block-title {
    height: 48px;
}
.block-table el-table {
    padding: 20px, 10px, 10px, 10px;
}

.block-title p {
    margin-left: 19px;
    padding-top: 11px;
}
.block-header {
    height: 72px;
}
.block-height2 {
  height: 400px;
}
.add_form_item_con .el-autocomplete {
  width: 100%;
}
.add_form_item{
	padding: 0;
	display: flex;
}
.add_form_label{
    width: 90px;
	margin: 6px 0;
	padding-right: 15px;
	text-align: right;
}
.add_form_item_con{
	flex: 1;
}

.main_block {
    padding: 10px ;
    border: 10px solid #EFF3F6;
    background: white;
}
</style>