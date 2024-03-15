package models

// PageData represents the data structure for the HTML template
type KVKey struct {
	Secret      string `json:"secret" redis:"secret"`
	Metadata    string `json:"metadata" redis:"metadata"`
	Keyvault    string `json:"keyvault" redis:"keyvault"`
	Expireddate string `json:"expireddate" redis:"expireddate"`
}

type PageInfo struct {
	PageNumber int
	Current    bool
}

type DynamicData struct {
	WarningMessage uint
	KVKey
}

type PageData struct {
	PageDataArray []DynamicData
	Pagination    []PageInfo
	SearchQuery   string
}
