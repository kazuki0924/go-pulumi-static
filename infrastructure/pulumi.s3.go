package main

import (
	"io/ioutil"

	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

const awsS3StaticBucketName = "pulumi-test-bucket"

func PulumiDeployS3ForStatic(ctx *pulumi.Context) (*s3.Bucket, error) {
	bucket, err := s3.NewBucket(ctx, awsS3StaticBucketName, &s3.BucketArgs{
		Website: s3.BucketWebsiteArgs{
			IndexDocument: pulumi.String("index.html"),
		},
	})
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func PulumiDeployStatic(ctx *pulumi.Context, bucket *s3.Bucket) error {
	htmlContent, err := ioutil.ReadFile("../frontend/index.html")
	if err != nil {
		return err
	}

	_, err = s3.NewBucketObject(ctx, "index.html", &s3.BucketObjectArgs{
		Acl:         pulumi.String("public-read"),
		ContentType: pulumi.String("text/html"),
		Bucket:      bucket.ID(),
		Content:     pulumi.String(string(htmlContent)),
	})
	if err != nil {
		return err
	}

	return nil
}
