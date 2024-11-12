package updater

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// PostTriggerUpdater trigger updater for special agent
func PostTriggerUpdater(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	TriggerUpdater(params["agent_name"])
	w.Write([]byte("Success to Trigger updater"))
}

// GetAllUpdaterInfo get all updater info
func GetAllUpdaterInfo(w http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(w).Encode(GetAllUpdaterAgent())
	if err != nil {
		return
	}

}
