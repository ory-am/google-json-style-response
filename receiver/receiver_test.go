package receiver

import (
	"bytes"
	"github.com/ory-am/google-json-style-response/responder"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DataMock struct {
	A string
	B string
}

func TestNew(t *testing.T) {
	r := New("1.0")
	assert.NotNil(t, r)
}

func TestGetResponse(t *testing.T) {
	rc := responder.New("1.0")
	o := rc.Success(DataMock{
		A: "a",
		B: "b",
	})
	b, err := o.Marshal()
	assert.Nil(t, err)

	r := New("1.0")
	reader := bytes.NewReader(b)
	result, err := r.GetResponse(reader)
	assert.Nil(t, err)

	do := result.Data.(map[string]interface{})
	assert.Equal(t, "a", do["A"])
}

func TestGetResponseFailsOnAPIVersionMismatch(t *testing.T) {
    rc := responder.New("1.0")
    o := rc.Success(DataMock{
        A: "a",
        B: "b",
    })
    b, err := o.Marshal()
    assert.Nil(t, err)

    r := New("1.1")
    reader := bytes.NewReader(b)
    _, err = r.GetResponse(reader)
    assert.NotNil(t, err)
}