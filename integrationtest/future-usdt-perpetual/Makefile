BYBIT_TEST_UPDATED=false

test:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/hirokisan/bybit/v2/integrationtest/future-usdt-perpetual -v -tags=integrationtestfutureusdtperpetual

test-spec:
	BYBIT_TEST_UPDATED=${BYBIT_TEST_UPDATED} BYBIT_TEST_KEY=${BYBIT_TEST_KEY} BYBIT_TEST_SECRET=${BYBIT_TEST_SECRET} go test github.com/hirokisan/bybit/v2/integrationtest/future-usdt-perpetual -v -tags=integrationtestfutureusdtperpetual -run ${BYBIT_TEST_METHOD}
