# ungrouped

Ungrouped apis

1. [](#1-)
2. [](#2-)
3. [用户登录接口](#3-用户登录接口)
4. [刷新token接口](#4-刷新token接口)
5. [](#5-)
6. [](#6-)
7. [](#7-)
8. [](#8-)
9. [](#9-)
10. [](#10-)
11. [](#11-)
12. [](#12-)
13. [](#13-)
14. [](#14-)
15. [](#15-)
16. [](#16-)
17. [](#17-)
18. [](#18-)
19. [](#19-)
20. [](#20-)
21. [](#21-)

## apis

### 1. 

```text
 /
```

### 2. 

```text
 /
```

### 3. 用户登录接口

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

### 4. 刷新token接口

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

### 5. 

```text
 /
```

### 6. 

```text
 /
```

### 7. 

```text
 /
```

### 8. 

```text
 /
```

### 9. 

```text
 /
```

### 10. 

```text
 /
```

### 11. 

```text
 /
```

### 12. 

```text
 /
```

### 13. 

```text
 /
```

### 14. 

```text
 /
```

### 15. 

```text
 /
```

### 16. 

```text
 /
```

### 17. 

```text
 /
```

### 18. 

```text
 /
```

### 19. 

```text
 /
```

### 20. 

```text
 /
```

### 21. 

```text
 /
```
