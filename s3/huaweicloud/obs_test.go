package huaweicloud

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"web-archiver/config"
)

func NewObsForTest(t *testing.T) (o *Obs) {
	o, err := NewObs(config.AK, config.SK, config.Endpoint, config.BucketName)
	assert.NoError(t, err)
	return
}

func TestObs_Upload(t *testing.T) {
	assert.NoError(t, NewObsForTest(t).Upload("hello.txt", []byte("hello-world")))
}
