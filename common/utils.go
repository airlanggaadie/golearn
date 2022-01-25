package common

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/mail"
	"time"
)

// loc variable instance to time location that store default timezone
var loc *time.Location

// SetTimezone /** Set default timezone or customize timezone on this app.
func SetTimezone(tz string) error {
	if tz == "" {
		tz = ViperEnvVariable(new(ViperParameters).KeyWithDefaultConfig("TIMEZONE",""))
	}
	location, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	loc = location
	return nil
}

// GetTime /** convert time to default timezone
func GetTime(t time.Time) time.Time {
	return t.In(loc)
}

func IsEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Bind Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

// CommonError My own Error type that will help return my customized Error info
//  {"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
//func NewValidatorError(err error) CommonError {
//	res := CommonError{}
//	res.Errors = make(map[string]interface{})
//	errs := err.(validator.ValidationErrors)
//	for _, v := range errs {
//		// can translate each error one at a time.
//		//fmt.Println("gg",v.NameNamespace)
//		if v.Param != "" {
//			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
//		} else {
//			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
//		}
//
//	}
//	return res
//}