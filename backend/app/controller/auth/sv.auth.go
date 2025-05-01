package auth

import (
	"app/app/enum"
	"app/app/model"
	"app/app/request"
	"app/app/util/jwt"
	"app/config"
	"app/internal/logger"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GenerateToken(ctx context.Context, authType string, user *model.User, isadmin bool) (string, error) {
	claims := jwtlib.MapClaims{
		"auth_type": authType,
		"id":        user.ID,
		"data": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"status":   user.Status,
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(7 * viper.GetDuration("TOKEN_DURATION_USER")).Unix(),
	}

	var secretKey string
	if !isadmin {
		secretKey = viper.GetString("TOKEN_SECRET_USER")
	} else {
		secretKey = viper.GetString("TOKEN_SECRET_ADMIN")
	}

	tokenString, err := jwt.CreateToken(claims, secretKey)
	if err != nil {
		logger.Infof("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}

func (s *Service) GenerateTokenGoogle(id string, userInfo map[string]interface{}) (string, error) {
	claims := jwtlib.MapClaims{
		"auth_type": "google",
		"id":        id,
		"data": map[string]interface{}{
			"id":    userInfo["id"],
			"name":  userInfo["name"],
			"email": userInfo["email"],
		},
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(7 * viper.GetDuration("TOKEN_DURATION_USER")).Unix(),
	}

	tokenString, err := jwt.CreateToken(claims, viper.GetString("TOKEN_SECRET_USER"))
	if err != nil {
		logger.Infof("[error]: %v", err)
		return "", err
	}
	return tokenString, nil
}

func (s *Service) Login(ctx context.Context, req model.User) (*model.User, error) {
	storedUser, err := s.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Log the stored hashed password and the incoming plain text password
	log.Printf("Stored password hash: %s", storedUser.Password)
	log.Printf("Incoming password: %s", req.Password)

	// Check if the provided password matches the stored password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return storedUser, nil
}

func (s *Service) LoginAdmin(ctx context.Context, req model.User) (*model.User, error) {
	storedUser, err := s.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Check if the user is an admin
	if storedUser.RoleID != 2 {
		return nil, errors.New("user is not an admin")
	}

	// Log the stored hashed password and the incoming plain text password
	log.Printf("Stored password hash: %s", storedUser.Password)
	log.Printf("Incoming password: %s", req.Password)

	// Check if the provided password matches the stored password
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return storedUser, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if username == "" {
		return nil, errors.New("username is required")
	}
	m := model.User{}
	logger.Infof("Searching for user with username: %s", username)
	if err := s.db.NewSelect().Model(&m).
		Where("username = ?", username).Scan(ctx); err != nil {
		logger.Infof("[error]: %v", err)
		return nil, err
	}
	return &m, nil
}

func (s *Service) GetUserDetailByToken(ctx context.Context, tokenString string) (*model.User, error) {
	// Parse the JWT token
	token, err := jwtlib.Parse(tokenString, func(token *jwtlib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(viper.GetString("TOKEN_SECRET_USER")), nil
	})
	if err != nil {
		return nil, err
	}

	// Extract user ID from token claims
	claims, ok := token.Claims.(jwtlib.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// data, ok := claims["data"].(map[string]interface{})
	// if !ok {
	// 	return userDetail, errors.New("invalid token payload: data field is missing")
	// }

	userID, ok := claims["id"]
	if !ok {
		return nil, errors.New("invalid token payload: id field is missing or not a number")
	}

	logger.Infof("User ID: %v", userID)

	// Define the query and execute it using Bun
	var user model.User
	err = s.db.NewSelect().
		Model(&user).
		Where("id = ?", userID).
		Scan(ctx)
	if err != nil {
		logger.Infof("[error]: Failed to fetch user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (s *Service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	m := model.User{}
	if err := s.db.NewSelect().Model(&m).
		Where("email = ?", email).Scan(ctx); err != nil {
		return nil, err
	}
	return &m, nil
}

func (s *Service) Create(ctx context.Context, user model.User) (*model.User, error) {
	m := model.User{
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DisplayName: user.DisplayName,
		RoleID:      1,
		Status:      enum.STATUS_ACTIVE,
	}

	// insert user
	if _, err := s.db.NewInsert().Model(&m).Exec(ctx); err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *Service) ResetPassword(ctx context.Context, req request.ResetPassword) error {
	user, err := s.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	// create UserReset
	userReset := model.UserReset{
		Email:     user.Email,
		DateStart: time.Now().Unix(),
		DateEnd:   time.Now().Add(time.Hour * 24).Unix(),
	}

	if _, err := s.db.NewInsert().Model(&userReset).Exec(ctx); err != nil {
		return err
	}

	err = config.SendEmail(
		user.Email,
		"Reset Password",
		"Reset Password",
		fmt.Sprintf(
			"Click the following link to reset your password: <a href=\"%s?token=%s\">Reset Password</a>",
			req.RedirectURL,
			userReset.ID,
		),
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ChangePasswordReset(ctx context.Context, req request.ChangePasswordReset) error {
	tn := time.Now().Unix()
	userReset := model.UserReset{}
	ex, err := s.db.NewSelect().Model(&userReset).Where("id = ?", req.Token).Where("date_end >= ?", tn).Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("token is invalid or expired")
	}

	err = s.db.NewSelect().Model(&userReset).Where("id = ?", req.Token).Scan(ctx)
	if err != nil {
		return err
	}

	user, err := s.GetUserByEmail(ctx, userReset.Email)
	if err != nil {
		return err

	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	if _, err := s.db.NewUpdate().Model(user).WherePK().Exec(ctx); err != nil {
		return err
	}

	return nil
}
