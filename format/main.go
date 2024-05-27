package main

import (
	"fmt"
	"os"
)

// Структура "точка"
type Point struct {
	x, y int
}

func main() {
	fmt.Println(" \n[ ФОРМАТИРОВАНИЕ ]\n ")

	/* Структура */

	p := Point{1, 2}

	fmt.Println("Структура:")

	fmt.Printf("struct1  [ %%v ]   %v\n", p)
	fmt.Printf("struct2  [ %%+v ]  %+v\n", p)
	fmt.Printf("struct3  [ %%#v ]  %#v\n", p)
	fmt.Printf("type     [ %%T ]   %T\n", p)

	/* Типы */

	fmt.Println()
	fmt.Println("Типы:")

	fmt.Printf("bool     [ %%t ]   %t\n", true)
	fmt.Printf("int      [ %%d ]   %d\n", 123)
	fmt.Printf("bin      [ %%b ]   %b\n", 14)
	fmt.Printf("char     [ %%c ]   %c\n", 33)
	fmt.Printf("hex      [ %%x ]   %x\n", 456)
	fmt.Printf("float1   [ %%f ]   %f\n", 78.9)
	fmt.Printf("float2   [ %%e ]   %e\n", 123400000.0)
	fmt.Printf("float3   [ %%E ]   %E\n", 123400000.0)
	fmt.Printf("str1     [ %%s ]   %s\n", "\"string\"")
	fmt.Printf("str2     [ %%q ]   %q\n", "\"string\"")
	fmt.Printf("str3     [ %%x ]   %x\n", "string")
	fmt.Printf("pointer  [ %%p ]   %p\n", &p)

	/* Выравнивание */

	fmt.Println()
	fmt.Println("Выравнивание:")

	fmt.Printf("right:   | %6d | %6d |\n", 12, 345)
	fmt.Printf("float:   | %6.2f | %6.2f |\n", 1.2, 3.45)
	fmt.Printf("left:    | %-6.2f | %-6.2f |\n", 1.2, 3.45)
	fmt.Printf("right:   | %6s | %6s |\n", "foo", "bar")
	fmt.Printf("left:    | %-6s | %-6s |\n", "foo", "bar")

	/* Разное */

	fmt.Println()
	fmt.Println("Разное:")

	fmt.Println(fmt.Sprintf("sprintf  [ %s ]", "string"))
	fmt.Fprintf(os.Stderr, "stderr   [ %s  ]\n", "error")
}
