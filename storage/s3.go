package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	s3bucket "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"path"
	"strings"
)

type S3 struct {
	session *session.Session
	uploader *s3manager.Uploader
	downloader *s3manager.Downloader
	bucket string
}

func NewS3(cfg *aws.Config, bucket string) (*S3, error) {
	sess, err := session.NewSession(cfg)
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)
	downloader := s3manager.NewDownloader(sess)


	return &S3{session: sess, uploader: uploader, downloader: downloader, bucket: bucket}, nil
}

func (s3 *S3) InitBucket() error {
	_, err := s3.uploader.S3.CreateBucket(&s3bucket.CreateBucketInput{Bucket: aws.String(s3.bucket), ACL: aws.String("public-read")})
	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == s3bucket.ErrCodeBucketAlreadyExists || err.Error() == s3bucket.ErrCodeBucketAlreadyOwnedByYou {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func (s3 *S3) StoreToken(token *Token) error {
	payload, zipErr := token.Gzip()
	if zipErr != nil {
		return fmt.Errorf("failed to upload file, %v", zipErr)
	}
	result, err := s3.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.bucket),
		Key: aws.String(path.Join(token.ContractAddress, token.UserAddress)),
		ACL: aws.String("public-read"),
		Body: strings.NewReader(payload),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("Uploaded file to, %s\n", aws.StringValue(&result.Location))
	return nil
}

func (s3 *S3) ListKeys() (keys []string, err error) {
	resp, err := s3.downloader.S3.ListObjects(&s3bucket.ListObjectsInput{
		Bucket:              aws.String(s3.bucket),
	})
	if err != nil {
		return nil, err
	}
	for _, object := range resp.Contents {
		keys = append(keys, *object.Key)
	}
	return keys, nil
}


func (s3 *S3) GetToken(contractAddress, userAddress string) (*Token, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := s3.downloader.Download(buf, &s3bucket.GetObjectInput{
		Bucket:                     aws.String(s3.bucket),
		Key:                        aws.String(path.Join(contractAddress, userAddress)),
	})
	if err != nil {
		return nil, err
	}
	token := &Token{}
	unZipErr := token.UnZip(buf.Bytes())
	return token, unZipErr
}