package main

import (
	"fmt"
	"introducao_a_testes/enderecos"
	"strings"
)

func main() {
	tipoEndereco := strings.Title(enderecos.TipoDeEndereco("Avenida Paulista"))
	fmt.Println(tipoEndereco)
	tipoEndereco = enderecos.TipoDeEndereco("Jesus Paulista")
	fmt.Println(strings.Title(tipoEndereco))
}
