package main

import (
	"encoding/json"
	"github.com/couriourc/supervisord-plus/config"
	"github.com/couriourc/supervisord-plus/process"
	"github.com/couriourc/supervisord-plus/types"
	"github.com/couriourc/supervisord-plus/updater"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// SupervisorRestful the restful interface to control the programs defined in configuration file
type SupervisorRestful struct {
	router     *mux.Router
	supervisor *Supervisor
}

// NewSupervisorRestful create a new SupervisorRestful object
func NewSupervisorRestful(supervisor *Supervisor) *SupervisorRestful {
	return &SupervisorRestful{router: mux.NewRouter(), supervisor: supervisor}
}

// CreateProgramHandler create http handler to process program related restful request
func (sr *SupervisorRestful) CreateProgramHandler() http.Handler {
	sr.router.HandleFunc("/program/list", sr.ListProgram).Methods("GET")
	sr.router.HandleFunc("/program/start/{name}", sr.StartProgram).Methods("POST", "PUT")
	sr.router.HandleFunc("/program/stop/{name}", sr.StopProgram).Methods("POST", "PUT")
	sr.router.HandleFunc("/program/log/{name}/stdout", sr.ReadStdoutLog).Methods("GET")
	sr.router.HandleFunc("/program/startPrograms", sr.StartPrograms).Methods("POST", "PUT")
	sr.router.HandleFunc("/program/stopPrograms", sr.StopPrograms).Methods("POST", "PUT")
	return sr.router
}

func (sr *SupervisorRestful) CreateUpdaterHandler() http.Handler {
	sr.router.HandleFunc("/updater/{agent_name}", updater.PostTriggerUpdater).Methods("PUT", "POST")
	sr.router.HandleFunc("/updater/list", updater.GetAllUpdaterInfo).Methods("GET")
	return sr.router
}

// CreateSupervisorHandler create http rest interface to control supervisor itself
func (sr *SupervisorRestful) CreateSupervisorHandler() http.Handler {
	sr.router.HandleFunc("/supervisor/config", sr.GetConfig).Methods("GET")
	sr.router.HandleFunc("/supervisor/config", sr.PostConfig).Methods("POST")
	sr.router.HandleFunc("/supervisor/summary", sr.GetSummary).Methods("GET")
	sr.router.HandleFunc("/supervisor/shutdown", sr.Shutdown).Methods("PUT", "POST")
	sr.router.HandleFunc("/supervisor/reload", sr.Reload).Methods("PUT", "POST")
	return sr.router
}
func (sr *SupervisorRestful) PostConfig(w http.ResponseWriter, req *http.Request) {

	tmpConfig := config.NewConfig("")
	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	if tmpConfig.CheckCanBeParse(body) {
		err := sr.supervisor.config.OverwriteConfigFile(body)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(struct {
				Error error `json:"error"`
			}{
				Error: err,
			})
			return
		}
		sr.GetConfig(w, req)
	} else {
		w.WriteHeader(400)

		json.NewEncoder(w).Encode(struct {
			Error string `json:"error"`
		}{
			Error: "invalid config",
		})
	}
}
func (sr *SupervisorRestful) GetConfig(w http.ResponseWriter, req *http.Request) {

	filename, content, err := sr.supervisor.config.GetConfigFileContent()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("BadGateWay"))
		return
	}
	result := types.SupervisordConfigInfo{
		Filename: filename,
		Content:  string(content),
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
	return
}
func (sr *SupervisorRestful) GetSummary(w http.ResponseWriter, req *http.Request) {
	result := types.SupervisordSummaryInfo{}

	sr.supervisor.procMgr.ForEachProcess(func(p *process.Process) {
		result.ProcessStatus = append(result.ProcessStatus, getProcessInfo(p).State)
	})

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(result)
}

// ListProgram list the status of all the programs
//
// json array to present the status of all programs
func (sr *SupervisorRestful) ListProgram(w http.ResponseWriter, req *http.Request) {
	result := struct{ AllProcessInfo []types.ProcessInfo }{make([]types.ProcessInfo, 0)}
	if sr.supervisor.GetAllProcessInfo(nil, nil, &result) == nil {
		json.NewEncoder(w).Encode(result.AllProcessInfo)
	} else {
		r := map[string]bool{"success": false}
		json.NewEncoder(w).Encode(r)
	}
}

type StartProgramBody struct {
	Args string `json:"args"`
}

// StartProgram start the given program through restful interface
func (sr *SupervisorRestful) StartProgram(w http.ResponseWriter, req *http.Request) {

	defer req.Body.Close()
	var body StartProgramBody

	params := mux.Vars(req)
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		body = StartProgramBody{Args: ""}
	}
	//log.Printf(body.Args)

	success, err := sr._startProgram(params["name"], body.Args)
	r := map[string]bool{"success": err == nil && success}
	json.NewEncoder(w).Encode(&r)
}

func (sr *SupervisorRestful) _startProgram(program string, args string) (bool, error) {
	startArgs := StartProcessArgs{Name: program, Wait: true, Args: args}
	result := struct{ Success bool }{false}
	err := sr.supervisor.StartProcess(nil, &startArgs, &result)
	return result.Success, err
}

// StartPrograms start one or more programs through restful interface
func (sr *SupervisorRestful) StartPrograms(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var b []byte
	var err error

	if b, err = ioutil.ReadAll(req.Body); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("not a valid request"))
		return
	}

	var programs []string
	if err = json.Unmarshal(b, &programs); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("not a valid request"))
	} else {
		for _, program := range programs {
			sr._startProgram(program, "")
		}
		w.Write([]byte("Success to start the programs"))
	}
}

// StopProgram stop a program through the restful interface
func (sr *SupervisorRestful) StopProgram(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	params := mux.Vars(req)
	success, err := sr._stopProgram(params["name"])
	r := map[string]bool{"success": err == nil && success}
	json.NewEncoder(w).Encode(&r)
}

func (sr *SupervisorRestful) _stopProgram(programName string) (bool, error) {
	stopArgs := StartProcessArgs{Name: programName, Wait: true}
	result := struct{ Success bool }{false}
	err := sr.supervisor.StopProcess(nil, &stopArgs, &result)
	return result.Success, err
}

// StopPrograms stop programs through the restful interface
func (sr *SupervisorRestful) StopPrograms(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var programs []string
	var b []byte
	var err error
	if b, err = ioutil.ReadAll(req.Body); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("not a valid request"))
		return
	}

	if err := json.Unmarshal(b, &programs); err != nil {
		w.WriteHeader(400)
		w.Write([]byte("not a valid request"))
	} else {
		for _, program := range programs {
			sr._stopProgram(program)
		}
		w.Write([]byte("Success to stop the programs"))
	}

}

// TriggerUpdater trigger updater for agent_name
func (sr *SupervisorRestful) TriggerUpdater(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	updater.TriggerUpdater(params["agent_name"])
	w.Write([]byte("Success to Trigger updater"))
}

// ReadStdoutLog read the stdout of given program
func (sr *SupervisorRestful) ReadStdoutLog(w http.ResponseWriter, req *http.Request) {
}

// Shutdown shutdown the supervisor itself
func (sr *SupervisorRestful) Shutdown(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	reply := struct{ Ret bool }{false}
	sr.supervisor.Shutdown(nil, nil, &reply)
	w.Write([]byte("Shutdown..."))
}

// Reload reload the supervisor configuration file through rest interface
func (sr *SupervisorRestful) Reload(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	reply := struct{ Ret bool }{false}
	sr.supervisor.Reload()
	r := map[string]bool{"success": reply.Ret}
	json.NewEncoder(w).Encode(&r)
}
