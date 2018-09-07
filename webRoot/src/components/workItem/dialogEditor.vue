<template>
	<el-dialog :visible.sync="dialogFormVisible" @close="close">
		<template slot="title">
			<div class="dialog-top"> 功能格式配置 </div>
		</template>
		<div class="dialog-con">
			<!-- 选项卡切换 -->
			<el-form ref="funcData" :data="funcData" label-width="80px">
				<el-form-item label="功能名称:" prop='funcName'>
					<el-input placeholder="请输入" size="small" v-model="funcData.funcName"></el-input>
				</el-form-item>
				<el-row style="line-height:40px;margin:0 10px;">
					<el-col :span="12" style="text-align:left;">
						<label>请求参数样例:</label>
					</el-col>
					<el-col :span="12" style="text-align:right">
						<el-button size='small' type="primary" @click="resFmtcode()">请求参数配置</el-button>
					</el-col>
				</el-row>
				<el-form-item prop='reqExample' class="elcs">
					<el-input type="textarea" style='height: 90px;margin-bottom:6px' v-model="funcData.reqExample"></el-input>
				</el-form-item>
				<el-row style="line-height:40px;margin:0 10px;">
					<el-col :span="12" style="text-align:left;">
						<label>返回参数样例:</label>
					</el-col>
					<el-col :span="12" style="text-align:right">
						<el-button size='small' type="primary" @click="reqFmtcode()">返回参数配置</el-button>
					</el-col>
				</el-row>
				<el-form-item class="elcs" prop='reqExample'>
					<el-input type="textarea" style='height: 90px;margin-bottom:6px' v-model="funcData.resExample"></el-input>
				</el-form-item>
			</el-form>
			
			<!-- 内层弹框 -->
			<el-dialog :visible.sync="innerVisible" append-to-body>
				<template slot="title">
					<div class="dialog-top">参数配置</div>
				</template>
				<div class="dialog-con">
					<el-form label-width="80px">
						<el-form-item label="格式ID:">
							<el-input placeholder="请输入" size='small' disabled v-model="fmtCode"></el-input>
						</el-form-item>
						<el-form-item label="格式类型:">
							<el-select size="small" style="width:200px" v-model="fmtType" clearable>
								<el-option label="json-json" value="json-json"></el-option>
								<el-option label="file-file" value="file-file"></el-option>
								<el-option label="ws-web" value="ws-web"></el-option>
								<el-option label="service" value="service"></el-option>
							</el-select>
						</el-form-item>
					</el-form>
					<div class="dialog-footer" style='text-align:right'>
						<el-button type="primary" size='small' @click="mainSave()">修改</el-button>
						<el-button @click="innerVisible = false;" size='small'>取消</el-button>
					</div>
					<div class="add-div">
						<el-button type="primary" size='small' style="width:70px" @click="dialogRequest()">新增</el-button>
					</div>
					<div class="page-table">
						<template>
							<el-table :data="tableReq" style="width: 100%" stripe>
								<el-table-column prop="fieldName" align="center" label="节点名称">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.fieldName" size='small'></el-input>
										<span v-else>{{scope.row.fieldName}}</span>
									</template>
								</el-table-column>
								<el-table-column prop="extInfo" align="center" label="父节点名称">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.extInfo" size='small'></el-input>
										<span v-else>{{scope.row.extInfo}}</span>
									</template>
								</el-table-column>
								<el-table-column prop="dataType" align="center" label="数据类型">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.dataType" size='small'></el-input>
										<span v-else>{{scope.row.dataType}}</span>
									</template>
								</el-table-column>
								<el-table-column prop="length" align="center" label="最大长度">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.length" size='small'></el-input>
										<span v-else>{{scope.row.length}}</span>
									</template>
								</el-table-column>
								<el-table-column prop="constrain" align="center" label="约束">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.constrain" size='small'></el-input>
										<span v-else>{{scope.row.constrain}}</span>
									</template>
								</el-table-column>
								<el-table-column prop="cmt" align="center" label="说明">
									<template slot-scope="scope">
										<el-input v-if="scope.row.type||scope.row.xgSave" v-model="scope.row.cmt" size='small'></el-input>
										<span v-else>{{scope.row.cmt}}</span>
									</template>
								</el-table-column>
								<el-table-column label="操作" align="center">
									 <template slot-scope="scope">
										<el-button type="text" size="small" v-if="!scope.row.type&&!scope.row.xgSave" @click="editReq(scope.$index,tableReq)">修改 |</el-button>
										<el-button type="text" size="small" v-if="scope.row.xgSave" @click="reqRevise(scope.$index,tableReq)">修改保存</el-button>
										<el-button type="text" size="small" v-if="!scope.row.type&&!scope.row.xgSave" @click="reqDelete(scope.$index,tableReq)">删除</el-button>
										<el-button type="text" size="small" v-if="scope.row.type" style="color:#658F34;" @click="reqSave(scope.$index,tableReq)">保存 |</el-button>
										<el-button type="text" size="small" v-if="scope.row.type" style="color:red;" @click="cancelReq(scope.$index,tableReq)">取消</el-button>
									</template>
								</el-table-column>
							</el-table>
						</template>
					</div>
				</div>
			</el-dialog>
		</div>
		<div slot="footer" class="dialog-footer">
			<el-button type="primary" @click="dialogSave()" size='small'>修改</el-button>
			<el-button @click="dialogFormVisible = false;" size='small'>取消</el-button>
		</div>
	</el-dialog>
