package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "os"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/gorilla/mux"
)

/*func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "sM4vGyIsdFwBJk50pu_XfzN5MW5UEh2AmNbO10eP3DI")
	return nil
}

func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}*/
type Image struct {
	id       string
	url      string
	selected bool
	pageNum  int
}

func main() {

	img0 := Image{"Qrjx2nTBHVo", "https://images.unsplash.com/photo-1593642532009-6ba71e22f468?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img1 := Image{"k9pQzz9UqKA", "https://images.unsplash.com/photo-1601642272829-92d46637a634?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img2 := Image{"JjSoCnYrpno", "https://images.unsplash.com/photo-1601647955514-79cbf7c88155?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img3 := Image{"5fgoGMbIjJo", "https://images.unsplash.com/photo-1601637825401-fd5347aa8570?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img4 := Image{"ZLjTS3ZHUZk", "https://images.unsplash.com/photo-1601637402905-22eb85fe3db5?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img5 := Image{"8YIYTuVykzw", "https://images.unsplash.com/photo-1601635401962-4c35344d238f?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img6 := Image{"3gnzTH68SeA", "https://images.unsplash.com/photo-1601635624324-ee44c9f86d32?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img7 := Image{"xoz4Wk5wnoA", "https://images.unsplash.com/photo-1601634755085-28ee563b7b3d?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}
	img8 := Image{"K81NdfnAdQI", "https://images.unsplash.com/photo-1601634499835-50158f60f9dc?ixlib=rb-1.2.1&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=400&fit=max&ixid=eyJhcHBfaWQiOjE3MTAzNH0", false, 11}

	images := [9]Image{img0, img1, img2, img3, img4, img5, img6, img7, img8}

	db, err := sql.Open("mysql", "nick:dog@tcp(127.0.0.1:3306)/animals")

	if err != nil {
		log.Fatal(err)
	}
	for _, elem := range images {
		sbool := ""

		if elem.selected == true {
			sbool = "true"
		} else {
			sbool = "false"
		}
		s := fmt.Sprintf("INSERT INTO images (id, url, selected, page_num) VALUES ( '%s', '%s', '%s', %d );", elem.id, elem.url, sbool, elem.pageNum)
		insert, err := db.Query(s)

		if err != nil {
			log.Fatal(err)
		}
		insert.Close()
	}

	/*client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", "https://api.unsplash.com/photos", nil)
	q := req.URL.Query()
	q.Add("page", "11")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Client-ID sM4vGyIsdFwBJk50pu_XfzN5MW5UEh2AmNbO10eP3DI")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(body)

	if err != nil {
		log.Fatal(err)
	}

	//_, err = os.Stdout.Write(body)

	if err != nil {
		log.Fatal(err)
	}*/

}
