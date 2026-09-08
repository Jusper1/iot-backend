package services

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"strconv"
	"time"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	AdminRepository        *repositories.AdminRepository
	PasswordResetRepository *repositories.PasswordResetRepository
}

func NewAuthService(
	adminRepository *repositories.AdminRepository,
	passwordResetRepository *repositories.PasswordResetRepository,
) *AuthService {

	return &AuthService{
		AdminRepository:         adminRepository,
		PasswordResetRepository: passwordResetRepository,
	}
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, *models.AdminUser, error) {

	user, err := s.AdminRepository.FindByEmail(email)

	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	if user.Flag != 1 {
		return "", nil, errors.New("akun tidak aktif")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil {
		return "", nil, errors.New("email atau password salah")
	}

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return "", nil, errors.New("JWT_SECRET belum dikofigurasi")
	}

	expireHours := 24

	if value := os.Getenv("JWT_EXPIRE_HOURS"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			expireHours = parsed
		}
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"used_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"iat":     now.Unix(),
		"exp":     now.Add(time.Duration(expireHours) * time.Hour).Unix(),

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims,)

	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "",nil,err
	}

	return signedToken, user, nil

}

func (s *AuthService) ForgotPassword(
	email string,
)(string, error) {

	user, err := s.AdminRepository.FindByEmail(email)

	if err != nil {
			
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}

		return "", err
	}

	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)

	if err != nil {
		return "", err
	}

	rawToken := hex.EncodeToString(tokenBytes)

	hash := sha256.Sum256([]byte(rawToken))

	tokenHash := hex.EncodeToString(hash[:])

	minutes := 30

	if value := os.Getenv("RESET_TOKEN_EXPIRE_MINUTES"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			minutes = parsed
		}
	}

	resetToken := &models.PasswordResetToken{
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(
			time.Duration(minutes) * time.Minute,
		),
	}

	err = s.PasswordResetRepository.Create(resetToken)

	if err != nil {
		return "", err
	}

	return rawToken, nil
}

func (s *AuthService) ResetPassword(
	rawToken string,
	newPassword string,
) error {

	hash := sha256.Sum256([]byte(rawToken))

		tokenHash := hex.EncodeToString(hash[:])

		resetToken, err := 
		s.PasswordResetRepository.FindValidToken(tokenHash)

		if err != nil {
				return errors.New(
					"token tidak valid atau sudah digunakan",
				)
			}

		passwordHash, err := 
		bcrypt.GenerateFromPassword(
			[]byte(newPassword),
			bcrypt.DefaultCost,
		)

		if err != nil {
			return err
		}

		err = s.AdminRepository.UpdatePassword(
			resetToken.UserID,
			string(passwordHash),
		)

		if err != nil {
			return err
		}

		return s.PasswordResetRepository.MarkUsed(
			resetToken.ID,
		)
	}