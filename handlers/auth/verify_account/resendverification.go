package handlers

import (
	dto "greet-auth-srv/dto/auth"
	res "greet-auth-srv/util/response"

	"github.com/labstack/echo/v4"
)

func (u *verifyHandler) ResendVerification(c echo.Context) error {
	var req dto.ResendVerificationRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	token := c.QueryParam("token")
	req.Token = token

	data, err := u.serviceVerify.ResendVerifyUserByToken(req)
	if err != nil {
		return c.HTML(400, "Verification failed: "+err.Error())
	}

	return res.SuccessResponse(data).Send(c)
}
