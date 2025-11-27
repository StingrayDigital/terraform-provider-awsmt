package main

import (
	"context"
	"log"
	"os"
	"slices"

	"terraform-provider-mediatailor/awsmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	err := providerserver.Serve(context.Background(), awsmt.New, providerserver.ServeOpts{
		Address: "terraform.stingray.tools/sre/awsmt",
		Debug:   slices.Contains(os.Args, "-debug"),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
