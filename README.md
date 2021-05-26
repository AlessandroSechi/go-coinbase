# Golang client for Coinbase V2 API

This client implements all [Coinbase v2 API](https://developers.coinbase.com/api/v2) endpoints.

It works with [API Keys authentication](https://developers.coinbase.com/api/v2#oauth2-coinbase-connect).

## Installation

`go get github.com/AlessandroSechi/go-coinbase`

## Initialization

```go
import "github.com/AlessandroSechi/go-coinbase"

// Create a new client
c := NewClient("<API Key>", "<API Secret>")
```

## Get current user’s public information

```go
user, err := c.GetUser(context.TODO())
```

## Lists current user’s accounts to which the authentication method has access to.

```go
// pagination contains data about pagination https://developers.coinbase.com/api/v2#pagination
accounts, pagination, err := c.ListAccounts(context.TODO())
```