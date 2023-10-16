package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func handleBindErr(ctx *gin.Context, err error) bool {
	if ctx.ContentType() != gin.MIMEJSON {
		ctx.JSON(http.StatusUnsupportedMediaType,
			ErrorResponse{
				Error:   "UNSUPPORTED_PAYLOAD_TYPE",
				Message: "Invalid content type",
			},
		)
		return true
	}

	if err == nil {
		return false
	}

	// Handle ValidationErrors
	if errs, ok := err.(validator.ValidationErrors); ok {
		invalidArgs := parseValidationErrors(errs)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"type":    "VALIDATION",
			"message": "Invalid input params",
			"args":    invalidArgs,
		})
		return true
	}

	ctx.JSON(http.StatusBadRequest,
		ErrorResponse{
			Error:   "VALIDATION",
			Message: err.Error(),
		},
	)
	return true
}

type invalidArgument struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

// parseValidationErrors converts BindJson errors to user friendly format
func parseValidationErrors(errs validator.ValidationErrors) []invalidArgument {
	var invalidArgs []invalidArgument

	for _, err := range errs {
		invalidArgs = append(invalidArgs, invalidArgument{
			err.Field(),
			err.Tag(),
		})
	}
	return invalidArgs
}

func errorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Message: err.Error(),
	}
}
