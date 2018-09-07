<template>
	<div class="page-box">
		<div style="margin:0px 10px" class="reginter-con">
			<div style="margin-bottom: 6px;" v-if='!publishCon'>
				<el-button size="small" @click='pubBut'>新增</el-button>
		  </div>
		  
			<el-tabs v-if='publishCon' v-model="activeName" type="card" editable  @tab-add="handleTabsEdit">
				<el-tab-pane name="first" >
				  <span slot="label">{{formInline.region}}</span>
					<div class="pagecon-box">
						<el-form ref="form" :model="ConfData" label-width="80px">
							<el-form-item label="接口名称:" style="margin-left:20px">
								<el-input placeholder="请输入" size="small" style="width:380px"></el-input>
							</el-form-item>
							<el-form-item label="接口ID:" style="margin-left:20px">
								<el-input placeholder="请输入" size="small" style="width:914px"></el-input>
							</el-form-item>
							<el-form-item label="接口说明:" style="margin-left:20px;margin-bottom: 0px;">
								<el-input type="textarea" style="width: 914px;"></el-input>
							</el-form-item>
							<el-row>
								<el-col :span="5">
									<el-form-item label="接口协议:" prop='inteType' style="margin:6px 0px 0px 20px;">
										<el-select v-model="ConfData.ip" size="small">
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
						</el-form>
					</div>
					<div class="page-botton">
						<el-button size="small" @click="$router.back(-1)">返回</el-button>
						<el-button type="primary" size="small">保存</el-button>
						<el-button type="danger" size="small">删除</el-button>
					</div>
					
					<div class="page-content-box">
						<div class="addCon" v-if="!addCon">
									<i class="iconfont icon-zengjia" @click="dialogStatus = true"></i>
								</div>
						<div class="addCon1" v-if="addCon">
								<i class="iconfont icon-zengjia" @click="dialogStatus = true"></i>
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
										<div class="con-button">
											<el-button type="text" size="small">编辑</el-button>
										</div>
									</div>
									<div class="con-right">
										<div class="right-content">请求参数	</div>
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
											<el-button type="text" size="small">编辑</el-button>
										</div>
									</div>
									<div class="con-right">
										<div class="right-content">返回参数	</div>
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
						</el-collapse>
					</div>
				</el-tab-pane>
			</el-tabs>
		</div>
		
		<!-- 新增开放协议 -->
		<el-dialog :visible.sync="dialogFormVisible" width="400px">
        <template slot="title">
          <div class="dialog-top">新增开放协议</div>
        </template>
        
        <div class="dialog-info">
           <el-form :model="formInline">
							<el-form-item label="开放协议">
								<el-select v-model="formInline.region" placeholder="请输入" size='small'>
									<el-option label="http" value="http"></el-option>
									<el-option label="flie" value="flie"></el-option>
								</el-select>
							</el-form-item>
					</el-form>
        </div>
        
				<div slot="footer" class="dialog-footer">
					<el-button type="primary" size='small' @click="dialogSelect">保存</el-button>
					<el-button @click="dialogFormVisible = false;" size='small'>取消</el-button>
				</div>
				
      </el-dialog>
	</div>
</template>

<script>
import dialogAdd from "./dialogAdd.vue";
import { mapState } from "vuex";
export default {
  components: {
    dialogAdd
},

data() {
 return {
  tableData: [],
  ConfData: {
    ip: "",
    port: "",
    path: ""
  },
  dialogStatus: "",
  dialogFormVisible: false,
  activeName: "first",
  addCon: false,
	publishCon:false,
	formInline: {
	  user: '',
	  region: ''
  }
  };
},

methods: {
  returnReginter() {
      this.$router.push({
        path: "/reginter"
      });
    },
		// 新增弹框切换
		pubBut(){
      this.dialogFormVisible = true;
		},
		
		handleTabsEdit(){
     this.dialogFormVisible = true;
		},
		
		// 点击保存 标签页显示
		dialogSelect(){
     this.publishCon = true;
		 this.dialogFormVisible = false;
		}
}
};
</script>

<style scoped>
.pagetable-top {
  margin-left: 8px;
}

.page-box {
  /* padding: 20px; */
}

.pagecon-box {
  height: 260px;
  border: none;
}

.reginter-con {
  background: white;
  padding: 20px;
  border-radius: 4px;
}
</style>