<template>
	<div class="page-box">
		<div class="pagecon-box">
			<el-form ref="form" :model="sinteData" label-width="80px">
				<div class="form-serch" style="margin-bottom: 20px;">
					<div>
						<span class="demonstration">用户名称:</span>
						<el-input placeholder="请输入" size="small" style="width: 200px;" v-model="sinteData.userCode"></el-input>
					</div>
					<div style="margin-left: 70px;">
						<span class="demonstration">接口名称:</span>
						<el-input placeholder="请输入" size="small" style="width: 200px;" v-model="sinteData.inteName"></el-input>
					</div>
				</div>
				<el-form-item label="接口说明:" style="width: 600px;margin-left:20px;margin-bottom: 0px;">
					<el-input type="textarea" v-model="sinteData.cmt"></el-input>
				</el-form-item>
				<div class="form-serch" :model="ConfData">
					<div>
						<span class="demonstration">接口协议:</span>
						<el-select v-model="sinteData.inteType " size="small">
							<el-option label="http" value="http"></el-option>
						</el-select>
					</div>
					<div class="seach-con">
						<span class="demonstration">IP</span>
						<el-input placeholder="请输入" size="small" style="width: 160px;" v-model="ConfData.ip"></el-input>
					</div>
          <div class="seach-con">
						<span class="demonstration">端口</span>
						<el-input placeholder="请输入" size="small" style="width: 160px;" v-model="ConfData.port"></el-input>
					</div>
          <div class="seach-con">
						<span class="demonstration">路径</span>
						<el-input placeholder="请输入" size="small" style="width: 160px;" v-model="ConfData.path"></el-input>
					</div>
				</div>
			</el-form>
       <div class="fun-save">
              <el-button type="primary" size="small" @click="sinteManage()">保存</el-button>         
              <el-button size="small" @click="$router.back(-1)">返回</el-button>
        </div>
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
                        <el-table :data="tableData" style="width: 100%" stripe>
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
                        <el-table :data="tableData" style="width: 100%" stripe>
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
            <el-form ref="form" :model="funcData" label-width="80px">
               <el-form-item label="功能名称:">
                   	<el-input placeholder="请输入" size="small"  v-model="funcData.funcName"></el-input>
                  </el-form-item>
                  <el-tabs v-model="activeName" type="card">
              <el-tab-pane label="请求参数" name="first">
                  <el-form-item label="样例:">
                    <el-input type="textarea" style='height: 90px;margin-bottom:6px'  v-model="funcData.reqExample"></el-input>
                  </el-form-item>
                  <el-form-item label="格式ID:">
                   	<el-input placeholder="请输入" size="small"></el-input>
                  </el-form-item>
                  <el-form-item label="格式类型:">
                   		<el-select v-model="formInline.region" size="small" style="width:200px">
                        <el-option label="http" value="http"></el-option>
                      </el-select>
                  </el-form-item>
                 <div class="add-div">
                   <el-button type="primary" size='small' style="width:70px" @click="dialogAdd()">新增</el-button>   
                 </div>
                 <div class="page-table">
                    <template>
                      <el-table :data="tableData" style="width: 100%" stripe>
                            <el-table-column prop="date" align="center" label="节点名称">
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
                            <el-table-column  label="操作" align="center">
                              <template slot-scope="scope">
                                <el-button type="text" size="small">修改 |</el-button>
                                <el-button type="text" size="small">删除</el-button>
                              </template>
                            </el-table-column>
                      </el-table>
                    </template>
                  </div>
              </el-tab-pane>
              <el-tab-pane label="返回参数" name="second">
                    <el-form-item label="样例:">
                      <el-input type="textarea" style='height: 90px;margin-bottom:6px' v-model="funcData.resExample"></el-input>
                    </el-form-item>
                    <el-form-item label="分隔符:">
                      <el-input placeholder="请输入" size="small" style="width:120px"></el-input>
                    </el-form-item>
                  <div class="add-div">
                    <el-button type="primary" size='small' style="width:70px"  @click="dialogAdd()">新增</el-button>   
                  </div>
                  <div class="page-table">
                      <template>
                        <el-table :data="tableData" style="width: 100%" stripe>
                              <el-table-column prop="date" align="center" label="字段名称">
                              </el-table-column>
                              <el-table-column  prop="name" align="center" label="数据类型">
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="最大长度">
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="约束">
                              </el-table-column>
                              <el-table-column prop="name" align="center" label="说明">
                              </el-table-column>
                              <el-table-column  label="操作" align="center">
                                <template slot-scope="scope">
                                  <el-button type="text" size="small">修改 |</el-button>
                                  <el-button type="text" size="small">删除</el-button>
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
import { funcManage, sinteManage } from "../../axios/axios.js";
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
      tableData: [
        {
          date: "2016-05-02",
          name: "王小虎",
          address: "上海市普陀区金沙江路 1518 弄"
        },
        {
          date: "2016-05-04",
          name: "王小虎",
          address: "上海市普陀区金沙江路 1517 弄"
        },
        {
          date: "",
          name: "",
          address: ""
        }
      ],
      dialogFormVisible: false,
      activeName: "first",
      formInline: {
        user: "",
        region: ""
      }
    };
  },
  methods: {
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
    // 弹框新增
    dialogAdd() {
      console.log(this.tableData);
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
    // 新增取消
    // addUserOff(formName) {
    //   this.addUser = false;
    //   this.$refs[formName].resetFields();
    // }
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
  text-align: left;
  margin: 20px 40px 0 40px;
}
</style>