<template>
  <div class="uipt_block uiptaddvue">
      <router-view></router-view> 
           <el-col :xs="8" :sm="8" :md="8" :lg="9" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">实例名称：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入" v-model="input3"></el-input>
                    </div>
                </div>   
            </el-col>
             <el-col :xs="8" :sm="8" :md="8" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">实例说明：</label>
                    <div class="add_form_item_con">
                         <el-input
                        type="textarea"
                        :rows="5"
                        placeholder="请输入内容"
                        v-model="textarea">
                        </el-input>
                    </div>
                </div>   
            </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">宿主IP地：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入" v-model="input1"></el-input>
                    </div>
                </div>   
            </el-col>
                <el-col :xs="8" :sm="8" :md="8" :lg="16" :xl="16"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">指令端口：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入" v-model="input2"></el-input>
                    </div>
                </div>   
            </el-col>
          <el-col :xs="8" :sm="8" :md="8" :lg="9" :xl="16"><div class="grid-content main_row"></div>
                  <el-button type="primary" @click="sub">保存</el-button>
                <el-button  @click="uiptBack()">返回</el-button>
            </el-col>

 
      </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      textarea: "",
      input1:"",
      input2:"",
      input3:""
    };
  },
  methods: {
    uiptBack() {
      this.$router.push({ path: "/uipt" });
    },
    sub(){
  var that=this;
  console.log(this.input1,this.input2,this.input3)


 axios.defaults.baseURL = "http://172.16.0.13:31425";
      axios
        .post("/instance/manage", {
          com:"POST",
          data: {
            "group_id": "YYGJUIP1",
            "inst_name": this.input3,
            "ip":  this.input1,
            "port": this.input2,
            "cmt": this.textarea
          }
        })
        .then(function(response) {
          console.log(response);
          if (response.data.code) {
            alert(response.data.msg);
          }
          that.$router.push({ path: "/uipt" });
        })
        .catch(function(error) {
          alert("新增失败");
          console.log(error);
        });



    }
  }
};
</script>
<style scoped>
.uiptaddvue {
  height: 500px;
}
.main_row {
  margin-top: 15px;
}
.uipt_block {
  padding: 10px;
  border: 10px solid #eff3f6;
}
</style>
