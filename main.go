package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    Message string `json:"message"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if req.Email == "admin@gmail.com" && req.Password == "pass321" {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(LoginResponse{Message: "Login successful"})
    } else {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(LoginResponse{Message: "Invalid email or password"})
    }
}

func main() {
    http.HandleFunc("/login", loginHandler)
    fmt.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
