package bybit

import (
	"encoding/json"
	"errors"

	"github.com/gorilla/websocket"
)

// SubscribeWallet :
func (s *V5WebsocketPrivateService) SubscribeWallet(
	f func(V5WebsocketPrivateWalletResponse) error,
) (func() error, error) {
	key := V5WebsocketPrivateParamKey{
		Topic: V5WebsocketPrivateTopicWallet,
	}
	if err := s.addParamWalletFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{V5WebsocketPrivateTopicWallet},
	}
	buf, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	if err := s.writeMessage(websocket.TextMessage, buf); err != nil {
		return nil, err
	}
	return func() error {
		param := struct {
			Op   string        `json:"op"`
			Args []interface{} `json:"args"`
		}{
			Op:   "unsubscribe",
			Args: []interface{}{V5WebsocketPrivateTopicWallet},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamWalletFunc(key)
		return nil
	}, nil
}

// V5WebsocketPrivateWalletResponse :
type V5WebsocketPrivateWalletResponse struct {
	ID           string                         `json:"id"`
	Topic        V5WebsocketPrivateTopic        `json:"topic"`
	CreationTime int64                          `json:"creationTime"`
	Data         []V5WebsocketPrivateWalletData `json:"data"`
}

// V5WebsocketPrivateWalletData :
type V5WebsocketPrivateWalletData struct {
	AccountIMRate          string                         `json:"accountIMRate"`
	AccountMMRate          string                         `json:"accountMMRate"`
	TotalEquity            string                         `json:"totalEquity"`
	TotalWalletBalance     string                         `json:"totalWalletBalance"`
	TotalMarginBalance     string                         `json:"totalMarginBalance"`
	TotalAvailableBalance  string                         `json:"totalAvailableBalance"`
	TotalPerpUPL           string                         `json:"totalPerpUPL"`
	TotalInitialMargin     string                         `json:"totalInitialMargin"`
	TotalMaintenanceMargin string                         `json:"totalMaintenanceMargin"`
	AccountType            AccountType                    `json:"accountType"`
	Coins                  []V5WebsocketPrivateWalletCoin `json:"coin"`
}

// V5WebsocketPrivateWalletCoin :
type V5WebsocketPrivateWalletCoin struct {
	Coin                Coin   `json:"coin"`
	Equity              string `json:"equity"`
	UsdValue            string `json:"usdValue"`
	WalletBalance       string `json:"walletBalance"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	AvailableToBorrow   string `json:"availableToBorrow"`
	BorrowAmount        string `json:"borrowAmount"`
	AccruedInterest     string `json:"accruedInterest"`
	TotalOrderIM        string `json:"totalOrderIM"`
	TotalPositionIM     string `json:"totalPositionIM"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
}

// Key :
func (r *V5WebsocketPrivateWalletResponse) Key() V5WebsocketPrivateParamKey {
	return V5WebsocketPrivateParamKey{
		Topic: r.Topic,
	}
}

// addParamWalletFunc :
func (s *V5WebsocketPrivateService) addParamWalletFunc(param V5WebsocketPrivateParamKey, f func(V5WebsocketPrivateWalletResponse) error) error {
	if _, exist := s.paramWalletMap[param]; exist {
		return errors.New("already registered for this param")
	}
	s.paramWalletMap[param] = f
	return nil
}

// removeParamWalletFunc :
func (s *V5WebsocketPrivateService) removeParamWalletFunc(key V5WebsocketPrivateParamKey) {
	delete(s.paramWalletMap, key)
}

// retrieveWalletFunc :
func (s *V5WebsocketPrivateService) retrieveWalletFunc(key V5WebsocketPrivateParamKey) (func(V5WebsocketPrivateWalletResponse) error, error) {
	f, exist := s.paramWalletMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
