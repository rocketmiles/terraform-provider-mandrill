package mandrill

type MandrillRequest struct {
	Key string `json:"key"`
}

type SendersCheckDomainRequest struct {
	MandrillRequest
	Domain string `json:"domain"`
}

type SendersCheckDomainResponse struct {
	Domain       string `json:"domain"`
	CreatedAt    string `json:"created_at"`
	LastTestedAt string `json:"last_tested_at"`
	Spf          struct {
		Valid      bool   `json:"valid"`
		ValidAfter string `json:"valid_after"`
		Error      string `json:"error"`
	} `json:"spf"`
	Dkim struct {
		Valid      bool   `json:"valid"`
		ValidAfter string `json:"valid_after"`
		Error      string `json:"error"`
	} `json:"dkim"`
	VerifiedAt   string `json:"verified_at"`
	ValidSigning bool   `json:"valid_signing"`
	VerifyTxtKey string `json:"verify_txt_key"`
}
