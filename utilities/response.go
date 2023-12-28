package utilities

const (
	SignupSuccess = "User has been signed up successfully"
	LoginSuccess  = "User has been logged in successfully"
	SignupFailed  = "Signup process has been failed"
	LoginFailed   = "Login process has been failed"
	LogoutFailed  = "Logout process has been failed"
	EmptyFIeldErr = "Email and Password are required"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

func ErrorResponse(message string, err error) errorResponse {
	return errorResponse{
		Success: false,
		Message: message,
		Error:   err.Error(),
	}
}

type successResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) successResponse {
	return successResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
}
