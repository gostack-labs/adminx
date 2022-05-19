# 基础接口

登录、注册、获取验证码、刷新token相关接口

1. [用户登录接口](#1-用户登录接口)
2. [刷新token接口](#2-刷新token接口)

## apis

### 1. 用户登录接口

```text
POST /signin
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "data": {  //object(api.logginUserResponse), 登录返回数据
    "access_token": "abc",  //string, accessToken
    "access_token_expires_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), accessToken 过期时间
    "refresh_token": "abc",  //string, 刷新token
    "refresh_token_expires_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time), 刷新token 过期时间
    "session_id": null,  //object, sessionID
    "user": {  //object(api.userResponse), 用户信息
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "email": "abc",  //string
      "full_name": "abc",  //string
      "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "phone": "abc",  //string
      "username": "abc"  //string
    }
  },
  "msg": "abc"  //string
}
```

---

### 2. 刷新token接口

```text
POST /tokens/renew_access
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "data": {  //object(api.renewAccessTokenResponse), 刷新token返回数据
    "access_token": "abc",  //string, accessToken
    "access_token_expires_at": "2022-05-16T16:47:48.741899+08:00"  //object(time.Time), accessToken 过期时间
  },
  "msg": "abc"  //string
}
```

---
