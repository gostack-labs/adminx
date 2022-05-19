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
  "data": {  //object(map[string]any)
    "abc": null  //any
  },
  "msg": "获取成功"  //string
}
```

---

### 2. 新增菜单接口

```text
POST /sys/menu
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
    "abc"
  ],
  "component": "abc",  //string, 组件路径
  "hyperlink": "abc",  //string, 超链接
  "icon": "abc",  //string, 图标
  "is_affix": false,  //bool, 是否固定在标签栏
  "is_hide": false,  //bool, 是否隐藏
  "is_iframe": false,  //bool, 是否内嵌窗口
  "is_keep_alive": false,  //bool, 是否缓存组件状态
  "name": "abc",  //string, required, 路由名称
  "path": "abc",  //string, 路径
  "redirect": "abc",  //string, 跳转路径
  "sort": 123,  //int32, 顺序
  "title": "abc"  //string, required, 标题
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
        "abc"
      ],
      "component": "abc",  //string
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "hyperlink": "abc",  //string
      "icon": "abc",  //string
      "id": 123,  //int64
      "is_affix": false,  //bool
      "is_hide": false,  //bool
      "is_iframe": false,  //bool
      "is_keep_alive": false,  //bool
      "name": "abc",  //string
      "parent": 123,  //int64
      "path": "abc",  //string
      "redirect": "abc",  //string
      "sort": 123,  //int32
      "title": "abc",  //string
      "type": 123  //int32
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
  "apis": [  //array[int64], required, api主键结合
    123
  ],
  "type": 123  //int, required, 操作类型 1:bind 2:unbind
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
    123
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
      "groups": 123,  //int64
      "id": 123,  //int64
      "method": "abc",  //string
      "remark": "abc",  //string
      "title": "abc",  //string
      "url": "abc"  //string
    }
  ],
  "msg": "获取成功"  //string
}
```

---
