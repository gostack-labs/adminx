# api

1. [分页获取api列表接口](#1-分页获取api列表接口)
2. [新增api接口](#2-新增api接口)
3. [更新api接口](#3-更新api接口)
4. [删除api接口](#4-删除api接口)
5. [批量删除api接口](#5-批量删除api接口)

## apis

### 1. 分页获取api列表接口

```text
GET /sys/api
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
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
  "msg": "abc"  //string
}
```

---

### 2. 新增api接口

```text
POST /sys/api
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "msg": "abc"  //string
}
```

---

### 3. 更新api接口

```text
PUT /sys/api/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

_body_:

```javascript
{  //object(api.updateApiRequest), 更新api请求参数
  "groups": 123,  //int64, required, 所属接口分组
  "method": "abc",  //string, required, 请求方式
  "remark": "abc",  //string, 备注
  "title": "abc",  //string, required, 标题
  "url": "abc"  //string, required, 接口地址
}
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "msg": "abc"  //string
}
```

---

### 4. 删除api接口

```text
DELETE /sys/api/single/:id
```

__Request__:

parameter|parameterType|dataType|required|validate|example|description
--|:-:|:-:|:-:|--|--|--
__id__|_param_|int64|true|required||主键ID

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "msg": "abc"  //string
}
```

---

### 5. 批量删除api接口

```text
DELETE /sys/api/batch
```

__Response__:

```javascript
//StatusCode: 200 
{  //object(resp.resultOK)
  "code": 123,  //int
  "msg": "abc"  //string
}
```

---
