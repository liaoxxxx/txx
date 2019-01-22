package controllers

type RegisterController struct {
	adminBaseController
}

func (c *RegisterController) Prepare() {
	c.adminBaseController.Prepare()
}
func (c *RegisterController) Get() {
	c.TplName = "admin/register/navbar.html"
}
