package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/ai/azopenai"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	azureOpenAIEndpoint := os.Getenv("AZURE_OPENAI_ENDPOINT")
	azureOpenAIDeployment := os.Getenv("AZURE_OPENAI_DEPLOYMENT")
	azureSearchEndpoint := os.Getenv("AZURE_SEARCH_ENDPOINT")
	azureSearchIndex := os.Getenv("AZURE_SEARCH_INDEX")
	managedIdentityResourceId := os.Getenv("MANAGED_IDENTITY_RESOURCE_ID")
	if azureOpenAIEndpoint == "" || azureOpenAIDeployment == "" {
		fmt.Println("Please set AZURE_OPENAI_ENDPOINT and AZURE_OPENAI_KEY environment variables")
		return
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		fmt.Println("Failed to create a credential")
		return
	}

	azureOpenAIClient, err := azopenai.NewClient(azureOpenAIEndpoint, credential, nil)
	if err != nil {
		fmt.Println("Failed to create Azure OpenAI client")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What do you need?")

	inputText, _ := reader.ReadString('\n')

	options := azopenai.ChatCompletionsOptions{
		Messages: []azopenai.ChatRequestMessageClassification{
			&azopenai.ChatRequestUserMessage{
				Content: azopenai.NewChatRequestUserMessageContent(inputText),
			},
		},
		DeploymentName: &azureOpenAIDeployment,
		AzureExtensionsOptions: []azopenai.AzureChatExtensionConfigurationClassification{
			&azopenai.AzureSearchChatExtensionConfiguration{
				Parameters: &azopenai.AzureSearchChatExtensionParameters{
					Endpoint:  &azureSearchEndpoint,
					IndexName: &azureSearchIndex,
					Authentication: &azopenai.OnYourDataUserAssignedManagedIdentityAuthenticationOptions{
						ManagedIdentityResourceID: &managedIdentityResourceId,
					},
				},
			},
		},
	}

	azureOpenAIClient.GetChatCompletionsStream(context.Background())

}
