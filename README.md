# go-echo

Simple stdout output functions for Go.

## Installation

```bash
go get github.com/Sha0ne/go-echo
```

## Usage

```go
package main

import "github.com/Sha0ne/go-echo"

func main() {
    echo.Echo("Hello World")   // with newline
    echo.Necho("No newline")   // without newline
}
```

### Dot import (no prefix)

```go
package main

import . "github.com/Sha0ne/go-echo"

func main() {
    Echo("Hello World")
    Necho("No newline")
}
```

### Lowercase wrappers

For even shorter syntax, add wrappers in your file:

```go
package main

import "github.com/Sha0ne/go-echo"

func echo(s string)  { echo.Echo(s) }
func necho(s string) { echo.Necho(s) }

func main() {
    echo("Hello World")
    necho("No newline")
}
```

## Functions

| Function | Description |
|----------|-------------|
| `Echo(s string)` | Writes s to stdout with a trailing newline |
| `Necho(s string)` | Writes s to stdout without a trailing newline |

## License

MIT