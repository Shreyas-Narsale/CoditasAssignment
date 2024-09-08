package login

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/jwt"
)

func GetTokens(userId string) (string, string, error) {

	globalConf := config.GetGlobalConfig()
	accessTokenString, err := jwt.GenerateJWT(userId, globalConf.JwtSecretKeyAccessToken, globalConf.AccessTokenExpTime)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := jwt.GenerateJWT(userId, globalConf.JwtSecretKeyRefreshToken, globalConf.RefreshTokenExpTime)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
