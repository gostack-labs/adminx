package code

import _ "embed"

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

const (
	ServerError             = 10101
	ParamBindError          = 10102
	SendEmailError          = 10103
	AuthHeaderNotExistError = 10104
	AuthHeaderFormatError   = 10105
	AuthTypeError           = 10106
	RBACError               = 10107

	ApiBindMenuExistError = 20101
	ApiDeleteError        = 20102
	ApiGroupNoRowError    = 20103
	ApiExistError         = 20104

	ApiGroupExistError  = 20201
	ApiGroupHasApiError = 20202

	MenuRootTypeError     = 20301
	MenuParentNoRowError  = 20302
	MenuNoButtonError     = 20303
	MenuNoDirOrMenuError  = 20304
	MenuNoChildNodeError  = 20305
	MenuHasChildNodeError = 20306

	RoleExistError   = 20401
	RoleHasUserError = 20402
	RoleHasMenuError = 20403

	UserNotExistError      = 20501
	UserPwdError           = 20502
	UserEmailExistError    = 20503
	UserPhoneExistError    = 20504
	UserUsernameExistError = 20505
	UserHasRoleError       = 20506

	TokenCreateError        = 20601
	RefreshTokenCreateError = 20602
	VerifyCodeError         = 20603
	SessionBlockedError     = 20604
	SessionError            = 20605
	SessionMismatchedError  = 20606
	SessionExpiredError     = 20607
	SessionNotExistError    = 20608
	TokenExpiredError       = 20609
	TokenInvalidError       = 20610
)

func Text(code int, lang string) string {
	switch lang {
	case "zh":
		return zhCNText[code]
	case "en":
		return enUSText[code]
	default:
		return zhCNText[code]
	}
}
