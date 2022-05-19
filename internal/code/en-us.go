package code

var enUSText = map[int]string{
	ServerError:             "Internal server error",
	ParamBindError:          "Parameter error",
	SendEmailError:          "Failed to send mail",
	AuthHeaderNotExistError: "authorization header is not provided",
	AuthHeaderFormatError:   "invalid authorization header format",
	AuthTypeError:           "unsupported authorization type",
	RBACError:               "No access",

	ApiBindMenuExistError: "menu bound to this api, please unbind it and delete it",
	ApiDeleteError:        "Failed to delete api",
	ApiGroupNoRowError:    "apiGroup does not exist",
	ApiExistError:         "api already exists",

	ApiGroupExistError:  "apiGroup already exists",
	ApiGroupHasApiError: "apiGroup contains api and cannot be deleted",

	MenuRootTypeError:     "root node type is incorrect",
	MenuParentNoRowError:  "parent node does not exist",
	MenuNoButtonError:     "button permissions cannot be created in the directory",
	MenuNoDirOrMenuError:  "cannot create directories or submenus under the menu",
	MenuNoChildNodeError:  "child nodes cannot be created under the button",
	MenuHasChildNodeError: "has child nodes,cannot be deleted",

	RoleExistError:   "role already exists",
	RoleHasUserError: "role is associated with other users and cannot be deleted",
	RoleHasMenuError: "role is associated with the menu and cannot be deleted",

	UserNotExistError:      "user does not exist",
	UserPwdError:           "incorrect password",
	UserEmailExistError:    "email already exists",
	UserPhoneExistError:    "phone already exists",
	UserUsernameExistError: "username already exists",
	UserHasRoleError:       "current user is associated with another role and cannot be deleted",

	TokenCreateError:        "Failed to create token",
	RefreshTokenCreateError: "Failed to create refreshToken",
	VerifyCodeError:         "verification code is incorrect or expired",
	SessionBlockedError:     "session is blocked",
	SessionError:            "incorrect session user",
	SessionMismatchedError:  "mismatched session token",
	SessionExpiredError:     "expired session",
	SessionNotExistError:    "session not exist",
	TokenExpiredError:       "token has Expired",
	TokenInvalidError:       "token is invalid",
}
