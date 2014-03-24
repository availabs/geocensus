package main


import (
	_ "github.com/lib/pq"
	//"database/sql"
	//"encoding/json"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/cors"
	"github.com/user/newmath"
	//"github.com/codegangsta/martini-contrib/binding"
	//"github.com/user/serverAid"
	//"strings"
	"fmt"
)

type geoCensusVar struct {
    GeoCenVar   string 	`form:"geoCen" json:"states"`
    GeoCenVar2  string  `form:"state" json:"states"`
}

func main() {
  m := martini.Classic()
   m.Use(cors.Allow(&cors.Options{
    AllowOrigins:     []string{"http://*"},
    AllowMethods:     []string{"PUT", "PATCH","POST"},
    AllowHeaders:     []string{"Origin","content-type"},
    ExposeHeaders:    []string{"Content-Length","content-type"},
    AllowCredentials: true,
  }))
  //m.Get("/",aid.helloWorld())
   fmt.Print("%s",newmath.helloWorld());

  //m.Get("/states", server_aid.statesInit())

  //m.Get("/states/:geoid", binding.Bind(geoCensusVar{}),server_aid.statesGeoid())
  //m.Post("/states/",binding.Bind(geoCensusVar{}),server_aid.statesGeoid())



	/*:geoid converts given part of url in a parameter to be potentially used. This param is used as params["x"] 
	where x is the given parameter*/

//m.Get("/states/:geoid/county",binding.Bind(geoCensusVar{}),server_aid.statesGeoCounty())
//m.Post("/states/:geoid/county",binding.Bind(geoCensusVar{}),server_aid.statesGeoCounty())

//m.Get("/states/:geoid/tracts", binding.Bind(geoCensusVar{}),server_aid.statesGeoTracts())
//m.Post("/states/:geoid/tracts", binding.Bind(geoCensusVar{}),server_aid.statesGeoTracts())

//m.Get("/states/:geoid/bg", binding.Bind(geoCensusVar{}),server_aid.statesGeoBG())
//m.Post("/states/:geoid/bg", binding.Bind(geoCensusVar{}),server_aid.statesGeoBG())

  m.Run()
}