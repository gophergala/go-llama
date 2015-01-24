package intchess

type APITypeOnly struct {
	Type string `json:"type"`
}

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

type APIGameResponse struct {
	Type     string `json:"type"`
	Response string `json:"response"`
}

type APIGameOutput struct {
	Type string    `json:"type"`
	Game ChessGame `json:"game"`
}
