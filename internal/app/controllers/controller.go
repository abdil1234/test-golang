package controllers

import (
	"strconv"
	"time"

	"github.com/abdil1234/test-golang/internal/app/commons"
	"github.com/abdil1234/test-golang/internal/app/commons/constants"
	"github.com/abdil1234/test-golang/internal/app/usecases"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RequestValue string

type ControllerOption struct {
	commons.Option
	*usecases.Services
}

func (ctl *ControllerOption) SetPostParams(c *gin.Context, params interface{}) error {
	if params == nil {
		return nil
	}
	b := binding.Default(c.Request.Method, c.ContentType())
	var i interface{} = b
	var err error
	bBody, ok := i.(binding.BindingBody)
	if ok {
		err = c.ShouldBindBodyWith(params, bBody)
	} else {
		err = c.ShouldBind(params)
	}

	if err != nil {
		_ = c.Error(err)
		return err
	}
	return nil
}

func (ctl *ControllerOption) GetParamValue(c *gin.Context, param string) (result RequestValue) {
	value := RequestValue(c.Request.URL.Query().Get(param))
	return value
}

func (ctl *ControllerOption) GetFormValue(c *gin.Context, param string) (result RequestValue) {
	value := RequestValue(c.PostForm(param))
	return value
}

func (r RequestValue) ToString() (result string) {
	return string(r)
}

func (r RequestValue) ToInt() (result int, err error) {
	return strconv.Atoi(string(r))
}

func (r RequestValue) ToDate() (result time.Time, err error) {
	return time.Parse(constants.DateFormat, string(r))
}
