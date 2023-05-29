package middlewares

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func getBodyParams(c *gin.Context, bodyParams map[string]interface{}) {
	if c.Request.Method == http.MethodGet || c.Request.Body == http.NoBody {
		return
	}

	b := binding.Default(c.Request.Method, c.ContentType())
	var i interface{} = b
	var err error
	bBody, ok := i.(binding.BindingBody)
	if ok {
		if bBody.Name() == "json" {
			err = c.ShouldBindBodyWith(&bodyParams, bBody) // application/json
		}
	} else if b == binding.FormMultipart {
		err = c.Request.ParseMultipartForm(0) // multipart/form-data
		assignMultipartForm(bodyParams, c.Request.MultipartForm)
	} else if b == binding.Form {
		err = c.Request.ParseForm() // application/x-www-form-urlencoded
		for key, value := range c.Request.PostForm {
			bodyParams[key] = value
		}
	} else {
		err = c.ShouldBind(&bodyParams)
	}

	if err != nil {
		_ = errors.New(fmt.Sprintf("Get body params error. Error: %v", err))
	}
}

func assignMultipartForm(bodyParams map[string]interface{}, multipartForm *multipart.Form) {
	for key, value := range multipartForm.Value {
		bodyParams[key] = value
	}
	for key, files := range multipartForm.File {
		var fileDetails []map[string]interface{}
		for _, file := range files {
			fileDetail := make(map[string]interface{})
			fileDetail["filename"] = file.Filename
			fileDetail["header"] = file.Header
			fileDetail["filesize"] = file.Size
			fileDetails = append(fileDetails, fileDetail)
		}
		bodyParams[key] = fileDetails
	}
}
