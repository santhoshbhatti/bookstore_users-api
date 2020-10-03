package errors

import "net/http"

/*RestErr represents errors in our rest layer returned to api clients*/
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

/*NewBadRequest This is the reusable Bad request */
func NewBadRequest(message string) *RestErr {
	restErr := RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}

	return &restErr
}

/*NewNotFoundError This is the reusable Not found request */
func NewNotFoundError(message string) *RestErr {
	restErr := RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}

	return &restErr
}
