package request

type CreatePlayerRequest struct {
	AgentID   string `json:"agentId" binding:"required"`
	UserName  string `json:"username" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CellPhone string `json:"cell_phone" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	BirthDate string `json:"birthdate" binding:"required"`
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
