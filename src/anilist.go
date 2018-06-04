package main

/*
This is a wrapper around the AniList GraphQL API.
https://anilist.github.io/ApiV2-GraphQL-Docs/query.doc.html
For primary use in the Ami anime client: https://github.com/Avuxo/Ami/
*/



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
	isAdult  bool
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
	// TODO
}

// fetch an anime list for a given user
func fetchAnimeList(userName string){
	// TODO
}
