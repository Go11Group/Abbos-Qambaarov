package main

import (
	"mymod/api"
)

func main()  {
	api.Routes().ListenAndServe()
}