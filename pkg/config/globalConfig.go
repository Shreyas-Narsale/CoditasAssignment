package config

import (
	"os"
	"strconv"
	"time"
)

type GlobalConfig struct {
	JwtSecretKeyAccessToken  []byte
	JwtSecretKeyRefreshToken []byte
	AccessTokenExpTime       time.Duration
	RefreshTokenExpTime      time.Duration
	Login                    LoginConfig
}

type LoginConfig struct {
	PasswordPepper string
}

var globalConf = &GlobalConfig{}

func GetGlobalConfig() *GlobalConfig {
	return globalConf
}

func LoadGlobalConfig() {
	globalConf.JwtSecretKeyAccessToken = []byte(os.Getenv("JWT_SECRET_KEY_ACCESS_TOKEN"))
	globalConf.JwtSecretKeyRefreshToken = []byte(os.Getenv("JWT_SECRET_KEY_REFRESH_TOKEN"))
	accessTokenExpTimeString, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP_TIME"))
	if err != nil {
		accessTokenExpTimeString = 30
		return
	}
	globalConf.AccessTokenExpTime = time.Duration(accessTokenExpTimeString) * time.Second
	refreshTokenExpTimeString, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP_TIME"))
	if err != nil {
		accessTokenExpTimeString = 2400
		return
	}
	globalConf.RefreshTokenExpTime = time.Duration(refreshTokenExpTimeString) * time.Second
	globalConf.Login.PasswordPepper = os.Getenv("PASSWORD_PEPPER")

}
