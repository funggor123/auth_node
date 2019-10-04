package request

type CreatePlayerRequest struct {
	AgentID   string `json:"agent_id" binding:"required"`
	UserName  string `json:"user_name" binding:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CellPhone string `json:"cellphone"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
	//Currency  string `json:"currency" binding:"required"`
	//CountryID string `json:"countryId" binding:"required"`
}

type GetTokenRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetGameStringRequest struct {
	Locale string `json:"locale" binding:"required"`
}

type LaunchGameRequest struct {
	PlayerID  string `json:"player_id" binding:"required"`
	ReturnURL string `json:"return_url" binding:"required"`
	GameID    string `json:"game_id" binding:"required"`
}

type DepositRequest struct {
	TransIDPlatform string  `json:"trans_id_platform" binding:"required"`
	PlayerID        string  `json:"player_id" binding:"required"`
	Money           float32 `json:"money" binding:"required"`
}

type WithDrawRequest struct {
	TransIDPlatform string  `json:"trans_id_platform" binding:"required"`
	PlayerID        string  `json:"player_id" binding:"required"`
	Money           float32 `json:"money" binding:"required"`
}

type BalanceRequest struct {
	PlayerID string `json:"player_id" binding:"required"`
}

type GetTransByPIDRequest struct {
	PlayerID        string `json:"player_id" binding:"required"`
	TransIDPlatform string `json:"trans_id_platform" binding:"required"`
}

type GetGameRecordByPIDRequest struct {
	PlayerID string `json:"player_id" binding:"required"`
}
