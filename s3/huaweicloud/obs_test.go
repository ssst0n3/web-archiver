package huaweicloud

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"io"
	"net/url"
	"strings"
	"testing"
	"web-archiver/config"
)

func NewObsForTest(t *testing.T) (o *Obs) {
	o, err := NewObs(config.AK, config.SK, config.Endpoint, config.BucketName)
	assert.NoError(t, err)
	return
}

func TestUrlEncode(t *testing.T) {
	log.Logger.Info(url.PathEscape("https://www.baidu.com"))
}

func TestObs_Upload(t *testing.T) {
	u, err := NewObsForTest(t).Upload(url.PathEscape("https://www.baidu.com"), io.NopCloser(strings.NewReader("hello-world")))
	assert.NoError(t, err)
	log.Logger.Info(u)
}
