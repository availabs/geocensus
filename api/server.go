package main


import (
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	"strings"
//	"fmt"
)

type geojson struct {
    Type        string
    Features 	feature
    //Point       Point
    //Line        Line
    //Polygon     Polygon
}

type feature struct {
	Type        string
	Geometry    geometry
	Properties  map[string]string
}

type geometry struct {
	Type        string
	//Coordinates []map[float]float{}
}
func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "Hey world!"
  })
  m.Get("/states", func() string {
    
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

	rows, err := db.Query("SELECT geoid, name FROM tl_2013_us_state ")
	if err != nil {
		return err.Error()
	}
	states :=[]map[string]string{}
	for rows.Next() {
			var name string
			var geoid string
            if err := rows.Scan(&geoid, &name); err != nil {
                    return err.Error()
            }
            state := map[string]string{
            	"GEOID":geoid,
            	"Name":name,
            }
            states = append(states, state)
    }
	b, err := json.Marshal(states)
	return string(b)
  })

  m.Get("/states/:geoid", func(params martini.Params) string {
  	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

	rows, err := db.Query("SELECT geoid, name,ST_AsGeoJSON(the_geom) as geo FROM tl_2013_us_state WHERE geoid = $1",params["geoid"])
	if err != nil {
		return err.Error()
	}
	states :=[]map[string]string{}
	for rows.Next() {
			var name string
			var geoid string
			var geo string
            if err := rows.Scan(&geoid, &name, &geo); err != nil {
                    return err.Error()
            }

            state := map[string]string{
            	"geometry": geo ,
            	"Properties": "{\"geoid\": "+geoid+"}",
            	"type":"Feature",
            }

            states = append(states, state)

    }
	b, err := json.Marshal(states)

	/*The Following lines of code are used to parse and remove unnecessary characters that get in the way of 
	creating proper geojson files*/

	c := strings.Replace(string(b), "\\", "", -1)
	c = strings.Replace(c, "\":\"", "\":", 1)
	c = strings.Replace(c, "\",\"g", ",\"g", 1)
	c = strings.Replace(c, "y\":\"{", "y\":{", 1)
	c = strings.Replace(c, "\",\"t", ",\"t", 1)
	d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
	return d
})

	/*:geoid converts given part of url in a parameter to be potentially used. This param is used as params["x"] 
	where x is the given parameter*/

m.Get("/states/:geoid/county", func(params martini.Params) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}

	/*The following code selects the information needed from the above host to be accessed and used here. 
	$1 is replaced with the given parameter*/

	rows, err := db.Query("SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county WHERE statefp = $1",params["geoid"])
	if err != nil {
		return err.Error()
	}

	/*The below variables are strings that contain the selected information from the above query*/

	var namelsad string
	var geom string
	var geoid string
	counties :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	county := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	counties = append(counties, county)
        }
        b, err := json.Marshal(counties)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

})

m.Get("/states/:geoid/tracts", func(params martini.Params) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	rows, err := db.Query("SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_tract")
	if err != nil {
		return err.Error()
	}	

	var namelsad string
	var geom string
	var geoid string
	tracts :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	tract := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	tracts = append(tracts, tract)
        }
        b, err := json.Marshal(tracts)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

	})

m.Get("/states/:geoid/bg", func(params martini.Params) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	rows, err := db.Query("SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_bg")
	if err != nil {
		return err.Error()
	}	

	var namelsad string
	var geom string
	var geoid string
	bgs :=[]map[string]string{}
	for rows.Next() {
        	if err := rows.Scan(&geom, &namelsad, &geoid); err != nil {
            	return err.Error()
        	}
        	bg := map[string]string{
            	"geometry": geom ,
            	"Properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	bgs = append(bgs, bg)
        }
        b, err := json.Marshal(bgs)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
        d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
    //result := "\n"+string(counties)+""
	return d//result

	})

  m.Run()
}