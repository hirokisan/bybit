package bybit

import (
	"errors"
	"fmt"
	"time"
)

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
