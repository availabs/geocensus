package serverAid


import (
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	//"github.com/martini-contrib/cors"
	//"github.com/codegangsta/martini-contrib/binding"
	"strings"
	"fmt"
	"log"
)

type GeoCensusVar struct {
    GeoCenVar   string 	`form:"geoCen" json:"states"`
    GeoCenVar2  string  `form:"state" json:"states"`
}

type ACSVar struct{
	Var1	string 		`form:"table_ID"`
}

type GeoCensusVar2 struct{
	States 			string	`form:"stateList" json:"states"`
	Counties		string 	`form:"countyList" json:"counties"`
	GeoCenVar3[]	string 	`form:"geoVar" json:"GeoVariable"`
}

type GeoCensusOutput struct {
	Geoid 			  string
	CensusVariables[] map[string]int  
}

type SeqAndGeoVar struct{
	SequenceNum 	string
	GeoVar 			string
}

func HelloWorld() string {
    return "Hey world!"
}

func StatesInit() string {
    
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
}

func StatesGeoCounty(params martini.Params,CountyPostA GeoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county"
	
	if CountyPostA.GeoCenVar != "" && CountyPostA.GeoCenVar2 != ""{
		fmt.Print(CountyPostA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+CountyPostA.GeoCenVar+ "AND statefp in "+CountyPostA.GeoCenVar2
		fmt.Print(sql_statement)
	} else{
		fmt.Print(CountyPostA.GeoCenVar+" Test\n")
		fmt.Print("This is empty! County With geoid:"+params["geoid"]+"\n")
		sql_statement += " WHERE statefp = '"+params["geoid"]+"'"
	}

	/*The following code selects the information needed from the above host to be accessed and used here. 
	$1 is replaced with the given parameter*/

	rows, err := db.Query(sql_statement)
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

}

func StatesGeoTracts(params martini.Params, GeoTractsA GeoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	if GeoTractsA.GeoCenVar2 != ""{
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+GeoTractsA.GeoCenVar2+"_tract"
	} else{
		sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_tract"
	} 
	if GeoTractsA.GeoCenVar != ""{
		fmt.Print(GeoTractsA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+GeoTractsA.GeoCenVar
		fmt.Print(sql_statement)
		
	} else{
		fmt.Print("This is empty! BG With geoid:"+GeoTractsA.GeoCenVar+"\n")
		//sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
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

	}

func StatesGeoBG(params martini.Params, GeoBGA GeoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	if GeoBGA.GeoCenVar2 != ""{
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+GeoBGA.GeoCenVar2+"_bg"
	} else {
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+params["geoid"]+"_bg"	
	}
	if GeoBGA.GeoCenVar != ""{
		fmt.Print(GeoBGA.GeoCenVar+" Test\n")
		sql_statement += " WHERE countyfp in "+GeoBGA.GeoCenVar
		fmt.Print(sql_statement)
	} else{
		fmt.Print("This is empty! BG With geoid:"+GeoBGA.GeoCenVar+"\n")
		//sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
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

	}

func StatesGeoid(params martini.Params,StatesPostA GeoCensusVar) string {
  	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	sql_statement = "SELECT geoid, name,ST_AsGeoJSON(the_geom) as geo FROM tl_2013_us_state"
	
	if StatesPostA.GeoCenVar != ""{
		fmt.Print(StatesPostA.GeoCenVar+" Test\n")
		sql_statement += " WHERE geoid in "+StatesPostA.GeoCenVar
		fmt.Print(sql_statement)
	} else{
		fmt.Print("This is empty! States With geoid:"+params["geoid"]+"\n")
		sql_statement += " WHERE geoid = '"+params["geoid"]+"'"
	}
	rows, err := db.Query(sql_statement)
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

	c := strings.Replace(string(b), "\\","",-1)
    c = strings.Replace(c, "Properties\":\"", "Properties\":", -1)
	c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
	c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
	c = strings.Replace(c, "\",\"type", ",\"type", -1)
    d := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+c+"}"
	return d
}

func Acs20105year() string{
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
		sql_statement := "SELECT table_ID, Stub FROM table_shell WHERE var_order = 0 and not stub like 'Universe%' and not COALESCE(stub, '') = ''"
	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}
	acs2010_5 :=[]map[string]string{}
	for rows.Next() {
			var table_ID string
			var stub string
            if err := rows.Scan(&table_ID, &stub); err != nil {
                    return err.Error()
            }

            acs2010 := map[string]string{
            	table_ID: stub,
            	
            	
            	//"type":"Feature",
            }

            acs2010_5 = append(acs2010_5, acs2010)

    }
	b, err := json.Marshal(acs2010_5)
		return string(b)
	}

func Acs20105yearTable_id(params martini.Params, TABLEA ACSVar) string{
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	sql_statement := "SELECT Unique_ID, Stub FROM table_shell WHERE var_order > 0 and table_ID = "+TABLEA.Var1
	rows, err := db.Query(sql_statement)
	if err != nil {
		return err.Error()
	}
	acs2010_5 :=[]map[string]string{}
	for rows.Next() {
			var table_ID string
			var stub string
            if err := rows.Scan(&table_ID, &stub); err != nil {
                    return err.Error()
            }

            acs2010 := map[string]string{
            	table_ID:stub,
            }
            acs2010_5 = append(acs2010_5, acs2010)

    }
	b, err := json.Marshal(acs2010_5)
		return string(b)
}

func Acs20105yearStates(params martini.Params, TABLE GeoCensusVar2) []byte{
	if TABLE.GeoCenVar3[0] != ""{
	arrayPos := 0
	var geoVarMaps []SeqAndGeoVar
	var outputArray []GeoCensusOutput
	seqNumCount := 0
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		log.Fatal("Database access error")
	}
	sql_statement := "SELECT unique_id, sequence_number FROM \"public\".table_shell JOIN \"public\".data_dictionary ON \"public\".table_shell.table_id = \"public\".data_dictionary.table_id AND \"public\".table_shell.var_order = \"public\".data_dictionary.line_number WHERE \"public\".table_shell.unique_id in "+TABLE.GeoCenVar3[0]
	rows, err := db.Query(sql_statement)
	if err != nil {
		log.Fatal("SQL error 1 "+err.Error())
	}
	var rowString string
	var geoVarSQL string
	for rows.Next(){
		if err := rows.Scan(&geoVarSQL,&rowString); err != nil {
                    log.Fatal("SQL error 2")
            }
        var geoVarMap SeqAndGeoVar
        	geoVarMap.GeoVar = geoVarSQL
        	geoVarMap.SequenceNum = rowString
        geoVarMaps = append(geoVarMaps,geoVarMap)
        seqNumCount++
    }
    for iterator := 0;iterator<seqNumCount;iterator++{
		sql_statement2 := "select state,"+geoVarMaps[iterator].GeoVar+" from \"acs2010_5yr\".seq00"+geoVarMaps[iterator].SequenceNum+" as a join \"acs2010_5yr\".geoheader as b ON a.logrecno = b.logrecno and a.stusab = b.stusab where b.sumlevel='040' and a.logrecno = 1 "//and b.state in "+TABLE.States//+" and b.county = '"+TABLE.Counties[0]+"'" 	
		rows2, err3 := db.Query(sql_statement2)
		if err3 != nil {
			log.Fatal("SQL error")
		}
		var rowString2 string
		var rowString3 string
		newItem := 1
		i := 0
		for rows2.Next(){
			var temp GeoCensusOutput
			if err := rows2.Scan(&rowString3,&rowString2); err != nil {
	           	        log.Fatal("SQL row error")
	       	    }
	       	    rowString2 = strings.Replace(rowString2, ".", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+06", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+07", "", -1)
	       	    if arrayPos == 0 {
	       	    	temp.Geoid = rowString3
	       	        var tempGeoid string = string(geoVarMaps[iterator].GeoVar)
	       	    	var tempValue int = int(rowString2)
	       	    	temp.CensusVariables = append(temp.CensusVariables,map[tempGeoid]tempValue)
	       	    	outputArray = append(outputArray,temp)
	       	    	arrayPos++
	       	    } else {
	       	    	newItem = 1
	       	    	i = 0
	       	    	for ;i<arrayPos;i++{
	       	    		if outputArray[i].Geoid == rowString3{
	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	       	    			tempValue := int(rowString2)
	       	    			outputArray[i].CensusVariables = append(outputArray[i].CensusVariables,map[tempGeoid]tempValue)
	       	    			newItem = 0
	       	    			break
	       	    		}
	       	    	}
	       	    	if newItem == 1{
	       	    			temp.Geoid = rowString3
	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	       	    			tempValue := int(rowString2)
	       	    			temp.CensusVariables = append(temp.CensusVariables,map[tempGeoid]tempValue)
	       	    			outputArray = append(outputArray,temp)
	       	    			arrayPos++
	       	    		}
	       	    	
	       	    }
	       	    //state := map[string]string{
	            //	"Geoid": rowString3,
	            // 	geoVarMaps[iterator].GeoVar: rowString2,
	            //}
	       	  	//states = append(states,state)
	       	    
	    	}
	}
	/*states :=[]map[string]string{}
		for i:=0;i<arrayPos-1;i++{
			state := make(map[string]string{})
			state["Geoid"] = GeoCensusVar2Arr[i].States
			
			//state = append(state, :GeoCensusVar2Arr[i].States)
			
			states = append(states,state)
		}*/
    	b, err3 := json.Marshal(outputArray)
    	if err3 != nil {
		log.Fatal("Marshal error")
		}
		return b
	}
	log.Fatal("Error with census field")
	return nil
}

func Acs20105yearStatesCounties(params martini.Params, TABLE GeoCensusVar2) []byte{
	if TABLE.GeoCenVar3[0] != ""{
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		log.Fatal("Database access error")
	}
	sql_statement := "SELECT unique_id, sequence_number FROM \"public\".table_shell JOIN \"public\".data_dictionary ON \"public\".table_shell.table_id = \"public\".data_dictionary.table_id AND \"public\".table_shell.var_order = \"public\".data_dictionary.line_number WHERE \"public\".table_shell.unique_id in "+TABLE.GeoCenVar3[0]
	rows, err := db.Query(sql_statement)
	if err != nil {
		log.Fatal("SQL error 1"+err.Error())
	}
	var rowString string
	var geoVarSQL string
	var geoVarMaps []SeqAndGeoVar
	arrayPos := 0
	states :=[]map[string]string{}
	seqNumCount := 0
	for rows.Next(){
		if err := rows.Scan(&geoVarSQL,&rowString); err != nil {
                    log.Fatal("SQL error 2")
            }
        var geoVarMap SeqAndGeoVar
        	geoVarMap.GeoVar = geoVarSQL
        	geoVarMap.SequenceNum = rowString
        geoVarMaps = append(geoVarMaps,geoVarMap)
        seqNumCount++
    }
    
    for iterator := 0;iterator < seqNumCount;iterator++{
		sql_statement2 := "select county,state,"+geoVarMaps[iterator].GeoVar+" from \"acs2010_5yr\".seq00"+geoVarMaps[iterator].SequenceNum+" as a join \"acs2010_5yr\".geoheader as b ON a.logrecno = b.logrecno and a.stusab = b.stusab where b.sumlevel='050'"//and b.state in "+TABLE.States//+" and b.county = '"+TABLE.Counties[0]+"'" 	
		rows2, err3 := db.Query(sql_statement2)
		if err3 != nil {
			log.Fatal("SQL error")
		}
		var rowString2 string
		var rowString3 string
		var rowString4 string
		for rows2.Next(){
			if err := rows2.Scan(&rowString4,&rowString3,&rowString2); err != nil {
	           	        log.Fatal("SQL row error")
	       	    }
	       	    rowString2 = strings.Replace(rowString2, ".", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+06", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+07", "", -1)
	       	    state := map[string]string{
	            	"Geoid": rowString3+rowString4,
	            	geoVarMaps[iterator].GeoVar: rowString2,
	            }
	       	    states = append(states,state)
	       	    arrayPos++
	    	}
    }
    	b, err3 := json.Marshal(states[:arrayPos])
    	if err3 != nil {
		log.Fatal("Marshal error")
		}
		return b
	}
	log.Fatal("Error with census field")
	return nil
}