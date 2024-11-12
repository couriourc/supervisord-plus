package updater

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"github.com/sanbornm/go-selfupdate/selfupdate"
	"sync"
)

type SelfUpdaterAgent struct {
	Updater   *selfupdate.Updater
	AgentName string
}

var lock = &sync.RWMutex{}

// go-selfupdate setup and config
var updaters = make([]*SelfUpdaterAgent, 0)
var signal = make(chan rxgo.Item)

func TriggerUpdater(agentName string) {
	signal <- rxgo.Of(agentName)
}
func (agent *SelfUpdaterAgent) GetAgentName() string {
	return agent.AgentName
}
func GetAllUpdaterAgent() []*UpdateAgentInfo {
	lock.RLock()
	agents := make([]*UpdateAgentInfo, 0)
	for _, agent := range updaters {
		agents = append(agents,
			&UpdateAgentInfo{
				AgentName:      agent.GetAgentName(),
				CurrentVersion: agent.Updater.CurrentVersion,
				ApiURL:         agent.Updater.ApiURL,
				BinURL:         agent.Updater.BinURL,
				DiffURL:        agent.Updater.DiffURL,
				Dir:            agent.Updater.Dir,
				CmdName:        agent.Updater.CmdName,
				ForceCheck:     agent.Updater.ForceCheck,
			})
	}
	lock.RUnlock()
	return agents
}
func SetupSelfUpdater(agent *SelfUpdaterAgent) {
	go func() {
		observable := rxgo.FromChannel(signal)
		for item := range observable.Observe() {
			switch item.V.(type) {
			case string:
				if item.V != agent.AgentName {
					break
				}
				go func() {
					err := agent.Updater.BackgroundRun()
					if err != nil {
						fmt.Println(err)
					}
				}()
			}
		}
	}()
}
func NewUpdater(cfg map[string]string) {
	lock.Lock()
	update := &SelfUpdaterAgent{
		Updater: &selfupdate.Updater{
			CurrentVersion: cfg["version"],               // Manually update the const, or set it using `go build -ldflags="-X main.VERSION=<newver>" -o hello-updater src/hello-updater/main.go`
			ApiURL:         cfg["api_url"],               // The server hosting `$CmdName/$GOOS-$ARCH.json` which contains the checksum for the binary
			BinURL:         cfg["bin_url"],               // The server hosting the zip file containing the binary application which is a fallback for the patch method
			DiffURL:        cfg["diff_url"],              // The server hosting the binary patch diff for incremental updates
			Dir:            cfg["dir"],                   // The directory created by the app when run which stores the cktime file
			CmdName:        cfg["cmd_name"],              // The app name which is appended to the ApiURL to look for an update
			ForceCheck:     cfg["force_check"] == "true", // For this example, always check for an update unless the version is "dev"
		},
		AgentName: cfg["agent_name"],
	}
	updaters = append(updaters, update)
	go SetupSelfUpdater(update)
	lock.Unlock()
}
