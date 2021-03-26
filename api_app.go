package hellosign

// APIAppList represent list of api apps response
type APIAppList struct {
	APIApps  []APIApp `json:"api_apps"`
	ListInfo ListInfo `json:"list_info"`
}

// APIApp represent api app response
type APIApp struct {
	APIApp APIAppDetail `json:"api_app"`
}

// APIAppDetail represent api app detail
type APIAppDetail struct {
	ClientID             string             `json:"client_id"`
	CreatedAt            int64              `json:"created_at"`
	Name                 string             `json:"name"`
	Domain               string             `json:"domain"`
	CallbackURL          string             `json:"callback_url"`
	IsApproved           bool               `json:"is_approved"`
	OwnerAccount         OwnerAccountDetail `json:"owner_account"`
	Options              OptionsDetail      `json:"options"`
	Oauth                OauthDetail        `json:"oauth_detail"`
	WhiteLabelingOptions map[string]string  `json:"white_labeling_options"`
}

// OwnerAccountDetail represent owner account detail
type OwnerAccountDetail struct {
	AccountID    string `json:"account_id"`
	EmailAddress string `json:"email_address"`
}

// OptionsDetail represent options detail
type OptionsDetail struct {
	CanInsertEverywhere bool `json:"can_insert_everywhere"`
}

// OauthDetail represent oatuh detail
type OauthDetail struct {
	CallbackURL  string   `json:"callback_url"`
	Secret       string   `json:"secret"`
	Scopes       []string `json:"scopes"`
	ChargesUsers bool     `json:"charges_users"`
}
