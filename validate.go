package tkit

// import (
// 	"bytes"
// 	"errors"
// 	"reflect"

// 	"github.com/go-playground/locales/zh"
// 	"github.com/go-playground/validator/v10"

// 	ut "github.com/go-playground/universal-translator"
// 	translations "github.com/go-playground/validator/v10/translations/zh"
// )

// type TkitValidator struct {
// 	Validate *validator.Validate
// 	Trans    ut.Translator
// }

// // GlobalValidator 全局验证器
// var GlobalValidator TkitValidator

// // InitValidator 初始化验证器
// func InitValidator() {
// 	zhs := zh.New()
// 	uni := ut.New(zhs, zhs)
// 	trans, ok := uni.GetTranslator("zh")
// 	if !ok {
// 		panic(ok)
// 	}

// 	validate := validator.New()

// 	// 收集结构体中的comment标签，用于替换英文字段名称，这样返回错误就能展示中文字段名称了
// 	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 		return fld.Tag.Get("comment")
// 	})

// 	// 注册中文翻译
// 	err := translations.RegisterDefaultTranslations(validate, trans)
// 	if err != nil {
// 		panic(err)
// 	}

// 	GlobalValidator.Validate = validate
// 	GlobalValidator.Trans = trans
// }

// // Check 验证器通用验证方法
// func (m *TkitValidator) Check(value interface{}) error {
// 	err := m.Validate.Struct(value)
// 	if err != nil {
// 		errs, ok := err.(validator.ValidationErrors)
// 		if !ok {
// 			return errors.New("验证异常")
// 		}

// 		// 将所有的参数错误进行翻译然后拼装成字符串返回
// 		errBuf := bytes.Buffer{}
// 		for i := 0; i < len(errs); i++ {
// 			errBuf.WriteString(errs[i].Translate(m.Trans) + " \n")
// 		}

// 		// 删除掉最后一个空格和换行符
// 		errStr := errBuf.String()
// 		return errors.New(errStr[:len(errStr)-2])
// 	}

// 	// 如果它实现了CanCheck接口，就进行自定义验证
// 	if v, ok := value.(CanCheck); ok {
// 		return v.Check()
// 	}
// 	return nil
// }

// // CanCheck 如果需要特殊校验，可以实现验证接口，或者通过自定义tag标签实现
// type CanCheck interface {
// 	Check() error
// }
