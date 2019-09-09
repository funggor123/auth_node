package node

import (
	"net/http"

	common "../common"
	e "../e"
	model "../model"
	request "../request"
	response "../response"
	"github.com/gin-gonic/gin"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Node struct {
	NodeType string `json:"node_type" binding:"required"`
}

var node *Node

func getNode() *Node {
	if node != nil {
		return node
	}
	node = new(Node)
	return node
}

// .. // User

func (node Node) CreatePlayer(c *gin.Context) {

	var createPlayerRequest request.CreatePlayerRequest
	err := c.BindJSON(&createPlayerRequest)
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request JSON -",
		createPlayerRequest)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.CreatePlayerErrorResponse{Message: err.Error()})
		return
	}

	agents, err:= model.GetAgentByIDFromDB(createPlayerRequest.AgentID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.CreatePlayerErrorResponse{Message: err.Error()})
		return
	}

	if (len(agents) == 0 ) {
		err := errors.New("Bad Agent ID")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.CreatePlayerErrorResponse{Message: err.Error()})
		return
	}

	wallet, err:= model.CreateWalletnInDB()

	player := model.Player{AgentID:  bson.ObjectIdHex(createPlayerRequest.AgentID),
		UserName:  createPlayerRequest.UserName,
		FirstName: createPlayerRequest.FirstName,
		LastName:  createPlayerRequest.LastName,
		Email:     createPlayerRequest.Email,
		Gender:    createPlayerRequest.Gender,
		BirthDate: createPlayerRequest.BirthDate,
		Currency:  createPlayerRequest.Currency,
		WalletID: wallet.ID,
		CountryID: createPlayerRequest.CountryID}

	player, err = model.CreatePlayerInDb(player)

	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.CreatePlayerErrorResponse{Message: err.Error()})
		return
	}


	c.JSON(http.StatusOK, response.CreatePlayerResponse{
		UserName: player.UserName,
		FirstName: player.UserName,
		LastName: player.LastName,
		Email: player.Email,
		Gender: player.Gender,
		BirthDate: player.BirthDate,
		Currency: player.Currency,
		ID : fmt.Sprintf("%x", string(player.ID)),
		CountryID: player.CountryID})
	return
}

//... // Agent

func (node Node) GetAgentToken(c *gin.Context) {
	username, password := c.Query("username"), c.Query("password")
	getTokenRequest := request.GetTokenRequest{UserName: username, Password: password}
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request -",
	getTokenRequest)

	agents, err := model.GetAgentFromDB(username, password)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.GetTokenErrorResponse{Error: err.Error()})
		return
	}

	if (len(agents) == 0 ) {
		err := errors.New("Bad Username/Password/Failed to Authenticate")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.GetTokenErrorResponse{Error: err.Error()})
		return
	}

	tokenString, err:= model.CreateTokenInDB(agents[0].ID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.GetTokenErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.GetTokenResponse{
		Token: "Bearer " + tokenString})
}
func (node Node) StartAPI(c *gin.Context) {

	
}

func (node Node) GetGameString(c *gin.Context) {
	locale := c.Query("locale")
	getGameStringRequest := request.GetGameStringRequest{Locale: locale}
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request -",
	getGameStringRequest)

	games, err := model.GetGameFromDB(locale)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, games)
}

func (node Node) LaunchGame(c *gin.Context) {
	var launchGameRequest request.LaunchGameRequest
	err := c.BindJSON(&launchGameRequest)
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request JSON -",
	launchGameRequest)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}

	games, err := model.GetGameByIDFromDB(launchGameRequest.GameID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}


	if (len(games) == 0 ) {
		err := errors.New("Bad Game id")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}

	players, err := model.GetPlayerByIDFromDB(launchGameRequest.PlayerID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}


	if (len(players) == 0 ) {
		err := errors.New("Bad Player id")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}

	play, err := model.CreatePlayInDB(model.Play{ PlayerID: players[0].ID, GameID: games[0].ID, ReturnURL: launchGameRequest.ReturnURL})
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.LaunchGameErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusInternalServerError, response.LaunchGameResponse{Token: play.TokenString, 
		PlayURL: common.GetConfiger().Configs.GameNodeAddress+"?game_id=" + fmt.Sprintf("%x", string(play.GameID)) + 
		"&player_id=" + fmt.Sprintf("%x", string(play.PlayerID)) +
		"&token="+ play.TokenString +
		"&return_url=" + launchGameRequest.ReturnURL })
}

