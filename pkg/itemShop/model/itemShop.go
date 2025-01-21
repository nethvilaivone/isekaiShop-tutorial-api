package model

import "time"

//model have to be a json
type (
	Item struct {
		ID          uint64 `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Picture     string `json:"picturs"`
		Price       uint   `json:"price"`
	}



	PlayerCoin struct {
		ID        uint64    `json:"id"`
		PlayerID  string    `json:"playnameid"`
		Amount    int64      `json:"amount"`
		CreatedAt time.Time   `json:"create_at"`
	}



	ItemFillter struct {
		Name string   `query:"name" validate:"omitempty, max=64"`
		Description string  `query:"description" validate:"omitempty, max= 126"`
	}
)






