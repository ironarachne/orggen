package main

import (
	"math/rand"

	"github.com/ironarachne/orggen"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("orggend")
	})

	app.Get("/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		rand.Seed(id)
		org := orggen.Generate()
		ctx.JSON(org)
	})

	app.Run(iris.Addr(":7135"))
}
