package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

// func getCommandoutput(command string, arguments ...string) string {
// 	cmd := exec.Command(command, arguments...)

// 	var out bytes.Buffer
// 	var stderr bytes.Buffer

// 	cmd.Stdout = &out
// 	cmd.Stderr = &stderr

// 	err := cmd.Start()
// 	if err != nil {
// 		log.Fatal(fmt.Sprint(err) + ":" + stderr.String())
// 	}

// 	err = cmd.Wait()
// 	if err != nil {
// 		log.Fatal(fmt.Sprint(err) + ":" + stderr.String())
// 	}

// 	return out.String()

// }
//refactor april 11,2023

func runCmd(w http.ResponseWriter, cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, fmt.Sprintf("err running command: %v", err), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, string(output))
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//fmt.Fprint(w, getCommandoutput("/usr/local/bin/go", "version"))

	path, err := exec.LookPath("go")
	if err != nil {
		http.Error(w, "go exec not found in PATH", http.StatusInternalServerError)
		return
	}

	// cmd := exec.Command(path, "version")
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("err running command: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// fmt.Fprint(w, string(output))

	runCmd(w, exec.Command(path, "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//fmt.Fprint(w, getCommandoutput("/bin/cat", params.ByName("name")))
	runCmd(w, exec.Command("/bin/cat", params.ByName("name")))
}

func main() {

	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	log.Fatal(http.ListenAndServe(":8000", router))

}
