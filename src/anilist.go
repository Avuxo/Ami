package main

/*
This is a wrapper around the AniList GraphQL API.
https://anilist.github.io/ApiV2-GraphQL-Docs/query.doc.html
For primary use in the Ami anime client: https://github.com/Avuxo/Ami/
*/


import (
	"context"
	"fmt"
	"log"
	"github.com/shurcooL/graphql"
)


/*
AnimeInfo
After a GQL query of AniList, this is the query data returned
for any anime.
*/
type AnimeInfo struct{
	ID       int64
	Episodes int32
	Title    string
	IsAdult  bool
	Status   string
	Genre    string
}

/*
MangaInfo
After a GQL query of AniList, this is the query data returned
for any manga.
*/
type MangaInfo struct{
	ID       int64
	Chapters int32
	IsAdult  bool
	Status   string
	Genre    string
}

/*
UserInfo
Query a user's basic info.
ID:   the user ID
Name: The user name
Bio:  The `about' section of the user.
Url:  The URL of their profile
*/
type UserInfo struct{
	ID   int64
	Name string
	Bio  string
	Url  string
}

/*
AnimeList
An AniList animelist 
*/
type AnimeList struct{
	Owner       UserInfo
	Watching    []AnimeInfo
	Completed   []AnimeInfo
	OnHold      []AnimeInfo
	PlanToWatch []AnimeInfo
	Dropped     []AnimeInfo
	
}

// fetch info on a given anime
func fetchAnimeInfo(ID int64) {
	// TODO
}

// fetch info on a given manga
func fetchMangaInfo(ID int64){
	// TODO
}

// fetch info on a given user
func fetchUserInfo(ID int64){

	// form the GQL query.
	var query struct{
		// User() query.
		User struct{
			// get the `name' field.
			Name graphql.String
		} `graphql:"User(id: $id)"`
	}

	client := graphql.NewClient("https://graphql.anilist.co", nil)

	// form the GQL variables with a map.
	variables := map[string]interface{}{
		"id": graphql.Int(ID),
	}
	
	// make the GQL query.
	err := client.Query(context.Background(), &query, variables)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(query.User.Name)

}
// fetch an anime list for a given user
func fetchAnimeList(userName string){
	// TODO
}
