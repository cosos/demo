package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"github.com/cosos/firefly/cloudlib/azurecloud"
)

func main() {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	option := arm.ClientOptions{Endpoint: arm.AzurePublicCloud}
	tiClient := armsecurityinsights.NewThreatIntelligenceIndicatorClient(os.Getenv("Subscription"), creds, &option)
	if azurecloud.CheckIndicator(tiClient, context.Background(), "Core", "Sentinel", "win.emotet") {
		log.Println("found it")
	}
	log.Println("Resorces")
}
