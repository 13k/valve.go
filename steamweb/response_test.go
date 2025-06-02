package steamweb_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/13k/valve.go/steamweb"
	"github.com/13k/valve.go/steamweb/schema"
)

func TestResponse_Getters(t *testing.T) {
	assert := assert.New(t)

	res := &steamweb.Response{}

	assert.Nil(res.Header())
	assert.Nil(res.Cookies())
	assert.Equal("", res.Status())
	assert.Equal(0, res.StatusCode())
	assert.False(res.IsError())
	assert.False(res.IsSuccess())
	assert.Nil(res.Result())
	assert.Equal(time.Duration(0), res.Duration())

	si := &schema.Interface{}
	sm := &schema.Method{}
	result := map[string]interface{}{}
	req := steamweb.NewRequest(si, sm).SetResult(result)

	res = &steamweb.Response{
		Request: req,
	}

	assert.Nil(res.Header())
	assert.Nil(res.Cookies())
	assert.Equal("", res.Status())
	assert.Equal(0, res.StatusCode())
	assert.False(res.IsError())
	assert.False(res.IsSuccess())
	assert.Exactly(result, res.Result())
	assert.Equal(time.Duration(0), res.Duration())

	header := http.Header{"X-Test": {"test"}}
	res = &steamweb.Response{
		Request: req,
		Raw: &http.Response{
			Header:     header,
			Status:     "299 Test OK",
			StatusCode: 299,
		},
	}

	assert.Equal(header, res.Header())
	assert.Empty(res.Cookies())
	assert.False(res.IsError())
	assert.True(res.IsSuccess())
	assert.Equal("299 Test OK", res.Status())
	assert.Equal(299, res.StatusCode())
	assert.Exactly(result, res.Result())
	assert.Equal(time.Duration(0), res.Duration())

	res = &steamweb.Response{
		Request: req,
		Raw: &http.Response{
			Header:     header,
			Status:     "400 Test Error",
			StatusCode: 400,
		},
	}

	assert.Equal(header, res.Header())
	assert.Empty(res.Cookies())
	assert.True(res.IsError())
	assert.False(res.IsSuccess())
	assert.Equal("400 Test Error", res.Status())
	assert.Equal(400, res.StatusCode())
	assert.Exactly(result, res.Result())
	assert.Equal(time.Duration(0), res.Duration())

	now := time.Now()
	res = &steamweb.Response{
		Request: req,
		Raw: &http.Response{
			Header:     header,
			Status:     "400 Test Error",
			StatusCode: 400,
		},
		Time: now,
	}

	assert.Equal(header, res.Header())
	assert.Empty(res.Cookies())
	assert.True(res.IsError())
	assert.False(res.IsSuccess())
	assert.Equal("400 Test Error", res.Status())
	assert.Equal(400, res.StatusCode())
	assert.Exactly(result, res.Result())
	assert.Equal(time.Duration(0), res.Duration())

	req.Time = now.Add(-13 * time.Second)

	res = &steamweb.Response{
		Request: req,
		Raw: &http.Response{
			Header:     header,
			Status:     "400 Test Error",
			StatusCode: 400,
		},
		Time: now,
	}

	assert.Equal(header, res.Header())
	assert.Empty(res.Cookies())
	assert.True(res.IsError())
	assert.False(res.IsSuccess())
	assert.Equal("400 Test Error", res.Status())
	assert.Equal(400, res.StatusCode())
	assert.Exactly(result, res.Result())
	assert.Equal(13*time.Second, res.Duration())
}
