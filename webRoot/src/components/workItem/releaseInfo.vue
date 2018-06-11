<template>
	<div class="page-box">
		<div style="margin:0px 10px" class="reginter-con">
            <el-tabs v-model="activeName" type="card">
                <el-tab-pane label="http" name="first">
                    <div class="pagecon-box">
                        <el-form ref="form" :model="formInline" label-width="80px">
                              <el-form-item label="接口名称:" style="margin-left:20px">
                                <el-input placeholder="请输入" size="small" style="width:380px"></el-input>
                              </el-form-item>
                              <el-form-item label="接口ID:" style="margin-left:20px">
                                <el-input placeholder="请输入" size="small" style="width:914px" ></el-input>
                              </el-form-item>
                            <el-form-item label="接口说明:" style="margin-left:20px;margin-bottom: 0px;">
                            <el-input type="textarea" style="width: 914px;"></el-input>
                            </el-form-item>
                            <div class="form-serch">
                            <div>
                                <span class="demonstration">接口协议:</span>
                                <el-select v-model="formInline.region" placeholder="请输入" size="small">
                                <el-option label="http" value="http"></el-option>
                                </el-select>
                            </div>
                            <div class="seach-con">
                                <span class="demonstration"></span>
                                <el-input placeholder="请输入" size="small" style="width: 160px;"></el-input>
                            </div>
                             <div class="seach-con">
                                <span class="demonstration"></span>
                                <el-input placeholder="请输入" size="small" style="width: 160px;"></el-input>
                            </div>
                             <div class="seach-con">
                                <span class="demonstration"></span>
                                <el-input placeholder="请输入" size="small" style="width: 160px;"></el-input>
                            </div>
                            </div>
                        </el-form>
                    </div>      
                    <div class="page-content-box">
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
                                    <div class="con-button"><el-button type="text" size="small">编辑</el-button></div>
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
                                    <div class="con-button"><el-button type="text" size="small">编辑</el-button></div>
                                </div>
                                <div class="con-right">
                                    <div class="right-content">
                                    返回参数
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
                        </el-collapse-item>
                        </el-collapse>
                        <div class="addCon">
                            <i class="iconfont icon-zengjia"  @click="dialogFormVisible = true"></i>
                        </div>
                    </div>
                    <div class="page-botton">
                            <el-button size="small" @click="$router.back(-1)">返回</el-button>
                            <el-button type="primary" size="small">保存</el-button>
                            <el-button type="danger" size="small">删除</el-button>
                        </div>
                </el-tab-pane>
                <el-tab-pane label="file" name="second">
                    23456
                </el-tab-pane>
            </el-tabs>
                
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
            <el-form ref="form" :model="formInline" label-width="80px">
               <el-form-item label="功能名称:">
                   	<el-input placeholder="请输入" size="small"></el-input>
                  </el-form-item>
                  <el-tabs v-model="activeName" type="card">
              <el-tab-pane label="请求参数" name="first">
                  <el-form-item label="样例:">
                    <el-input type="textarea" style='height: 90px;margin-bottom:6px'></el-input>
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
                   <el-button type="primary" size='small' style="width:70px">新增</el-button>   
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
                      <el-input type="textarea" style='height: 90px;margin-bottom:6px'></el-input>
                    </el-form-item>
                    <el-form-item label="分隔符:">
                      <el-input placeholder="请输入" size="small" style="width:120px"></el-input>
                    </el-form-item>
                  <div class="add-div">
                    <el-button type="primary" size='small' style="width:70px">新增</el-button>   
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
            <el-button type="primary" @click="dialogFormVisible = false" size='small'>保存</el-button>            
            <el-button @click="dialogFormVisible = false" size='small'>取消</el-button>
          </div>
        </el-dialog>
	</div>
</template>
<script>
export default {
  data() {
    return {
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