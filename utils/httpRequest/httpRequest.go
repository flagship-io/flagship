package httprequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Chadiii/flagship-mock/models"
	"github.com/Chadiii/flagship-mock/utils"
	"github.com/spf13/viper"
)

func HttpCreateProject(name string) {
	projectRequest := models.ProjectRequest{
		Name: name,
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("POST", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects", bytes.NewBuffer(projectRequestJSON))
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s \n", body)
}

func HttpDeleteProject(id string) {

	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("DELETE", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects/"+id, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}

func HttpListProject() {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s \n", body)
}

func HttpToken(client_id, client_secret, scope, grant_type string) {

	var authenticationResponse models.AuthenticationResponse

	authRequest := models.AuthenticationRequest{
		Client_id:     client_id,
		Client_secret: client_secret,
		Scope:         scope,
		Grant_type:    grant_type,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	c := http.Client{Timeout: time.Duration(100) * time.Second}
	req, err := http.NewRequest("POST", utils.HostAuth+"/"+viper.GetViper().GetString("account_id")+"/token?expires_in=0", bytes.NewBuffer(authRequestJSON))
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&authenticationResponse)

	fmt.Println("token: " + authenticationResponse.Access_token)

	if authenticationResponse.Access_token == "" {
		fmt.Println("required valid client_id and client_secret")
		return
	}

	viper.Set("token", authenticationResponse.Access_token)
	dir_err := viper.WriteConfigAs("config.yaml")
	if dir_err != nil {
		fmt.Println(dir_err)
	}
}
