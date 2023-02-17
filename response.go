package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type checkResponseBodyFunc func([]byte) error

func checkResponseBody(body []byte) error {
	var commonResponse CommonResponse
	if err := json.Unmarshal(body, &commonResponse); err != nil {
		return err
	}

	switch {
	case commonResponse.RetCode == 10006:
		rateLimitError := &RateLimitError{}
		if err := json.Unmarshal(body, rateLimitError); err != nil {
			return err
		}
		return rateLimitError
	case commonResponse.RetCode != 0:
		return &ErrorResponse{
			RetCode: commonResponse.RetCode,
			RetMsg:  commonResponse.RetMsg,
		}
	default:
		return nil
	}
}

func checkV3ResponseBody(body []byte) error {
	var commonResponse CommonV3Response
	if err := json.Unmarshal(body, &commonResponse); err != nil {
		return err
	}

	switch {
	case commonResponse.RetCode != 0:
		return &ErrorResponse{
			RetCode: commonResponse.RetCode,
			RetMsg:  commonResponse.RetMsg,
		}
	default:
		return nil
	}
}

func checkV5ResponseBody(body []byte) error {
	var commonResponse CommonV5Response
	if err := json.Unmarshal(body, &commonResponse); err != nil {
		return err
	}

	switch {
	case commonResponse.RetCode == 10006, commonResponse.RetCode == 10018:
		rateLimitError := &RateLimitError{}
		if err := json.Unmarshal(body, rateLimitError); err != nil {
			return err
		}
		return rateLimitError

	case commonResponse.RetCode != 0:
		return &ErrorResponse{
			RetCode: commonResponse.RetCode,
			RetMsg:  commonResponse.RetMsg,
		}
	default:
		return nil
	}
}

// CommonResponse :
type CommonResponse struct {
	RetCode          int    `json:"ret_code"`
	RetMsg           string `json:"ret_msg"`
	ExtCode          string `json:"ext_code"`
	ExtInfo          string `json:"ext_info"`
	TimeNow          string `json:"time_now"`
	RateLimitStatus  int    `json:"rate_limit_status"`
	RateLimitResetMs int    `json:"rate_limit_reset_ms"`
	RateLimit        int    `json:"rate_limit"`
}

// CommonV3Response :
type CommonV3Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

// CommonV5Response :
type CommonV5Response struct {
	RetCode    int         `json:"retCode"`
	RetMsg     string      `json:"retMsg"`
	RetExtInfo interface{} `json:"retExtInfo"`
	Time       int         `json:"time"`
}

// ErrorResponse :
type ErrorResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
}

// Error :
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d, %s", r.RetCode, r.RetMsg)
}

// RateLimitError :
type RateLimitError struct {
	*CommonResponse `json:",inline"`
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%s, %s", r.RetMsg, time.Until(time.Unix(int64(r.RateLimitResetMs/1000), 0)))
}

var (
	// ErrPathNotFound : Request path not found
	ErrPathNotFound = errors.New("path not found")
	// ErrAccessDenied : Access denied
	ErrAccessDenied = errors.New("access denied")
)
