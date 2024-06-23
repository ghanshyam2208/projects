package handlers

import (
	"banking_app2/cmd/internals/dto"
	h "banking_app2/cmd/utils/helpers"
	"banking_app2/cmd/utils/logger"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (r *UserHandlers) GetAllUsers(ctx echo.Context) error {
	page, stdErr := strconv.Atoi(ctx.QueryParam("page"))
	if stdErr != nil {
		logger.Error("validation failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}
	pageSize, stdErr := strconv.Atoi(ctx.QueryParam("pageSize"))
	if stdErr != nil {
		logger.Error("validation failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}
	if page < 1 || pageSize != 10 {
		logger.Error("validation failed for page and pageSize ")
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: "page must be greater than 1 and page size must be 10",
			ErrorData: ErrorToMap(errors.New("page must be greater than 1 and page size must be 10")),
		})
	}
	users, stdErr := r.service.GetAllUsers(page, pageSize)
	if stdErr != nil {
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}

	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error: false,
		Code:  http.StatusOK,
		Data: map[string]interface{}{
			"account": users,
		},
		Message: "accounts fetched successfully",
	})
}

func (r *UserHandlers) CreateUser(ctx echo.Context) error {
	var createUserRequest dto.CreateUserDto
	stdErr := ctx.Bind(&createUserRequest)

	if stdErr != nil {
		logger.Error("binding request failed " + stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusBadRequest,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}

	// Validate the request struct
	stdErr = r.validator.Struct(&createUserRequest)
	if stdErr != nil {
		return handleValidationError(ctx, stdErr)
	}

	// call service
	user, stdErr := r.service.CreateUser(createUserRequest)
	if stdErr != nil {
		logger.Error(stdErr.Error())
		return h.WriteErrorApiResponse(ctx, h.ErrorApiResponse{
			Error:     true,
			Code:      http.StatusInternalServerError,
			ErrorInfo: stdErr.Error(),
			ErrorData: ErrorToMap(stdErr),
		})
	}

	return h.WriteSuccessApiResponse(ctx, h.SuccessApiResponse{
		Error: false,
		Code:  http.StatusOK,
		Data: map[string]interface{}{
			"user": user,
		},
		Message: "user created successfully",
	})

}
