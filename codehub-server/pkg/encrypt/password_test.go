package encrypt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		expected string
	}{
		{"Case1", "123456", "3132333435366a3b6df9c11836f95addead09676f347"},
		{"Case2", "password123", "70617373776f72643132336a3b6df9c11836f95addead09676f347"},
		// 添加更多的测试用例
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetPassword(tc.password)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		name        string
		password    string
		newPassword string
		expected    bool
	}{
		{"Case1", "3132333435366a3b6df9c11836f95addead09676f347", "123456", true},
		{"Case2", "70617373776f72643132336a3b6df9c11836f95addead09676f347", "password123", true},
		{"Case3", "InvalidHash", "password", false}, // Incorrect hash should return false
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := CheckPassword(tc.password, tc.newPassword)
			require.Equal(t, tc.expected, result)
		})
	}
}
