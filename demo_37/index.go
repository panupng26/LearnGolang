package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Csummarys struct {
	Province 	Provinces 
	Agegroup 	AgeGroups 	
}

type Data struct {
	Data []User
}

type User struct {
	ConfirmDate		string   `json:"ConfirmDate"`
	No 				string 	 `json:"No"`
	Age  			int 	 `json:"Age"`
	Gender       	string   `json:"Gender"`
	GenderEn     	string	 `json:"GenderEn"`
	Nation			string	 `json:"Nation"`
	NationEn		string	 `json:"NationEn"`
	Province		string   `json:"Province"`
	ProvinceId		int		 `json:"ProvinceId"`
	District		string 	 `json:"District"`
	ProvinceEn		string 	 `json:"ProvinceEn"`
	StatQuarantine	int 	 `json:"StatQuarantine"`
}

type Provinces struct {
	Samutsakhon int 	`json:"Samut Sakhon"`
	Bangkok		int 	`json:"Bangkok"`
}

type AgeGroups struct {
	Lenone 		int `json:"0-30"`
	Lentwo 		int `json:"31-60"`
	Lenthree	int `json:"61+"`
	Lenfour 	int `json:"N/A"`
}
var client *http.Client
var user Data
var respone Csummarys

func getUser() {
	url := "http://static.wongnai.com/devinterview/covid-cases.json"
	err := getJson(url, &user)

	if err != nil {
		fmt.Printf("error getting can't got it: %s\n", err.Error())
		return
	} else {
		len1,len2,len3,len4,lenpro1,lenpro2 := 0,0,0,0,0,0
		
		for i, x := range user.Data {
			fmt.Println(i)
			if(x.Age >= 0 && x.Age <= 30) {
				len1++;
			} else if(x.Age >= 31 && x.Age <= 60){
				len2++;
			} else if(x.Age > 60) {
				len3++;
			}else {
				len4++;
			}
			if(x.Province == "Samut Sakhon") {
				lenpro1++;
			}else if(x.Province == "Bangkok") {
				lenpro2++;
			}
		}
		respone.Province.Samutsakhon = lenpro1
		respone.Province.Bangkok = lenpro2
		respone.Agegroup.Lenone = len1
		respone.Agegroup.Lentwo = len2
		respone.Agegroup.Lenthree = len3
		respone.Agegroup.Lenfour = len4
	}
}

func getJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	
	client = &http.Client{Timeout: 10 * time.Second}
	getUser()

	r := gin.Default()
	r.GET("/covid/summary", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Data" : respone,
			"success": http.StatusOK,
			"message": "Success",
		})
	})
	r.Run()
}