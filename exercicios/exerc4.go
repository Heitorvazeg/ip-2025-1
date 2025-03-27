package main

import "fmt"

func main() {
	var (
		salario         float32
		kilow           float32
		custoporkw      float32
		custoresidencia float32
		custodesconto   float32
	)

	fmt.Println("Digite o valor do seu salário: ")
	fmt.Scan(&salario)
	fmt.Println("Digite o quanto você gasta em kw: ")
	fmt.Scan(&kilow)

	custoporkw = salario * 70 / 100 / 100
	custoresidencia = custoporkw * kilow
	custodesconto = custoresidencia - (custoresidencia * 10 / 100)

	fmt.Printf("Custo por kW: R$ %.2f\n", custoporkw)
	fmt.Printf("Custo por consumo: R$ %.2f\n", custoresidencia)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custodesconto)
}
