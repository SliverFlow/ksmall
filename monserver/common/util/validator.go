package util

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// ProcessErr go validator参数校验器自定义规则及提示
func ProcessErr(u interface{}, err error) string {
	if err == nil { //如果为nil 说明校验通过
		return ""
	}
	var invalid *validator.InvalidValidationError
	ok := errors.As(err, &invalid) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return "输入参数错误：" + invalid.Error()
	}
	var validationErrs validator.ValidationErrors
	errors.As(err, &validationErrs) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		typeOf := reflect.TypeOf(u)
		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		field, ok := typeOf.FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("reg_error_info") // 获取field对应的reg_error_info tag值
			return fieldName + ":" + errorInfo           // 返回错误
		} else {
			return "缺失reg_error_info"
		}
	}
	return ""
}

// ValidaMsg go validator参数校验器自定义规则及提示
func ValidaMsg(err error, obj any) string {
	getObt := reflect.TypeOf(obj)
	//将err断言为具体类型
	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		//断言成功
		for _, e := range errs {
			//根绝报错字段名，获取结构体具体字段
			if f, exists := getObt.Elem().FieldByName(e.Field()); exists {
				msg := f.Tag.Get("message")
				if msg == "" {
					msg = "参数错误"
				}
				return msg
			}
		}
	}
	return "参数错误"
}
