package models

import "time"

//Country provides entities for an Country structure
type Country struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name" sql:",notnull,unique"`
	Code      string    `json:"code" sql:",notnull,unique"`
	DialCode  string    `json:"dialCode" sql:",notnull,unique"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"-" sql:",default:now()"`
	UpdatedAt time.Time `json:"-" sql:",default:now()"`
}

// type CountryMinified struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name" sql:",notnull,unique"`
// }
