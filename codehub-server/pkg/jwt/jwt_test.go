package jwt

import (
	"bluebell/conf"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGenToken(t *testing.T) {
	testCases := []struct {
		name           string
		userID         int64
		username       string
		expectedErrMsg string
	}{
		{"Case1", 123, "user1", ""},
		{"Case2", 456, "user2", ""},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			token, err := GenToken(tc.userID, tc.username)
			if tc.expectedErrMsg != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErrMsg)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, token)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	//读取配置文件
	_ = os.Chdir("../..") // 更改工作目录到项目的根目录
	const defaultConfFile = "./conf/config.yaml"

	confFile := defaultConfFile
	conf.Init(confFile)

	// 生成一个长有效期的 token 用于测试 Case1
	userID := int64(789)
	username := "user3"
	tokenString, err := GenToken(userID, username)
	require.NoError(t, err)

	// 生成一个格式错误的 token 用于测试 Case2
	invalidToken := "InvalidToken"

	testCases := []struct {
		name           string
		tokenString    string
		expectedClaims *MyClaims
		expectedErrMsg string
	}{
		{"Case1", tokenString, &MyClaims{UserID: userID, Username: username}, ""},
		{"Case2", invalidToken, nil, "invalid token"},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			claims, err := ParseToken(tc.tokenString)
			if tc.expectedErrMsg != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedErrMsg)
				require.Nil(t, claims)
			} else {
				require.NoError(t, err)
				require.NotNil(t, claims)
				require.Equal(t, tc.expectedClaims.UserID, claims.UserID)
				require.Equal(t, tc.expectedClaims.Username, claims.Username)
				// Additional checks for other claims if needed
			}
		})
	}
}
