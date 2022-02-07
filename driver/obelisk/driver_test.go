package obelisk

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArchive(t *testing.T) {
	content, err := Archive("https://www.baidu.com")
	assert.NoError(t, err)
	log.Logger.Info(string(content))
}
