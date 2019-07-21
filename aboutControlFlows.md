# if-else
Basically the same as other languages, without parenthesis 
```go
i := 0
if i == 0 {
	fmt.Println("Zero")
} else {
  fmt.Println("Not zero")
}
```
Go can have assignments before the actual evaluations:
```go
if j:= 0; i == j {
	fmt.Println("They are equal!")
}
```

# for loops
Basically the same as other languages; `for start; end; incre`
```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
```
Go can emit start, incre, which becomes a while:
```go
for i < 10 {
	fmt.Println(i)
}
```

# Defer makes a stack!