package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account--function
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit can manipulate balance--method
// 리시버는 반드시 구조체의 첫 글자를 따서 소문자로 지어야함
// Account의 복사본을 쓰지말고 *을 붙여서 호출한 account를 사용해야 값 변경 가능.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

// String() is Automatically called
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas : ", a.Balance())
}
