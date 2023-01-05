package types

type ErrorHttpResponse struct {
	Code    int    `json:"statusCode"`
	Message string `json:"message"`
}
