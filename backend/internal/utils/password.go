package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// bcryptのコスト設定（DefaultCostは推奨値）
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// パスワード検証関数
func CheckPassword(hash, password string) error {
	// パスワードとハッシュ値の比較
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
