package auth

type TwitchAuthBody struct{
  // names specified by twitch api: https://dev.twitch.tv/docs/api/reference/#get-users
  AccessToken string `json:"access_token" xml:"access_token" form:"access_token"`
  RefreshToken string `json:"refresh_token" xml:"refresh_token" form:"refresh_token"`
}

type TwitchClaims struct{
  // names specified by twitch api: https://dev.twitch.tv/docs/authentication/getting-tokens-oidc/#getting-claims-information-from-an-access-token
  Id string `json:"sub"`
  DisplayName string `json:"preferred_username"`
  ImgUrl string `json:"picture"`
}

type AccountAuthBody struct{
  Id string `json:"id" xml:"id" form:"id"`
  AccessToken string `json:"access_token" xml:"access_token" form:"access_token"`
}
