package controllers

import (
	"context"
	"finance-utility/configs"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// listFiles lists objects within specified bucket.
func ListFiles() error {
	bucket := configs.BUCKET_NAME
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()

	query := &storage.Query{Prefix: configs.QUERY_PREFIX}
	it := client.Bucket(bucket).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("Bucket(%q).Objects: %v", bucket, err)
		}
		// fmt.Fprintln(w, attrs.Name)
		fmt.Printf("%v \n", attrs.Name)
	}
	return nil
}
