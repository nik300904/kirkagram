package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// Получаем имя бакета из аргумента командной строки
	//bucketName := flag.String("b", "", "The name of the bucket")
	//flag.Parse()

	//if */bucketName == "" {
	//	fmt.Println("You must supply the name of a bucket (-b BUCKET)")
	//return
	//}

	bucketName := "kirkagram"

	// Подгружаем конфигурацию из ~/.aws/*
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Создаем клиента для доступа к хранилищу S3
	client := s3.NewFromConfig(cfg)

	// Запрашиваем список всех файлов в бакете
	result, err := client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, object := range result.Contents {
		log.Printf("object=%s size=%d Bytes last modified=%s", aws.ToString(object.Key), aws.ToInt64(object.Size), object.LastModified.Local().Format("2006-01-02 15:04:05 Monday"))
	}
}
