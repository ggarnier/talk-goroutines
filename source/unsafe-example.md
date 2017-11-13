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
	go deposit(50)
	go draw(20)
	go deposit(120)
	go draw(35)

	time.Sleep(100 * time.Millisecond)
	fmt.Println(currentBalance()) // 115, 170, 150, 135, ...
}
```
