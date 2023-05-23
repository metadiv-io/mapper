# mapper

## Installation

```bash
go get -u github.com/metadiv-io/mapper
```

## Usage - Mapping Models

`mapper` provides two generic methods for mapping models: `Map2Model` and `Map2Models`. The generic model is the model that is used to map the data to.

### Map2Model

```go
type User struct {
    ID uint
    Name string
}

type UserDTO struct {
    ID uint
    Name string
}

func main() {
    user := User{Name: "John Doe"}
    dto := mapper.Map2Model[UserDTO](&user)
}
```

### Map2Models

```go
type User struct {
    ID uint
    Name string
}

type UserDTO struct {
    ID uint
    Name string
}

func main() {
    users := []User{{Name: "John Doe"}, {Name: "Jane Doe"}}
    dtos := mapper.Map2Models[UserDTO](users)
}
```

## Base Mapper

Embed the `BaseMapper` struct in your mapper struct to get the `Map2Model` and `Map2Models` methods.

```go
type UserMapper[T any] struct {
    mapper.BaseMapper[T]
}
```
