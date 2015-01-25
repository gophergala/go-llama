package intchess

type APITypeOnly struct {
	Type string `json:"type"`
}

/*
{
"type":"authentication_request",
"username":"test",
"user_token":"test"
}
*/
type APIAuthenticationRequest struct {
	Type      string `json:"type"`
	Username  string `json:"username"`
	UserToken string `json:"user_token"`
}

type APIAuthenticationResponse struct {
	Type     string `json:"type"`
	Response string `json:"response"`
	User     *User  `json:"user,omitempty"` //This will thus be blank (not NULL) if they are not logged in
}

type APIGameRequest struct {
	Type     string `json:"type"`
	Opponent *User  `json:"opponent"`
}

/*
{
"type":"game_response",
"response":"ok"
}
*/
type APIGameResponse struct {
	Type     string `json:"type"`
	Response string `json:"response"`
}

type APIGameOutput struct {
	Type string     `json:"type"`
	Game *ChessGame `json:"game"`
}

type APISignupRequest struct {
	Type      string `json:"type"`
	Username  string `json:"username"`
	UserToken string `json:"user_token"`
	IsAi      bool   `json:"is_ai"`
	VersesAi  bool   `json:"verses_ai"`
}
