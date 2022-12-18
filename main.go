package main

import (
	"flag"
	"fmt"
	"FastApiCaller/events/src"
)

type stringArray []string

func (s *stringArray) String() string {
	return fmt.Sprint(*s)
}

func (s *stringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var duration int
var urls stringArray
var file string
var Name string
func init() {
	// Definimos a flag "duration" para o tempo de execução do script em segundos.
	flag.IntVar(&duration, "d", 60, "tempo de execução do script em segundos")

	// Definimos a flag "urls" para as URLs da API.
	flag.Var(&urls, "u", "URLs da API")

	// Definimos a flag "file" para o tipo de arquivo que será gerado.
	flag.StringVar(&file, "f", "json", "tipo de arquivo que será gerado")

	flag.StringVar(&Name, "n", "result", "nome do arquivo que será gerado")
	// Analisamos as flags.
	flag.Parse()
}

func main() {
	//chamar a funcao que faz a chamada da api
	apiCaller.ApiCallerMainFunction(duration, urls,file,Name)
}