</template>

<script>
	import {
		funcManage,
		fmtManageSub,
		funcManageAll,
		fmtManageMain,
		fmtSubAll
	} from "../../axios/axios.js";
	
	import global_ from '../Global.vue';
	
	export default {
		props: ["func", "dialogStat"],
		
		data() {
			return {
				groupId: global_.GroupId,
				operCode: "8542",
				funcData: {
					funcName: "",
					reqExample: "",
					resExample: "",
					reqDataProp: "",
					resDataProp: "",
					cmt: ""
				},
				resCode: {},
				tableReq: [],
				activeName: "first",
				innerVisible: false,
				dialogFormVisible: false,
				xgSave: false,
				fmtCode: "",
				fmtType: ""
			};
		},

		methods: {
			// 关闭弹框
			close() {
				this.$emit("chgeDialog", false);
				this.$refs["funcData"].resetFields();
			},
			
			//  请求参数编辑
			editReq(index, rows) {
				rows[index].xgSave = true;
			},
			
			//  请求参数取消
			cancelReq(index, rows) {
				this.tableReq.splice(index, 1);
			},
			
			// 请求参数id
			resFmtcode() {
				this.fmtCode = this.funcData.reqFmtCode;
				this.innerVisible = true;
				this.funcAllsnb();
			},
			
			// 返回参数id
			reqFmtcode() {
				this.fmtCode = this.funcData.resFmtCode;
				this.innerVisible = true;
				this.funcAllsnb();
			},
			
			returnReginter() {
				this.$router.push({
					path: "/reginter"
				});
			},
			
			// 弹框保存
			dialogSave() {
				let params = JSON.stringify({
					com: "PUT",
					data: this.funcData
				});
				console.log(params);
				let that = this;
				funcManage(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
						this.$message({
							message: "修改成功",
							type: "success"
						});
						this.dialogFormVisible = false;
					} else {
						console.log(res.msg);
						this.$message({
							message: "修改失败",
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
				console.log(this.resCode);
			},
			
			// 参数保存每一行
			reqSave(index, rows) {
			  console.log("in reqSave() groupId = " + this.groupId);
				let params = JSON.stringify({
					com: "POST",
					data: {
						groupId: this.groupId,
						fmtCode: this.fmtCode,
						operCode: this.operCode,
						constrain: rows[index].constrain,
						dataType: rows[index].dataType,
						fieldName: rows[index].fieldName,
						length: rows[index].length,
						extInfo: rows[index].extInfo,
						cmt: rows[index].cmt
					}
				});

				fmtManageSub(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
						this.funcAllsnb();
					} else {
						this.$message({
							message: "增加失败",
							type: "error",
							center: true
						});
					}
				});
			},
			
			// 参数修改每一行
			reqRevise(index, rows) {
				let params = JSON.stringify({
					com: "PUT",
					data: {
						groupId: rows[index].groupId,
						fmtCode: rows[index].fmtCode,
						operCode: rows[index].operCode,
						dtlCode: rows[index].dtlCode,
						constrain: rows[index].constrain,
						dataType: rows[index].dataType,
						fieldName: rows[index].fieldName,
						length: rows[index].length,
						extInfo: rows[index].extInfo,
						cmt: rows[index].cmt
					}
				});
				fmtManageSub(params).then(res => {
					if(res.code == 200000) {
						this.funcAllsnb();
					} else {
						this.$message({
							message: "修改失败",
							type: "error",
							center: true
						});
					}
				});
			},
			
			// 参数删除每一行
			reqDelete(index, rows) {
				console.log(rows);
				let that = this;
				let params = JSON.stringify({
					com: "DELETE",
					data: {
						groupId: rows[index].groupId,
						fmtCode: rows[index].fmtCode,
						dtlCode: rows[index].dtlCode
					}
				});
				fmtManageSub(params).then(res => {
					if(res.code == 200000) {
						that.tableReq.splice(index, 1);
					} else {
						console.log(res.msg);
					}
				});
			},
			
			// 数据格式主表修改后保存到数据库
			mainSave() {
				let params = {
					com: "POST",
					data: {
						groupId: this.groupId,
						fmtCode: this.fmtCode,
						fmtType: this.fmtType,
						operCode: this.operCode
					}
				};
				console.log(params);
				fmtManageMain(params).then(res => {
					console.log(res);
					if(res.code == 200000) {
						this.$message({
							message: "修改保存成功",
							type: "success"
						});
					} else {
						console.log(res.msg);
						this.$message({
							message: "修改保存失败",
							type: "error"
						});
					}
				});
			},
			
			// 渲染功能里的数据格式列表
			funcAllsnb() {
				let params = {
					com: "search",
					groupId: this.funcData.groupId,
					fmtCode: this.fmtCode
				};
				let that = this;
				fmtSubAll(params).then(res => {
					if(res.code == 200000) {
						console.log(res);
						  res.data.fields.forEach(function(c){
							c.xgSave = false;
						  })
						that.tableReq = res.data.fields;
						that.fmtType = res.data.fmtType;
					} else {
						console.log(res.msg);
						this.tableReq = [];
					}
				});
			},
		},
		
		// 监听弹框显示隐藏
		watch: {
			dialogStat() {
				this.dialogFormVisible = this.dialogStat;
				this.funcData = this.func;
			}
		},
	};
</script>

<style scoped>
	.saveCss {
		position: absolute;
		right: 80px;
		bottom: 15px;
	}
</style>
