package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type DockerPayload struct {
	Callback_Url string
	Repository   struct {
		Comment_Count string
		Description   string
		Is_Trusted    bool
		Is_Private    bool
		Status        string
		Repo_Url      string
		Repo_Name     string
	}
}

func dockerPull(r string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	cmd := string(dir) + "/dockerPull.sh"
	if err := exec.Command(cmd, r).Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Docker ran")
}

func handler(w http.ResponseWriter, req *http.Request) {
	var p DockerPayload
	if req.Body == nil {
		http.Error(w, "Please provide a request body", 400)
	}
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(p.Callback_Url)
	dockerPull(p.Repository.Repo_Name)
}

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":3333", nil)
}
