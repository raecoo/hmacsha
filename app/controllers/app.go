package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Calculate() revel.Result {
	key := []byte(c.Params.Form.Get("secret"))
	h := hmac.New(sha256.New, key)
	h.Write([]byte(c.Params.Form.Get("body")))
	result := base64.StdEncoding.EncodeToString(h.Sum(nil))
	c.Flash.Success(result)
	return c.Redirect("/")
}
