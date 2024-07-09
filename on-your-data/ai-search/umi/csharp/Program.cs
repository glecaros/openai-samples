using System;
using System.ClientModel;
using System.IO;
using Azure.AI.OpenAI;
using Azure.AI.OpenAI.Chat;
using Azure.Identity;
using OpenAI.Chat;

static void InitEnvironment()
{
    string filePath = Path.Join(Environment.CurrentDirectory, ".env");
    if (File.Exists(filePath))
    {
        var lines = File.ReadAllLines(filePath);
        foreach (var line in lines)
        {
            var parts = line.Split("=");
            if (parts.Length == 2)
            {
                Environment.SetEnvironmentVariable(parts[0], parts[1].Trim('\"'));
            }
        }
    }
}

InitEnvironment();

var azureOpenAIEndpoint = Environment.GetEnvironmentVariable("AZURE_OPENAI_ENDPOINT")!;
var azureOpenAIDeployment = Environment.GetEnvironmentVariable("AZURE_OPENAI_DEPLOYMENT")!;
var azureSearchEndpoint = Environment.GetEnvironmentVariable("AZURE_SEARCH_ENDPOINT")!;
var azureSearchIndex = Environment.GetEnvironmentVariable("AZURE_SEARCH_INDEX")!;
var managedIdentityResourceId = Environment.GetEnvironmentVariable("MANAGED_IDENTITY_RESOURCE_ID")!;

AzureOpenAIClient client = new AzureOpenAIClient(new(azureOpenAIEndpoint), new DefaultAzureCredential(), new()
{
});
var chatClient = client.GetChatClient(azureOpenAIDeployment);

ChatCompletionOptions options = new ChatCompletionOptions();

options.AddDataSource(new AzureSearchChatDataSource()
{
    Endpoint = new(azureSearchEndpoint),
    IndexName = azureSearchIndex,
    Authentication = DataSourceAuthentication.FromUserManagedIdentity(managedIdentityResourceId),
});

try
{
    var response = chatClient.CompleteChat([
        new SystemChatMessage("You are an assistant that helps users get information about their Contoso Electronics benefits."),
        new UserChatMessage("tell me about my benefits")
    ], options);
    ChatCompletion completion = response;
    foreach (var content in completion.Content)
    {
        Console.WriteLine(content.Text);

    }
}
catch (ClientResultException ex)
{
    var response = ex.GetRawResponse();
    Console.WriteLine(response?.Content.ToString());

}


Console.WriteLine("Hello, World!");
