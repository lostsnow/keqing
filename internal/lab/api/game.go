package api

type GameRecord struct {
	GameID     int              `json:"game_id"`
	Level      int              `json:"level"`
	RegionName string           `json:"region_name"`
	Region     string           `json:"region"`
	GameRoleID string           `json:"game_role_id"`
	NickName   string           `json:"nickname"`
	Data       []GameRecordData `json:"data"`
}

type GameRecordData struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}