func (node Node) Deposit(c *gin.Context) {
	var depositRequest request.DepositRequest
	err := c.BindJSON(&depositRequest)
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request JSON -",
	depositRequest)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.DepositErrorResponse{})
		return
	}
	
	players, err := model.GetPlayerByIDFromDB(depositRequest.PlayerID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.DepositErrorResponse{})
		return
	}

	if (len(players) == 0 ) {
		err := errors.New("Bad Player id")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.DepositErrorResponse{})
		return
	}

	wallets, err := model.GetWalletByIDFromDB(players[0].WalletID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.DepositErrorResponse{})
		return
	}

	if (len(wallets) == 0 ) {
		err := errors.New("No Wallet Found")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.DepositErrorResponse{})
		return
	}

	if depositRequest.Money < 0 {
		err := errors.New("Deposit Money cannot be negative")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.DepositErrorResponse{})
		return 
	}

	remainMoney, err := model.UpdateWalletMoneyInDB(wallets[0].ID, depositRequest.Money)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.DepositErrorResponse{})
		return
	}
	
	trans, err := model.CreateTransactionInDB(model.Transaction{MoneyExchange: depositRequest.Money,
		MoneyRemain: remainMoney, PlayerID: players[0].ID})

	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}	

	c.JSON(http.StatusInternalServerError, response.DepositResponse{PlayerID: depositRequest.PlayerID, TransactionID: 
		fmt.Sprintf("%x", string(trans.ID))})
		
	return 
}

func (node Node) GetBalance(c *gin.Context) {
	var balanceRequest request.BalanceRequest
	err := c.BindJSON(&balanceRequest)
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request JSON -",
	&balanceRequest)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.GetBalanceErrorResponse{})
		return
	}
	
	players, err := model.GetPlayerByIDFromDB(balanceRequest.PlayerID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.GetBalanceErrorResponse{})
		return
	}

	if (len(players) == 0 ) {
		err := errors.New("Bad Player id")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.GetBalanceErrorResponse{})
		return
	}

	wallets, err := model.GetWalletByIDFromDB(players[0].WalletID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.GetBalanceErrorResponse{})
		return
	}

	if (len(wallets) == 0 ) {
		err := errors.New("No Wallet Found")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.DepositErrorResponse{})
		return
	}

	c.JSON(http.StatusInternalServerError, response.GetBalanceResponse{Amount: wallets[0].Money})
	return 
}

func (node Node) WithDraw(c *gin.Context) {
	var withDrawRequest request.WithDrawRequest
	err := c.BindJSON(&withDrawRequest)
	common.GetLogger().Log(e.TRACK, "Request Host - ", common.GetCurrentIP(*c.Request), "|", "Request JSON -",
	withDrawRequest)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}
	
	players, err := model.GetPlayerByIDFromDB(withDrawRequest.PlayerID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}

	if (len(players) == 0 ) {
		err := errors.New("Bad Player id")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.WithDrawErrorResponse{})
		return
	}

	wallets, err := model.GetWalletByIDFromDB(players[0].WalletID)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}

	if (len(wallets) == 0 ) {
		err := errors.New("No Wallet Found")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.WithDrawErrorResponse{})
		return
	}

	
	if withDrawRequest.Money < 0 {
		err := errors.New("WithDraw Money cannot be negative")
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
		err.Error())
		c.JSON(http.StatusBadRequest, response.WithDrawErrorResponse{})
		return 
	}

	remainMoney, err := model.UpdateWalletMoneyInDB(wallets[0].ID, -withDrawRequest.Money)
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}	
	
	trans, err := model.CreateTransactionInDB(model.Transaction{MoneyExchange: -withDrawRequest.Money,
		MoneyRemain: remainMoney, PlayerID: players[0].ID})
	if err != nil {
		common.GetLogger().Log(e.ERROR, "Request Host -", common.GetCurrentIP(*c.Request), "|", "Error -",
			err.Error())
		c.JSON(http.StatusInternalServerError, response.WithDrawErrorResponse{})
		return
	}	

	c.JSON(http.StatusInternalServerError, response.DepositResponse{PlayerID: withDrawRequest.PlayerID, TransactionID: 
		fmt.Sprintf("%x", string(trans.ID))})
}



