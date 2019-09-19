package response

type CreatePlayerResponse struct {
	AgentName      string `json:"agentName" binding:"required"`
	ActiveSessions string `json:"active_sessions" binding:"required"`
	FirstName      string `json:"firstName" binding:"required"`
	LastName       string `json:"lastName" binding:"required"`

	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender" binding:"required"`

	CellPhone string `json:"cellPhone" binding:"required"`

	BirthDate string `json:"birthdate" binding:"required"`
	CountryID string `json:"countryId" binding:"required"`

	Currency string `json:"currency" binding:"required"`
	WalletID string `json:"walletId" binding:"required"`
	ID       string `json:"id" binding:"required"`
	Status   int    `json:"status" binding:"required"`
	Msg      string `json:"msg" binding:"required"`
}

type GetTokenResponse struct {
	Token    string `json:"token" binding:"required"`
	ExpireAt string `json:"expire_at" binding:"required"`
	Status   int    `json:"status" binding:"required"`
	Msg      string `json:"msg" binding:"required"`
}

type LaunchGameResponse struct {
	PlayURL string `json:"play_url" binding:"required"`
	Token   string `json:"token" binding:"required"`
	Status  int    `json:"status" binding:"required"`
	Msg     string `json:"msg" binding:"required"`
}

type DepositResponse struct {
	TransactionID string `json:"transaction_id" binding:"required"`
	PlayerID      string `json:"player_id" binding:"required"`
	Status        int    `json:"status" binding:"required"`
	Msg           string `json:"msg" binding:"required"`
}

type GetBalanceResponse struct {
	Amount float32 `json:"money" binding:"required"`
	Status int     `json:"status" binding:"required"`
	Msg    string  `json:"msg" binding:"required"`
}
