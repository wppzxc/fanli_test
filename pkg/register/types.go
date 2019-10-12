package register

type ActivationCode struct {
	ExpireTimestamp int64 `json:"expireTimestamp"`
	CertificateInfo string `json:"certificateInfo"`
	MainVersion string `json:"mainVersion"`
}
