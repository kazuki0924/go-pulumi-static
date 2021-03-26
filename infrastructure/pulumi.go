package main

import (
	"log"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucket, err := PulumiDeployS3ForStatic(ctx)
		if err != nil {
			log.Fatal(err)
		}

		err = PulumiDeployStatic(ctx, bucket)
		if err != nil {
			log.Fatal(err)
		}

		ctx.Export("bucketEndpoint", pulumi.Sprintf("http://%s", bucket.WebsiteEndpoint))

		return nil
	})
}
