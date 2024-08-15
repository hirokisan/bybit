package bybit

import (
	"context"
)

// TimeService :
type TimeService struct {
	client *Client
}

// TimeServiceI :
type TimeServiceI interface {
	GetServerTime(ctx context.Context) (*GetServerTimeResponse, error)
}

// GetServerTimeResponse :
type GetServerTimeResponse struct {
	CommonResponse `json:",inline"`
	Result         GetServerTimeResult `json:"result"`
}

type GetServerTimeResult struct {
	TimeSecond string `json:"timeSecond"`
	TimeNano   string `json:"timeNano"`
}

// GetServerTime :
func (s *TimeService) GetServerTime(ctx context.Context) (*GetServerTimeResponse, error) {
	var res GetServerTimeResponse

	if err := s.client.getPublicly(ctx, "/v3/public/time", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
