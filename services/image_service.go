// image_service.go
// this will handle coordination between bouldertracker and the third
// party image hosting API, currently aws S3
package services

import (
	"io"
    "fmt"
    "os"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
)

// creates the credentials for the s3 auth, makes
// the file uploader and uploads the file according to
// env vars to the right bucket and key
// returns success bool and the public url for the image
func UploadProfilePicture(filename string, header string, image_file io.Reader) (bool, string) {
	secretKey := os.Getenv("S3_SECRET_KEY")
	accessKey := os.Getenv("S3_ACCESS_KEY")
	regionName := os.Getenv("S3_REGION")
	bucketName := os.Getenv("S3_BUCKET")
	fileKey := os.Getenv("S3_KEY")
	acl := os.Getenv("S3_ACL")
	imageKey := fmt.Sprintf("%s/%s", fileKey, filename)

	// S3 Upload Manager
	uploader := s3manager.NewUploader(session.New(&aws.Config{
		Region: aws.String(regionName), 
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageKey),
		ContentType: aws.String(header),
		ACL: aws.String(acl),
		Body:   image_file,
	})
	if err != nil {
		fmt.Println(err.Error())
		return false, err.Error()
	}
	return true, result.Location
}

