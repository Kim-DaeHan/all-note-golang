package dto

// APIResponse info
// @Description api response type
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
} //@name APIResponse
