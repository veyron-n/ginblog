<template>
  <div>
    <div>
      <div class='input-container'>
        <el-input style='width: 200px' size='default' v-model='username' placeholder='请输入用户名' clearable @clear='clearInput'>
          <template #prefix>
            <el-icon class='el-icon--left'>
              <Search />
            </el-icon>
          </template>
        </el-input>
        <el-button style='margin-left: 5px' size='default' type='primary' @click='getUserList'>搜索</el-button>
        <el-button style='margin-left: 5px' size='default' type='success' @click='dialogVisible = true'>添加用户
        </el-button>
        <el-dialog style='width: 500px; height: 500px' v-model='dialogVisible' title='添加用户'>
          <el-form ref="openUserForm" :model="dialog.openUser.form" :rules="dialog.openUser.formRules"
                   label-width="100px">
            <el-form-item label='登录账号' prop='account'>
              <el-input style='width: 300px' placeholder='请输入登录账号' maxlength='20' show-word-limit v-model='dialog.openUser.form.account' clearable>
                <template #prefix>
                  <el-icon class="el-icon--left">
                    <User />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label='登录密码' prop='pwd' :show-password='true'>
              <el-input style='width: 300px' type='password' show-password placeholder='请输入登录密码'
                        v-model='dialog.openUser.form.pwd' clearable>
                <template #prefix>
                  <el-icon class="el-icon--left">
                    <Lock />
                  </el-icon>
                </template>
              </el-input>
            </el-form-item>
            <el-form-item label="账号角色" prop="roleIds">
              <el-select style='width: 300px' multiple v-model="dialog.openUser.form.rid" placeholder="请选择角色">
                <el-option
                  v-for="item in dialog.openUser.form.roles"
                  :key="item.id"
                  :label="item.roleName"
                  :value="item.id">
                </el-option>
              </el-select>
            </el-form-item>
          </el-form>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="dialogVisible = false">取消</el-button>
              <el-button type="primary" @click="dialogVisible = false">确定</el-button>
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
        <el-table-column fixed='left' :show-overflow-tooltip=true prop='username' label='用户名' align='center'
                         width='330' />
        <el-table-column prop='role' label='角色' align='center' width='200' :formatter='handleRole'/>
        <el-table-column prop='CreatedAt' sortable label='创建时间' align='center' width='330'
                         :formatter='formatTimestamp' />
        <el-table-column prop='UpdatedAt' sortable label='修改时间' align='center' width='330'
                         :formatter='formatTimestamp' />
        <el-table-column label='操作' align='center' width='350'>
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
        :total=userInfo.length>
      </el-pagination>
    </div>

  </div>
</template>
<script>
import { User, Lock, Search } from '@element-plus/icons-vue'

export default {
  name: 'LoginView',
  components: { User, Lock, Search },
  computed: {
    pageData() {
      const start = (this.page.current - 1) * this.page.size
      const end = start + this.page.size
      return this.userInfo.slice(start, end)
    }
  },
  mounted() {
    this.getUserList()
  },
  data() {
    return {
      username: '',
      dialogVisible: false,
      dialog: {
        openUser: {
          /*表单*/
          form: {
            account: '',
            pwd: '',
            rid: undefined,
            roles: [
              { id: 1, roleName: '管理者' },
              { id: 2, roleName: '订阅者' },
              { id: 3, roleName: '...' },
            ],
          },
          /*表单验证*/
          formRules: {
            account: [
              { required: true, message: '请输入登录账号', trigger: 'blur' },
              { min: 4, max: 20, message: '长度在 4 - 20 个字符', trigger: 'blur' }
            ],
            pwd: [
              { required: true, message: '请输入登录密码', trigger: 'blur' },
              {
                pattern: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,20}$/,
                message: '密码必须由数字、字母两种字符组成，长度在6-20位之间',
                trigger: 'blur'
              }
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
      userInfo: []
    }
  },
  methods: {
    handleRole(row, column, cellValue) {
      let roleStr;
      if(cellValue === 1) {
        roleStr = '管理者'
      } else if (cellValue === 2) {
        roleStr = '订阅者'
      } else {
        roleStr = '-'
      }
      return roleStr
    },
    async getUserList() {
      const res = await this.$http.get('users', {
        params: {
          username: this.username,
          pageSize: this.page.size,
          pageNum: this.page.current
        }
      })
      this.userInfo = res.data.data
    },
    clearInput() {
      this.username = ''
      this.getUserList()
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