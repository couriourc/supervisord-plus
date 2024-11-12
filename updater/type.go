package updater

type UpdateAgentInfo struct {
	CurrentVersion string `json:"version"`
	ApiURL         string `json:"api_url"`
	BinURL         string `json:"bin_url"`
	DiffURL        string `json:"diff_url"`
	Dir            string `json:"dir"`
	CmdName        string `json:"cmd_name"`
	ForceCheck     bool   `json:"force_check"`
	AgentName      string `json:"agent_name"`
}
