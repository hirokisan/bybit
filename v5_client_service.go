package bybit

// V5ServiceI :
type V5ServiceI interface {
	Market() V5MarketServiceI
	Order() V5OrderServiceI
	Position() V5PositionServiceI
	Execution() V5ExecutionServiceI
	Account() V5AccountServiceI
	SpotLeverageToken() V5SpotLeverageTokenServiceI
	SpotMarginTrade() V5SpotMarginTradeServiceI
	Asset() V5AssetServiceI
	User() V5UserServiceI
}

// V5Service :
type V5Service struct {
	client *Client
}

// Market :
func (s *V5Service) Market() V5MarketServiceI {
	return &V5MarketService{s.client}
}

// Order :
func (s *V5Service) Order() V5OrderServiceI {
	return &V5OrderService{s.client}
}

// Position :
func (s *V5Service) Position() V5PositionServiceI {
	return &V5PositionService{s.client}
}

// Execution :
func (s *V5Service) Execution() V5ExecutionServiceI {
	return &V5ExecutionService{s.client}
}

// Account :
func (s *V5Service) Account() V5AccountServiceI {
	return &V5AccountService{s.client}
}

// SpotLeverageToken :
func (s *V5Service) SpotLeverageToken() V5SpotLeverageTokenServiceI {
	return &V5SpotLeverageTokenService{s.client}
}

// SpotMarginTrade :
func (s *V5Service) SpotMarginTrade() V5SpotMarginTradeServiceI {
	return &V5SpotMarginTradeService{s.client}
}

// Asset :
func (s *V5Service) Asset() V5AssetServiceI {
	return &V5AssetService{s.client}
}

// User :
func (s *V5Service) User() V5UserServiceI {
	return &V5UserService{s.client}
}

// V5 :
func (c *Client) V5() V5ServiceI {
	return &V5Service{c.withCheckResponseBody(checkV5ResponseBody)}
}
