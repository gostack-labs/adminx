# 菜单管理相关接口

菜单管理相关接口

1. [获取菜单树接口](#1-获取菜单树接口)
2. [新增菜单接口](#2-新增菜单接口)

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
