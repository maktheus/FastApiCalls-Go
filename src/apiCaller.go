package apiCaller

import (
	"fmt"
	"sync"
	"time"
	"net/http"
)


func ApiCallerMainFunction(duration int) {

	//print o tempo de duracao do script em  00:00:00

	var tempoTotal =   time.Duration(duration) * time.Second 

	fmt.Println("Tempo de execução do script: ", tempoTotal)
	
	
	endTime := time.Now().Add(time.Duration(duration) * time.Second)

	// Criamos um grupo de espera para garantir que todas as threads
	// sejam concluídas antes de terminarmos o programa.
	var wg sync.WaitGroup

	// Criamos um slice de URLs que queremos chamar.
	urls := []string{
		"http://example.com/api/v1/foo",
		"http://example.com/api/v1/bar",
		"http://example.com/api/v1/baz",
	}

	for time.Now().Before(endTime){

		// Vamos percorrer cada URL e criar uma thread para fazer a chamada.
		
		for _, url := range urls {
			// Adicionamos uma unidade ao contador de threads no grupo de espera.
			wg.Add(1)

			// Iniciamos uma thread e passamos a URL e o grupo de espera como argumentos.
			go func(url string, wg *sync.WaitGroup) {
				// Quando a thread for finalizada, removemos uma unidade do contador no grupo de espera.
				defer wg.Done()
				
				// Aqui, você pode colocar o código para fazer a chamada API.
				// Por exemplo, você pode usar o pacote "http" do Go para fazer a chamada.
				_, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
				}else{
					fmt.Println("Sucesso")
				}
				// Imprimimos a URL que estamos chamando para mostrar que está funcionando.
				fmt.Println("Fazendo chamada para:", url)
			}(url, &wg)
		}
	time.Sleep(1 * time.Second)
	}
	
	// Aqui, nós esperamos até que todas as threads tenham sido concluídas.
	wg.Wait()
}
