<template>
	<div class="page-box">
		<div class="pagecon-box">
			<div class="form-serch">
				<div>
					<span class="demonstration">接口查询:</span>
					<el-input placeholder="请输入" size="small" style="width: 200px;"></el-input>
				</div>
				<div class="pagetable-top">
					<el-button type="primary" size="small">查询</el-button>
					<el-button size="small" @click="addOpeninter">新增</el-button>
				</div>
			</div>
			<div class="page-table-box">
				<!-- 固定列表格 -->
				<div class="page-table">
					<template>
						<el-table :data="tableData" style="width: 100%" stripe>
							<el-table-column prop="date" align="center" label="接口ID">
							</el-table-column>
							<el-table-column prop="name" align="center" label="操作名称">
							</el-table-column>
							<el-table-column prop="name" align="center" label="接口说明">
							</el-table-column>
							<el-table-column label="操作" align="center">
								<template slot-scope="scope">
									<el-button type="text" size="small" @click="editOpeninter">编辑 |</el-button>
									<el-button type="text" size="small">发布 |</el-button>
									<el-button type="text" size="small">删除</el-button>
								</template>
							</el-table-column>
						</el-table>
					</template>
				</div>
				<!-- <div class="paginat-el">
          <el-pagination
            background
            layout="prev, pager, next"
            :total="1000">
          </el-pagination>
        </div>   -->
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		data() {
			return {
				tableData: [{
						date: "2016-05-02",
						name: "王小虎",
						address: "上海市普陀区金沙江路 1518 弄"
					},
					{
						date: "2016-05-04",
						name: "王小虎",
						address: "上海市普陀区金沙江路 1517 弄"
					}
				],
				
				formInline: {
					user: "",
					region: ""
				}
			};
		},
		
		methods: {
			addOpeninter() {
				this.$router.push({
					path: "/reginterAdd"
				});
			},
			
			editOpeninter() {
				this.$router.push({
					path: "/reginterEditor"
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
	
	.pagecon-box {
		border-radius: 6px;
		height: 664px;
		background: white;
	}
	
	.page-table {
		border: 1px solid #dde2e3;
		border-radius: 6px;
		margin: 10px 20px 20px 20px;
		min-height: 520px;
	}
</style>