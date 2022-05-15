package domain

type User struct {
	Login    string  `json:"login"`
	Password *string `json:"password,omitempty"`
}

type RegistryResponse struct {
	Token string `json:"token"`
}
