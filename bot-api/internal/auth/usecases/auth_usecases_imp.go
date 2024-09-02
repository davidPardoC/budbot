package usecases

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/davidPardoC/budbot/config"
	"github.com/davidPardoC/budbot/constants"
	"github.com/davidPardoC/budbot/internal/auth/delivery/dtos"
	"github.com/davidPardoC/budbot/internal/auth/models"
	userModels "github.com/davidPardoC/budbot/internal/users/models"
	userRepository "github.com/davidPardoC/budbot/internal/users/repository"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthUseCasesImp struct {
	userRepository userRepository.IUserRepository
	config         config.Config
}

func NewAuthUseCasesImp(userRepository userRepository.IUserRepository, config config.Config) IAuthUseCases {
	return &AuthUseCasesImp{userRepository: userRepository, config: config}
}

func (auc *AuthUseCasesImp) Login(dto dtos.TelegramCallbackDto, query map[string][]string) (*models.Credentials, error) {
	parsedId, err := strconv.ParseInt(dto.Id, 10, 64)

	if err != nil {
		return nil, err
	}

	checkString := generateDataCheckString(query)

	keyHash := sha256.New()
	keyHash.Write([]byte(auc.config.Telegram.Token))
	secretkey := keyHash.Sum(nil)

	hash := hmac.New(sha256.New, secretkey)
	hash.Write([]byte(checkString))
	hashstr := hex.EncodeToString(hash.Sum(nil))

	if hashstr != dto.Hash {
		return nil, errors.New("hash does not match")
	}

	var user *userModels.User

	user, err = auc.userRepository.FindByChatID(parsedId)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		user, _ = auc.userRepository.CreateUser(parsedId, "", dto.FirstName, dto.LastName, "user", dto.PhotoUrl)
	} else if err != nil {
		return nil, err
	}

	user.PhotoUrl = dto.PhotoUrl
	auc.userRepository.UpdateUser(user)

	tokenClaims := models.TokenCustomClaims{
		ChatID:   user.ChatID,
		UserId:   user.ID,
		PhotoUrl: user.PhotoUrl,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(constants.DayInMinutes)).Unix(),
		},
	}

	refreshTokenClaims := models.TokenCustomClaims{
		ChatID:   user.ChatID,
		UserId:   user.ID,
		PhotoUrl: user.PhotoUrl,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(constants.DayInMinutes*7)).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessTokenStr, _ := accessToken.SignedString([]byte(auc.config.Auth.JwtSecret))

	refreshTokn := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenStr, _ := refreshTokn.SignedString([]byte(auc.config.Auth.RefresJwtSecret))

	credentials := &models.Credentials{
		Token:        accessTokenStr,
		RefreshToken: refreshTokenStr,
	}

	return credentials, nil
}

func generateDataCheckString(query map[string][]string) string {
	var data []string
	for key, values := range query {
		if key != "hash" {
			data = append(data, fmt.Sprintf("%s=%s", key, values[0]))
		}
	}

	sort.Strings(data)

	return strings.Join(data, "\n")
}
