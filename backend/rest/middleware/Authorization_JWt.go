package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func (m *Middlewares) AuthorizationJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized: Missing Header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Unauthorized: Invalid Header Format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		tokenParts := strings.Split(token, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized: Invalid Token Format", http.StatusUnauthorized)
			return
		}

		headerB64 := tokenParts[0]
		payloadB64 := tokenParts[1]
		signature := tokenParts[2]

		message := headerB64 + "." + payloadB64

		// 2. Verify Signature (Raw byte comparison to prevent timing attacks)
		h := hmac.New(sha256.New, []byte(m.cnf.JWTSecretKey))
		h.Write([]byte(message))
		expectedMac := h.Sum(nil)

		// ক্লায়েন্ট থেকে আসা সিগনেচার ডিকোড করুন
		actualMac, err := base64.RawURLEncoding.DecodeString(signature)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid Signature Base64", http.StatusUnauthorized)
			return
		}

		// hmac.Equal দিয়ে র' বাইট তুলনা (শতভাগ নিরাপদ)
		if !hmac.Equal(actualMac, expectedMac) {
			http.Error(w, "Unauthorized: Signature Mismatch", http.StatusUnauthorized)
			return
		}

		// 3. Decode Payload
		payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadB64)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid Payload Base64", http.StatusUnauthorized)
			return
		}

		var claims struct {
			Sub int64 `json:"sub"`
			Exp int64 `json:"exp"`
		}

		if err := json.Unmarshal(payloadBytes, &claims); err != nil {
			http.Error(w, "Unauthorized: Malformed Claims", http.StatusUnauthorized)
			return
		}

		// 4. Expiry Check
		if time.Now().Unix() > claims.Exp {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		// 5. Inject User into Context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.Sub)
		r = r.WithContext(ctx)

		// 6. Continue
		next.ServeHTTP(w, r)
	})
}
