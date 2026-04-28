package exercicios

import "fmt"

type Conta interface {
	Depositar(valor float64)
	Sacar(valor float64)
	VerSaldo() float64
}

type ContaCorrente struct {
	Titular string
	Saldo   float64
}

type ContaPoupanca struct {
	Titular string
	Saldo   float64
}

func (c *ContaCorrente) Depositar(valor float64) {
	c.Saldo += valor
}

func (c *ContaCorrente) VerSaldo() float64 {
	return c.Saldo
}

func (c *ContaCorrente) Sacar(valor float64) {
	if valor > c.Saldo {
		fmt.Println("Saldo insuficiente")
	} else {
		c.Saldo -= valor
	}
}

func (c *ContaPoupanca) Depositar(valor float64) {
	c.Saldo += valor
}

func (c *ContaPoupanca) VerSaldo() float64 {
	return c.Saldo
}

func (c *ContaPoupanca) Sacar(valor float64) {
	taxa := valor * 0.05
	if valor+taxa > c.Saldo {
		fmt.Println("Saldo insuficiente")
	} else {
		c.Saldo -= (valor + taxa)
	}
}

// ExercicioStructs retorna uma string com o resultado do teste das contas
func ExercicioStructs() string {
	res := "=== Teste de Contas ===\n"
	
	corrente := ContaCorrente{Titular: "Amauri", Saldo: 100}
	corrente.Depositar(50)
	corrente.Sacar(30)
	res += fmt.Sprintf("Corrente (Amauri): Saldo Final %.2f\n", corrente.VerSaldo())

	poupanca := ContaPoupanca{Titular: "Amauri", Saldo: 500}
	poupanca.Depositar(100)
	poupanca.Sacar(200)
	res += fmt.Sprintf("Poupanca (Amauri): Saldo Final %.2f\n", poupanca.VerSaldo())

	return res
}
