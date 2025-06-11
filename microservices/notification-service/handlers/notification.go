package handlers

import (
    "encoding/json"
    "log"
    "net/http"
)

type NotificationRequest struct {
    UserID  string `json:"user_id"`
    Message string `json:"message"`
    Type    string `json:"type"` // email or sms
}

func SendNotification(w http.ResponseWriter, r *http.Request) {
    var req NotificationRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    log.Printf("Sending %s notification to user %s: %s", req.Type, req.UserID, req.Message)

    // TODO: Здесь будет реальная логика отправки email/SMS
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Notification sent"))
}
