package object

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateFood struct {
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	Stock    bool   `json:"in_stock"`
	Quantity int64  `json:"quantity"`
	Unit     string `json:"unit"`
}

type GetFoods struct {
	UserId int64 `json:"user_id"`
}

type CreateAlert struct {
	UserId   int64  `json:"user_id"`
	Message  string `json:"message"`
	Severity int8   `json:"severity"`
}

type GetAlerts struct {
	UserId int64 `json:"user_id"`
}

type CreateMission struct {
}

type GetMissions struct {
	UserId int64 `json:"user_id"`
}

type CreateEquipement struct {
}

type GetEquipements struct {
	UserId int64 `json:"user_id"`
}
