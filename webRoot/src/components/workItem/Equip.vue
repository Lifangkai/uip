<template>
      <div class="main_div main_block">
        <!-- <router-view></router-view>  -->
        <div class="main_header">
            <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6"><div class="grid-content">
                <div class="add_form_item">
                    <label class="add_form_label">终端查询：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入内容" v-model="input"></el-input>
                    </div></div>
                </div>   
            </el-col>
            <el-col :xs="12" :sm="12" :md="6" :lg="6" :xl="6"><div class="grid-content">
            <!-- <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content"> -->
                <el-button size="small" type="primary" @click="queryAllEquip">查询</el-button>
                <el-button size="small" @click="gotoAddEquip()">新增</el-button>
                </div>   
            </el-col>
        </div>
        <div class="main_table">
            <el-table :data="tableData" height="500" border style="width: 95%">
                <el-table-column prop="termCode" label="终端ID" style="width: 20%"></el-table-column>
                <el-table-column prop="termName" label="终端名称" style="width: 25%"></el-table-column>
                <el-table-column prop="cmt" label="终端说明" style="width: 25%"></el-table-column>
                <el-table-column label="操作" style="width: 25%">
                <template slot-scope="scope">
                  <el-button
                    size="mini"
                    @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                  <el-button
                    size="mini"
                    type="danger"
                    @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                </template>
              </el-table-column>
            </el-table>
        </div>  
        </div>
</template>
<script>
import axios from 'axios'

export default {
    data() {
        return {
            input: '',//查询内容
            tableData:[] //表格数据
        }
    },
    methods:{
        gotoAddEquip(){
            this.$router.push({ path :'/equipAdd' })
        },
        //查询-全文检索
        queryAllEquip(){
            let _this = this;
            axios.defaults.baseURL = "http://172.16.0.13:31425"
            axios.get('/terminal/manage/ftsearch', {
                params: {
                    groupId: '12345678',
                    fttext: this.input          
                }
            })
            .then(function (response) {
                console.log(response);
                if (response.data.code === "200000") {
                    _this.tableData = response.data.data;
                } else {
                    _this.tableData = [];
                    alert(response.data.msg);
                }
            })
            .catch(function (error) {
                console.log(error);
                _this.tableData = [];
            });
        },
        //删除终端
        handleDelete(index, row) {
            console.log(index);
            console.log(row);
            axios.defaults.baseURL = "http://172.16.0.13:31425"
            axios.post('/terminal/manage', {
                com:'DELETE',
                data: {
                groupId: '12345678',
                termCode: row.termCode
                }          
            })
            .then(function (response) {
                console.log(response);
                if (response.data.code) {
                    alert(response.data.msg);
                }
            })
            .catch(function (error) {
                alert('删除失败');
                console.log(error);
            });
        }
        //点击图标刷新页面
        // reload(){
        //     window.location.reload();
        // }
  }
}
</script>
<style>
.equip_blcok {
    height: 550px;
}
.main_block {
    padding: 10px;
    border: 10px solid #EFF3F6;
}
.main_header {
    margin-top:20px;
    margin-bottom:20px;
}
.main_table {
    margin-top:20px;
}
.main_div { 
    /* margin-left: 10px; */
    background: white;
}
/* 输入框 */
.add_form_item{
	/* 
	height: 28px;
	line-height: 28px; */
	/* margin: 6px 0; */
  /* font-size: 1.4rem; */
	padding: 0;
	display: flex;
}
.add_form_label{
	/* font-weight: normal;
	height: 28px;
	line-height: 28px; */
    width: 90px;
	margin: 6px 0;
	padding-right: 15px;
	text-align: right;
}
.add_form_item_con{
	flex: 1;
}

</style>
