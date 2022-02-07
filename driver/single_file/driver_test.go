package single_file

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/web-archiver/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSingleFile(t *testing.T) {
	config.SingleFileUrl = "http://172.19.0.3"
	reader, err := SingleFile("https://www.baidu.com/robots.txt")
	assert.NoError(t, err)
	content, err := ioutil.ReadAll(reader)
	assert.NoError(t, err)
	log.Logger.Info(string(content))
	assert.NoError(t, reader.Close())
}
