<template>
	<div class="page-box">
		<div class="pagecon-box">
      <el-form ref="sinteData" :model="sinteData" label-width="80px">
        <el-row style="margin-top:10px;">
          <el-col :span="6">
            <el-form-item label="用户名称:" prop='userCode'  style="margin:6px 0px 0px 20px;">
              <el-input placeholder="请输入" size="small" style="width: 200px;" v-model="sinteData.userCode"></el-input>
            </el-form-item>
            </el-col>
            <el-col :span="6">
              <el-form-item label="接口名称:" prop='inteName'  style="margin:6px 0px 0px 20px;">
               <el-input placeholder="请输入" size="small" style="width: 200px;" v-model="sinteData.inteName"></el-input>
              </el-form-item>
            </el-col>
        </el-row>
				<el-form-item label="接口说明:" prop='cmt' style="width: 600px;margin:6px 0px 0px 20px;">
					  <el-input type="textarea" v-model="sinteData.cmt"></el-input>
				</el-form-item>
         <el-row>
          <el-col :span="5">
            <el-form-item label="接口协议:" prop='inteType' style="margin:6px 0px 0px 20px;">
                <el-select v-model="sinteData.inteType " size="small">
                  <el-option label="http" value="http"></el-option>
                </el-select>
            </el-form-item>
            </el-col>
            <el-col :span="4">
                <div class="seach-con">
                  <span class="demonstration">IP</span>
						      <el-input placeholder="请输入" size="small" v-model="ConfData.ip"></el-input>
                </div>
            </el-col>
             <el-col :span="4">
                 <div class="seach-con">
                  <span class="demonstration">端口</span>
                  <el-input placeholder="请输入" size="small"  v-model="ConfData.port"></el-input>
                </div>
            </el-col>
             <el-col :span="4">
                 <div class="seach-con">
                 	<span class="demonstration">路径</span>
					      	<el-input placeholder="请输入" size="small" v-model="ConfData.path"></el-input>
                </div>
            </el-col>
        </el-row>
        <el-row>
        <el-form-item>
          <div class="fun-save">
            <el-button type="primary" size="small" @click="sinteManage()">保存</el-button>         
            <el-button size="small" @click="$router.back(-1)">返回</el-button>
          </div>
        </el-form-item>
         </el-row>
			</el-form>
		</div>      
			<div class="page-content-box">
        <div class="addCon">
             <i class="iconfont icon-zengjia"  @click="dialogFormVisible = true"></i>
          </div>
           <el-collapse accordion>
          <el-collapse-item>
            <template slot="title">
             	  <span class="span">功能一:新增序列</span>
               <div class="btn">
                 <el-button type="text" size="small">刷新 |</el-button>
                <el-button type="text" size="small">删除</el-button>
               </div>
            </template>
            <div class="coll-con">
                <div class="con-left">请求参数
                </div>
                <div class="con-right">
                    <div class="right-content">
                       请求参数
                    </div>
                    <div class="page-table1">
                        <template>
                        <el-table :data="tableReq" style="width: 100%" stripe>
                              <el-table-column fixed prop="date"  align="center" label="节点名称">
                              </el-table-column>
                              <el-table-column  prop="name" align="center" label="父节点名称">
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="数据类型">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="最大长度">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="约束">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="说明">
                              </el-table-column>
                        </el-table>
                      </template>
                   </div>
                </div>
            </div>
            <div class="coll-con">
                <div class="con-left">返回参数
                </div>
                <div class="con-right">
                    <div class="right-content">
                       返回参数
                    </div>
                    <div class="page-table1">
                        <template>
                        <el-table :data="tableRes" style="width: 100%" stripe>
                              <el-table-column fixed prop="date"  align="center" label="字段名称">
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="数据类型">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="最大长度">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="约束">
                              </el-table-column>
                               <el-table-column prop="name" align="center" label="说明">
                              </el-table-column>
                        </el-table>
                      </template>
                   </div>
                </div>
            </div>
          </el-collapse-item>
        </el-collapse>
			</div>
       
       <!-- 新增弹框 -->
        <el-dialog :visible.sync="dialogFormVisible">
           <template slot="title">
             	 <div class="dialog-top">
                   功能格式配置
                </div>
            </template>
          <div class="dialog-con">
            <!-- 选项卡切换 -->
            <el-form ref="funcData" :model="funcData" label-width="80px">
               <el-form-item label="功能名称:" prop='funcName'>
                   	<el-input placeholder="请输入" size="small"  v-model="funcData.funcName"></el-input>
                  </el-form-item>
                  <el-tabs v-model="activeName" type="card">
              <el-tab-pane label="请求参数" name="first">
                  <el-form-item label="样例:" prop='reqExample'>
                    <el-input type="textarea" style='height: 90px;margin-bottom:6px'  v-model="funcData.reqExample"></el-input>
                  </el-form-item>
                  <el-form-item label="格式ID:" prop='reqFtmCode'>
                   	<el-input placeholder="请输入" size="small"  v-model="funcData.reqFtmCode"></el-input>
                  </el-form-item>
                  <el-form-item label="格式类型:" prop='reqDataProp'>
                   		<el-select size="small" style="width:200px"  v-model="funcData.reqDataProp">
                        <el-option label="http" value="http" ></el-option>
                      </el-select>
                  </el-form-item>
                 <div class="add-div">
                   <el-button type="primary" size='small' style="width:70px" @click="dialogRequest()">新增</el-button>   
                 </div>
                 <div class="page-table">
                    <template>
                      <el-table :data="tableReq" style="width: 100%" stripe>
                            <el-table-column prop="date" align="center" label="节点名称">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.fieldName"  size='small'></el-input>
                                  <span v-else>{{scope.row.fieldName}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column  prop="name" align="center" label="父节点名称">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.extInfo"  size='small'></el-input>
                                  <span v-else>{{scope.row.extInfo}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column prop="name" align="center" label="数据类型">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.dataType"  size='small'></el-input>
                                  <span v-else>{{scope.row.dataType}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column prop="name" align="center" label="最大长度">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.length"  size='small'></el-input>
                                  <span v-else>{{scope.row.length}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column prop="name" align="center" label="约束">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.constrain"  size='small'></el-input>
                                  <span v-else>{{scope.row.constrain}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column prop="name" align="center" label="说明">
                              <template slot-scope="scope">
                                  <el-input v-if="scope.row.type" v-model="scope.row.cmt"  size='small'></el-input>
                                  <span v-else>{{scope.row.cmt}}</span>
                              </template>
                            </el-table-column>
                            <el-table-column  label="操作" align="center">
                               <template slot-scope="scope">
                                  <el-button type="text" size="small"  v-if="!scope.row.type" @click="edit(scope.$index,tableReq)">修改 |</el-button>
                                  <el-button type="text" size="small"  v-if="!scope.row.type">删除</el-button>
                                  <el-button type="text" size="small" v-if="scope.row.type" style="color:#658F34;" @click="reqSave(scope.$index,tableReq)">保存 |</el-button>
                                  <el-button type="text" size="small" v-if="scope.row.type" style="color:red;" @click="cancelReq(scope.$index,tableReq)">取消</el-button>
                                </template>
                            </el-table-column>
                      </el-table>
                    </template>
                  </div>
              </el-tab-pane>
              <el-tab-pane label="返回参数" name="second">
                    <el-form-item label="样例:" prop='resExample'>
                      <el-input type="textarea" style='height: 90px;margin-bottom:6px' v-model="funcData.resExample"></el-input>
                    </el-form-item>
                    <el-form-item label="分隔符:" prop='resFmtCode'>
                      <el-input placeholder="请输入" size="small" style="width:120px" v-model="funcData.resFmtCode"></el-input>
                    </el-form-item>
                  <div class="add-div">
                    <el-button type="primary" size='small' style="width:70px"  @click="dialogReturn()">新增</el-button>   
                  </div>
                  <div class="page-table">
                      <template>
                        <el-table :data="tableRes" style="width: 100%" stripe>
                              <el-table-column prop="date" align="center" label="字段名称">
                                    <template slot-scope="scope">
                                        <el-input v-if="scope.row.type" v-model="scope.row.fieldName"  size='small'></el-input>
                                        <span v-else>{{scope.row.fieldName}}</span>
                                    </template>
                                </el-table-column>
                              </el-table-column>
                              <el-table-column  prop="name" align="center" label="数据类型">
                                    <template slot-scope="scope">
                                      <el-input v-if="scope.row.type" v-model="scope.row.dataType"  size='small'></el-input>
                                      <span v-else>{{scope.row.dataType}}</span>
                                    </template>
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="最大长度">
                                    <template slot-scope="scope">
                                        <el-input v-if="scope.row.type" v-model="scope.row.length" size='small'></el-input>
                                        <span v-else>{{scope.row.length}}</span>
                                    </template>
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="约束">
                                 <template slot-scope="scope">
                                        <el-input v-if="scope.row.type" v-model="scope.row.constrain" size='small'></el-input>
                                        <span v-else>{{scope.row.constrain}}</span>
                                    </template>
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="说明">
                                 <template slot-scope="scope">
                                        <el-input v-if="scope.row.type" v-model="scope.row.cmt" size='small'></el-input>
                                        <span v-else>{{scope.row.cmt}}</span>
                                    </template>
                              </el-table-column>
                              <el-table-column  label="操作" align="center">
                                <template slot-scope="scope">
                                  <el-button type="text" size="small"  v-if="!scope.row.type" @click="editRes(scope.$index,tableRes)">修改 |</el-button>
                                  <el-button type="text" size="small"  v-if="!scope.row.type">删除</el-button>
                                  <el-button type="text" size="small" v-if="scope.row.type" style="color:#658F34;"  @click="resSave(scope.$index,tableRes)">保存 |</el-button>
                                  <el-button type="text" size="small" v-if="scope.row.type" style="color:red;" @click="cancelRes(scope.$index,tableRes)">取消</el-button>
                                </template>
                              </el-table-column>
                        </el-table>
                      </template>
                    </div>
                </el-tab-pane>
              </el-tabs>
             </el-form>
          </div>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="dialogSave()" size='small'>保存</el-button>            
            <el-button @click="dialogFormVisible = false;" size='small'>取消</el-button>
          </div>
        </el-dialog>
	</div>
</template>
<script>
import { funcManage, sinteManage, fmtManageSub } from "../../axios/axios.js";
export default {
  data() {
    return {
      groupId: "YYGJUIP1",
      sinteData: {
        inteName: "",
        inteType: "",
        userCode: "",
        cmt: "",
        funcList: "[]"
      },
      ConfData: {
        ip: "",
        port: "",
        path: ""
      },
      funcData: {
        funcName: "",
        reqExample: "",
        resExample: "",
        reqDataProp: "",
        resDataProp: "",
        reqFtmCode: "",
        resFmtCode: "",
        cmt: ""
      },
      tableRes: [],
      tableReq: [],
      dialogFormVisible: false,
      activeName: "first",
      formInline: {
        user: "",
        region: ""
      }
    };
  },
  methods: {
    //  请求参数编辑
    editReq(index, rows) {
      this.tableReq[index].type = true;
    },
    //  请求参数编辑
    editRes(index, rows) {
      this.tableRes[index].type = true;
    },
    //  请求参数取消
    cancelReq(index, rows) {
      this.tableReq.splice(index,1) 
    },
    //  返回参数取消
    cancelRes(index, rows) {
       this.tableRes.splice(index,1)
    },
    returnReginter() {
      this.$router.push({ path: "/reginter" });
    },
    // 接口登记保存
    sinteManage() {
      this.sinteData.groupId = this.groupId;
      this.sinteData.connConf = JSON.stringify(this.ConfData);
      let params = JSON.stringify({
        com: "POST",
        data: this.sinteData
      });
      console.log(params);
      sinteManage(params).then(res => {
        console.log(res);
        if (res.code == 200000) {
          this.ConfData = {
            ip: "",
            port: "",
            path: ""
          };
          this.$refs["sinteData"].resetFields();
          this.$message({
            message: "新增成功",
            type: "success"
          });
        } else {
          console.log(res.msg);
          this.$message({
            message: "新增失败",
            type: "error"
          });
        }
      });
    },
    // 弹框保存
    dialogSave() {
      this.funcData.groupId = this.groupId;
      let params = JSON.stringify({
        com: "POST",
        data: this.funcData
      });
      console.log(params);
      funcManage(params).then(res => {
        console.log(res);
        if (res.code == 200000) {
          this.$message({
            message: "新增成功",
            type: "success"
          });
          this.$refs["funcData"].resetFields();
          this.dialogFormVisible = false;
        } else {
          console.log(res.msg);
          this.$message({
            message: "新增失败",
            type: "error"
          });
        }
      });
    },
    // 增加请求参数
    dialogRequest() {
      var data = {
        type: true
      };
      this.tableReq.push(data);
    },
    // 请求参数保存每一行
    reqSave(index, rows) {
      let params = JSON.stringify({
        com: "POST",
        data: {
          groupId: this.groupId,
          fmtCode: "123",
          dtlInfo: "UIPFMT2",
          constrain: rows[index].constrain,
          dataType: rows[index].dataType,
          fieldName: rows[index].fieldName,
          length: rows[index].length,
          extInfo: "",
          cmt: rows[index].cmt
        }
      });
      console.log(params);
      fmtManageSub(params).then(res => {
        console.log(res);
        // if (res.code == 200000) {
        //   this.tableData[index].type = false;
        // } else {
        //   this.$message({
        //     message: res.msg,
        //     type: "error",
        //     center: true
        //   });
        // }
      });
    },
    // 增加返回参数
    dialogReturn() {
      var data = {
        type: true
      };
      this.tableRes.push(data);
    },
    // 返回参数保存每一行
    resSave(index, rows) {
      let params = JSON.stringify({
        com: "POST",
        data: {
          groupId: this.groupId,
          fmtCode: "123",
          dtlInfo: "UIPFMT2",
          constrain: rows[index].constrain,
          dataType: rows[index].dataType,
          fieldName: rows[index].fieldName,
          length: rows[index].length,
          extInfo: "",
          cmt: rows[index].cmt
        }
      });
      console.log(params);
      fmtManageSub(params).then(res => {
        console.log(res);
        // if (res.code == 200000) {
        //   this.tableData[index].type = false;
        // } else {
        //   this.$message({
        //     message: res.msg,
        //     type: "error",
        //     center: true
        //   });
        // }
      });
    }
  }
};
</script>
<style scoped>
.pagetable-top {
  margin-left: 8px;
}
.page-box {
  padding: 10px;
}
.fun-save {
  margin: 0px 0px 0 -40px;
}
</style>