<<template>
<div>
<div class="uipt_block uipt_hight1">
 
      <el-col :xs="8" :sm="8" :md="8" :lg="13" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">实例编码：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
 <el-col :xs="8" :sm="8" :md="8" :lg="24" :xl="8"><div class="grid-content main_row"></div>
                
            </el-col>

               <el-col :xs="8" :sm="8" :md="8" :lg="6" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">实例名称：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
   <el-col :xs="8" :sm="8" :md="8" :lg="19" :xl="8"><div class="grid-content main_row"></div>
               
            </el-col>

               <el-col :xs="8" :sm="8" :md="8" :lg="13" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label size="mini" class="add_form_label">实例说明：</label>
                    <div class="add_form_item_con">
                          <el-input
                        type="textarea"
                        :rows="4"
                        placeholder="请输入内容"
                        v-model="textarea">
                        </el-input>
                    </div>
                </div>   
            </el-col>
<el-col :xs="8" :sm="8" :md="8" :lg="19" :xl="8"><div class="grid-content main_row"></div>
</el-col>
<el-col :xs="8" :sm="8" :md="8" :lg="6" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">宿主Ip地：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
            <el-col :xs="8" :sm="8" :md="8" :lg="19" :xl="8"><div class="grid-content main_row"></div>
</el-col>
<el-col :xs="8" :sm="8" :md="8" :lg="6" :xl="8"><div class="grid-content main_row"></div>
                <div class="add_form_item">
                    <label class="add_form_label">指令端口：</label>
                    <div class="add_form_item_con">
                        <el-input size="small" placeholder="请输入"></el-input>
                    </div>
                </div>   
            </el-col>
         
    </div>
          <div class="uipt_block uipt_hight2">
         <el-tabs v-model="editableTabsValue" type="card" editable @edit="handleTabsEdit">
  <el-tab-pane
    :key="item.name"
    v-for="(item, index) in editableTabs"
    :label="item.title"
    :name="item.name"
  >
    {{item.content}}
  </el-tab-pane>
</el-tabs>
      </div>
    <div class="uipt_block uipt_hight3">




<el-col :span="6" class="anniu">
 
  <el-button size="mini" type="primary">保存</el-button>
 
  <el-button size="mini" plain>返回</el-button>
</el-col>
  </div>
          
  </div>
   
</template>
<script>
  export default {
    data() {
      return {
        textarea:"",





        editableTabsValue: '2',
        editableTabs: [{
          title: 'Tab 1',
          name: '1',
          content: 'Tab 1 content'
        }, {
          title: 'Tab 2',
          name: '2',
          content: 'Tab 2 content'
        }],
        tabIndex: 2
      }
    },
    methods: {

      handleTabsEdit(targetName, action) {
        if (action === 'add') {
          let newTabName = ++this.tabIndex + '';
          this.editableTabs.push({
            title: 'New Tab',
            name: newTabName,
            content: 'New Tab content'
          });
          this.editableTabsValue = newTabName;
        }
        if (action === 'remove') {
          let tabs = this.editableTabs;
          let activeName = this.editableTabsValue;
          if (activeName === targetName) {
            tabs.forEach((tab, index) => {
              if (tab.name === targetName) {
                let nextTab = tabs[index + 1] || tabs[index - 1];
                if (nextTab) {
                  activeName = nextTab.name;
                }
              }
            });
          }
          
          this.editableTabsValue = activeName;
          this.editableTabs = tabs.filter(tab => tab.name !== targetName);
        }
      }
    },
    created(){



      this.textarea="2kjhkjhkjhkjh222";

    }
  }
</script>
<style scoped>
.uiptaddvue {
    height: 800px;
}
.uipt_block {
    padding: 10px;
    border: 10px solid #EFF3F6;
}
.uipt_hight1{
  height:360px;
}
.main_row{
    margin-top: 10px;
}
.uipt_hight3{
    height:30px;
}
.anniu{
  padding:5px;
}
.uipt_hight2{
  height:360px;

}

</style>

