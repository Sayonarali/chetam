package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (rt Router) GetApiV1User(w http.ResponseWriter, r *http.Request) {
	token := TrimPrefix(r.Header.Get("Authorization"), "Bearer ")

	fmt.Println(token)
	claims, err := rt.authService.ValidateToken(token)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	user, err := rt.userService.GetUserByLogin(claims.Login)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func TrimPrefix(s, prefix string) string {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):]
	}
	return s
}
