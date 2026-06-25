package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"gin-todo/routes"

	"github.com/gin-gonic/gin"
)

// GetTodos の成功ケースをテストする
func TestGetTodos_Success(t *testing.T) {
	// Gin をテストモードにする
	gin.SetMode(gin.TestMode)

	// Router を作成
	r := routes.SetupRouter()

	// テスト用リクエスト作成
	req, err := http.NewRequest(
		http.MethodGet,
		"/todos",
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// レスポンス記録用
	w := httptest.NewRecorder()

	// API実行
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusOK,
			w.Code,
		)
	}

}

// CreateTodo の成功ケースをテストする
func TestCreateTodo_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	jsonBody := []byte(`{"title":"Test todo","completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusCreated,
			w.Code,
		)
	}
}

// 無効なリクエストボディに対して、入力検証エラーとして 400 を返す
func TestCreateTodo_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// title が文字列ではなく数値になっている。その結果、ShouldBindJSON のバインドに失敗する。
	jsonBody := []byte(`{"title":123,"completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

// タイトルが重複する場合、409 Conflict を返すことを確認するテスト
func TestCreateTodo_DuplicateTitle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// まず、タイトルが "Duplicate todo" の Todo を作成する
	jsonBody := []byte(`{"title":"Duplicate todo","completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 2回目のリクエストで同じタイトルの Todo を作成しようとする
	req2, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusConflict,
			w2.Code,
		)
	}
}

// GetTodoByID の成功ケースをテストする
func TestGetTodoByID_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// まず、Todo を作成する
	jsonBody := []byte(`{"title":"GetByID todo","completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 作成した Todo の ID を取得する
	var createdTodoResponse struct {
		Code string `json:"code"`
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	// レスポンスボディを構造体にデコードする
	if err := json.Unmarshal(w.Body.Bytes(), &createdTodoResponse); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 作成した Todo の ID を使って GET リクエストを送信する
	req2, err := http.NewRequest(
		http.MethodGet,
		"/todos/"+strconv.Itoa(createdTodoResponse.Data.ID),
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusOK,
			w2.Code,
		)
	}
}

// GetTodoByID の失敗ケースをテストする(存在しない ID)
func TestGetTodoByID_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 存在しない ID で GET リクエストを送信する
	req, err := http.NewRequest(
		http.MethodGet,
		"/todos/9999", // 存在しない ID
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}

// GetTodoByID の失敗ケースをテストする（無効な ID）
func TestGetTodoByID_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 無効な ID で GET リクエストを送信する
	req, err := http.NewRequest(
		http.MethodGet,
		"/todos/invalid", // 無効な ID
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

// UpdateTodo の成功ケースをテストする
func TestUpdateTodo_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// まず、Todo を作成する
	jsonBody := []byte(`{"title":"Update todo","completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 作成した Todo の ID を取得する
	var createdTodoResponse struct {
		Code string `json:"code"`
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	// レスポンスボディを構造体にデコードする
	if err := json.Unmarshal(w.Body.Bytes(), &createdTodoResponse); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 作成した Todo の ID を使って PUT リクエストを送信する
	updateBody := []byte(`{"title":"Updated todo","completed":true}`)
	req2, err := http.NewRequest(
		http.MethodPut,
		"/todos/"+strconv.Itoa(createdTodoResponse.Data.ID),
		bytes.NewBuffer(updateBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusOK,
			w2.Code,
		)
	}
}

// UpdateTodo の失敗ケースをテストする（存在しない ID）
func TestUpdateTodo_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 存在しない ID で PUT リクエストを送信する
	updateBody := []byte(`{"title":"Updated todo","completed":true}`)
	req, err := http.NewRequest(
		http.MethodPut,
		"/todos/9999", // 存在しない ID
		bytes.NewBuffer(updateBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}

// UpdateTodo の失敗ケースをテストする（無効な ID）
func TestUpdateTodo_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 無効な ID で PUT リクエストを送信する
	updateBody := []byte(`{"title":"Updated todo","completed":true}`)

	req, err := http.NewRequest(
		http.MethodPut,
		"/todos/invalid", // 無効な ID
		bytes.NewBuffer(updateBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

// UpdateTodo の失敗ケースをテストする（タイトル重複）
func TestUpdateTodo_DuplicateTitle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// まず、2つの Todo を作成する
	jsonBody1 := []byte(`{"title":"First todo","completed":false}`)
	req1, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody1),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	r.ServeHTTP(w1, req1)

	// 作成した最初の Todo の ID を取得する
	var createdTodoResponse1 struct {
		Code string `json:"code"`
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	if err := json.Unmarshal(w1.Body.Bytes(), &createdTodoResponse1); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 2つ目の Todo を作成する
	jsonBody2 := []byte(`{"title":"Second todo","completed":false}`)
	req2, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody2),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	// 作成した2つ目の Todo の ID を取得する
	var createdTodoResponse2 struct {
		Code string `json:"code"`
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	if err := json.Unmarshal(w2.Body.Bytes(), &createdTodoResponse2); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 2つ目の Todo のタイトルを、1つ目の Todo のタイトルと同じにして更新しようとする
	updateBody := []byte(`{"title":"First todo","completed":true}`)
	req3, err := http.NewRequest(
		http.MethodPut,
		"/todos/"+strconv.Itoa(createdTodoResponse2.Data.ID),
		bytes.NewBuffer(updateBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req3.Header.Set("Content-Type", "application/json")

	w3 := httptest.NewRecorder()

	r.ServeHTTP(w3, req3)

	if w3.Code != http.StatusConflict {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusConflict,
			w3.Code,
		)
	}
}

// DeleteTodo の成功ケースをテストする
func TestDeleteTodo_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// まず、Todo を作成する
	jsonBody := []byte(`{"title":"Delete todo","completed":false}`)

	req, err := http.NewRequest(
		http.MethodPost,
		"/todos",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// 作成した Todo の ID を取得する
	var createdTodoResponse struct {
		Code string `json:"code"`
		Data struct {
			ID int `json:"id"`
		} `json:"data"`
	}

	// レスポンスボディを構造体にデコードする
	if err := json.Unmarshal(w.Body.Bytes(), &createdTodoResponse); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	// 作成した Todo の ID を使って DELETE リクエストを送信する
	req2, err := http.NewRequest(
		http.MethodDelete,
		"/todos/"+strconv.Itoa(createdTodoResponse.Data.ID),
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w2 := httptest.NewRecorder()

	r.ServeHTTP(w2, req2)

	if w2.Code != http.StatusOK {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusOK,
			w2.Code,
		)
	}
}

// DeleteTodo の失敗ケースをテストする（存在しない ID）
func TestDeleteTodo_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 存在しない ID で DELETE リクエストを送信する
	req, err := http.NewRequest(
		http.MethodDelete,
		"/todos/9999", // 存在しない ID
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusNotFound,
			w.Code,
		)
	}
}

// DeleteTodo の失敗ケースをテストする（無効な ID）
func TestDeleteTodo_InvalidID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := routes.SetupRouter()

	// 無効な ID で DELETE リクエストを送信する
	req, err := http.NewRequest(
		http.MethodDelete,
		"/todos/invalid", // 無効な ID
		nil,
	)

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"Expected status %d, but got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}
