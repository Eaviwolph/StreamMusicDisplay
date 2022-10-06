package structs

import "fmt"

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func (t AccessToken) String() string {
	return fmt.Sprintf("{\n\taccess_token: %v\n\ttoken_type: %v\n\texpires_in: %v\n\trefresh_token: %v\n\tscope: %v\n}", t.AccessToken, t.TokenType, t.ExpiresIn, t.RefreshToken, t.Scope)
}
