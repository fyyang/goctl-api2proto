# goctl-api2proto

### 1. 编译goctl-api2proto插件

```
GOPROXY=https://goproxy.cn/,direct
go install github.com/fyyang/goctl-api2proto@latest
```

### 2. 配置环境

将$GOPATH/bin中的goctl-api2proto添加到环境变量

### 3. 使用姿势

* 创建api文件

    ```go
    info(
     title: "type title here"
     desc: "type desc here"
     author: "type author here"
     email: "type email here"
     version: "type version here"
    )
    
    
    type (
     RegisterReq {
      Username string `json:"username"`
      Password string `json:"password"`
      Mobile string `json:"mobile"`
     }
     
     LoginReq {
      Username string `json:"username"`
      Password string `json:"password"`
     }
     
     UserInfoReq {
      Id string `path:"id"`
     }
     
     UserInfoReply {
      Name string `json:"name"`
      Age int `json:"age"`
      Birthday string `json:"birthday"`
      Description string `json:"description"`
      Tag []string `json:"tag"`
     }
     
     UserSearchReq {
      KeyWord string `form:"keyWord"`
     }
    )
    
    service user-api {
     @doc(
      summary: "注册"
     )
     @handler register
     post /api/user/register (RegisterReq)
     
     @doc(
      summary: "登录"
     )
     @handler login
     post /api/user/login (LoginReq)
     
     @doc(
      summary: "获取用户信息"
     )
     @handler getUserInfo
     get /api/user/:id (UserInfoReq) returns (UserInfoReply)
     
     @doc(
      summary: "用户搜索"
     )
     @handler searchUser
     get /api/user/search (UserSearchReq) returns (UserInfoReply)
    }
    ```

* 生成user.proto 文件

    ```shell script
    goctl api plugin -plugin goctl-api2proto="api2proto -filename user.proto" -api user.api -dir .
    ```