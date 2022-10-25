package bybit

// DerivativeUnifiedMarginServiceI :
type DerivativeUnifiedMarginServiceI interface{}

// DerivativeUnifiedMarginService :
type DerivativeUnifiedMarginService struct {
	client *Client

	*DerivativeCommonService
}
