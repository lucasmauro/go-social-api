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
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		AuthRequired: true,
	},
}