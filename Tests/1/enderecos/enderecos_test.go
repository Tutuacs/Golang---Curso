// package enderecos_test É possivel usar um package diferente para testes
package enderecos

import (
	"fmt"
	"testing"
)

type test struct {
	input  string
	output string
}

func TestTipoDeEndereco(t *testing.T) {

	tests := builder_test()
	// enderecoParaTeste := "Avenida Paulista"

	// tipodeEnderecoEsperado := "avenida"

	for _, test := range tests {
		returned := TipoDeEndereco(test.input)
		fmt.Println('"', returned, '"')
		if returned != test.output {
			t.Errorf("O tipo recebido é diferente do esperado!\nEsperado: %s\nRecebido: %v", test.output, returned)
		}
	}

	// tipoRecebido := TipoDeEndereco(enderecoParaTeste)
	// if tipoRecebido != tipodeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido é diferente do esperado!\nEsperado: %s\nRecebido: %v", tipodeEnderecoEsperado, tipoRecebido)
	// }
}

func builder_test() (slice []test) {
	slice = []test{
		{
			input:  "Rua antônio",
			output: "Rua",
		},
		{
			input:  "Avenida antônio",
			output: "Avenida",
		},
		{
			input:  "Estrada de Brasília",
			output: "Estrada",
		},
		{
			input:  "Rodovia BR 116",
			output: "Rodovia",
		},
		{
			input:  "RODOVIA BR 116",
			output: "RODOVIA",
		},
		{
			input:  "ERRADO BR 116",
			output: "Tipo Inválido",
		},
	}

	return
}
