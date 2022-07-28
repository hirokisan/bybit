# bybit client integration test

There are tests so that we can get to know the changes of bybit api response.

Before run test, export testnet api key and secret
```console
$ export BYBIT_TEST_KEY=xxx
$ export BYBIT_TEST_SECRET=xxx
```

Also be sure to prepare assets on your testnet wallet

Test with updating golden file
```console
$ make test BYBIT_TEST_UPDATED=true
```

Test
```console
$ make test
```

Test specific method
```
$ make test-spec BYBIT_TEST_METHOD=TestBalance
```
