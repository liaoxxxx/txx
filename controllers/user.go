package controllers

type UserController struct {
	indexBaseController
}

func (c *UserController) Prepare() {
}
func (c *UserController) Get() {
	c.TplName = "index/register/navbar.html"
}

func (c *UserController) Login() {
	c.TplName = "index/user/login.html"
}
