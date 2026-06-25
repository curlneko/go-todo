package successes

type SuccessResponse struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(data any) SuccessResponse {
	return SuccessResponse{
		Code: "SUCCESS",
		Data: data,
	}
}
