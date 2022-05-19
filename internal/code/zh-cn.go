package code

var zhCNText = map[int]string{
	ServerError:             "内部服务器错误",
	ParamBindError:          "参数信息错误",
	SendEmailError:          "发送邮件失败",
	AuthHeaderNotExistError: "没有提供授权标头",
	AuthHeaderFormatError:   "无效的授权头格式",
	AuthTypeError:           "不支持的授权类型",
	RBACError:               "暂无访问权限",

	ApiBindMenuExistError: "有菜单绑定了该接口，请解绑后再删除",
	ApiDeleteError:        "接口删除失败",
	ApiGroupNoRowError:    "接口分组不存在",
	ApiExistError:         "接口已存在",

	ApiGroupExistError:  "接口分组已存在",
	ApiGroupHasApiError: "接口分组下存在接口，不允许删除",

	MenuRootTypeError:     "根节点类型有误",
	MenuParentNoRowError:  "父节点不存在",
	MenuNoButtonError:     "目录下不允许创建按钮权限",
	MenuNoDirOrMenuError:  "菜单下不允许创建目录或子菜单",
	MenuNoChildNodeError:  "按钮下不允许创建子节点",
	MenuHasChildNodeError: "存在子节点不允许删除",

	RoleExistError:   "角色已存在",
	RoleHasUserError: "角色被其他用户关联，无法删除",
	RoleHasMenuError: "角色存在与菜单的关联，无法删除",

	UserNotExistError:      "用户不存在",
	UserPwdError:           "密码不正确",
	UserEmailExistError:    "该邮箱已存在！",
	UserPhoneExistError:    "该手机号码已存在",
	UserUsernameExistError: "该用户名已存在",
	UserHasRoleError:       "当前用户关联了其他角色，无法直接删除",

	TokenCreateError:        "令牌生成失败",
	RefreshTokenCreateError: "刷新令牌生成失败",
	VerifyCodeError:         "验证码输入错误或已过期",
	SessionBlockedError:     "会话已屏蔽",
	SessionError:            "不正确的用户会话",
	SessionMismatchedError:  "不匹配的会话令牌",
	SessionExpiredError:     "会话已过期",
	SessionNotExistError:    "会话不存在",
	TokenExpiredError:       "令牌过期",
	TokenInvalidError:       "令牌无效",
}
