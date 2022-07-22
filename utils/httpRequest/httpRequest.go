package httprequest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Chadiii/flagship/models"
	"github.com/Chadiii/flagship/utils"
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
	projectsModel := models.ProjectResponse{}

	err = json.Unmarshal(body, &projectsModel)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s \n", body)
	//fmt.Println(projectsModel.Items)
}

func HttpGetProject(id string) {

	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects/"+id, nil)
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

func HttpEditProject(id, name string) {
	projectRequest := models.ProjectRequest{
		Name: name,
	}
	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("PATCH", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects/"+id, bytes.NewBuffer(projectRequestJSON))
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

func HttpCreateToken(client_id, client_secret, grant_type, scope, expiration string) (string, error) {

	var authenticationResponse models.AuthenticationResponse

	authRequest := models.AuthenticationRequest{
		Client_id:     client_id,
		Client_secret: client_secret,
		Scope:         scope,
		Grant_type:    grant_type,
	}
	authRequestJSON, err := json.Marshal(authRequest)
	if err != nil {
		//fmt.Printf("%s", err)
		return "", err
	}

	c := http.Client{Timeout: time.Duration(100) * time.Second}
	req, err := http.NewRequest("POST", utils.HostAuth+"/"+viper.GetViper().GetString("account_id")+"/token?expires_in="+expiration, bytes.NewBuffer(authRequestJSON))

	if err != nil {
		//fmt.Printf("error %s", err)
		return "", err
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/json`)
	resp, err := c.Do(req)
	if err != nil {
		//fmt.Printf("error %s", err)
		return "", err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&authenticationResponse)

	return authenticationResponse.Access_token, err
}

func HttpCheckToken(token string) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", utils.HostAuth+"/token?access_token="+token, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `*/*`)
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s \n", body)
}

func HttpToggleProject(id, state string) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}

	projectRequest := models.ProjectToggleRequest{
		State: state,
	}

	projectRequestJSON, err := json.Marshal(projectRequest)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	req, err := http.NewRequest("PATCH", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/projects/"+id+"/toggle", bytes.NewBuffer(projectRequestJSON))
	if err != nil {
		fmt.Printf("error %s", err)
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	defer resp.Body.Close()
	fmt.Println("status: " + resp.Status)
}

func HttpToggleCampaign(id, state string) {
	c := http.Client{Timeout: time.Duration(100) * time.Second}

	campaignToggleRequest := models.CampaignToggleRequest{
		State: state,
	}

	campaignToggleRequestJSON, err := json.Marshal(campaignToggleRequest)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	req, err := http.NewRequest("PATCH", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id+"/toggle", bytes.NewBuffer(campaignToggleRequestJSON))
	if err != nil {
		fmt.Printf("error %s", err)
	}
	req.Header.Add("Accept", `*/*`)
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	req.Header.Add("Accept-Encoding", `gzip, deflate, br`)

	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		fmt.Println("Campaign is " + state + " successfully.")
	}

}

func HttpListCampaign() {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
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

func HttpCreateCampaign(data string) {

	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("POST", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns", bytes.NewBuffer([]byte(data)))
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
	fmt.Printf("\n%s \n", body)
}

func HttpGetCampaign(id string) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("GET", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
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

func HttpDeleteCampaign(id string) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("DELETE", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+viper.GetViper().GetString("token"))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 204 {
		fmt.Println("Campaign deleted successfully.")
	}
}

func HttpEditCampaign(id, data string) {

	c := http.Client{Timeout: time.Duration(10) * time.Second}
	req, err := http.NewRequest("PATCH", utils.Host+"/v1/accounts/"+viper.GetViper().GetString("account_id")+"/account_environments/"+viper.GetViper().GetString("account_environment_id")+"/campaigns/"+id, bytes.NewBuffer([]byte(data)))

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
	fmt.Printf("\n%s \n", body)
}
