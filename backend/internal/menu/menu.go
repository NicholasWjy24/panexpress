package menu

type Menu struct {
	ID           int    `json:"id"`
	MenuName     string `json:"menu_name"`
	MinRoleLevel int    `json:"min_role_level"`
	MaxRoleLevel int    `json:"max_role_level"`
	Route        string `json:"route"`
}
