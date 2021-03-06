package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetPosts,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.GetPost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodGet,
		Function:     controllers.FindPost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetUserPosts,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/up-vote",
		Method:       http.MethodPost,
		Function:     controllers.UpVotePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/down-vote",
		Method:       http.MethodPost,
		Function:     controllers.DownVotePost,
		AuthRequired: true,
	},
}
