package model

type Company struct {
	Code_id     int    `db:"code_id"`
	Code        string `db:"code"`
	Name        string `db:"name"`
	Code_type   int    `db:"code_type"`
	Market_type int    `db:"market_type"`
	Stop        bool   `db:"stop"`
	//Name string
}
type CompanyDetail struct {
	Company `db:"company"`

	Full_code                 string `db:"full_code"`
	Full_name_kr              string `db:"full_name_kr"`
	Full_name_eng             string `db:"full_name_eng"`
	Listing_date              string `db:"listing_date"`
	Market                    string `db:"market"`
	Securities_classification string `db:"securities_classification"`
	Department                string `db:"department"`
	Stock_type                string `db:"stock_type"`
	Face_value                string `db:"face_value"`
	Listed_stocks             string `db:"listed_stocks"`
}

type CompanyState struct {
	Company `db:"company"`

	Clear         bool `db:"clear"`
	Managed       bool `db:"managed"`
	Ventilation   bool `db:"ventilation"`
	Unfaithful    bool `db:"unfaithful"`
	Low_liquidity bool `db:"low_liquidity"`
	Lack_listed   bool `db:"lack_listed"`
	Overheated    bool `db:"overheated"`
	Caution       bool `db:"caution"`
	Warning       bool `db:"warning"`
	Risk          bool `db:"risk"`
}
