package featureLink

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeFromId(t *testing.T) {
	assert.Equal(t, "b", codeFromId(1))
	assert.Equal(t, "B", codeFromId(27))
	assert.Equal(t, "9", codeFromId(61))
	assert.Equal(t, "ba", codeFromId(62))
	assert.Equal(t, "baa", codeFromId(3844))
	assert.Equal(t, "eQPpme", codeFromId(4294967296))
	assert.Equal(t, "v8QrISNYGaG", codeFromId(1.8446744e+19))
}

func TestIdFromCode(t *testing.T) {
	assert.Equal(t, uint(1), idFromCode("b"))
	assert.Equal(t, uint(27), idFromCode("B"))
	assert.Equal(t, uint(61), idFromCode("9"))
	assert.Equal(t, uint(62), idFromCode("ba"))
	assert.Equal(t, uint(3844), idFromCode("baa"))
	assert.Equal(t, uint(4294967296), idFromCode("eQPpme"))
	assert.Equal(t, uint(1.8446744e+19), idFromCode("v8QrISNYGaG"))
}
