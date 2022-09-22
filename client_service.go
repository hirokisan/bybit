package bybit

// SpotService :
type SpotService struct {
	V1 *SpotV1Service
	V3 *SpotV3Service
}

// Spot :
func (c *Client) Spot() *SpotService {
	return &SpotService{
		V1: &SpotV1Service{c},
		V3: &SpotV3Service{c},
	}
}

// FutureService :
type FutureService struct {
	InversePerpetual *FutureInversePerpetualService
	USDTPerpetual    *FutureUSDTPerpetualService
	InverseFuture    *FutureInverseFutureService
}

// Future :
func (c *Client) Future() *FutureService {
	return &FutureService{
		InversePerpetual: &FutureInversePerpetualService{
			client:              c,
			FutureCommonService: &FutureCommonService{c},
		},
		USDTPerpetual: &FutureUSDTPerpetualService{
			client:              c,
			FutureCommonService: &FutureCommonService{c},
		},
		InverseFuture: &FutureInverseFutureService{
			client:              c,
			FutureCommonService: &FutureCommonService{c},
		},
	}
}

// Derivative :
func (c *Client) Derivative() *DerivativeService {
	return &DerivativeService{c}
}

// AccountAsset :
func (c *Client) AccountAsset() *AccountAssetService {
	return &AccountAssetService{c}
}

// CopyTrading :
func (c *Client) CopyTrading() *CopyTradingService {
	return &CopyTradingService{c}
}

// USDCContractService :
type USDCContractService struct {
	Option    *USDCContractOptionService
	Perpetual *USDCContractPerpetualService
}

// USDCContract :
func (c *Client) USDCContract() *USDCContractService {
	return &USDCContractService{
		Option:    &USDCContractOptionService{c},
		Perpetual: &USDCContractPerpetualService{c},
	}
}
