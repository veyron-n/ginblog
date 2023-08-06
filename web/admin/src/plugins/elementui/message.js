/**
 * 封装message 消息提示
 */
import { ElMessage  } from 'element-plus'

export default {
  success(msg){
    ElMessage ({
      showClose: true,
      message: msg,
      type: 'success'
    })
  },

  error(msg){
    ElMessage ({
      showClose: true,
      message: msg,
      type: 'error'
    })
  },

  warning(msg){
    ElMessage ({
      showClose: true,
      message: msg,
      type: 'warning'
    })
  },

  defaultMsg(msg){
    ElMessage ({
      showClose: true,
      message: msg
    })
  }
}
