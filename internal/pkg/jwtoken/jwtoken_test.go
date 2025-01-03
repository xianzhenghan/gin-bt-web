package jwtoken

import (
	"fmt"
	"testing"
	"time"

	"gin-api-mono/internal/proposal"
)

const secret = "i1ydX9RtHyuJTrw7frcu"

func TestSign(t *testing.T) {
	sessionUserInfo := proposal.SessionUserInfo{
		Id:       1001,
		UserName: "gin-api-mono",
		NickName: "mono",
	}

	tokenString, err := New(secret).Sign(sessionUserInfo, 24*time.Hour)

	fmt.Println(tokenString, err)
	if err != nil {
		t.Error("sign error", err)
		return
	}

	t.Log(tokenString)
}

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMSwidXNlcm5hbWUiOiJnaW4tYXBpLW1vbm8iLCJuaWNrbmFtZSI6Im1vbm8iLCJleHAiOjE3MzU4NjgwNTAsIm5iZiI6MTczNTc4MTY1MCwiaWF0IjoxNzM1NzgxNjUwfQ.wu9FHJUmgifsQVwAO9nqvyenywM6kKXiX6lsQqfxn3I"
	jwtInfo, err := New(secret).Parse(tokenString)
	if err != nil {
		t.Error("parse error", err)
		return
	}

	t.Log(string(jwtInfo.Marshal()))
}
