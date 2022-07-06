## FEN: fiber enhance

add some enhance feature for golang fiber web framework

```golang
var (
	BusErr_UserNotFound = NewBusError(http.StatusNotFound, 1001, "User not found: %s")
	BusErr_GetParamFailed = NewBusError(http.StatusBadRequest, 1002, "get param %s failed")
)

func main() {
	svc := fiber.New()
	svc.Get("/user/:name", fen.Func(serve))
	svc.Get("/data", fen.DataFunc(serveData))
	svc.Post("/user/:id", fen.DataFunc2(serveUserUpdate,fen.Int("id", BusErr_GetParamFailed), fen.Bind(&User{})))
	svc.Listen(":888")
}

func serve(ctx *gin.Context) error {
	return BusErr_UserNotFound
}

func serveData(ctx *gin.Context) (User, error) {
	return User{Name: "John"}, nil
}

func serveUserUpdate(ctx *gin.Context, id int, user *User) (User, error) {
	return user, nil
}
```