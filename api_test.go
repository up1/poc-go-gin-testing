package q1_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"q1"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestClaimV4(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// request := q1.Request{
	// 	Address_id: "0x123DFPER321XCAAZ",
	// }
	// jsonStr, _ := json.Marshal(request)

	reader := strings.NewReader("address_id=0x123DFPER321XCAAZ")
	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/claim", reader)
	c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r.POST("/claim", q1.ClaimRewards)
	r.ServeHTTP(res, c.Request)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response q1.Response
	json.Unmarshal(res.Body.Bytes(), &response)
	if response.Message != "0x123DFPER321XCAAZ" {
		t.Errorf("Message: got %v want %v",
			response.Message, "0x123DFPER321XCAAZ")
	}
}

func TestClaimV4_with_JSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	request := q1.Request{
		Address_id: "0x123DFPER321XCAAZ",
	}
	jsonStr, _ := json.Marshal(request)

	res := httptest.NewRecorder()
	c, r := gin.CreateTestContext(res)
	c.Request = httptest.NewRequest(http.MethodPost, "/claim/json", bytes.NewBuffer(jsonStr))
	c.Request.Header.Set("Content-Type", "application/json")
	
	r.POST("/claim/json", q1.ClaimRewardsJson)
	r.ServeHTTP(res, c.Request)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response q1.Response
	json.Unmarshal(res.Body.Bytes(), &response)
	if response.Message != "0x123DFPER321XCAAZ" {
		t.Errorf("Message: got %v want %v",
			response.Message, "0x123DFPER321XCAAZ")
	}
}
