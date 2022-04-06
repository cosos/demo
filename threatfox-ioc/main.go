package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"github.com/cosos/firefly/cloudlib/azurecloud"
)

func main() {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	options := arm.ClientOptions{
		Host: armcore.AzurePublicCloud
	}
	tiClient := armsecurityinsights.NewThreatIntelligenceIndicatorClient(os.Getenv("Subscription"), creds, &options)
	if azurecloud.CheckIndicator(tiClient, context.Background(), "Core", "win.emotet") {
		log.Println("found it")
	}
	log.Println("Resorces")
}
