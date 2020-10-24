package hellosign

const (
	// WhiteLabelingOptionsPageBackgroundColor is label for page background color
	WhiteLabelingOptionsPageBackgroundColor = "page_background_color"
	// WhiteLabelingOptionsHeaderBackgroundColor is label for header background color
	WhiteLabelingOptionsHeaderBackgroundColor = "header_background_color"
	// WhiteLabelingOptionsTextColor1 is label for options text color1
	WhiteLabelingOptionsTextColor1 = "text_color1"
	// WhiteLabelingOptionsTextColor2 is label for options text color2
	WhiteLabelingOptionsTextColor2 = "text_color2"
	// WhiteLabelingOptionsLinkColor is label for link color
	WhiteLabelingOptionsLinkColor = "link_color"
	// WhiteLabelingOptionsPrimaryButtonColor is label for primary button color
	WhiteLabelingOptionsPrimaryButtonColor = "primary_button_color"
	// WhiteLabelingOptionsPrimaryButtonTextColor is label for primary button text color
	WhiteLabelingOptionsPrimaryButtonTextColor = "primary_button_text_color"
	// WhiteLabelingOptionsPrimaryButtonColorHover is label for primary button color hover
	WhiteLabelingOptionsPrimaryButtonColorHover = "primary_button_color_hover"
	// WhiteLabelingOptionsPrimaryButtonTextColorHover is label for primary button text color hover
	WhiteLabelingOptionsPrimaryButtonTextColorHover = "primary_button_text_color_hover"
	// WhiteLabelingOptionsSecondaryButtonColor is label for secondary button color
	WhiteLabelingOptionsSecondaryButtonColor = "secondary_button_color"
	// WhiteLabelingOptionsSecondaryButtonTextColor is label for secondary button text color
	WhiteLabelingOptionsSecondaryButtonTextColor = "secondary_button_text_color"
	// WhiteLabelingOptionsSecondaryButtonColorHover is label for secondary button color hover
	WhiteLabelingOptionsSecondaryButtonColorHover = "secondary_button_color_hover"
	// WhiteLabelingOptionsSecondaryButtonTextColorHover is label for secondary button text color hover
	WhiteLabelingOptionsSecondaryButtonTextColorHover = "secondary_button_text_color_hover"
)

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
