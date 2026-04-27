package main

import "fmt"

type conta interface {
	Depositar(valor float64)
	Sacar(valor float64)
	VerSaldo() float64
}

type contaCorrente struct {
	Titular string
	Saldo   float64
}
type contaPoupanca struct {
	Titular string
	Saldo   float64
}

func main() {
	conta := contaCorrente{Titular: "Amauri", Saldo: 100}
	conta.Depositar(50)
	conta.Sacar(30)

	fmt.Printf("Titular: %s - Saldo Final: %.2f\n", conta.Titular, conta.Saldo)

	poupanca := contaPoupanca{Titular: "Amauri", Saldo: 500}
	poupanca.Depositar(100)
	poupanca.Sacar(200)
	exibirSaldo(&poupanca)
}

func (c *contaCorrente) Depositar(valor float64) {
	c.Saldo += valor
}
func (c *contaCorrente) VerSaldo() float64 {
	return c.Saldo
}
func (c *contaCorrente) Sacar(valor float64) {
	if valor > c.Saldo {
		fmt.Println("Saldo insuficiente para o saque de: !", valor)
	} else {
		c.Saldo -= valor
		fmt.Println("Saque de ", valor, "realizado com sucesso!")
	}
}

func (c *contaPoupanca) Depositar(valor float64) {
	c.Saldo += valor
}
func (c *contaPoupanca) Sacar(valor float64) {
	taxa := valor * 0.05
	if valor > c.Saldo {
		fmt.Println("Saldo Insuficiente!")
	} else {
		totalDebito := valor + taxa
		c.Saldo -= totalDebito
		fmt.Println("Saque de ", valor, "realizado com sucesso!")
	}
}
func (c *contaPoupanca) VerSaldo() float64 {
	return c.Saldo
}
func exibirSaldo(c conta) {
	fmt.Printf("Saldo: %.2f\n", c.VerSaldo())
}
