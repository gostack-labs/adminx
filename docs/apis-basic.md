# 基础接口

登录、注册、获取验证码、刷新token相关接口

1. [获取错误列表接口](#1-获取错误列表接口)
2. [仪表盘接口](#2-仪表盘接口)
3. [用户登录接口](#3-用户登录接口)
4. [用户注册接口](#4-用户注册接口)
5. [发送邮箱验证码接口](#5-发送邮箱验证码接口)
6. [刷新token接口](#6-刷新token接口)

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

### 2. 仪表盘接口

```text
GET /dashboard
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.viewResponse)
    "CpuCores": 0,  //int32, cpu 核心
    "CpuName": "string",  //string, cpu 名字
    "CpuUsedPercent": 1.23,  //float64, cpu使用百分比
    "DiskTotal": "string",  //string, 存储总量
    "DiskUsed": "string",  //string, 存储使用量
    "DiskUsedPercent": 1.23,  //float64, 存储使用百分比
    "Env": "string",  //string, 环境变量
    "GoArch": "string",  //string, GOARCH
    "GoOS": "string",  //string, GOOS
    "GoPath": "string",  //string, GOPATH
    "Goroutine": 0,  //int, go 协程
    "GoVersion": "string",  //string, go 版本
    "Host": "string",  //string, 主机
    "HostName": "string",  //string, 主机名
    "HostOS": "string",  //string, 主机系统
    "MemTotal": "string",  //string, 内存总量
    "MemUsed": "string",  //string, 内存使用量
    "MemUsedPercent": 1.23,  //float64, 内存使用百分比
    "PostgresqlVersion": "string",  //string, pgsql 版本
    "ProjectPath": "string",  //string, 项目路径
    "ProjectVersion": "string",  //string, 项目版本
    "RedisVersion": "string"  //string, redis 版本
  },
  "msg": "获取成功"  //string
}
```

---

### 3. 用户登录接口

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
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 创建时间
      "email": "string",  //string, 邮箱
      "full_name": "string",  //string, 全名
      "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 密码修改时间
      "phone": "string",  //string, 手机号
      "username": "string"  //string, 用户名
    }
  },
  "msg": "string"  //string
}
```

---

### 4. 用户注册接口

```text
POST /signup
```

_body_:

```javascript
{  //object(api.signupRequest), 用户注册请求参数
  "email": "string",  //string, validate:"required_without=Phone,omitempty,email", 邮箱
  "full_name": "string",  //string, validate:"required", 全名
  "password": "string",  //string, validate:"required,min=6", 密码
  "phone": "string",  //string, validate:"required_without=Email,omitempty,phone", 手机号
  "username": "string",  //string, validate:"required,alphanum", 用户名
  "verify_code": "string"  //string, validate:"required,alphanum", 验证码
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.userResponse), 用户注册返回数据
    "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 创建时间
    "email": "string",  //string, 邮箱
    "full_name": "string",  //string, 全名
    "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 密码修改时间
    "phone": "string",  //string, 手机号
    "username": "string"  //string, 用户名
  },
  "msg": "操作成功"  //string
}
```

---

### 5. 发送邮箱验证码接口

```text
GET /signup/sendUsingEmail
```

_body_:

```javascript
{  //object(api.verifyCodeEmailRequest), 发送邮箱验证码请求参数
  "email": "string"  //string, validate:"required,email", 邮箱
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "操作成功"  //string
}
```

---

### 6. 刷新token接口

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
