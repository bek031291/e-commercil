package api

import (
	"bytes"
	"ecommerce/internal/db"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

// RegisterHandler godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param request body User true "Данные пользователя"
// @Success 200
// @Router /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var req User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error122121212121": err.Error()})
		return
	}
	// 1. Получить access_token от Keycloak клиента
	token, err := GetAdminToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get admin token"})
		return
	}
	// 2. Отправить запрос в Keycloak для создания пользователя
	keycloakID, err := CreateKeycloakUser(token, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	user := User{
		KeycloakID: keycloakID,
		UserName:   req.UserName,
		Email:      req.Email,
		FullName:   req.FullName,
		IsActive:   true,
	}
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save user in DB"})
		return
	}

	// 3. Отправить письмо
	//_ = SendRegistrationEmail(req.Email, req.Username, req.Password)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginHandler godoc
// @Summary Регистрация пользователя
// @Description Регистрирует нового пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param request body Login true "Данные пользователя"
// @Success 200
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var req Login

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("client_id", "admin-cli")
	data.Set("client_secret", "KG1yDiz8Gp1d3CWWBJ4HxprPsE2GsCnP")
	data.Set("username", req.Username)
	data.Set("password", req.Password)

	resp, err := http.PostForm("http://localhost:8280/realms/master/protocol/openid-connect/token", data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при авторизации"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(resp.StatusCode, gin.H{"error": "неверный логин или пароль"})
		return
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		TokenType   string `json:"token_type"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка обработки ответа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": result.AccessToken,
		"expires_in":   result.ExpiresIn,
		"token_type":   result.TokenType,
	})
}

func GetAdminToken() (string, error) {
	values := url.Values{
		"grant_type":    {"password"},
		"client_id":     {"admin-cli"},
		"client_secret": {"KG1yDiz8Gp1d3CWWBJ4HxprPsE2GsCnP"},
		"username":      {"admin"},
		"password":      {"admin"},
	}

	resp, err := http.PostForm("http://localhost:8280/realms/master/protocol/openid-connect/token", values)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func CreateKeycloakUser(token string, req User) (string, error) {
	user := map[string]interface{}{
		"firstName":     req.FullName,
		"lastName":      req.UserName,
		"username":      req.UserName,
		"email":         req.Email,
		"enabled":       true,
		"emailVerified": true,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     req.Password,
				"temporary": false,
			},
		},
	}

	body, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "http://localhost:8280/admin/realms/master/users", bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode >= 300 {
		return "", fmt.Errorf("failed to create user: %v", resp.Status)
	}
	defer resp.Body.Close()

	// 🆔 Извлекаем ID из Location header
	location := resp.Header.Get("Location")
	parts := strings.Split(location, "/")
	keycloakID := parts[len(parts)-1]

	return keycloakID, nil
}

type User struct {
	KeycloakID string `gorm:"uniqueIndex" json:"keycloak_id"`
	UserName   string `json:"user_name"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	IsActive   bool   `json:"is_active"`
	Password   string `gorm:"-" json:"password"`
}
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
