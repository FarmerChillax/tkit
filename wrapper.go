package tkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自动校验请求参数并将其绑定到 DTO 中
// 返回 DTO 与 错误信息
func shouldBindAndValid[T any](c *gin.Context, v T) (T, error) {
	err := c.ShouldBind(v)
	if err != nil {
		// errs, ok := err.(validator.ValidationErrors)
		// if !ok {
		// 	// 尝试从 路由参数 绑定
		// 	if err = c.ShouldBindUri(v); err == nil {
		// 		return v, nil
		// 	}
		// 	return v, errors.New("验证异常")
		// }

		// // 将所有的参数错误进行翻译然后拼装成字符串返回
		// var errMsgs []string
		// for _, value := range errs.Translate(GlobalValidator.Trans) {
		// 	errMsgs = append(errMsgs, value)
		// }

		// return v, errors.New(strings.Join(errMsgs, " \n"))
	}

	return v, err
}

// 该函数利用范型自动绑定请求参数并传递到 controller 中，如果有错误会自动记录错误日志
// 该函数的作用是减少重复代码，提高代码复用率
// Cook Book:
//
//	type AddUserRequest struct {
//		Name  string `json:"name,omitempty" binding:"required" validate:"required"`
//		Phone string `json:"phone,omitempty"`
//	}

//	type AddUserResponse struct {
//		ID int64 `json:"id,omitempty"`
//	}
//
//	func (s *Service) addUser(ctx *gin.Context, req *AddUserRequest) (*AddUserResponse, error) {
//	    // 处理请求...
//	}
//
//	func main() {
//	    router := gin.Default()
//	    router.POST("/user/add", Wrap(s.addUser))
//	}
func Wrap[Request any, Response any](handler func(c *gin.Context, requestDTO *Request) (Response, error)) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		requestDTO := new(Request)

		requestDTO, err := shouldBindAndValid(ctx, requestDTO)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}

		response, err := handler(ctx, requestDTO)
		if err != nil {
			// 检查错误是否为框架定义的错误
			// 如果是框架定义的错误，就返回通用格式的错误详情
			// 如果不是框架定义的错误，就返回 500 错误
			if e, ok := err.(Error); ok {
				ResultError(ctx, e)
			} else {
				ResultError(ctx, NewError(http.StatusInternalServerError, 0, err.Error()))
			}
		}
		// 嵌入通用响应格式并返回结果
		if !ctx.Writer.Written() {
			ResultData(ctx, response)
		}
	}
}
