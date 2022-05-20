# 菜单管理相关接口

菜单管理相关接口

1. [获取菜单树接口](#1-获取菜单树接口)
2. [新增菜单接口](#2-新增菜单接口)
3. [更新菜单接口](#3-更新菜单接口)
4. [删除菜单接口](#4-删除菜单接口)
5. [批量删除菜单接口](#5-批量删除菜单接口)
6. [获取菜单对应的按钮接口](#6-获取菜单对应的按钮接口)
7. [菜单绑定/解绑接口](#7-菜单绑定/解绑接口)
8. [获取菜单绑定的接口接口](#8-获取菜单绑定的接口接口)
9. [获取菜单对应的接口接口](#9-获取菜单对应的接口接口)

## apis

### 1. 获取菜单树接口

```text
GET /sys/menu/tree
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": {  //object(api.menuTreeResponse)
    "Button": {  //object(map[int64]&{%!s(token.Pos=2875) <nil> %!s(*ast.StarExpr=&{2877 0x14000282b58})})
      "0": [  //array[db.Menu]
        {  //object(db.Menu)
          "auth": [  //array[string]
            "string"
          ],
          "component": "string",  //string
          "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
          "hyperlink": "string",  //string
          "icon": "string",  //string
          "id": 0,  //int64
          "is_affix": false,  //bool
          "is_hide": false,  //bool
          "is_iframe": false,  //bool
          "is_keep_alive": false,  //bool
          "name": "string",  //string
          "parent": 0,  //int64
          "path": "string",  //string
          "redirect": "string",  //string
          "sort": 0,  //int32
          "title": "string",  //string
          "type": 0  //int32
        }
      ]
    },
    "Menu": [  //array[api.MenuValue]
      {  //object(api.MenuValue)
        "children": [  //array[api.MenuValue]

        ],
        "component": "string",  //string
        "id": 0,  //int64
        "meta": null,  //object
        "name": "string",  //string
        "parent": 0,  //int64
        "path": "string",  //string
        "sort": 0,  //int32
        "title": "string",  //string
        "type": 0  //int32
      }
    ]
  },
  "msg": "获取成功"  //string
}
```

---

### 2. 新增菜单接口

```text
POST /sys/menu
```

_body_:

```javascript
{  //object(api.createMenuRequest), 新增菜单请求参数
  "auth": [  //array[string], 权限粒子
    "string"
  ],
  "component": "string",  //string, 组件路径
  "hyperlink": "string",  //string, 超链接
  "icon": "string",  //string, 图标
  "is_affix": false,  //bool, 是否固定在标签栏
  "is_hide": false,  //bool, 是否隐藏
  "is_iframe": false,  //bool, 是否内嵌窗口
  "is_keep_alive": false,  //bool, 是否缓存组件状态
  "name": "string",  //string, validate:"required", 路由名称
  "parent": 0,  //int64, validate:"required,numeric", 父级
  "path": "string",  //string, 路径
  "redirect": "string",  //string, 跳转路径
  "sort": 0,  //int32, 顺序
  "title": "string",  //string, validate:"required", 标题
  "type": 0  //int32, validate:"oneof=1 2 3", 类型：1 目录，2 菜单，3 按钮
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

### 3. 更新菜单接口

```text
PUT /sys/menu/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

_body_:

```javascript
{  //object(api.updateMenuRequest)
  "auth": [  //array[string], 权限粒子
    "string"
  ],
  "component": "string",  //string, 组件路径
  "hyperlink": "string",  //string, 超链接
  "icon": "string",  //string, 图标
  "is_affix": false,  //bool, 是否固定在标签栏
  "is_hide": false,  //bool, 是否隐藏
  "is_iframe": false,  //bool, 是否内嵌窗口
  "is_keep_alive": false,  //bool, 是否缓存组件状态
  "name": "string",  //string, validate:"required", 路由名称
  "path": "string",  //string, 路径
  "redirect": "string",  //string, 跳转路径
  "sort": 0,  //int32, 顺序
  "title": "string"  //string, validate:"required", 标题
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

### 4. 删除菜单接口

```text
DELETE /sys/menu/single/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required,numeric||主键ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "修改成功"  //string
}
```

---

### 5. 批量删除菜单接口

```text
DELETE /sys/menu/batch
```

_body_:

```javascript
{  //object(api.batchDeleteMenuRequest), 批量删除菜单请求参数
  "ids": [  //array[int64], validate:"required", 主键ID集合
    0
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

### 6. 获取菜单对应的按钮接口

```text
GET /sys/menu/button/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required,numeric||主键ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[db.Menu]
    {  //object(db.Menu)
      "auth": [  //array[string]
        "string"
      ],
      "component": "string",  //string
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "hyperlink": "string",  //string
      "icon": "string",  //string
      "id": 0,  //int64
      "is_affix": false,  //bool
      "is_hide": false,  //bool
      "is_iframe": false,  //bool
      "is_keep_alive": false,  //bool
      "name": "string",  //string
      "parent": 0,  //int64
      "path": "string",  //string
      "redirect": "string",  //string
      "sort": 0,  //int32
      "title": "string",  //string
      "type": 0  //int32
    }
  ],
  "msg": "获取成功"  //string
}
```

---

### 7. 菜单绑定/解绑接口

```text
POST /sys/menu/api/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

_body_:

```javascript
{  //object(api.menuBindApiRequest), 菜单绑定/解绑api请求参数
  "apis": [  //array[int64], validate:"required", api主键结合
    0
  ],
  "type": 0  //int, validate:"required,oneof=1 2", 操作类型 1:bind 2:unbind
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

### 8. 获取菜单绑定的接口接口

```text
GET /sys/menu/api/:menu
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__menu__|_param_|int64|false|||菜单主键ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[int64]
    0
  ],
  "msg": "获取成功"  //string
}
```

---

### 9. 获取菜单对应的接口接口

```text
GET /sys/menu/api-list/:menu
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__menu__|_param_|int64|true|required||

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[db.Api]
    {  //object(db.Api)
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "groups": 0,  //int64
      "id": 0,  //int64
      "method": "string",  //string
      "remark": "string",  //string
      "title": "string",  //string
      "url": "string"  //string
    }
  ],
  "msg": "获取成功"  //string
}
```

---
