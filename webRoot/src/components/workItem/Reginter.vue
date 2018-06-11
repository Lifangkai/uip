<template>
  <div class="page-box">
    <div class="pagecon-box">
       <div class="form-serch">
          <div>
              <span class="demonstration">接口查询:</span>
              <el-input placeholder="请输入" size="small" style="width: 200px;" v-model="searchKey"></el-input>
          </div>
          <div class="pagetable-top">
              <el-button type="primary" size="small"  @click="regSearch">查询</el-button>
              <el-button size="small" @click="addReginter">新增</el-button>
          </div>
        </div>
      <div class="page-table-box">
        <!-- 固定列表格 -->
        <div class="page-table">
          <template>
           <el-table :data="tableData" style="width: 100%" stripe>
                <el-table-column prop="inteCode"  align="center" label="接口ID">
                </el-table-column>
                <el-table-column  prop="inteName" align="center" label="操作名称">
                </el-table-column>
                <el-table-column  prop="cmt" align="center" label="接口说明">
                </el-table-column>
                <el-table-column  prop="inteCode" label="操作" align="center">
                  <template slot-scope="scope">
                    <el-button type="text" size="small" @click="editReginter">编辑 |</el-button>
                    <el-button type="text" size="small">发布 |</el-button>
                    <el-button type="text" size="small">删除</el-button>
                  </template>
                </el-table-column>
          </el-table>
        </template>
        </div>
        <div class="paginat-el">
          <el-pagination
            background
            layout="prev, pager, next"
            :total="1000">
          </el-pagination>
        </div>  
      </div>
    </div>
  </div>
</template>
<script>
import { sinteSearch } from "../../axios/axios.js";
export default {
  data() {
    return {
      groupId: "YYGJUIP1",
      searchKey: "",
      tableData: [],
      formInline: {
        user: "",
        region: ""
      }
    };
  },
  methods: {
    addReginter() {
      this.$router.push({ path: "/reginterAdd" });
    },
    editReginter() {
      this.$router.push({ path: "/reginterEditor" });
    },
    // 查询信息
    regSearch() {
      let params = {
        com: "key",
        groupId: this.groupId,
        fettext: this.searchKey
      };
      sinteSearch(params).then(res => {
        console.log(res);
        if (res.code == 200000) {
          this.tableData = res.data;
          this.searchKey = '';
          console.log(this.tableData);
        } else {
          console.log(res.msg);
          this.$message({
            message: "新增失败",
            type: "error"
          });
        }
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