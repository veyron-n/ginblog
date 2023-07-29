package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSPWORD_WRONG  = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_USER_NO_RIGHT    = 1004
	ERROR_TOKEN_NOT_EXIST  = 1005
	ERROR_TOKEN_RUNTIME    = 1006
	ERROR_TOKEN_WRONG      = 1007
	ERROR_TOKEN_TYPE_WRONG = 1008
	// code = 2000... 文章模块的错误
	ERROR_ARTICLE_USED      = 2001
	ERROR_ARTICLE_NOT_EXIST = 2002
	// code = 3000... 分类模块的错误
	ERROR_CATEGORY_USED     = 3001
	ERROR_CATEGORY_NOT_EXIT = 3002
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSPWORD_WRONG:  "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",

	ERROR_CATEGORY_USED:     "分类已存在",
	ERROR_CATEGORY_NOT_EXIT: "分类不存在",

	ERROR_ARTICLE_USED:      "文章已存在",
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}