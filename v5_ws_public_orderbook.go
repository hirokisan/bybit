package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// SubscribeOrderBook :
func (s *V5WebsocketPublicService) SubscribeOrderBook(
	key V5WebsocketPublicOrderBookParamKey,
	f func(V5WebsocketPublicOrderBookResponse) error,
) (func() error, error) {
	if err := s.addParamOrderBookFunc(key, f); err != nil {
		return nil, err
	}
	param := struct {
		Op   string        `json:"op"`
		Args []interface{} `json:"args"`
	}{
		Op:   "subscribe",
		Args: []interface{}{key.Topic()},
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
			Args: []interface{}{key.Topic()},
		}
		buf, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if err := s.writeMessage(websocket.TextMessage, []byte(buf)); err != nil {
			return err
		}
		s.removeParamOrderBookFunc(key)
		return nil
	}, nil
}

// V5WebsocketPublicOrderBookParamKey :
type V5WebsocketPublicOrderBookParamKey struct {
	Depth  int
	Symbol SymbolV5
}

// Topic :
func (k *V5WebsocketPublicOrderBookParamKey) Topic() string {
	return fmt.Sprintf("%s.%d.%s", V5WebsocketPublicTopicOrderBook, k.Depth, k.Symbol)
}

// V5WebsocketPublicOrderBookResponse :
type V5WebsocketPublicOrderBookResponse struct {
	Topic     string                         `json:"topic"`
	Type      string                         `json:"type"`
	TimeStamp int64                          `json:"ts"`
	Data      V5WebsocketPublicOrderBookData `json:"data"`
}

// V5WebsocketPublicOrderBookData :
type V5WebsocketPublicOrderBookData struct {
	Symbol   SymbolV5                       `json:"s"`
	Bids     V5WebsocketPublicOrderBookBids `json:"b"`
	Asks     V5WebsocketPublicOrderBookAsks `json:"a"`
	UpdateID int                            `json:"u"`
	Seq      int                            `json:"seq"`
}

// V5WebsocketPublicOrderBookBids :
type V5WebsocketPublicOrderBookBids []struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

// UnmarshalJSON :
func (b *V5WebsocketPublicOrderBookBids) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := make(V5WebsocketPublicOrderBookBids, len(parsedData))
	for i, item := range parsedData {
		item := item
		if len(item) != 2 {
			return errors.New("so far len(item) must be 2, please check it on documents")
		}
		items[i].Price = item[0]
		items[i].Size = item[1]
	}
	*b = items
	return nil
}

// V5WebsocketPublicOrderBookAsks :
type V5WebsocketPublicOrderBookAsks []struct {
	Price string `json:"price"`
	Size  string `json:"size"`
}

// UnmarshalJSON :
func (b *V5WebsocketPublicOrderBookAsks) UnmarshalJSON(data []byte) error {
	parsedData := [][]string{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}
	items := make(V5WebsocketPublicOrderBookAsks, len(parsedData))
	for i, item := range parsedData {
		item := item
		if len(item) != 2 {
			return errors.New("so far len(item) must be 2, please check it on documents")
		}
		items[i].Price = item[0]
		items[i].Size = item[1]
	}
	*b = items
	return nil
}

// Key :
func (r *V5WebsocketPublicOrderBookResponse) Key() V5WebsocketPublicOrderBookParamKey {
	topic := r.Topic
	arr := strings.Split(topic, ".")
	if arr[0] != V5WebsocketPublicTopicOrderBook.String() || len(arr) != 3 {
		return V5WebsocketPublicOrderBookParamKey{}
	}
	depth, err := strconv.Atoi(arr[1])
	if err != nil {
		return V5WebsocketPublicOrderBookParamKey{}
	}
	symbol := SymbolV5(arr[2])
	return V5WebsocketPublicOrderBookParamKey{
		Depth:  depth,
		Symbol: symbol,
	}
}

// addParamOrderBookFunc :
func (s *V5WebsocketPublicService) addParamOrderBookFunc(key V5WebsocketPublicOrderBookParamKey, f func(V5WebsocketPublicOrderBookResponse) error) error {
	if _, exist := s.paramOrderBookMap[key]; exist {
		return errors.New("already registered for this param")
	}
	s.paramOrderBookMap[key] = f
	return nil
}

// removeParamTradeFunc :
func (s *V5WebsocketPublicService) removeParamOrderBookFunc(key V5WebsocketPublicOrderBookParamKey) {
	delete(s.paramOrderBookMap, key)
}

// retrievePositionFunc :
func (s *V5WebsocketPublicService) retrieveOrderBookFunc(key V5WebsocketPublicOrderBookParamKey) (func(V5WebsocketPublicOrderBookResponse) error, error) {
	f, exist := s.paramOrderBookMap[key]
	if !exist {
		return nil, errors.New("func not found")
	}
	return f, nil
}
