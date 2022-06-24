package model

type ServerLog struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	CreatedAt int64  `json:"createdAt"`
}
