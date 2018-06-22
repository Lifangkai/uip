<template>
	<div class="page-box">
		<div class="nav_box">
			<template>
				<el-tabs v-model="activeName1">
					<el-tab-pane label="登录接口信息" name="first">
						<div style="margin:0px 10px" class="reginter-con">
							<div class="pagecon-box">
								<el-form ref="sinteData" :model="sinteData" label-width="80px">
									<el-row style="margin-top:10px;">
										<el-col :span="6">
											<el-form-item label="用户名称:" prop='userCode' style="margin:6px 0px 0px 20px;">
												<el-input placeholder="请输入" size="small" style="width: 200px;" v-model="sinteData.userCode"></el-input>
											</el-form-item>
										</el-col>
										<el-col :span="6">
											<el-form-item label="接口名称:" prop='inteName' style="margin:6px 0px 0px 20px;">
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
											<el-form-item style="margin-left:-60px;">
												<el-input placeholder="请输入内容" v-model="ConfData.ip" size="small">
													<template slot="prepend">IP</template>
												</el-input>
											</el-form-item>
										</el-col>
										<el-col :span="4">
											<el-form-item style="margin-left:-60px;">
												<el-input placeholder="请输入内容" v-model="ConfData.port" size="small">
													<template slot="prepend">端口</template>
												</el-input>
											</el-form-item>
										</el-col>
										<el-col :span="4">
											<el-form-item style="margin-left:-60px;">
												<el-input placeholder="请输入内容" v-model="ConfData.path" size="small">
													<template slot="prepend">路径</template>
												</el-input>
											</el-form-item>
										</el-col>
									</el-row>
									<el-row>
										<el-form-item>
											<div class="fun-save">
												<el-button type="primary" size="small" @click="sinteEditor">修改</el-button>
												<el-button size="small" @click="$router.back(-1)">返回</el-button>
											</div>
										</el-form-item>
									</el-row>
								</el-form>
							</div>
							<div class="page-content-box">
								<div class="addCon" v-if="!addCon">
									<i class="iconfont icon-zengjia" @click="dialogStatus = true"></i>
								</div>
								<div class="addCon1" v-if="addCon">
									<i class="iconfont icon-zengjia" @click="dialogStatus = true"></i>
								</div>
								<el-collapse accordion v-if="!addCon">
									<div v-for="(item,index) in funcManage" :key="index" @click="onclick(index)">
										<el-collapse-item>
											<template slot="title">
												<span class="span">{{item.funcName}}</span>
												<div class="btn">
													<el-button type="text" size="small">刷新 |</el-button>
													<el-button type="text" size="small">删除</el-button>
												</div>
											</template>
											<div class="coll-con">
												<div class="con-left">请求参数
													<div class="con-button">
														<el-button type="text" size="small" @click="dialogFormVisible = true">编辑</el-button>
													</div>
												</div>
												<div class="con-right">
													<div class="right-content">
														{{item.reqExample}}
													</div>
													<div class="page-table1">
														<template>
															<el-table :data="tableData" style="width: 100%" stripe>
																<el-table-column fixed prop="date" align="center" label="节点名称">
																</el-table-column>
																<el-table-column prop="name" align="center" label="父节点名称">
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
													<div class="con-button">
														<el-button type="text" size="small" @click="dialogFormVisible = true">编辑</el-button>
													</div>
												</div>
												<div class="con-right">
													<div class="right-content">
														{{item.resExample}}
													</div>
													<div class="page-table1">
														<template>
															<el-table :data="tableData" style="width: 100%" stripe>
																<el-table-column fixed prop="date" align="center" label="节点名称">
																</el-table-column>
																<el-table-column prop="name" align="center" label="父节点名称">
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
									</div>
								</el-collapse>
							</div>
						</div>
					</el-tab-pane>
					<el-tab-pane label="发布信息" name="second">
						<!-- 发布信息页面 -->
						<releaseInfo></releaseInfo>
					</el-tab-pane>
				</el-tabs>
			</template>
		</div>

		<!-- 新增弹框 -->
			<dialogAdd @changeDialog="changeDialog" :dialogStatus='dialogStatus'></dialogAdd>
	</div>
</template>
<script>
import { sinteManage, funcManageAll, fmtSubAll } from "../../axios/axios.js";
import releaseInfo from "./releaseInfo.vue";
import dialogAdd from "./dialogAdd.vue";
import { mapState } from "vuex";
export default {
  components: {
    releaseInfo,
    dialogAdd
  },
  data() {
    return {
      groupId: "",
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
      funcManage: "",
      tableData: [],
      reqFmtNub: [],
      resFmtNub: [],
      dialogStatus: "",
      addCon: false,
      activeName: "first",
      activeName1: "first"
    };
  },
  methods: {
    // 切换弹框显示
    changeDialog() {
      this.dialogStatus = false;
    },
    returnReginter() {
      this.$router.push({
        path: "/reginter"
      });
    },
    sinteEditor() {
      this.sinteData.groupId = this.groupId;
      this.sinteData.connConf = JSON.stringify(this.ConfData);
      let params = JSON.stringify({
        com: "PUT",
        data: this.sinteData
      });
      console.log(params);
      sinteManage(params).then(res => {
        console.log(res);
        if (res.code == 200000) {
          this.$message({
            message: "修改成功",
            type: "success"
          });
        } else {
          console.log(res.msg);
          this.$message({
            message: "修改失败",
            type: "error"
          });
        }
      });
    },
    // 渲染功能列表
    funcAll() {
      let params = {
        com: "all",
        groupId: this.sinteData.groupId,
        inteCode: this.sinteData.inteCode
      };
      let that = this;
      console.log(params);
      funcManageAll(params).then(res => {
        console.log(res);
        if (res.code == 200000) {
          this.funcManage = res.data;
          console.log(this.funcManage);
          if (this.funcManage == null) {
            this.addCon = true;
          } else {
            this.addCon = false;
          }
        } else {
          console.log(res.msg);
          this.$message({
            message: "新增失败",
            type: "error"
          });
        }
      });
    },
    onclick: function(index) {
      console.log(index);
      console.log("123");
    }
    // funcListTab() {
    //   let that = this;
    //   console.log(this.reqFmtNub, this.resFmtNub);
    //   this.reqFmtNub.forEach(function(obj) {
    //     console.log(obj);
    //     let params = {
    //       com: "search",
    //       groupId: that.sinteData.groupId,
    //       fmtCode: obj
    //     };
    //     fmtSubAll(params).then(res => {
    //       console.log(res);
    //     });
    //   });
    //   this.resFmtNub.forEach(function(obj1) {
    //     console.log(obj1);
    //     let params = {
    //       com: "search",
    //       groupId: that.sinteData.groupId,
    //       fmtCode: obj1
    //     };
    //     fmtSubAll(params).then(res => {
    //       console.log(res);
    //     });
    //   });
    // }
  },
  // 页面刚进来就加载
  mounted() {
    this.$nextTick(function() {
      this.sinteData = this.$route.params;
      this.groupId = this.sinteData.groupId;
      this.ConfData = JSON.parse(this.$route.params.connConf);
      this.funcAll();
    });
  }
};
</script>
<style scoped>
.pagetable-top {
  margin-left: 8px;
}

.nav_box > div {
  width: 100%;
}

.nav_box {
  width: 100%;
  height: 40px;
  line-height: 40px;
  background: #fff;
  text-align: left;
  display: flex;
}

.fun-save {
  margin: 0px 0px 0 -40px;
}
</style>