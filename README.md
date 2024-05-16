# wx-transfer

提供第三方调用微信云函数能力的中转站后端

## 快速开始

```bash
git clone git@github.com:Ninohana/wx-transfer.git
cd wx-transfer
```

在`global.go`中配置Appid、AppSecret和EnvId（云环境ID）

编译、执行：

```bash
go build main.go
main.exe
# or
go run main.go
```

### 使用方式

POST方式请求`/invoke`，并携带对应参数。

```javascript
// example
var data = {
    method: "login",
    params: {
        username: "admin",
        password: "admin"
    }
}
// func is your cloud function name on the weixin cloud function
fetch("http://localhost:1126/invoke?func=cloud", {
        method: "POST",
        body: JSON.stringify(data),
      }).then(res => console.log(res))
        .catch((e) => console.error(e));
```
