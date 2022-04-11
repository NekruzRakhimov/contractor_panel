package dto

import "contractor_panel/domain/model"

type CredentialsDto struct {
	UserLogin    string `json:"userLogin"`
	UserPassword string `json:"userPassword"`
}

type TokensDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func ConvertTokenDetailsToDto(td *model.TokenDetails) TokensDto {
	return TokensDto{AccessToken: td.AccessToken, RefreshToken: td.RefreshToken}
}

func ConvertCredentialsDtoToEntity(dto *CredentialsDto) model.Credentials {
	return model.Credentials{UserLogin: dto.UserLogin, UserPassword: dto.UserPassword}
}
