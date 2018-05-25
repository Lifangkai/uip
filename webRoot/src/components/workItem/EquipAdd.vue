<template>
    <div>
        <router-view></router-view>
        <div class="main_block block-height1">
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端SN：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端分类：</label>
                    <div class="add_form_item_con">
                        <el-select size="small" v-model="termType" placeholder="请选择">
                            <el-option
                            v-for="item in termType_options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">终端名称：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">说明：</label>
                    <div class="add_form_item_con">
                        <el-input
                        type="textarea"
                        :rows="2"
                        placeholder="请输入内容"
                        v-model="textarea">
                        </el-input>
                    </div>
                </div>   
            </el-col>
        </div>
        <div class="main_block block-height1">
            <div>
                <p>宿主信息</p>
                <hr style="height:1px;border:none;border-top:1px solid #EFF3F6;" />
                <el-col :xs="8" :sm="8" :md="12" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">实例名称：</label>
                        <div class="add_form_item_con">
                            <!-- <el-input size="small" placeholder="请输入"></el-input> -->
                            <el-autocomplete
                                size="small"
                                v-model="queryString"
                                :fetch-suggestions="querySearchAsync"
                                placeholder="请输入内容"
                                @select="handleSelect"></el-autocomplete>
                        </div>
                    </div>   
                </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><p>&nbsp;</p></el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">IP地址：</label>
                        <div class="add_form_item_con">
                            <el-input size="small" placeholder="请输入"></el-input>
                        </div>
                    </div>   
                </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                    <div class="add_form_item">
                        <label class="add_form_label">服务端口：</label>
                        <div class="add_form_item_con">
                            <el-input size="small" placeholder="请输入"></el-input>
                        </div>
                    </div>   
                </el-col>
            </div>
        </div>
        <div class="main_block block-height2">
                <p>驱动信息</p>
                <hr style="height:1px;border:none;border-top:1px solid #EFF3F6;" />
                 <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">驱动类型：</label>
                    <div class="add_form_item_con">
                        <el-select size="small" v-model="drivType" placeholder="请选择">
                            <el-option
                            v-for="item in drivType_options"
                            :key="item.value"
                            :label="item.label"
                            :value="item.value">
                            </el-option>
                        </el-select>
                    </div>
                </div>
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row">
                    <el-button type="primary" size="small">新增</el-button></div>
            </el-col>
            <el-table :data="tableData" height="200" border style="width: 95%">
                <el-table-column prop="termCode" label="参数名称" style="width: 20%"></el-table-column>
                <el-table-column prop="termName" label="参数编码" style="width: 25%"></el-table-column>
                <el-table-column prop="cmt" label="参数值" style="width: 25%"></el-table-column>
                <el-table-column label="操作" style="width: 25%">
                    <template slot-scope="scope">
                        <el-button
                        size="mini"
                        type="danger"
                        @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                        <el-button
                        size="mini"
                        @click="handleSave(scope.$index, scope.row)">保存</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-col :xs="16" :sm="16" :md="16" :lg="16" :xl="16"><p>&nbsp;</p></el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="8" :xl="8"><div class="grid-content main_row">
                <el-button type="primary" size="small" @click="equipAdd">保存</el-button></div>
            </el-col>
        </div>
    </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      termType_options: [
        {
          value: "选项1",
          label: "黄金糕"
        },
        {
          value: "选项2",
          label: "双皮奶"
        },
        {
          value: "选项3",
          label: "蚵仔煎"
        },
        {
          value: "选项4",
          label: "龙须面"
        },
        {
          value: "选项5",
          label: "北京烤鸭"
        }
      ],
      drivType_options: [
        {
          value: "驱动1",
          label: "驱动1"
        }
      ],
      drivType: "",
      termType: "",
      instances: [],
      // state4: '',
      timeout: null,
      value: "",
      textarea: "",
      tableData: [],
      queryString: ""
    };
  },
  methods: {
    //新增终端
    equipAdd() {
      axios.defaults.baseURL = "http://172.16.0.13:31425";
      axios
        .post("/terminal/manage", {
          com: "POST",
          data: {
            groupId: "12345678",
            termName: "终端名4",
            termType: "终端类型4",
            instCode: "sd3434002",
            drivType: "驱动4",
            drivParam: "",
            cmt: "说明4"
          }
        })
        .then(function(response) {
          console.log(response);
          if (response.data.code) {
            alert(response.data.msg);
          }
        })
        .catch(function(error) {
          console.log(error);
          if (response.data.code) {
            alert(response.data.msg);
          }
        });
    },
    loadAll() {
      return [];
    },
    querySearchAsync(queryString, cb) {
      console.log(this.queryString);
      let _this = this;
      axios.defaults.baseURL = "http://172.16.0.13:31425";
      axios
        .get("/instance/manage/ftsearch", {
          params: {
            groupId: "12345678",
            fttext: _this.queryString
          }
        })
        .then(function(response) {
          console.log(response);
          if (response.data.code === "200000") {
            // _this.tableData = response.data.data;
          } else {
          }
        })
        .catch(function(error) {
          console.log(error);
          // _this.tableData = [];
        });
    },
    handleSelect(item) {
      console.log(item);
    }
  },
  mounted() {
    this.instances = this.loadAll();
  }
};
</script>

<style>
.main_row {
  margin-top: 15px;
}
.block-height1 {
  height: 180px;
}
.block-height2 {
  height: 400px;
}
.add_form_item_con .el-autocomplete {
  width: 100%;
}
</style>