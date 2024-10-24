// Code generated by goctl. DO NOT EDIT.
package types

import "github.com/tickstep/aliyunpan-api/aliyunpan"

type GenerateQrcodeReq struct {

}

type GenerateQrcodeResp struct {
	Qrcode string `json:"qrCode"`
	T      int64  `json:"t"`
	Ck     string `json:"ck"`
}

type QrcodeStateReq struct {
	T  int64  `json:"t"`
	Ck string `json:"ck"`
}

type QrcodeStateResp struct {
	UserName     string `json:"userName"`
	UserID       string `json:"userId"`
	TokenType    string `json:"tokenType"`
	RefreshToken string `json:"refreshToken"`
}

type CreatInstanceReq struct {
	RefreshToken string `json:"refreshToken"`
}

type CreatInstanceResp  = aliyunpan.UserInfo
