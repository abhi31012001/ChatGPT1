package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	apiKey, err := os.LookupEnv("API_KEY")
	if !err {
		log.Fatalln("Miaaing API KEY")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	resp, errr := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"what is golang"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if errr != nil {
		log.Fatalln(errr.Error())
	}
	fmt.Println(resp.Choices[0].Text)

}
