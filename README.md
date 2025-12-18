# go-echo

Simple stdout output functions for Go.

## Installation

```bash
go get github.com/Sha0ne/go-echo
```

## Usage

### Standard usage (recommended)

```go
package main

import "github.com/Sha0ne/go-echo"

func main() {
    echo.Echo("Hello World")   // with newline
    echo.Necho("No newline")   // without newline
}
```

### With import alias (to avoid naming conflicts)

```go
package main

import e "github.com/Sha0ne/go-echo"

func main() {
    e.Echo("Hello World")
    e.Necho("No newline")
}
```

### Lowercase wrappers (for shorter syntax)

If you want to use lowercase function names like in shell scripting:

```go
package main

import e "github.com/Sha0ne/go-echo"  // Use alias to avoid conflict

func echo(s string)  { e.Echo(s) }     // Wrapper functions
func necho(s string) { e.Necho(s) }

func main() {
    echo("Hello World")     // Clean, lowercase syntax
    necho("No newline")
}
```

### Dot import (not recommended, but possible)

```go
package main

import . "github.com/Sha0ne/go-echo"

func main() {
    Echo("Hello World")    // Direct access without prefix
    Necho("No newline")
}
```

⚠️ **Note:** Dot imports can cause naming conflicts and make code less readable. Use with caution.

## Functions

| Function | Description |
|----------|-------------|
| `Echo(s string)` | Writes s to stdout with a trailing newline |
| `Necho(s string)` | Writes s to stdout without a trailing newline |

## Why use wrappers?

The wrapper approach allows you to:
- Keep the clean, Unix-like lowercase syntax (`echo` instead of `Echo`)
- Avoid naming conflicts (the package name `echo` doesn't clash with your function `echo`)
- Maintain code clarity (it's obvious where the functions come from)

## License

MIT