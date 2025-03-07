package utils

import (
	"testing"
	"time"
)

func TestJwt(t *testing.T) {
	// Test GenerateTokens
	accessString, refreshToken, err := GenerateTokens(1)
	if err != nil {
		t.Errorf("GenerateTokens error: %v", err)
	}
	t.Logf("Access Token: %v", accessString)
	t.Logf("Refresh Token: %v", refreshToken)

	// Sleep 16 seconds
	time.Sleep(16 * time.Second)

	// Test ParseToken
	claims, err := ParseToken(accessString)
	if err != nil {
		t.Errorf("ParseToken error: %v", err)
	}
	t.Logf("Claims: %v", claims)
}
