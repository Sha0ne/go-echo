/*
To use lowercase function names, add these wrappers in your main.go:
	
	func echo(s string)  { echo.Echo(s) }
	func necho(s string) { echo.Necho(s) }

Before:
	Echo("hello")
	Necho("world")

After:
	echo("hello")
	necho("world")
*/
package echo