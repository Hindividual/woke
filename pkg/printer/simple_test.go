package printer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple_Print(t *testing.T) {
	p := NewSimple()
	res := generateFileResult()
	got := captureOutput(func() {
		assert.NoError(t, p.Print(res))
	})
	expected := fmt.Sprintf("foo.txt:1:6: [warn] %s\n", res.Results[0].Reason())
	assert.Equal(t, expected, got)
}
