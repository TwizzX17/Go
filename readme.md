# Stock Data Service

This service provides endpoints for managing stock symbols, including fetching symbols from a remote endpoint and saving them to the database.

## Endpoints

### 1. Get Symbols

- **URL**: `/api/symbols`
- **Method**: `GET`
- **Description**: Retrieves a list of stock symbols from the database.
- **Response**: A JSON array of symbol objects.

#### Example Request
```
sh curl -X GET http://localhost:8080/api/symbols
```

#### Example Response
```
[
    {
        "id": 1,
        "symbol": "AAPL",
        "name": "Apple Inc.",
        "exchange": "NASDAQ",
        "assetType": "Stock",
        "status": true,
        "createdTs": "2023-01-01T00:00:00Z",
        "deletedTs": null
    },
    {
        "id": 2,
        "symbol": "GOOGL",
        "name": "Alphabet Inc.",
        "exchange": "NASDAQ",
        "assetType": "Stock",
        "status": true,
        "createdTs": "2023-01-01T00:00:00Z",
        "deletedTs": null
    }
]
```


#### Fetch Symbols to DB
- **URL**: `/api/scheduledSync`
- **Method**: `POST`
- **Description**: Triggers a scheduled sync to fetch the latest symbols from a remote endpoint and save them to the database.
- **Response**: A JSON object indicating the success or failure of the operation.


## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/stock_data.git
    cd stock_data
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

3. Set the necessary environment variables:
    ```sh
    export API_KEY=your_api_key
    export BASE_URI=https://www.alphavantage.co

    export DB_HOST=`insert value`
    export DB_PORT=`insert value`
    export DB_USER=`insert value`
    export DB_PASSWORD=`insert value`
    export DB_NAME=`insert value`
    ```

## Running the Service

To run the service, use the following command:

```sh
go run main.go