<h1 align="center"><a href="https://www.pulsesecurity.org/">Pulse Security</a></h1>
<p align="center">
<img src="https://avatars.githubusercontent.com/u/161549711?s=200&v=4"/>
</p>
<h1 align="center">Golang SDK</h1>

## Installation

```sh
$ go get github.com/pulsesec/pulse-sdk-go
```

## Example

```go
import (
	"errors"
	"os"

	pulse "github.com/pulsesec/pulse-sdk-go"
)

var (
	client = pulse.New(os.Getenv("PULSE_SITE_KEY"), os.Getenv("PULSE_SECRET_KEY"))
)

func classify(ctx context.Context, token string) bool {
	isBot, err := client.Classify(ctx, token)
	if err != nil {
		if errors.Is(err, pulse.ErrTokenNotFound) {
			panic("Token not found")
		}

		if errors.Is(err, pulse.ErrTokenUsed) {
			panic("Token already used")
		}

		if errors.Is(err, pulse.ErrTokenExpired) {
			panic("Token expired")
		}

		panic(err)
	}

	return isBot
}
```

## Index

- [Variables](#variables)
- [type Client](#Client)
  - [func New\(siteKey, siteSecret string\) \*Client](#New)
  - [func \(c \*Client\) Classify\(ctx context.Context, token string\) \(bool, error\)](#Client.Classify)
- [type Error](#Error)
  - [func \(e Error\) Error\(\) string](#Error.Error)
  - [func \(e Error\) Unwrap\(\) error](#Error.Unwrap)

## Variables

<a name="ErrTokenNotFound"></a>

```go
var (
	ErrTokenNotFound = errors.New("user token not found")
	ErrTokenUsed     = errors.New("user token already used")
	ErrTokenExpired  = errors.New("user token expired")
	ErrOther         = errors.New("user classification failed")
)
```

<a name="Client"></a>

## type Client

```go
type Client struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>

### func New

```go
func New(siteKey, siteSecret string) *Client
```

<a name="Client.Classify"></a>

### func \(\*Client\) Classify

```go
func (c *Client) Classify(ctx context.Context, token string) (bool, error)
```

<a name="Error"></a>

## type Error

```go
type Error struct {
    Message string `json:"error"`
    Code    string `json:"code"`
}
```

<a name="Error.Error"></a>

### func \(Error\) Error

```go
func (e Error) Error() string
```

<a name="Error.Unwrap"></a>

### func \(Error\) Unwrap

```go
func (e Error) Unwrap() error
```
