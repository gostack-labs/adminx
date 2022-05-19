# 角色管理相关接口

角色管理相关接口

1. [分页获取角色列表接口](#1-分页获取角色列表接口)
2. [创建角色接口](#2-创建角色接口)
3. [更新角色接口](#3-更新角色接口)
4. [删除角色接口](#4-删除角色接口)
5. [批量删除角色接口](#5-批量删除角色接口)
6. [角色授权/解除菜单权限接口](#6-角色授权/解除菜单权限接口)
7. [获取角色授权的菜单权限接口](#7-获取角色授权的菜单权限接口)
8. [角色授权/解除菜单权限接口](#8-角色授权/解除菜单权限接口)
9. [获取角色授权的接口权限接口](#9-获取角色授权的接口权限接口)

## apis

### 1. 分页获取角色列表接口

```text
GET /sys/role
```

_body_:

```javascript
{  //object(api.listRoleRequest), 分页获取角色列表请求参数
  "key": "string",  //string, 标识
  "name": "string",  //string, 名称
  "page_id": 0,  //int32, validate:"required,min=1", 页码
  "page_size": 0  //int32, validate:"required,max=50", 页尺寸
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[db.Role]
    {  //object(db.Role)
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "id": 0,  //int64
      "is_disable": false,  //bool
      "key": "string",  //string
      "name": "string",  //string
      "remark": "string",  //string
      "sort": 0  //int32
    }
  ],
  "msg": "获取成功"  //string
}
```

---

### 2. 创建角色接口

```text
POST /sys/role
```

_body_:

```javascript
{  //object(api.createRoleRequest), 创建角色请求参数
  "is_disable": false,  //bool, 是否禁用
  "key": "string",  //string, validate:"required", 标识
  "name": "string",  //string, validate:"required", 名称
  "remark": "string",  //string, 备注
  "sort": 0  //int32, 排序
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "创建成功"  //string
}
```

---

### 3. 更新角色接口

```text
PUT /sys/role/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||

_body_:

```javascript
{  //object(api.updateRoleRequest), 更新角色请求参数
  "is_disable": false,  //bool
  "key": "string",  //string, validate:"required"
  "name": "string",  //string, validate:"required"
  "remark": "string",  //string
  "sort": 0  //int32
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "更新成功"  //string
}
```

---

### 4. 删除角色接口

```text
DELETE /sys/role/single/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "删除成功"  //string
}
```

---

### 5. 批量删除角色接口

```text
DELETE /sys/role/batch
```

_body_:

```javascript
{  //object(api.batchDeleteRoleRequest), 批量删除角色请求参数
  "ids": [  //array[int64], validate:"required", 主键集合
    0
  ]
}
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

### 6. 角色授权/解除菜单权限接口

```text
POST /sys/role/permission/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||角色ID

_body_:

```javascript
{  //object(api.updateRolePermissionRequest), 角色授权/解除菜单权限请求参数
  "role_menus": [  //array[db.CreateRoleMenuParams], 角色菜单集合
    {  //object(db.CreateRoleMenuParams)
      "menu": 0,  //int64
      "role": 0,  //int64
      "type": 0  //int32
    }
  ]
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

### 7. 获取角色授权的菜单权限接口

```text
GET /sys/role/permission/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||角色ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.getRolePermissionResponse), 获取角色授权的菜单权限
    "button": {  //object(map[int64]&{%!s(token.Pos=9842) <nil> int64})
      "0": [  //array[int64]
        0
      ]
    },
    "menu": [  //array[int64]
      0
    ]
  },
  "msg": "获取成功"  //string
}
```

---

### 8. 角色授权/解除菜单权限接口

```text
POST /sys/role/api/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||角色ID

_body_:

```javascript
{  //object(api.roleApiPermissionRequest), 角色授权/解除接口权限请求参数
  "api": [  //array[int64], validate:"required", 接口ID集合
    0
  ],
  "type": 0  //int, validate:"required,oneof=0 1", 操作类型 0:解除api权限 1:绑定api权限
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

### 9. 获取角色授权的接口权限接口

```text
GET /sys/role/api/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||角色ID

_body_:

```javascript
{  //object(api.getRoleApiRequest)
  "menu": 0  //int64, validate:"required", 菜单ID
}
```

__Response__:

```javascript
//StatusCode: 200 接口ID集合
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[int64]
    0
  ],
  "msg": "获取成功"  //string
}
```

---
