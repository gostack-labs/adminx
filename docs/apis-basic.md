# 基础接口

登录、注册、获取验证码、刷新token相关接口

1. [获取错误列表接口](#1-获取错误列表接口)
2. [用户登录接口](#2-用户登录接口)
3. [刷新token接口](#3-刷新token接口)

## apis

### 1. 获取错误列表接口

```text
GET /code
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.codeResponse), 获取错误列表
    "BusinessCodes": [  //array[api.codes], 业务错误
      {  //object(api.codes)
        "code": 0,  //int, 错误码
        "message": "string"  //string, 错误描述
      }
    ],
    "SystemCodes": [  //array[api.codes], 系统错误
      {  //object(api.codes)
        "code": 0,  //int, 错误码
        "message": "string"  //string, 错误描述
      }
    ]
  },
  "msg": "获取成功"  //string
}
```

---

### 2. 用户登录接口

```text
POST /signin
```

_body_:

```javascript
{  //object(api.logginUserRequest), 登录请求参数
  "password": "string",  //string, validate:"required,min=6", 密码
  "username": "string"  //string, validate:"required", 用户名，邮箱，手机号
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 0,  //int
  "data": {  //object(api.logginUserResponse), 登录返回数据
    "access_token": "string",  //string, accessToken
    "access_token_expires_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), accessToken 过期时间
    "refresh_token": "string",  //string, 刷新token
    "refresh_token_expires_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 刷新token 过期时间
    "session_id": null,  //object, sessionID
    "user": {  //object(api.userResponse), 用户信息
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "email": "string",  //string
      "full_name": "string",  //string
      "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "phone": "string",  //string
      "username": "string"  //string
    }
  },
  "msg": "string"  //string
}
```

---

### 3. 刷新token接口

```text
POST /tokens/renew_access
```

_body_:

```javascript
{  //object(api.renewAccessTokenRequest), 刷新token请求数据
  "refresh_token": "string"  //string, validate:"required", 刷新token
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 0,  //int
  "data": {  //object(api.renewAccessTokenResponse), 刷新token返回数据
    "access_token": "string",  //string, accessToken
    "access_token_expires_at": "2022-05-16T16:47:48.741899+08:00"  //object(time.Time), accessToken 过期时间
  },
  "msg": "string"  //string
}
```

---
