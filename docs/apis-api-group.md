# 接口分组管理相关接口

接口分组管理相关接口

1. [分页获取接口分组接口](#1-分页获取接口分组接口)
2. [新增接口分组接口](#2-新增接口分组接口)
3. [更新接口分组接口](#3-更新接口分组接口)
4. [删除接口分组接口](#4-删除接口分组接口)
5. [新增接口分组接口](#5-新增接口分组接口)

## apis

### 1. 分页获取接口分组接口

```text
GET /sys/api-group
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "data": [  //array[db.ApiGroup]
    {  //object(db.ApiGroup)
      "created_at": "2022-05-16T16:47:48.741899+08:00",  //object(time.Time)
      "id": 123,  //int64
      "name": "abc",  //string
      "remark": "abc"  //string
    }
  ],
  "msg": "获取成功"  //string
}
```

---

### 2. 新增接口分组接口

```text
POST /sys/api-group
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

### 3. 更新接口分组接口

```text
PUT /sys/api-group/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

_body_:

```javascript
{  //object(api.updateApiGroupRequest), 更新接口分组请求参数
  "name": "abc",  //string, required, 接口分组名称
  "remark": "abc"  //string, required, 备注
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

### 4. 删除接口分组接口

```text
DELETE /sys/api-group/single/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 10000,  //int
  "msg": "删除成功"  //string
}
```

---

### 5. 新增接口分组接口

```text
DELETE /sys/api-group/batch
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
