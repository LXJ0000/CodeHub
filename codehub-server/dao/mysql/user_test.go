package mysql

import (
	"bluebell/conf"
	"bluebell/models"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 在所有测试开始之前执行初始化操作
	_ = os.Chdir("../..")
	const confFile = "./conf/config_test.yaml" // 请替换为你的配置文件路径
	conf.Init(confFile)
	Init()

	// 执行测试
	exitCode := m.Run()

	// 在所有测试结束后执行清理操作（如果需要）

	// 退出测试
	os.Exit(exitCode)
}

func TestUserDao_CheckUserExist(t *testing.T) {
	// 准备测试用例
	testCases := []struct {
		name          string
		username      string
		expectedExist bool
		expectedUser  *models.UserModel
	}{
		{name: "ExistingUser", username: "lxj", expectedExist: true, expectedUser: &models.UserModel{UserName: "lxj"}},
		{name: "NonExistingUser", username: "lxjj", expectedExist: false, expectedUser: nil},
	}
	userDaoTest := NewUserDao()
	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			exist, user := userDaoTest.CheckUserExist(tc.username)
			// 使用require库进行断言
			require.Equal(t, tc.expectedExist, exist, "Existence mismatch")
			if tc.expectedExist {
				require.NotNil(t, user, "User should not be nil for existing user")
				require.Equal(t, tc.expectedUser.UserName, user.UserName, "User model mismatch")
			} else {
				require.Nil(t, user, "User should be nil for non-existing user")
			}
		})
	}
}

func TestUserDao_Create(t *testing.T) {

	// 准备测试用例
	testCases := []struct {
		name           string
		user           *models.UserModel
		expectedErrMsg string
	}{
		{name: "ValidUser", user: &models.UserModel{UserName: "test_user"}},
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			err := userDao.Create(tc.user)

			require.NoError(t, err, "Expected no error")
		})
	}

	// 清理测试数据
	for _, tc := range testCases {
		if err := db.Where(&models.UserModel{}).Delete(tc.user).Error; err != nil {
			log.Printf("Error cleaning up test data: %v", err)
		}
	}
}

func TestUserDao_GetInfo(t *testing.T) {
	// 准备测试用例
	testCases := []struct {
		name           string
		uid            int64
		expectedUser   *models.UserInfoResp
		expectedErrMsg string
	}{
		{
			name: "ValidUser",
			uid:  127299289373216768,
			expectedUser: &models.UserInfoResp{
				UserResp: &models.UserResp{
					UserName: "jannan",
				},
				Gender: true,
			},
			expectedErrMsg: "",
		},
		{
			name:           "NonExistingUser",
			uid:            1272992893732167689,
			expectedUser:   nil,
			expectedErrMsg: "record not found",
		},
		// 添加更多测试用例
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			user, err := userDao.GetInfo(tc.uid)

			// 使用require库进行断言
			if tc.expectedErrMsg == "" {
				// 期望没有错误
				require.NoError(t, err, "Expected no error")
				require.Equal(t, tc.expectedUser.UserName, user.UserName, "User info mismatch")
				require.Equal(t, tc.expectedUser.Gender, user.Gender, "User info mismatch")
			} else {
				// 期望有错误，并检查错误消息
				require.Error(t, err, "Expected error")
				require.Contains(t, err.Error(), tc.expectedErrMsg, "Error message mismatch")
				require.Nil(t, user, "User info should be nil")
			}
		})
	}
}

func TestUserDao_GetUserName(t *testing.T) {
	testCases := []struct {
		name           string
		uid            int64
		expectedName   string
		expectedErrMsg string
	}{
		{
			name:           "ValidUser",
			uid:            127299690419982336,
			expectedName:   "lxj",
			expectedErrMsg: "",
		},
		{
			name:           "NonExistingUser",
			uid:            999,
			expectedName:   "",
			expectedErrMsg: "record not found",
		},
		// 添加更多测试用例
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			username, err := userDao.GetUserName(tc.uid)

			// 使用require库进行断言
			if tc.expectedErrMsg == "" {
				// 期望没有错误
				require.NoError(t, err, "Expected no error")
				require.Equal(t, tc.expectedName, username, "Username mismatch")
			} else {
				// 期望有错误，并检查错误消息
				require.Error(t, err, "Expected error")
				require.Contains(t, err.Error(), tc.expectedErrMsg, "Error message mismatch")
				require.Equal(t, "", username, "Username should be empty")
			}
		})
	}
}
