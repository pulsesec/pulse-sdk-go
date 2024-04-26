# Pulse Security - Go SDK

```go
import "github.com/pulsesec/pulse-sdk-go"
```

## Index

- [Variables](#variables)
- [type Client](#Client)
  - [func New\(siteKey, siteSecret string\) \*Client](#New)
  - [func \(c \*Client\) Classify\(token string\) \(bool, error\)](#Client.Classify)
- [type Error](#Error)
  - [func \(e Error\) Error\(\) string](#Error.Error)
  - [func \(e Error\) Unwrap\(\) error](#Error.Unwrap)

## Variables

<a name="ErrTokenNotFound"></a>

```go
var (
    ErrTokenNotFound = errors.New("user token not found")
    ErrTokenUsed     = errors.New("user token already used")
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
func (c *Client) Classify(token string) (bool, error)
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
