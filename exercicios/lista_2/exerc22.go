package main

import "fmt"

func main() {
	var (
		id                    int
		horasTrabalhadas       float64
		horasExtrasTrabalhadas float64
		salarioBruto, salarioLiquido, inss, impostoDeRenda float64
	)
	const (
		salarioMinimo    float64 = 788.0
		horaExtraValor   float64 = 10.0
	)

	fmt.Scan(&id, &horasTrabalhadas, &horasExtrasTrabalhadas)

	salarioHoraExtra := horaExtraValor * horasExtrasTrabalhadas
	salarioBruto = (3 * salarioMinimo) + salarioHoraExtra

	if salarioBruto > 1500.0 {
		inss = salarioBruto * (12.0 / 100.0)
		salarioLiquido = salarioBruto - inss

		if salarioBruto > 2000.0 {
			impostoDeRenda = salarioBruto * (20.0 / 100.0)

			salarioLiquido = salarioBruto - (inss + impostoDeRenda)
		}
	}

	salarioLiquido = salarioBruto

	fmt.Printf("id: %d\n", id)
	fmt.Printf("Salário bruto: %.2f\n", salarioBruto)
	fmt.Printf("Salário líquido: %.2f\n", salarioLiquido)
}
