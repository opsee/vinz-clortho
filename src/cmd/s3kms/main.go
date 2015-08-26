package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/codegangsta/cli"
)

const (
	version = "0.0.1"
)

var (
	awsConfig *aws.Config
)

func getS3Client() *s3.S3 {
	return s3.New(awsConfig)
}

func get(c *cli.Context) {
	bucket := c.String("bucket")
	object := c.String("object")

	s3Svc := getS3Client()

	s3Params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	}
	s3Resp, err := s3Svc.GetObject(s3Params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	plaintext, err := ioutil.ReadAll(s3Resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	os.Stdout.Write(plaintext)
}

func put(c *cli.Context) {
	bucket := c.String("bucket")
	object := c.String("object")

	s3Svc := getS3Client()

	plaintext, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s3Params := &s3.PutObjectInput{
		Bucket:               aws.String(bucket),
		Key:                  aws.String(object),
		SSEKMSKeyId:          aws.String(c.String("key")),
		ServerSideEncryption: aws.String("aws:kms"),
		Body:                 bytes.NewReader(plaintext),
		ContentType:          aws.String("text/plain"),
	}
	//s3Params.GrantRead = aws.String(fmt.Sprintf("id=%s", c.String("read")))

	s3Resp, err := s3Svc.PutObject(s3Params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(s3Resp)
}

func main() {
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.EnvProvider{},
			&ec2rolecreds.EC2RoleProvider{ExpiryWindow: 5 * time.Minute},
		})

	region := aws.String(os.Getenv("AWS_DEFAULT_REGION"))

	awsConfig = &aws.Config{Credentials: creds, Region: region}

	app := cli.NewApp()
	app.Name = "s3kms"
	app.Usage = "Manage keys encrypted with KMS stored in S3."
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "region, r",
			EnvVar: "AWS_DEFAULT_REGION",
			Usage:  "AWS Region name",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "get",
			Usage:  "get a key from S3 and decrypt it",
			Action: get,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket, b",
					Usage: "S3 bucket name",
				},
				cli.StringFlag{
					Name:  "object, o",
					Usage: "S3 object name",
				},
			},
		},
		{
			Name:   "put",
			Usage:  "put a key into S3 and encrypt it",
			Action: put,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bucket, b",
					Usage: "S3 bucket name",
				},
				cli.StringFlag{
					Name:   "key, k",
					EnvVar: "AWS_KMS_KEY_ARN",
					Usage:  "AWS KMS key ARN",
				},
				cli.StringFlag{
					Name:  "object, o",
					Usage: "S3 object name",
				},
				cli.StringFlag{
					Name:   "read, r",
					EnvVar: "AWS_ACCOUNT_ID",
					Usage:  "AWS account ID for S3 read access",
				},
			},
		},
	}

	app.Run(os.Args)
}
