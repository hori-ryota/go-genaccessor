# go-genaccessor

Accsessor generator for Go.

## Usage

```go
    type Foo struct {
        key string `getter:"[alias,]..." setter:"[alias,]..."`
    }
```

with `go generate` command

```go
    //go:generate go-genaccessor
```

### Example

def
[./_example/person.go](./_example/person.go)

generated
[./_example/example_accessor_gen.go](./_example/example_accessor_gen.go)


## Installation

```sh
$ go get github.com/hori-ryota/go-genaccessor
```
