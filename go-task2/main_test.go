package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCalculateHandler_Success(t *testing.T) {
    req := Request{Numbers: []int{1, 2, 3, 4, 5}}
    body, _ := json.Marshal(req)

    r := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader(body))
    w := httptest.NewRecorder()

    calculateHandler(w, r)

    if w.Code != http.StatusOK {
        t.Errorf("Статус: ожидался 200, получен %d", w.Code)
    }

    var resp Response
    json.NewDecoder(w.Body).Decode(&resp)
    
    if resp.Sum != 55 {
        t.Errorf("Сумма: ожидалась 55, получена %d", resp.Sum)
    }
}

func TestCalculateHandler_InvalidMethod(t *testing.T) {
    r := httptest.NewRequest(http.MethodGet, "/calculate", nil)
    w := httptest.NewRecorder()

    calculateHandler(w, r)

    if w.Code != http.StatusMethodNotAllowed {
        t.Errorf("Статус: ожидался 405, получен %d", w.Code)
    }
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
    r := httptest.NewRequest(http.MethodPost, "/calculate", bytes.NewReader([]byte("invalid")))
    w := httptest.NewRecorder()

    calculateHandler(w, r)

    if w.Code != http.StatusBadRequest {
        t.Errorf("Статус: ожидался 400, получен %d", w.Code)
    }
}