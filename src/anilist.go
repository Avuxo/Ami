package main

/*
This is a wrapper around the AniList GraphQL API.
https://anilist.github.io/ApiV2-GraphQL-Docs/query.doc.html
For primary use in the Ami anime client: https://github.com/Avuxo/Ami/
*/


import (
	"context"
	"log"
	"fmt"
	"github.com/shurcooL/graphql"
	"golang.org/x/oauth2"
)

/*
NATIVE STRUCTURES
The native structures allow for storage of native datatypes.
If just the query structures were used then the `graphql'
types would all be used.
*/

/*
animeInfo
After a GQL query of AniList, this is the query data returned
for any anime.
*/
type animeInfo struct{
	ID       int64
	Episodes int32
	Title    string
	IsAdult  bool
	Status   string
	Genres   []string
}

/*
mangaInfo
After a GQL query of AniList, this is the query data returned
for any manga.
*/
type mangaInfo struct{
	ID       int64
	Chapters int32
	IsAdult  bool
	Status   string
	Genres   []string
}

/*
userInfo
Query a user's basic info.
ID:   the user ID
Name: The user name
Bio:  The `about' section of the user.
Url:  The URL of their profile
*/
type userInfo struct{
	ID   int64
	Name string
	Bio  string
	Url  string
}

/*
animeList
An AniList animelist 
*/
type animeList struct{
	Owner       userInfo
	Watching    []animeInfo
	Completed   []animeInfo
	OnHold      []animeInfo
	PlanToWatch []animeInfo
	Dropped     []animeInfo
	
}

/* 
QUERY STRUCTS 
These are all structured GraphQL queries using the schemas
defined at https://anilist.github.io/ApiV2-GraphQL-Docs/.
For each individual internal structure, there is a corresponding
query structure.
*/
type userQuery struct{
	Name    graphql.String
	About   graphql.String
	SiteUrl graphql.String
} 

type mediaQuery struct{
	IsAdult  bool
	Episodes graphql.Int
	Genres   []graphql.String
	Status   graphql.String
	Title struct{
		Romaji graphql.String
	}
	
}

// struct used inside the listQuery structure 
type list struct{
	Name graphql.String
	Entries struct{
		id graphql.Int
	}
}

type listQuery struct{
	Lists []list
}

// fetch info on a given anime
func fetchAnimeInfo(ID int64) (info animeInfo){
	var query struct{
		Media mediaQuery `graphql:"Media(id: $id type:ANIME)"`
	}
	client := graphql.NewClient("https://graphql.anilist.co", nil)

	// configure the `id' variable into the passed var.
	variables := map[string]interface{}{
		"id": graphql.Int(ID),
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil{
		log.Fatal(err)
	}
	//convert []graphql.String to []string.
	convertedGenres := make([]string, len(query.Media.Genres))
	for i := range query.Media.Genres {
		convertedGenres[i] = string(query.Media.Genres[i])
	}

	// parse the query into an internal structure.
	info = animeInfo{
		int64(ID),
		int32(query.Media.Episodes),
		string(query.Media.Title.Romaji),
		query.Media.IsAdult,
		string(query.Media.Status),
		convertedGenres }

	return info
}

// fetch info on a given manga
func fetchMangaInfo(ID int64){
}

// fetch info on a given user
func fetchUserInfo(ID int64) (info userInfo){

	// form the GQL query.
	var query struct{
		// User() query.
		User userQuery `graphql:"User(id: $id)"`
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

	// convert the query struct into the return struct.
	info = userInfo{
		int64(ID),
		string(query.User.Name),
		string(query.User.About),
		string(query.User.SiteUrl) }

	return info
}
// fetch an anime list for a given user
func fetchAnimeList(userName string){
	var query struct{
		MediaListCollections listQuery `graphql:"MediaListCollection(userName: $name type:ANIME)"`
	}

	client := graphql.NewClient("https://graphql.anilist.co", nil)

	// form the GQL variables according to a map
	variables := map[string]interface{}{
		"name": graphql.String(userName),
	}
	
	err := client.Query(context.Background(), &query, variables)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(query)
}



/*
Mutation
The section of the API that handles things like updating resources.
All mutation functions require an OAuth token. This is provided by `config.json'.
All mutation functions return a boolean.
  true if it passed with no erroes
  false if an error was encountered during the mutation.
*/



// add 1 to the given show's episode's watched (episodesWatched++)
// showID is the ID of the show being updated.
// progress is the current progress of the user.
// this is not count checked, so it's the duty of the frontend to check. [TEMPORARY (hopefully)]
// TODO: countcheck episodes.
func incEpisodesWatched(showID int32, progress int32, OAuthToken string) (bool){
	
	// load the OAuth2 token and instantiate the GQL client with the authenticated token.
	token := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: OAuthToken})
	oauthClient := oauth2.NewClient(context.Background(), token)
	client := graphql.NewClient("https://graphql.anilist.co", oauthClient)
	
	// the mutation struct for SaveMediaListEntry
	// https://anilist.github.io/ApiV2-GraphQL-Docs/mutation.doc.html
	var mutation struct{
		SaveMediaListEntry struct{
			mediaId graphql.Int
			progress graphql.Int
		} `graphql:"SaveMediaListEntry(mediaId: $id progress: $progress)"`
	}

	// map the arguments to the mutation variables.
	variables := map[string]interface{}{
		"id":       graphql.Int(showID),
		"progress": graphql.Int(progress + 1), // 1 episode past previous.
	}

	// shoot off the mutation request (with OAuth2)
	err := client.Mutate(context.Background(), &mutation, variables)
	if err != nil {
		fmt.Println(err)
		return false // error encountered
	}
	
	return true // no errors encountered.
}
