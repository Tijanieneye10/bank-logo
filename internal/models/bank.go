package models

type Bank struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Logo string `json:"logo"`
}



var ListOfBanks = []Bank{
	{
		Name: "Bank 1",
		Code: "1",
		Logo: "https://cdn.jsdelivr.net/npm/banks/gtbank",
	},
	{
		Name: "Bank 2",
		Code: "2",
		Logo: "https://cdn.jsdelivr.net/npm/banks/access",
	},
}