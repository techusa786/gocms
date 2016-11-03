package main

import (
	"github.com/kataras/iris"
)

func main() {
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello %s", "iris")
	})

	// person route
	iris.Get("/person/:firstname/:lastname", func(ctx *iris.Context) {
		firstName := ctx.Param("firstname")
		lastName := ctx.Param("lastname")
		ctx.Write("route:%s\n", "/person/:firstname/:lastname")
		ctx.Write("firstname:%s\tlastname:%s", firstName, lastName)
	})

	iris.Get("/person", func(ctx *iris.Context) {
		id := ctx.URLParam("id")
		firstName := ctx.URLParam("firstname")
		lastName := ctx.URLParam("lastname")
		ctx.Write("route:%s\n", "/person")
		ctx.Write(" id:%s\n firstname:%s \n lastname:%s \n", id, firstName, lastName)
	})

	iris.Listen(":3000")
}

//route:  /details?color=blue&weight=20
//func details(ctx *iris.Context){
//	color:= ctx.URLParam("color")
//	weight:= ctx.URLParamInt("weight")
//}
