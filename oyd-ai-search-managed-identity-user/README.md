
1. Login using the Azure CLI

```bash
az login
```

1. Create a resource group

```bash
az group create --name MyResourceGroup --location eastus
```

1. Create an Azure OpenAI resource

```bash
az cognitiveservices account create --name openai-samples --resource-group openai-samples --location eastus --kind OpenAI --sku s0
```

1. Get the endpoint

```bash
az cognitiveservices account show --name openai-samples --resource-group  openai-samples  | jq -r .properties.endpoint
```

Save it to your `.env` file in the `AZURE_OPENAI_ENDPOINT` variable

1. Deploy a model

```bash
az cognitiveservices account deployment create --name openai-samples --resource-group  openai-samples --deployment-name gpt-4o --model-name gpt-4o --model-version "2024-05-13" --model-format OpenAI --sku-capacity "1" --sku-name "Standard"
```

Save your deployment in your `.env` file in the `AZURE_OPENAI_DEPLOYMENT` variable.

1. Assign the `Cognitive Services OpenAI Contributor` role to your account.

- Find the id of the role

```bash
az role definition list --query "[?roleName=='Cognitive Services OpenAI Contributor'].{name:roleName, id:id}"
```

- Find the id of your resource

```bash
az resource show --name openai-samples --resource-group openai-samples --resource-type Microsoft.CognitiveServices/accounts | jq -r .id
```

- Create the role assigment
```bash
az role assignment create --assignee <UserID> --role "<RoleID>" --scope "<ResourceID>"

```

1. Create your Azure AI Search resource

```bash
az search service create --name openai-samples-search --resource-group openai-samples --sku Standard --partition-count 1 --replica-count 1
```
