package auth

type TwitchAuthBody struct {
	// names specified by twitch api: https://dev.twitch.tv/docs/api/reference/#get-users
	AccessToken  string `json:"access_token" xml:"access_token" form:"access_token"`
	RefreshToken string `json:"refresh_token" xml:"refresh_token" form:"refresh_token"`
}

type TwitchClaims struct {
	// names specified by twitch api: https://dev.twitch.tv/docs/authentication/getting-tokens-oidc/#getting-claims-information-from-an-access-token
	Id          string `json:"sub"`
	DisplayName string `json:"preferred_username"`
	ImgUrl      string `json:"picture"`
}

type AccountAuthBody struct {
	Id           string `query:"id"`
	RefreshToken string `query:"refresh_token"`
}

type inputsNames struct {
	Id           string
	AccessToken  string
	RefreshToken string
}

var HTML_INPUTS_NAMES = inputsNames{
	Id:           "id",
	AccessToken:  "access_token",
	RefreshToken: "refresh_token",
}
