package auth

import (
	common "../common"
) 

func GenerateToken() string {
	randomToken := common.RandStringRunes(10) 
	return randomToken
}