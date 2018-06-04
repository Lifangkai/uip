<template>
 <section>
    <div class="page-content">
         <div class="page-form">
           <el-form :inline="true" :model="formInline" class="demo-form-inline">
              <el-form-item label="接口查询">
                <el-input v-model="formInline.user" placeholder="请输入"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary">查询</el-button>
                 <el-button>新增</el-button>
              </el-form-item>
            </el-form>
        </div>
        <div class="page-table-box">
            <!-- 固定列表格 -->
            <template>
            <el-table :data="tableData" style="width: 100%">
                <el-table-column fixed prop="date"  align="center" label="接口ID">
                </el-table-column>
                <el-table-column  prop="name" align="center" label="操作名称">
                </el-table-column>
                <el-table-column prop="name" align="right" label="接口说明">
                </el-table-column>
                <el-table-column  label="操作" align="center" min-width="120">
                  <template slot-scope="scope">
                    <el-button type="text" size="small">编辑 |</el-button>
                    <el-button type="text" size="small">发布 |</el-button>
                    <el-button type="text" size="small">删除</el-button>
                  </template>
                </el-table-column>
          </el-table>
            </template>
        </div>  
    </div>      
 </section>
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
      formInline: {
        user: "",
        region: ""
      }
    };
  }
};
</script>
<style scoped>
.el-table th {
  text-align: center !important;
}
.fl {
  float: left;
}
.page-content {
  border: 10px solid #e0e5e6;
  background: white;
}
</style>