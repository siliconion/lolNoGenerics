# About variable declaration
Go declaring variable as: `var variable_name data_type`
```go
var i int
```

Go can also "short assign" with `:=`, which automatically declares the variable and assigns the type. 
```go
i := 1
```

Go can also declare multiple variables:
```go
var i, j, k int
i, j, k := 1, 2, 3
```

Pointers can also be declared this way
```go
i := 1
p := *i
```

# About function declaration
Go declaring functions as: `func func_name(input_var input_type) output_type`
```go
func main(argc int, argv []string) int {
	// func body
}
```