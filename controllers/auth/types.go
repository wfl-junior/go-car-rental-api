package controllers

type RegisterBody struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type LoginBody struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}
