# Sirius API

## Introduction

Sirius is an RESTful API that allows developers to integrate the Stellar blockchain into your application, supporting the following features:

- Account creation
- Account balance retrieval
- Account transaction history retrieval
- Account transaction submission
- Account transaction status retrieval
- Account transaction fee retrieval
- Address validation

## Table of Contents

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20 or higher

### Installation

1. Clone the repo

```sh
git clone
```

2. Install Go dependencies

```sh
go mod download
```

3. Run the server

```sh
go run main.go
```

## Usage

### Account Creation

#### Request

`POST /account`

```json
{
  "network": "testnet"
}
```

#### Response

```
{
  "address": "***"
}
```

### Account Balance Retrieval

#### Request

`GET /account/balance`

```json
{
  "network": "testnet",
  "address": "***"
}
```

#### Response

```
{
  "balance": "***"
}
```

### Account Transaction History Retrieval

#### Request

`GET /account/transactions`

```json
{
  "network": "testnet",
  "address": "***"
}
```

#### Response

```

```

### Account Transaction Submission

#### Request

`POST /account/transaction`

```json
{
  "network": "testnet",
  "address": "***",
  "destination": "***",
  "amount": "***",
  "memo": "***"
}
```

#### Response

```
{
  "hash": "***"
}
```

### Account Transaction Status Retrieval

#### Request

`GET /account/transaction`

```json
{
  "network": "testnet",
  "address": "***",
  "hash": "***"
}
```

#### Response

```
{
  "status": "***"
}
```

### Account Transaction Fee Retrieval

#### Request

`GET /account/transaction/fee`

```json
{
  "network": "testnet",
  "address": "***"
}
```

#### Response

```
{
  "fee": "***"
}
```

### Address Validation

#### Request

`GET /address/validate`

```json
{
  "network": "testnet",
  "address": "***"
}
```

#### Response

```
{
  "valid": true
}
```

## Tests

To run tests, run the following command

```sh
go test
```

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request
