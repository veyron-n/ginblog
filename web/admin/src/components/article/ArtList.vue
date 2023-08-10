<template>
  <div>
    <div>
      <div class='input-container'>
        <el-input style='width: 200px' size='default' v-model='username' placeholder='请输入文章名' clearable
                  @clear='clearInput'>
          <template #prefix>
            <el-icon class='el-icon--left'>
              <Search />
            </el-icon>
          </template>
        </el-input>
        <el-button style='margin-left: 5px' size='default' type='primary' @click='search'>搜索</el-button>
        <el-button style='margin-left: 5px' size='default' type='success' @click='dialogFormVisible = true'>添加文章
        </el-button>
        <el-dialog style='width: 800px; height: 600px' v-model='dialogFormVisible' title='添加文章'>
          <el-form ref="openUserForm" :model="dialog.openArt.form" :rules="dialog.openArt.formRules"
                   label-width="100px">
            <el-form-item label='文章标题' prop='title'>
              <el-input style='width: 300px' placeholder='请输入文章标题' maxlength='50' show-word-limit v-model='dialog.openArt.form.account' clearable>
                <template #prefix>
                  <el-icon class="el-icon--left">
                    <User />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label='文章描述' prop='desc'>
              <el-input style='width: 300px' placeholder='请输入文章描述' maxlength='100' show-word-limit v-model='dialog.openArt.form.desc' clearable>
                <template #prefix>
                  <el-icon class="el-icon--left">
                    <Lock />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label='文章正文' prop='context'>
              <el-input style='width: 300px' placeholder='请输入文章正文' maxlength='1000' show-word-limit v-model='dialog.openArt.form.context' clearable>
                <template #prefix>
                  <el-icon class="el-icon--left">
                    <Lock />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
          </el-form>
          <template #footer>
            <span class='dialog-footer'>
              <el-button @click='dialogFormVisible = false'>取消</el-button>
              <el-button type='primary' @click='dialogFormVisible = false'>
                提交
              </el-button>
            </span>
          </template>
        </el-dialog>
      </div>
      <el-table
        style='margin-top: 5px;width: 100%'
        height='73vh'
        :row-style="{height:'40px'}"
        :cell-style="{padding:'0px'}"
        border
        :data='pageData'
      >
        <el-table-column fixed='left' type='selection' align='center' width='40' />
        <el-table-column fixed='left' :show-overflow-tooltip=true prop='title' label='文章标题' align='center'
                         width='330' />
        <el-table-column prop='desc' label='描述' align='center' width='210' />
        <el-table-column prop='context' label='正文' align='center' width='400' />
        <el-table-column prop='CreatedAt' sortable label='创建时间' align='center' width='200'
                         :formatter='formatTimestamp' />
        <el-table-column prop='UpdatedAt' sortable label='修改时间' align='center' width='200'
                         :formatter='formatTimestamp' />
        <el-table-column label='操作' align='center' width='200'>
          <el-button type="primary" circle plain>
            <el-icon>
              <Edit />
            </el-icon>
          </el-button>

          <el-popconfirm
            width="220"
            confirm-button-text="好的"
            cancel-button-text="不用了"
            title="这是一段内容确定删除吗？"
          >
            <template #reference>
              <el-button type="danger" circle plain>
                <el-icon>
                  <Delete />
                </el-icon>
              </el-button>
            </template>
          </el-popconfirm>

        </el-table-column>
      </el-table>
    </div>
    <div class='pagination-container'>
      <el-pagination class='pagination-wrapper'
        @size-change='pageSizeChange'
        @current-change='pageCurrentChange'
        :current-page='page.current'
        :page-sizes='[10, 30, 50, 100, 200, 300]'
        :page-size=page.size
        layout='total, sizes, prev, pager, next, jumper'
        :total=articlesInfo.length>
      </el-pagination>
    </div>

  </div>
</template>
<script>
import { User, Lock, Search, Edit, Delete } from '@element-plus/icons-vue'
import { isNull } from '@/utils/service'

export default {
  name: 'LoginView',
  components: { User, Lock, Search, Edit, Delete },
  computed: {
    pageData() {
      const start = (this.page.current - 1) * this.page.size
      const end = start + this.page.size
      return this.articlesInfo.slice(start, end)
    }
  },
  mounted() {
    this.getArtList()
  },
  data() {
    return {
      username: '',
      dialogFormVisible: false,
      dialog: {
        openArt: {
          /*表单*/
          form: {
            title: '',
            desc: '',
            context: ''
          },
          /*表单验证*/
          formRules: {
            title: [
              { required: true, message: '请输入文章标题', trigger: 'blur' },
              { min: 4, max: 20, message: '长度在 5 - 50 个字符', trigger: 'blur' }
            ],
            desc: [
              { required: true, message: '请输入文章描述', trigger: 'blur' },
              { min: 4, max: 20, message: '长度在 5 - 100 个字符', trigger: 'blur' }
            ],
            context: [
              { required: true, message: '请输入文章正文', trigger: 'blur' },
              { min: 1, max: 1000, message: '长度在 5 - 1000 个字符', trigger: 'blur' }
            ]
          }
        }
      },
      /*分页对象*/
      page: {
        selectRow: [],
        size: 30,
        current: 1
      },
      articlesInfo: []
    }
  },
  methods: {
    /** 获取文章列表 */
    async getArtList() {
      const res = await this.$http.get('articles', {
        params: {
          pageSize: this.page.size,
          pageNum: this.page.current
        }
      })
      this.articlesInfo = res.data.data
    },
    /** 搜索文章 */
    search() {
      if (isNull(this.username)) {
        this.$message.warning('请输入文章名')
      }
    },
    /** 清空输入框 */
    clearInput() {
      console.log('clearInput...')
    },
    /** 移除 */
    removeTemplate() {
      console.log('removeTemplate...')
    },
    pageSizeChange(v) {
      this.page.size = v
    },
    pageCurrentChange(v) {
      this.page.current = v
    },
    // 格式化时间
    formatTimestamp(row, column, cellValue) {
      const dateObject = new Date(cellValue)
      const year = dateObject.getFullYear()
      const month = (dateObject.getMonth() + 1).toString().padStart(2, '0')
      const day = dateObject.getDate().toString().padStart(2, '0')
      const hours = dateObject.getHours().toString().padStart(2, '0')
      const minutes = dateObject.getMinutes().toString().padStart(2, '0')
      const seconds = dateObject.getSeconds().toString().padStart(2, '0')
      return `${year}/${month}/${day} ${hours}:${minutes}:${seconds}`
    },
  }
}
</script>
<style>
.input-container {
  display: flex;
  align-items: center;
}

.dialog-footer {
  margin-top: 165px;
  margin-right: 10px;
  display: flex;
  justify-content: flex-end; /* 按钮右对齐 */
}

.pagination-container {
  margin-top: 22px;
  display: flex;
  flex-direction: column-reverse; /* 设置为从下往上的弹性容器 */
}

.pagination-wrapper {
  align-self: flex-end; /* 对齐到底部 */
}
</style>