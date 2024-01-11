package controllers

import (
	"bluebell/pkg/types"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostController_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, PostController{}.Create)

	body := `{
				"title":"测试ing",
				"content":"测试ing",
				"community_id":5
			}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	////方法1 判断响应内容中是否包含指定字符串
	//assert.Contains(t, w.Body.String(), "无效的token")

	//	方法2
	res := new(types.Response)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatal("ERROR")
	}
	assert.Equal(t, res.Code, types.CodeInvalidToken)
}
