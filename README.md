# go_tricks

## file in.go examples

```var (
	_err_1 = fmt.Errorf("sql: sql no rows")
	_err_2 = fmt.Errorf("database ping failed")
)

func main() {
	// int in string
	fmt.Println("int in string. exists 1 in b,c,d -", in(1, "b", "c", "d"))
	// string in int
	fmt.Println("string in int. exists a in 1,2,3 -", in("a", 1, 2, 3))
	// string in any
	fmt.Println("string in any. exists a in 1,2.0,error,4,true -", in("a", 1, 2.0, _err_1, "4", true))

	// string
	fmt.Println("string. exists a in b,c,d -", in("a", "b", "c", "d"))
	fmt.Println("string. exists a in a,b,c,d -", in("a", "a", "b", "c", "d"))
	// int
	fmt.Println("int. exists 1 in 2,3,4,5 -", in(1, 2, 3, 4, 5))
	fmt.Println("int. exists 1 in 1,2,3,4,5", in(1, 1, 2, 3, 4, 5))
	// float
	fmt.Println("float. exists 1.2 in 2.3, 3.4, 4.5, 5.6 -", in(1.2, 2.3, 3.4, 4.5, 5.6))
	fmt.Println("float. exists 1.2 in 1.2, 2.3, 3.4, 4.5, 5.6", in(1.2, 1.2, 2.3, 3.4, 4.5, 5.6))
	// error
	fmt.Println("error. exists _err_1 in _err_2 -", in(_err_1, _err_2))
	fmt.Println("error. exists _err_1 in _err_1, _err_2", in(_err_1, _err_1, _err_2))
}
```