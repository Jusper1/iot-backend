package handlers

import (
	"iot-backend/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService,
	) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

type LoginRequest struct {
	Email string `json:"Email"`
	Password string `json:"Password"`
}

func (h *AuthHandler) Login(c *gin.Context) {

	var input LoginRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Format request tidak valid", 
		})
		return
	}

	input.Email = strings.TrimSpace(
	strings.ToLower(input.Email))

	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Email dan password harus diisi",
		})
		return
	}

	token, user, err :=
	h.Service.Login(
		input.Email,
		input.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login berhasil",
		"data": gin.H{
			"token": token,
			"user": gin.H{
				"id": user.ID,
				"nama": user.Nama,
				"email": user.Email,
				"Role": user.Role,
			},
		},
	})
}

type ForgotPasswordRequest struct {
	Email string `json:"Email"`
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {

	var input ForgotPasswordRequest

	if err :=c.ShouldBindJSON(&input); err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Format request tidak valid",
		})
		return
	}

	input.Email = strings.TrimSpace(
	strings.ToLower(input.Email),
	)
	
	token, err := h.Service.ForgotPassword(input.Email)

	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Gagal Memproses lupa password",
		})
		return
	}
	
	Response := gin.H{
		"success": true,
		"message": "Jika email Terdaftar, link reset password telah dikirim",
	}

	if token != "" {
		Response["reset_token_dev"] = token
	}

	c.JSON(http.StatusOK, Response)
}

type ResetPasswordRequest struct {
	Token string `json:"ResetToken"`
	NewPassword string `json:"New_Password"`
}

func (h *AuthHandler) ResetPassword(c *gin.Context){

	var input ResetPasswordRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Format request tidak valid",
		})
		return
	}

	if len(input.Token) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Password minimal 8 karakter",
		})
		return
	}

	err := h.Service.ResetPassword(
		input.Token,
		input.NewPassword,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Password berhasil diubah",
	})
}