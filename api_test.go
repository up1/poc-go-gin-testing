package q1_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"q1"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestClaimV4(t *testing.T) {
	expected := "0x123DFPER321XCAAZ"
	gin.SetMode(gin.TestMode)
	reader := strings.NewReader(fmt.Sprintf("address_id=%s", expected))
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/claim", reader)
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.POST("/claim", q1.ClaimRewards)
	r.ServeHTTP(res, c.Request)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response q1.Response
	json.Unmarshal(res.Body.Bytes(), &response)
	if response.Message != expected {
		t.Errorf("Message: got %v want %v", response.Message, expected)
	}
}

func TestClaimV4_with_JSON(t *testing.T) {
	expected := "0x123DFPER321XCAAZ"
	gin.SetMode(gin.TestMode)
	request := q1.Request{
		Address_id: expected,
	}
	jsonStr, _ := json.Marshal(request)

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/claim/json", bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("Content-Type", "application/json")

	r.POST("/claim/json", q1.ClaimRewardsJson)
	r.ServeHTTP(res, c.Request)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response q1.Response
	json.Unmarshal(res.Body.Bytes(), &response)
	if response.Message != expected {
		t.Errorf("Message: got %v want %v", response.Message, expected)
	}
}
