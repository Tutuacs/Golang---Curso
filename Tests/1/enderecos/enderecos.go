package enderecos

import (
	"strings"
)

// TipoDeEndereco verifica se um endereço é válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	firstWord := strings.Split(endereco, " ")[0]

	if validator(tiposValidos, firstWord) {
		return firstWord
	}

	return "Tipo Inválido"
}

func validator(tiposValidos []string, firstWord string) bool {
	for _, tipo := range tiposValidos {
		if strings.EqualFold(tipo, firstWord) {
			return true
		}
	}
	return false
}
