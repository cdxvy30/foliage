package domain

import "time"

type StockData struct {
	MsgArray []struct {
		TV              string `json:"tv"`
		PS              string `json:"ps"`
		Price           string `json:"pz"` // 當前成交價
		ZPrice          string `json:"z"`
		Name            string `json:"n"`     // 股票名稱
		Code            string `json:"c"`     // 股票代號
		DataUpdatedTime string `json:"tlong"` // 資料更新時間
		DocUpdatedTime  string
		Time            time.Time
	} `json:"msgArray"`
	// Referer   string `json:"referer"`
	// UserDelay int    `json:"userDelay"`
	// RtCode    string `json:"rtcode"`
	// QueryTime struct {
	// 	SysDate           string `json:"sysDate"`
	// 	StockInfoItem     int    `json:"stockInfoItem"`
	// 	StockInfo         int    `json:"stockInfo"`
	// 	SessionStr        string `json:"sessionStr"`
	// 	SysTime           string `json:"sysTime"`
	// 	ShowChart         bool   `json:"showChart"`
	// 	SessionFromTime   int    `json:"sessionFromTime"`
	// 	SessionLatestTime int    `json:"sessionLatestTime"`
	// } `json:"queryTime"`
	// RtMessage string `json:"rtmessage"`
}
