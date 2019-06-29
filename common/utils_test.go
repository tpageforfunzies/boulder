//utils_test.go
package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	message := Message(true, "message")
	shouldBe := map[string]interface{}{"status": true, "message": "message"}
	assert.Equal(t, shouldBe, message)
}
