package bybit

// SpotServiceI :
type SpotServiceI interface {
	V1() SpotV1ServiceI
	V3() *SpotV3Service
}

// SpotService :
type SpotService struct {
	client *Client
}

// V1 :
func (s *SpotService) V1() SpotV1ServiceI {
	return &SpotV1Service{s.client}
}

// V3 :
func (s *SpotService) V3() *SpotV3Service {
	return &SpotV3Service{s.client}
}

// Spot :
func (c *Client) Spot() SpotServiceI {
	return &SpotService{c}
}

// NewTimeService :
func (c *Client) NewTimeService() TimeServiceI {
	return &TimeService{c}
}

// FutureServiceI :
type FutureServiceI interface {
	InversePerpetual() FutureInversePerpetualServiceI
	USDTPerpetual() FutureUSDTPerpetualServiceI
	InverseFuture() FutureInverseFutureServiceI
}

// FutureService :
type FutureService struct {
	client *Client
}

// InversePerpetual :
func (s *FutureService) InversePerpetual() FutureInversePerpetualServiceI {
	return &FutureInversePerpetualService{
		client:              s.client,
		FutureCommonService: &FutureCommonService{s.client},
	}
}

// USDTPerpetual :
func (s *FutureService) USDTPerpetual() FutureUSDTPerpetualServiceI {
	return &FutureUSDTPerpetualService{
		client:              s.client,
		FutureCommonService: &FutureCommonService{s.client},
	}
}

// InverseFuture :
func (s *FutureService) InverseFuture() FutureInverseFutureServiceI {
	return &FutureInverseFutureService{
		client:              s.client,
		FutureCommonService: &FutureCommonService{s.client},
	}
}

// Future :
func (c *Client) Future() FutureServiceI {
	return &FutureService{c}
}

// DerivativeServiceI :
type DerivativeServiceI interface {
	UnifiedMargin() DerivativeUnifiedMarginServiceI
	Contract() DerivativeContractServiceI
}

// DerivativeService :
type DerivativeService struct {
	client *Client
}

// UnifiedMargin :
func (s *DerivativeService) UnifiedMargin() DerivativeUnifiedMarginServiceI {
	return &DerivativeUnifiedMarginService{
		client:                  s.client,
		DerivativeCommonService: &DerivativeCommonService{s.client},
	}
}

// Contract :
func (s *DerivativeService) Contract() DerivativeContractServiceI {
	return &DerivativeContractService{
		client:                  s.client,
		DerivativeCommonService: &DerivativeCommonService{s.client},
	}
}

// Derivative :
func (c *Client) Derivative() DerivativeServiceI {
	return &DerivativeService{c.withCheckResponseBody(checkV3ResponseBody)}
}

// AccountAsset :
func (c *Client) AccountAsset() *AccountAssetService {
	return &AccountAssetService{c}
}

// CopyTrading :
func (c *Client) CopyTrading() *CopyTradingService {
	return &CopyTradingService{c}
}

// USDCContractServiceI :
type USDCContractServiceI interface {
	Option() *USDCContractOptionService
	Perpetual() *USDCContractPerpetualService
}

// USDCContractService :
type USDCContractService struct {
	client *Client
}

// Option :
func (s *USDCContractService) Option() *USDCContractOptionService {
	return &USDCContractOptionService{s.client}
}

// Perpetual :
func (s *USDCContractService) Perpetual() *USDCContractPerpetualService {
	return &USDCContractPerpetualService{s.client}
}

// USDCContract :
func (c *Client) USDCContract() USDCContractServiceI {
	return &USDCContractService{c}
}
