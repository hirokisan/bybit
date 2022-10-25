package bybit

// DerivativeContractServiceI :
type DerivativeContractServiceI interface{}

// DerivativeContractService :
type DerivativeContractService struct {
	client *Client

	*DerivativeCommonService
}
