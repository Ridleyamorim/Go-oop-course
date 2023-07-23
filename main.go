package main

import(
	"fmt"
	"github.com/ridleyamorim/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
    conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
    Sacar(valor float64) string
}

func main() {
	contaDoRidley := contas.ContaPoupanca{}

	contaDoRidley.Depositar(100)
	PagarBoleto(&contaDoRidley, 60)

	fmt.Println(contaDoRidley.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)
	
	fmt.Println(contaDaLuisa.ObterSaldo())
}