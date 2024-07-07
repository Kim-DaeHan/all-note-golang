package dto

// APIResponse info
// @Description api response type
type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
} //@name APIResponse

// APIResponseWithoutData info
// @Description api response type without data field
type APIResponseWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
} //@name APIResponseWithoutData
