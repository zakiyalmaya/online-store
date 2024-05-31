package model

type ResponseSystem struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HTTPSuccessResponse(res interface{}) ResponseSystem {
	return ResponseSystem{
		Message: "success",
		Data:    res,
	}
}

func HTTPErrorResponse(errMsg string) ResponseSystem {
	return ResponseSystem{
		Message: errMsg,
	}
}
