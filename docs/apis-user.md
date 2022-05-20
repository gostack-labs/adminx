# 

1. [分页获取用户列表接口](#1-分页获取用户列表接口)
2. [获取用户详情接口](#2-获取用户详情接口)
3. [通过用户获取用户详情接口](#3-通过用户获取用户详情接口)
4. [新增用户接口](#4-新增用户接口)
5. [更新用户信息接口](#5-更新用户信息接口)
6. [删除用户接口](#6-删除用户接口)
7. [批量删除用户接口](#7-批量删除用户接口)

## apis

### 1. 分页获取用户列表接口

```text
GET /sys/user
```

_body_:

```javascript
{  //object(api.listUserRequest), 分页获取用户列表请求参数
  "email": "string",  //string, 邮箱
  "full_name": "string",  //string, 全名
  "page_id": 0,  //int32, validate:"required,min=1", 页码
  "page_size": 0,  //int32, validate:"required,max=50", 页尺寸
  "phone": "string",  //string, 手机号
  "username": "string"  //string, 用户名
}
```

__Response__:

```javascript
//StatusCode: 200 用户集合
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[db.User]
    {  //object(db.User)
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "email": "string",  //string
      "full_name": "string",  //string
      "hashed_password": "string",  //string
      "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "phone": "string",  //string
      "username": "string"  //string
    }
  ],
  "msg": "获取成功"  //string
}
```

---

### 2. 获取用户详情接口

```text
GET /sys/user/info
```

__Response__:

```javascript
//StatusCode: 200 用户详情
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.userInfoResponse)
    "button": [  //array[any], 按钮列表

    ],
    "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "email": "string",  //string
    "full_name": "string",  //string
    "hashed_password": "string",  //string
    "page": [  //array[any], 菜单列表

    ],
    "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "phone": "string",  //string
    "role": [  //array[string], 角色列表
      "string"
    ],
    "username": "string"  //string
  },
  "msg": "获取成功"  //string
}
```

---

### 3. 通过用户获取用户详情接口

```text
GET /sys/user/info/:username
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__username__|_param_|string|true|required||用户（用户名/手机号/邮箱）

__Response__:

```javascript
//StatusCode: 200 用户详情
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.userInfoByIDResponse), 通过用户获取用户详情返回数据
    "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "email": "string",  //string
    "full_name": "string",  //string
    "hashed_password": "string",  //string
    "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "phone": "string",  //string
    "Role": [  //array[string], 角色集合
      "string"
    ],
    "username": "string"  //string
  },
  "msg": "获取成功"  //string
}
```

---

### 4. 新增用户接口

```text
POST /sys/user
```

_body_:

```javascript
{  //object(api.createUserRequest), 新增用户请求参数
  "email": "string",  //string, validate:"required_without=Phone,omitempty,email", 邮箱
  "full_name": "string",  //string, validate:"required", 全名
  "password": "string",  //string, validate:"required,min=6", 密码
  "phone": "string",  //string, validate:"required_without=Email,omitempty,phone", 手机号
  "role": [  //array[string], 角色
    "string"
  ],
  "username": "string"  //string, validate:"required,alphanum", 用户名
}
```

__Response__:

```javascript
//StatusCode: 200 用户详情
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(db.User)
    "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "email": "string",  //string
    "full_name": "string",  //string
    "hashed_password": "string",  //string
    "password_change_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
    "phone": "string",  //string
    "username": "string"  //string
  },
  "msg": "创建成功"  //string
}
```

---

### 5. 更新用户信息接口

```text
PUT /sys/user/:username
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__username__|_param_|string|true|required||用户名

_body_:

```javascript
{  //object(api.updateUserRequest), 更新用户信息请求参数
  "email": "string",  //string, validate:"required_without=Phone,omitempty,email", 手机号
  "full_name": "string",  //string, validate:"required", 全名
  "phone": "string",  //string, validate:"required_without=Email,omitempty,phone", 密码
  "role": [  //array[string], 角色集合
    "string"
  ]
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "修改成功"  //string
}
```

---

### 6. 删除用户接口

```text
DELETE /sys/user/single/:username
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "删除成功"  //string
}
```

---

### 7. 批量删除用户接口

```text
DELETE /sys/user/batch
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "删除成功"  //string
}
```

---
