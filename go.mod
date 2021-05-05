module modanisa/todo

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/sobaidat/todoapp/handlers v0.0.0-00010101000000-000000000000
	github.com/sobaidat/todoapp/utils v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210503195802-e9a32991a82e // indirect
	golang.org/x/net v0.0.0-20210505024714-0287a6fb4125 // indirect
	golang.org/x/sys v0.0.0-20210503173754-0981d6026fa6 // indirect
)

replace github.com/sobaidat/todoapp/handlers => ./handlers

replace github.com/sobaidat/todoapp/utils => ./utils

replace github.com/sobaidat/todoapp/models => ./models
