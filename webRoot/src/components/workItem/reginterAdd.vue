<template>
	<div class="page-box">
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
							<el-select v-model="sinteData.inteType " size="small" clearable>
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
							<el-button type="primary" size="small" @click="sinteManage()">保存</el-button>
							<el-button size="small" @click="$router.back(-1)">返回</el-button>
						</div>
					</el-form-item>
				</el-row>
			</el-form>
		</div>
		<div class="page-content-box">
			<div class="addCon" v-if="addCon">
				<i class="iconfont icon-zengjia" @click="sinteSns()"></i>
			</div>
			<div class="addCon1" v-if="!addCon">
				<i class="iconfont icon-zengjia" @click="sinteSns()"></i>
			</div>
			<div v-for="(item,index) in childList" :key="index">
				<el-collapse accordion v-if="addCon" @change="changeColl(item,index)">
					<el-collapse-item>
						<template slot="title">
							<span class="span">{{item.funcName}}</span>
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

		<!-- 新增弹框 -->
		<dialogAdd v-bind:inte='inteCode' v-bind:inteSns='snsCode' v-on:funcCon="list" @changeDialog="changeDialog" :dialogStatus='dialogStatus'></dialogAdd>
	</div>
</template>
<script>
	import {
		funcManage,
		sinteSns,
		sinteManage,
		fmtManageSub,
		fmtSubAll
	} from "../../axios/axios.js";
	import dialogAdd from "./dialogAdd.vue";
	import { mapState } from "vuex";
	export default {
		components: {
			dialogAdd
		},
		data() {
			return {
				groupId: "SNS16494",
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
				tableRes: [],
				tableReq: [],
				inteCode: "",
				addCon: false,
				activeName: "first",
				childList: '',
				dialogStatus: "",
				snsCode: ""
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
					if(res.code == 200000) {
						this.inteCode = res.data.inteCode;
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
			// 点击产生sns
			sinteSns() {
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
			list: function(somedata) {
				this.childList = somedata;
				console.log(this.childList);
				if(this.childList) {
					this.addCon = true;
				} else {
					this.addCon = false;
				}
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