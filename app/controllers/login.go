package controllers

import "github.com/revel/revel"

type Login struct {
	*revel.Controller
}

func (c Login) LoginPage() revel.Result {
	return c.Render()
}

func (c Login) DoLogin() revel.Result {

	return c.Render(c.Request.Form)
}
