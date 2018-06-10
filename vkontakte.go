package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"log"
	"time"
)
//-37008294
//-160714694 mm
const owner_id = "-37008294"
const access_token = "13d3de0213d3de0213d3de029b13b3c292113d313d3de0249e43129fa9241a0bc1fa3e5"
const v = "5.52"



func Zapros_For_Posts(Posts string) PresponseJson{
	responseP, err := http.Get(Posts)
	if err != nil {
		log.Fatal(err)
	}
	defer responseP.Body.Close()
	contentsP, err := ioutil.ReadAll(responseP.Body)
	if err != nil {
		log.Fatal(err)
	}
	var unmarshP= PresponseJson{}
	respP := json.Unmarshal(contentsP, &unmarshP)
	if respP != nil {
		log.Fatal(respP)
	}
	return unmarshP
}
func Poluchit_DATA_iz_Post(Posts PresponseJson) []int  {
	dataPost := []int{}
	for _, rowP := range Posts.Response.Items{
		dataPost = append(dataPost, rowP.Date)
	}
	return dataPost
}
func Pechat_Posts_and_Comments(Posts PresponseJson)[]int{
	ID := []int{}
	for _, rowP := range Posts.Response.Items{
		fmt.Println("POST" ,rowP.Text)
		ID = append(ID, rowP.ID)
		Zapros_For_Comments(rowP.ID)
	}
	return ID
}
func Zapros_Two_Poluchit_DATA_Comment(Posts PresponseJson)[]int {
	dataComment := []int{}
	for _, u := range Posts.Response.Items{
		zapr2 := fmt.Sprintf("https://api.vk.com/method/wall.getComments?&owner_id=%v&post_id=%v&count=4&sort=desc&access_token=%v&v=%v", owner_id, u.ID, access_token, v)
		responseC, err := http.Get(zapr2)
		if err != nil {
			log.Fatal(err)
		}
		defer responseC.Body.Close()
		contentsC, err := ioutil.ReadAll(responseC.Body)
		if err != nil {
			log.Fatal(err)
		}
		var unmarshC = CresponseJson{}
		respC := json.Unmarshal(contentsC, &unmarshC)
		if respC != nil {
			log.Fatal(respC)
		}
		for _, j := range unmarshC.Response.Items{
			dataComment = append(dataComment, j.Date)
		}
	}
	return dataComment
}
func Zapros_For_Comments(ID int)[]int{
	zapr2 := fmt.Sprintf("https://api.vk.com/method/wall.getComments?&owner_id=%v&post_id=%v&count=4&sort=desc&access_token=%v&v=%v", owner_id, ID, access_token, v)
	responseC, err := http.Get(zapr2)
	if err != nil {
		log.Fatal(err)
	}
	defer responseC.Body.Close()
	contentsC, err := ioutil.ReadAll(responseC.Body)
	if err != nil {
		log.Fatal(err)
	}
	var unmarshC= CresponseJson{}
	respC := json.Unmarshal(contentsC, &unmarshC)
	if respC != nil {
		log.Fatal(respC)
	}
	dataComments := []int{}
	for _, t := range unmarshC.Response.Items{
		fmt.Println("COMMENT", t.Text)
		dataComments = append(dataComments, t.Date)
	}
	return dataComments
}
func Zapros_For_Posts_Proverka()PostStructJson {
	zapr3 := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%v&count=10&offset=0&access_token=%v&v=%v", owner_id, access_token, v)
	OtvetP, err := http.Get(zapr3)
	if err != nil {
		log.Fatal(err)
	}
	defer OtvetP.Body.Close()
	SoderjimoeP, err := ioutil.ReadAll(OtvetP.Body)
	if err != nil {
		log.Fatal(err)
	}
	var UnmarshP = PostStructJson{}
	respP := json.Unmarshal(SoderjimoeP, &UnmarshP)
	if respP != nil {
		log.Fatal(respP)
	}
	return UnmarshP
}
func Obrabotka(Data_Post, Data_Comment []int)  {
	sort.Ints(Data_Post)
	sort.Ints(Data_Comment)
	Post := []int{1, 1}
	for i := 1; i <= 1; i++{
		Post = append(Post, Data_Post[len(Data_Post)-1])
	}
	sort.Ints(Post)
	Comment := []int{2, 2}
	for y := 1; y <= 1; y++{
		Comment = append(Comment, Data_Comment[len(Data_Comment)-1])
	}
	sort.Ints(Comment)
	ID := []int{}
	Post_Prov := Zapros_For_Posts_Proverka()
	for _, i := range Post_Prov.Response.Items {
		ID = append(ID, i.ID)
	}
	for i := 0; i <= 5; i++{
		timer := time.NewTimer(time.Second * 10)
		<- timer.C
		print("Таймер истек!", "\n")
		sort.Ints(Comment)
		Post_Prov := Zapros_For_Posts_Proverka()
		for _, i := range Post_Prov.Response.Items{
			ID = append(ID, i.ID)
		}
		for _, p := range Post_Prov.Response.Items {
			if p.Date > Post[len(Post)-1] {
				fmt.Println("Текст поста", p.Text, "Текст поста")
				Post = append(Post, p.Date)
			} else {
				continue
			}
		}
		for i:= 0; i < len(ID); i++{
			zapros := fmt.Sprintf("https://api.vk.com/method/wall.getComments?&owner_id=%v&post_id=%v&count=4&sort=desc&access_token=%v&v=%v", owner_id,
				ID[i], access_token, v)
			OtvetC, err := http.Get(zapros)
			if err != nil {
				log.Fatal(err)
			}
			defer OtvetC.Body.Close()
			SoderjimoeC, err := ioutil.ReadAll(OtvetC.Body)
			if err != nil {
				log.Fatal(err)
			}
			var UnmarshC= CommentsStructJson{}
			respCom := json.Unmarshal(SoderjimoeC, &UnmarshC)
			if respCom != nil {
				log.Fatal(respCom)
			}
			for _, c := range UnmarshC.Response.Items{
				if c.Date > Comment[len(Comment)-1] {
					fmt.Println("Текст комментария", c.Text, "Текст комментария")
					Comment = append(Comment, c.Date)
				}else {
					continue
				}
			}
		}
	}
}
func pecat(Start string){
	Posts := Zapros_For_Posts(Start)
	Pechat_Posts_and_Comments(Posts)
	Data_Comment := Zapros_Two_Poluchit_DATA_Comment(Posts)
	Data_Post := Poluchit_DATA_iz_Post(Posts)
	Obrabotka(Data_Post, Data_Comment)
}
func main() {
	Start := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=%v&count=10&offset=0&access_token=%v&v=%v", owner_id, access_token, v)
	pecat(Start)
}





