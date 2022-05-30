package user

type RegUserData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}

type LoginData struct {
	ID    int    `json:"user_id"`
	Token string `json:"token"`
}
