<template>
	<div class="page-box">
		<div class="nav_box">
			<template>
				<el-tabs v-model="activeName1">
					<el-tab-pane label="登录接口信息" name="first">
						<div style="margin:0px 10px" class="reginter-con">
							<div class="pagecon-box">
								<el-form ref="sinteData" :data="sinteData" label-width="80px">
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
									<i class="iconfont icon-zengjia" @click="funcSns()"></i>
								</div>
								<div class="addCon1" v-if="addCon">
									<i class="iconfont icon-zengjia" @click="funcSns()"></i>
								</div>
								<div v-for="(item,index) in funcManage" :key="index">
									<el-collapse accordion v-if="!addCon" @change="changeColl(item,index)">
										<el-collapse-item name='1'>
											<template slot="title">
												<span class="span">{{item.funcName}}</span>
												<div class="btn">
													<el-button type="text" size="small" @click = 'funcBj(item,index)'>编辑 |</el-button>
													<el-button type="text" size="small" @click="funcDel(item,index)">删除</el-button>
												</div>
											</template>
											<div class="coll-con">
												<div class="con-left">请求参数
												</div>
												<div class="con-right">
													<div class="right-content">
														{{item.reqExample}}
													</div>
													<div class="page-table1">
														<template>
															<el-table :data="tableReq" style="width: 100%" stripe>
																<el-table-column fixed prop="fieldName" align="center" label="节点名称">
																</el-table-column>
																<el-table-column prop="extInfo" align="center" label="父节点名称">
																</el-table-column>
																<el-table-column prop="dataType" align="center" label="数据类型">
																</el-table-column>
																<el-table-column prop="length" align="center" label="最大长度">
																</el-table-column>
																<el-table-column prop="constrain" align="center" label="约束">
																</el-table-column>
																<el-table-column prop="cmt" align="center" label="说明">
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
														{{item.resExample}}
													</div>
													<div class="page-table1">
														<template>
															<el-table :data="tableRes" style="width: 100%" stripe>
																<el-table-column fixed prop="fieldName" align="center" label="节点名称">
																</el-table-column>
																<el-table-column prop="extInfo" align="center" label="父节点名称">
																</el-table-column>
																<el-table-column prop="dataType" align="center" label="数据类型">
																</el-table-column>
																<el-table-column prop="length" align="center" label="最大长度">
																</el-table-column>
																<el-table-column prop="constrain" align="center" label="约束">
																</el-table-column>
																<el-table-column prop="cmt" align="center" label="说明">
																</el-table-column>
															</el-table>
														</template>
													</div>
												</div>
											</div>
										</el-collapse-item>
									</el-collapse>
								</div>
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
		<dialogAdd @changeDialog="changeDialog" v-on:funcCon="list" :dialogStatus='dialogStatus' v-bind:inteSns='snsCode' v-bind:inte='inteCode'></dialogAdd>
	  <!-- 编辑弹框 -->
    <dialogEditor :dialogStat='dialogStat' @chgeDialog="chgeDialog"  v-bind:func='func'></dialogEditor>
  </div>
</template>
<script>
	import { sinteManage, funcManageAll, fmtSubAll, funcManage, sinteSns } from "../../axios/axios.js";
	import releaseInfo from "./releaseInfo.vue";
	import dialogAdd from "./dialogAdd.vue";
  import dialogEditor from "./dialogEditor.vue";
	import { mapState } from "vuex";
	export default {
		components: {
			releaseInfo,
			dialogAdd,
      dialogEditor
		},
		data() {
			return {
				groupId: "",
				sinteData: '',
				ConfData: '',
				funcManage: "",
				tableRes: [],
				tableReq: [],
				reqFmtNub: [],
				resFmtNub: [],
				dialogStatus: "",
        dialogStat: "",
				addCon: false,
				activeName1: "first",
				snsCode: "",
				inteCode: "",
        func:''
			};
		},
		methods: {
			// 切换弹框显示
			changeDialog() {
				this.dialogStatus = false;
			},
      chgeDialog() {
				this.dialogStat = false;
			},
			returnReginter() {
				this.$router.push({
					path: "/reginter"
				});
			},
			funcSns() {
				let params = {
					com: "LIST",
					groupId: "SNS16494",
					snId: "UIPFMT01",
					operatorId: "12000",
					bs: "2"
				};
				sinteSns(params).then(res => {
					console.log(res);
					if(res.Code == 200000) {
						this.snsCode = res.Data;
						this.dialogStatus = true;
					} else {
						console.log(res.Msg);
					}
				});
			},
			// 修改登记页面
			sinteEditor() {
				this.sinteData.connConf = JSON.stringify(this.ConfData);
				let params = JSON.stringify({
					com: "PUT",
					data: this.sinteData
				});
				console.log(params);
				sinteManage(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
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
				funcManageAll(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
						this.funcManage = res.data;
						console.log(this.funcManage);
						if(this.funcManage == null) {
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
			// 渲染数据格式
			changeColl(item, index) {
				let that = this;
				let params = {
					com: "search",
					groupId: that.groupId,
					fmtCode: item.reqFmtCode
				};
				fmtSubAll(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
						that.tableReq = res.data.fields;
					} else {
						console.log(res.msg);
						that.tableReq = [];
					}
					let param = {
						com: "search",
						groupId: that.groupId,
						fmtCode: item.resFmtCode
					};
					fmtSubAll(param).then(res => {
						if(res.code == 200000) {
							that.tableRes = res.data.fields;
						} else {
							console.log(res.msg);
							that.tableRes = [];
						}
					});
				});
			},
      // func编辑
      funcBj(item, index){
        this.func = item;
  	    this.dialogStat = true;
      },
			// func表删除
			funcDel(item, index) {
				console.log(item);
				let params = JSON.stringify({
					com: "DELETE",
					data: {
						groupId: item.groupId,
						inteCode: item.inteCode,
						funcCode: item.funcCode
					}
				});
				console.log(params);
				funcManage(params).then(res => {
					if(res.code == 200000) {
						this.funcAll();
					} else {
						console.log(res.msg);
					}
				});
			},
			list: function(somedata) {
				this.funcManage = somedata;
				if(this.funcManage == null) {
					this.addCon = true;
				} else {
					this.addCon = false;
				}
			}
		},
		// 页面刚进来就加载
		mounted() {
			this.$nextTick(function() {
				this.sinteData = this.$route.params;
				this.groupId = this.sinteData.groupId;
				this.inteCode = this.sinteData.inteCode;
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
	
	.nav_box>div {
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