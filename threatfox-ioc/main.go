package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"github.com/cosos/firefly/cloudlib/azurecloud"
)

func main() {
	creds, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	tiClient := armsecurityinsights.NewThreatIntelligenceIndicatorClient(os.Getenv("Subscription"), creds, nil)
	if azurecloud.CheckIndicator(tiClient, context.Background(), "Core", "win.emotet") {
		log.Println("found it")
	}
	log.Println("Resorces")
}
