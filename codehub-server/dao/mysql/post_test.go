package mysql

import (
	"bluebell/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostDao_Create(t *testing.T) {
	testCases := []struct {
		name           string
		post           *models.PostModel
		expectedErrMsg string
	}{
		{
			name: "ValidPost",
			post: &models.PostModel{
				Title:   "Test Post",
				Content: "Test Content",
				// 添加其他必要的字段
			},
			expectedErrMsg: "",
		},
		// 添加更多测试用例
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			err := postDao.Create(tc.post)

			// 使用require库进行断言
			if tc.expectedErrMsg == "" {
				// 期望没有错误
				require.NoError(t, err, "Expected no error")
			} else {
				// 期望有错误，并检查错误消息
				require.Error(t, err, "Expected error")
				require.Contains(t, err.Error(), tc.expectedErrMsg, "Error message mismatch")
			}
		})

		result := db.Exec("DELETE FROM post WHERE title LIKE ?", "%Test%")
		require.NoError(t, result.Error, "Expected no error When Delete test case")
	}
}

func TestPostDao_GetList(t *testing.T) {
	testCases := []struct {
		name            string
		condition       map[string]interface{}
		page            *models.Page
		expectedList    []*models.PostResp
		expectedListLen int
		expectedErrMsg  string
	}{
		{
			name:      "EmptyConditionAndPage",
			condition: map[string]interface{}{},
			page: &models.Page{
				Num:  1,
				Size: 5,
			},
			expectedListLen: 5,
			expectedErrMsg:  "",
		},
		{
			name: "ValidConditionWithResults",
			condition: map[string]interface{}{
				"author_id": 127299690419982336,
			},
			page: &models.Page{
				Num:  1,
				Size: 10,
			},
			expectedListLen: 4,
			expectedErrMsg:  "",
		},
		{
			name:      "NonExistingPage",
			condition: map[string]interface{}{},
			page: &models.Page{
				Num:  10, // 不存在的页数
				Size: 10,
			},
			expectedList:   make([]*models.PostResp, 0), // 期望返回一个空列表
			expectedErrMsg: "",
		},
		{
			name: "InvalidCondition",
			condition: map[string]interface{}{
				"invalid_field": "invalid_value", // 不存在的字段
			},
			page: &models.Page{
				Num:  1,
				Size: 10,
			},
			expectedList:   nil, // 期望返回 nil，因为条件无效
			expectedErrMsg: "Unknown column 'invalid_field'",
		},
	}

	// 执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 执行测试逻辑
			list, err := postDao.GetList(tc.condition, tc.page)

			// 使用require库进行断言
			if tc.expectedErrMsg == "" {
				// 期望没有错误
				require.NoError(t, err, "Expected no error")
				require.Equal(t, tc.expectedListLen, len(list), "Post list's len mismatch")
				if tc.expectedList != nil {
					require.Equal(t, tc.expectedList, list, "Post list mismatch")
				}
			} else {
				// 期望有错误，并检查错误消息
				require.Error(t, err, "Expected error")
				require.Contains(t, err.Error(), tc.expectedErrMsg, "Error message mismatch")
				require.Nil(t, list, "Post list should be nil")
			}
		})
	}

}
