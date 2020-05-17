package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func HandleCommonError(err error) map[string]string {
	errs := err.(validator.ValidationErrors)
	res := make(map[string]string)
	for _, er := range errs {
		f := strings.ToLower(er.Field())
		fmt.Println(f)
		res[f] = er.Tag()
	}
	return res
}
