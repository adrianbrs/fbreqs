package sender

import (
	"caixa-falso/client"
	"caixa-falso/console"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Sender struct {
	ID          int
	Response    *http.Response
	ElapsedTime time.Duration
	Error       error
}

// Run começa o envio das requisições
func Run(client *client.Client) {
	var req = client.Request

	// Display starting request header
	console.Header().SetLogger(console.InfoLogger).SetBar("-").SetValues(
		fmt.Sprintf("Iniciando o envio de %d requests...", req.Amount),
	).Render()

	// Create a channel to the length of request amount
	reqChannel := make(chan *Sender, req.Amount)

	// Create a wait group for waiting all goroutines finish
	var waitGroup sync.WaitGroup

	// Começa a enviar os dados
	for i := 0; i < req.Amount; i++ {
		// Add goroutine to wait group
		waitGroup.Add(1)

		// Log starting request info
		console.Log("Enviando requisição número: #%d...", i)
		startTime := time.Now()

		// Cria goroutines para enviar as requisições em paralelo
		go func(reqId int, wg *sync.WaitGroup) {
			// Create Sender instance
			sender := &Sender{
				ID: reqId,
			}
			defer func(sender *Sender, wg *sync.WaitGroup) {
				// Done wait gorup
				wg.Done()

				// Send channel when quit from goroutine
				reqChannel <- sender
			}(sender, wg)

			// Get HTTP Client with correct configuration
			httpClient, err := GetHTTPClient(req)

			// HttpClient error
			if err != nil {
				sender.Error = err
				return
			}

			// Faz a requisição
			sender.Response, err = httpClient.PostForm(req.URL, GetData())

			if err != nil {
				sender.Error = err
				return
			}

			defer sender.Response.Body.Close()

			// Pega a resposta da requisição
			var result map[string]interface{}
			json.NewDecoder(sender.Response.Body).Decode(&result)

			// Calc elapsed time since request starts
			sender.ElapsedTime = time.Since(startTime)
		}(i, &waitGroup)
	}

	// Create goroutin to wait wait group
	go func(wg *sync.WaitGroup, reqChan chan *Sender) {
		// Wait for wait group to continue
		waitGroup.Wait()

		// Close reqChannel after wait group wait
		close(reqChannel)
	}(&waitGroup, reqChannel)

	// Count of successfully requests
	successRequests := 0

	// Quit from sender when all requests are done
	for sender := range reqChannel {
		// Check for error
		if sender.Error != nil {
			console.Error("Request error: %s", sender.Error)
			continue
		}

		// Get response
		resp := sender.Response
		successRequests++

		// Log success message
		console.Success("Requisição #%d retornou %s (%s)",
			sender.ID, resp.Status, sender.ElapsedTime)

	}

	// Log final state
	console.Separator()
	console.Info("%d requisições enviadas com sucesso de um total de %d", successRequests, req.Amount)
}
