package huaweicloud

import (
	"bytes"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
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

func (o Obs) Upload(key string, content []byte) (err error) {
	input := &obs.PutObjectInput{}
	input.Bucket = o.bucketName
	input.Key = key
	input.Body = bytes.NewReader(content)
	_, err = o.obsClient.PutObject(input)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
