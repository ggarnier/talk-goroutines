```go
var balance int

func deposit(value int) {
	balance += value
}

func draw(value int) {
	if balance >= value {
		balance -= value
	}
}

func currentBalance() int {
	return balance
}

func main() {
	deposit(50)
	draw(20)
	deposit(120)
	draw(35)

	fmt.Println(currentBalance()) // 115
}
```
