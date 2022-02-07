package huaweicloud

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/web-archiver/config"
	"io"
	"net/url"
)

type Obs struct {
	bucketName string
	obsClient  *obs.ObsClient
}

func NewObs(ak, sk, endpoint, bucketName string) (o *Obs, err error) {
	obsClient, err := obs.New(ak, sk, endpoint)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	o = &Obs{bucketName: bucketName, obsClient: obsClient}
	return
}

func (o Obs) Upload(key string, reader io.ReadCloser) (obsAddress string, err error) {
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: o.bucketName,
				Key:    key,
				ACL:    obs.AclPublicRead,
			},
			ContentType: "text/html",
		},
		Body: reader,
	}
	_, err = o.obsClient.PutObject(input)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = reader.Close()
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	obsAddress = fmt.Sprintf("https://%s.%s/%s", config.BucketName, config.Endpoint, url.QueryEscape(key))
	return
}
