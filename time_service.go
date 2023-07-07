package bybit

// TimeService :
type TimeService struct {
	client *Client
}

// TimeServiceI :
type TimeServiceI interface {
	GetServerTime() (*GetServerTimeResponse, error)
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
func (s *TimeService) GetServerTime() (*GetServerTimeResponse, error) {
	var res GetServerTimeResponse

	if err := s.client.getPublicly("/v3/public/time", nil, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
