package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	balance   float64
}

// 存款
func (account *Account) Deposit(money float64, pwd string) {

	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 {
		fmt.Println("money error")
		return
	}

	account.balance += money

	fmt.Printf("存款 %v 元成功\n", money)
}

// 提款
func (account *Account) Withdraw(money float64, pwd string) {

	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}
	if money <= 0 {
		fmt.Println("money error")
		return
	}
	if money > account.balance {
		fmt.Println("money error")
		return
	}

	account.balance -= money

	fmt.Printf("提款 %v 元成功\n", money)
}

// 查詢餘額
func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("pwd error")
		return
	}

	fmt.Printf("帳號＝%v, 餘額＝%v\n", account.AccountNo, account.balance)
}

func main() {

	// &Account 這樣用是因為上面的方法都是傳指標過去
	var account = &Account{
		"台灣銀行",
		"66666666",
		100,
	}

	account.Query("66666666")
	account.Deposit(300, "66666666")
	account.Withdraw(10, "66666666")
	account.Query("66666666")

}
