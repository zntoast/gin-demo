package utils

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// 传参校验
func GetValidMsg(err error, obj interface{}) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg")
			}
		}
	}
	return err.Error()
}

// 判断是否为手机号
func VerifyPhoneNumber(phone string) error {
	res := regexp.MustCompile(`1[3456789]\d{9}`).FindAllStringSubmatch(phone, -1)
	if len(res) == 0 {
		return fmt.Errorf("手机号不合法")
	}
	return nil
}
