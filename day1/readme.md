# 1. Fundamental REST API
**Objectives** :

- Create **Postman collection** (JSON file)
- Create **Postman environment**
- Implement **HTTP Method** (GET, POST, UPDATE, DELETE)

**Tasks** :

Do request to the following API target by using Postman environment, save the result using "Save Response" (Save as example), then export collection:

[https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0](https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0)

| Method | Endpoint | Requirement |
| --- | --- | --- |
| GET | /book | Get all book |
| GET | /book/{id} | Get book by id |
| POST | /book | Create new book |
| UPDATE | /book/{id} | Update book by id |
| DELETE | /book/{id} | Delete book by id |

Disclaimer : this is only a mock-up, you don't need to worry the JSON output, just make sure the response code is OK (200). See the swagger documentation [here](https://app.swaggerhub.com/apis-docs/sepulsa/RentABook-API/1.0.0).

[https://testnet.binance.vision](https://testnet.binance.vision/)

| Method | Endpoint | Requirement | Reference |
| --- | --- | --- | --- |
| GET | /api/v1/klines | Get recent 1000 BTCUSDT klines data with 1 minute interval | [https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-data](https://binance-docs.github.io/apidocs/spot/en/#kline-candlestick-data) |
| GET | /api/v1/klines | Get BTCUSDT klines data with 1 day interval, start from 1 September 2022 to 7 September 2022 (UTC) |
| GET | /api/v3/account | Get information of your account | [https://binance-docs.github.io/apidocs/spot/en/#account-information-user\_data](https://binance-docs.github.io/apidocs/spot/en/#account-information-user_data) |

You can create API Key and Secret [here](https://testnet.binance.vision/) (Need login with Github first), by clicking "Generate HMAC\_SHA256 Key". Make sure you save the API Key and Secret, because it only shows once.

The last endpoint need some setup in order to get correct result

- Create an environment with variables:
    - api\_key, api\_secret, timestamp, signature
- put "X-MBX-APIKEY" as header key and api key as value
- copy paste this [code](https://pastebin.com/09SAB4ew) to Pre-request Script