package objects

type InstallInfo struct {
	Winget string `json:"winget"`
	Choco  string `json:"choco"`
}

type Applications map[string]InstallInfo
