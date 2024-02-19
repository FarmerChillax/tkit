# 错误处理

错误处理的设计上主要是为了框架核心的 Wrap 函数进行配合，包含「http 错误」和「业务错误码」两个方面的思考，下面具体展开讲讲：

## 设计

在错误码上，框架层主要分为 3 个部分，其定义为：
- statusCode: 表示 http 状态码
- code: 表示业务码
- msg: 表示错误信息描述

此外还提供了 id 字段用于可选的返回 RequestID

### 使用设计
在实际开发中，同一个错误码会被反复使用，但其内部的信息可能在某些时候是需要临时修改的（比如 Msg）
因此在 error 的设计上，框架提供了多个 `WithXXX` 方法，这些方法可以临时的将错误码内部信息进行修改，而不影响错误码本身
使用方法：
```go
func (us *UserService) Register(ctx context.Context, req *api.UserRegisterRequest) (*api.EmptyResponse, error) {
	db := tkit.Database.Get(ctx)
	userRepo := repository.NewUserRepo(db)
	err := userRepo.Create(&mysql.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		tkit.Logger.Errorf(ctx, "create user failed: %v", err)
		return nil, errcode.DatabaseError.WithMsg("临时的错误信息")
	}

	return &api.EmptyResponse{}, nil
}
```


## 定义错误码
框架并不提供任何预设的错误码，因此用户需要自己声明错误码（当然也可以使用 layout 内置的）

接口声明如下：
```go
func NewError(httpStatusCode, businessCode int, msg string) Error
```

用例：
```go
var (
    tkit.NewError(400, 10000001, "入参错误")
    // 更多的错误码...
)
```
