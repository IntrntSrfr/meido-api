package entity

type Aquarium struct {
	UserID    string `db:"user_id" json:"user_id"`
	Common    int    `db:"common" json:"common"`
	Uncommon  int    `db:"uncommon" json:"uncommon"`
	Rare      int    `db:"rare" json:"rare"`
	SuperRare int    `db:"super_rare" json:"super_rare"`
	Legendary int    `db:"legendary" json:"legendary"`
}
