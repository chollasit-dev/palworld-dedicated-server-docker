package palworld

type Player struct {
	AccountName   string  `json:"accountName,omitzero"`
	BuildingCount int     `json:"building_count,omitzero"`
	IP            string  `json:"ip,omitzero"`
	Level         int     `json:"level,omitzero"`
	LocationX     float64 `json:"location_x,omitzero"`
	LocationY     float64 `json:"location_y,omitzero"`
	Name          string  `json:"name,omitzero"`
	Ping          float64 `json:"ping,omitzero"`
	PlayerID      string  `json:"playerId,omitzero"`
	UserID        string  `json:"userId,omitzero"`
}
