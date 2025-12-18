/*
Package echo provides simple stdout output functions similar to shell commands.

Basic usage:

	import "github.com/Sha0ne/go-echo"
	
	func main() {
		echo.Echo("hello")   // Outputs: hello\n
		echo.Necho("world")  // Outputs: world
	}

To use lowercase function names (Unix-style), add these wrappers in your main.go:
	
	import e "github.com/Sha0ne/go-echo"  // Use alias to avoid naming conflict
	
	func echo(s string)  { e.Echo(s) }
	func necho(s string) { e.Necho(s) }
	
	func main() {
		echo("hello")    // Clean lowercase syntax
		necho("world")
	}

The alias is necessary because:
- The package name is "echo" (from "package echo" in the source)
- You want to create a function also named "echo"
- Without the alias, there would be a naming conflict

Alternative approaches:

1. Direct usage (most explicit):
	echo.Echo("hello")
	
2. With custom alias:
	import myecho "github.com/Sha0ne/go-echo"
	myecho.Echo("hello")
	
3. Dot import (not recommended):
	import . "github.com/Sha0ne/go-echo"
	Echo("hello")  // No prefix needed but can cause conflicts
*/
package echo