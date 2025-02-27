package serverAid


import (
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"github.com/codegangsta/martini"
	"strings"
	"fmt"
	"log"
	"strconv"
	"io"
	"os"
	"os/exec"
	//"bytes"
)

type GeoCensusVar struct {
    GeoCenVar   	string 	`form:"geoCen" json:"states"`
    GeoCenVar2  	string  `form:"state" json:"states"`
    GeoCenVar3		string 	`form:"data" json:states"`
}

type ACSVar struct{
	Var1	string 		`form:"table_ID"`
}

type GeoCensusVar2 struct{
	States 			string	`form:"stateList" json:"states"`
	Counties		string 	`form:"countyList" json:"counties"`
	Tract 			string 	`form:"tractList" json:tract"`
	GeoCenVar3[]	string 	`form:"geoVar" json:"GeoVariable"`
	StateFips		string  `form:"state" json:"states"`
}

type GeoCensusOutput struct {
	Geoid 			  string
	Tract 			  string
	CensusVariables[] map[string]int  
}

type SeqAndGeoVar struct{
	SequenceNum 	string
	GeoVar 			string
}

func StatesGeoTracts(params martini.Params, GeoTractsA GeoCensusVar) string {

	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		return err.Error()
	}
	var sql_statement string
	//fmt.Println(GeoTractsA.GeoCenVar+" "+GeoTractsA.GeoCenVar2)
	if GeoTractsA.GeoCenVar == ""{
	sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+GeoTractsA.GeoCenVar2+"_tract"
	} else{
		sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county WHERE geoid = '"+GeoTractsA.GeoCenVar2+"'"
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
            	"properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\"}",
            	"type":"Feature",
            }
        	tracts = append(tracts, tract)
        }
        b, err := json.Marshal(tracts)
        c := strings.Replace(string(b), "\\","",-1)
        c = strings.Replace(c, "properties\":\"", "properties\":", -1)
		c = strings.Replace(c, "\",\"geometry", ",\"geometry", -1)
		c = strings.Replace(c, "geometry\":\"{", "geometry\":{", -1)
		c = strings.Replace(c, "\",\"type", ",\"type", -1)
		c = strings.Replace(c,"]]]]}\",\"properties\":{\"","]]]]},\"properties\":{\"",-1)
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
            fmt.Println(table_ID)
            fmt.Println(stub)
            acs2010 := map[string]string{
            	table_ID: stub,
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
	db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
	if err != nil {
		log.Fatal("Database access error")
	}
	var sql_statement2 string
	if TABLE.States == "050"{
		sql_statement2 = "select state,name,county from \"acs2010_5yr\".geoheader where sumlevel = '050'"
	} else{
		sql_statement2 = "select state,name,county from \"acs2010_5yr\".geoheader where logrecno = 1"
	}
		rows2, err3 := db.Query(sql_statement2)
		if err3 != nil {
			//fmt.Println("Test "+sql_statement2)
			//log.Fatal("SQL error")
			return nil
		}
		var rowString2 sql.NullString
		var rowString3 sql.NullString
		var rowString6 sql.NullString
		rowString4 := ""
		rowString5 := ""
		states :=[]map[string]string{}
		for rows2.Next(){
			if err := rows2.Scan(&rowString3,&rowString2,&rowString6); err != nil {
	           	        //log.Fatal("SQL row error")
						//fmt.Println("Error!")
	           	        return nil
	       	    }
	       	    if rowString2.Valid{
	       	    	rowString4 = rowString2.String
	       	    }
	       	    if rowString3.Valid{
	       	    	rowString5 = rowString3.String
	       	    }
	       	    if TABLE.States == "050" && rowString6.Valid{
	       	    	rowString5 = rowString5+rowString6.String
	       	    }
	       	    state := map[string]string{
	            	rowString4: rowString5,
	             	
	            	}
	       	  	states = append(states,state)
	       	  	
	       	  }
	       	  b, err3 := json.Marshal(states)
    	if err3 != nil {
    		//fmt.Println("Error2!")
			//log.Fatal("Marshal error")
			return nil
		}
		return b
}




func Acs20105yearQuery(params martini.Params, TABLE GeoCensusVar2) []byte{
	
	if TABLE.GeoCenVar3[0] != ""{
		db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
		if err != nil {
			//log.Fatal("Database access error")
			return nil
		}
		
		sql_statement := "SELECT unique_id, sequence_number FROM \"public\".table_shell JOIN \"public\".data_dictionary ON \"public\".table_shell.table_id = \"public\".data_dictionary.table_id AND \"public\".table_shell.var_order = \"public\".data_dictionary.line_number WHERE \"public\".table_shell.unique_id in "+TABLE.GeoCenVar3[0]
		rows, err := db.Query(sql_statement)
		if err != nil {
			//fmt.Println(sql_statement)
			//log.Fatal("SQL error 1 "+err.Error())
			return nil
		}
		var rowString string
		var geoVarSQL string
		var geoVarMaps []SeqAndGeoVar
		arrayPos := 0
		var outputArray []GeoCensusOutput
		seqNumCount := 0
		for rows.Next(){
			if err := rows.Scan(&geoVarSQL,&rowString); err != nil {
                //log.Fatal("SQL error 2")
				return nil
            }
        var geoVarMap SeqAndGeoVar
        geoVarMap.GeoVar = geoVarSQL
        geoVarMap.SequenceNum = rowString
        geoVarMaps = append(geoVarMaps,geoVarMap)
        seqNumCount++
    	}
    	for iterator := 0;iterator<seqNumCount;iterator++{
    	stateSubStrArr := SubStringArray(TABLE.States)
    	for si := 0; si < len(stateSubStrArr); si++ {
    	
    	checkStr, errS := strconv.ParseInt(geoVarMaps[iterator].SequenceNum, 10, 0)
    	seqStr := ""
    	if checkStr < 10 && errS == nil {
    		seqStr = "0"+geoVarMaps[iterator].SequenceNum
    	} else{
    		seqStr = geoVarMaps[iterator].SequenceNum
    	}
    	//fmt.Println(geoVarMaps[iterator].SequenceNum)
    	sql_statement2 := "select geoid,"+geoVarMaps[iterator].GeoVar+" from \"acs2010_5yr\".seq00"+seqStr+" as a join \"acs2010_5yr\".geoheader as b ON a.logrecno = b.logrecno and a.stusab = b.stusab where b.sumlevel='"+TABLE.Counties+"' and b.geoid LIKE '"+TABLE.Counties+"00US"+stateSubStrArr[si]+"%'" 	
		rows2, err3 := db.Query(sql_statement2)
		if err3 != nil {
			//log.Fatal("SQL error "+err3.Error())
			return nil
		}
		var rowString4 sql.NullString
		var rowString5 sql.NullString
		rowString2 := ""
		rowString3 := ""
		newItem := 1
		i := 0
		rowString2 = "0"
		for rows2.Next(){
			var temp GeoCensusOutput
			if err := rows2.Scan(&rowString5,&rowString4); err != nil {
	           	        //log.Fatal("SQL row error "+err.Error())
					return nil
	       	    }
	       	    if rowString4.Valid{
	       	    rowString2 = rowString4.String
	       	    rowString2 = strings.Replace(rowString2, ".", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+06", "", -1)
	       	    rowString2 = strings.Replace(rowString2, "e+07", "", -1)
	       		}
	       		if rowString5.Valid{
	       			rowString3 = rowString5.String
	       	    	rowString3 = strings.Replace(rowString3,TABLE.Counties+"00US","",-1)
	       		}

	       	    if arrayPos == 0 {
	       	    	temp.Geoid = rowString3
	       	        var tempGeoid string = string(geoVarMaps[iterator].GeoVar)
	       	 	  	tempValue, err := strconv.Atoi(rowString2)
	       	 	  	if err != nil {
        				//log.Fatal("Conversion Error")
        				tempValue = 0
    				}
	       	 	  	tempMap := map[string]int{
	       	 	  		tempGeoid:tempValue,
	       	 	  	}
	       	    	temp.CensusVariables = append(temp.CensusVariables,tempMap)
	       	    	outputArray = append(outputArray,temp)
	       	    	arrayPos++
	       	    } else {
	       	    	newItem = 1
	       	    	i = 0
	       	    	for ;i<arrayPos;i++{
	       	    		if outputArray[i].Geoid == rowString3{
	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	       	    			tempValue, err := strconv.Atoi(rowString2)
			       	 	  	if err != nil {
		        				//log.Fatal("Conversion Error")
		        				tempValue = 0
		        				//return nil
    						}

	       	    			tempMap := map[string]int{
	       	 	  				tempGeoid:tempValue,
	       	 	  			}
	       	    			outputArray[i].CensusVariables = append(outputArray[i].CensusVariables,tempMap)
	       	    			newItem = 0
	       	    			break
	       	    		}
	       	    	}
	       	    	if newItem == 1{
	       	    			temp.Geoid = rowString3
	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	       	    			tempValue, err := strconv.Atoi(rowString2)
	       	       	 	  	if err != nil {
		        				//log.Fatal("Conversion Error")
		        				tempValue = 0
    						}

	       	    			tempMap := map[string]int{
	       	 	  				tempGeoid:tempValue,
	       	 	  			}
	       	    			temp.CensusVariables = append(temp.CensusVariables,tempMap)
	       	    			outputArray = append(outputArray,temp)
	       	    			arrayPos++
	       	    		}
	       	    	
	       	    }
	       	}
	       }
}
	b, err3 := json.Marshal(outputArray)
    if err3 != nil {
		//log.Fatal("Marshal error")
		return nil
	}
	return b
	
	}
//log.Fatal("Error with census field")
return nil

}

func SubStringArray(s string) []string{

	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	ssA := strings.Split(s,",")
	return ssA
}



func Acs20105yearQuerySpecial(params martini.Params, TABLE GeoCensusVar2) string{
	
	if TABLE.GeoCenVar3[0] != ""{
		db, err := sql.Open("postgres", "host=lor.availabs.org password=transit user=postgres dbname=geocensus sslmode=require")
		if err != nil {

			return "A"
		}
		//parsedVar := "('B01003_001E,B12006_001E,B12006_005E,B12006_010E,B12006_016E,B12006_021E,B12006_027E,B12006_032E,B12006_038E,B12006_043E,B12006_049E,B12006_054E,B12006_006E,B12006_011E,B12006_017E,B12006_022E,B12006_028E,B12006_033E,B12006_039E,B12006_044E,B12006_050E,B12006_055E,B08301_001E,B08301_002E,B08301_010E,B08301_016E,B08301_017E,B08301_018E,B08301_019E,B08301_020E,B08301_021E,B08301_011E,B08301_013E,B08301_014E,B08126_001E,B08126_002E,B08126_003E,B08126_004E,B08126_005E,B08126_006E,B08126_007E,B08126_008E,B08126_009E,B08126_010E,B08126_011E,B08126_012E,B08126_013E,B08126_014E,B08126_015E,B19001_001E,B19001_002E,B19001_003E,B19001_004E,B19001_005E,B19001_006E,B19001_007E,B19001_008E,B19001_009E,B19001_010E,B19001_011E,B19001_012E,B19001_013E,B19001_014E,B19001_015E,B19001_016E,B19001_017E,B19013_001E,B17001_002E,B14003_003E,B14003_012E,B14003_031E,B14003_040E,B23006_002E,B23006_009E,B23006_016E,B23006_023E,B05006_001E,B06007_005E,B06007_008E,B01001_002E,B01001_026E,B01001_004E,B01001_005E,B01001_006E,B01001_007E,B01001_008E,B01001_009E,B01001_010E,B01001_011E,B01001_012E,B01001_013E,B01001_014E,B01001_015E,B01001_016E,B01001_017E,B01001_018E,B01001_019E,B01001_020E,B01001_021E,B01001_022E,B01001_023E,B01001_024E,B01001_025E,B01001_027E,B01001_028E,B01001_029E,B01001_030E,B01001_031E,B01001_032E,B01001_033E,B01001_034E,B01001_035E,B01001_036E,B01001_037E,B01001_038E,B01001_039E,B01001_040E,B01001_041E,B01001_042E,B01001_043E,B01001_044E,B01001_045E,B01001_046E,B01001_047E,B01001_048E,B01001_049E,B02001_002E,B02001_003E,B02001_004E,B02001_005E,B02001_006E,B02001_007E,B02001_008E,B25002_001E,B25002_002E,B25002_003E,B25024_002E,B25024_003E,B25024_004E,B25024_005E,B25024_006E,B25024_007E,B25024_008E,B25024_009E,B25024_010E,B25024_011E,B25003_002E,B25003_003E,B08014_002E,B08014_003E,B08014_004E,B08014_005E,B08014_006E,B08014_007E,B08132_002E,B08132_003E,B08132_004E,B08132_005E,B08132_006E,B08132_007E,B08132_008E,B08132_009E,B08132_010E,B08132_011E,B08132_012E,B08132_013E,B08132_014E,B08132_015E,B08132_046E,B08132_047E,B08132_048E,B08132_049E,B08132_050E,B08132_051E,B08132_052E,B08132_053E,B08132_054E,B08132_055E,B08132_056E,B08132_057E,B08132_058E,B08132_059E,B08132_060E,B08133_001E,B08133_002E,B08133_003E,B08133_004E,B08133_005E,B08133_006E,B08133_007E,B08133_008E,B08133_009E,B08133_010E,B08133_011E,B08133_012E,B08133_013E,B08133_014E,B08133_015E,B08122_001E,B08122_002E,B08122_003E,B08122_004E,B08122_005E,B08122_006E,B08122_007E,B08122_008E,B08122_009E,B08122_010E,B08122_011E,B08122_012E,B08122_013E,B08122_014E,B08122_015E,B08122_016E,B08122_017E,B08122_018E,B08122_019E,B08122_020E,B08122_021E,B08122_022E,B08122_023E,B08122_024E,B08122_025E,B08122_026E,B08122_027E,B08122_028E,B08136_001E,B08136_002E,B08136_003E,B08136_004E,B08136_005E,B08136_006E,B08136_007E,B08136_008E,B08136_009E,B08136_010E,B08136_011E,B08136_012E,B08126_001E,B08126_002E,B08126_003E,B08126_004E,B08126_005E,B08126_006E,B08126_007E,B08126_008E,B08126_009E,B08126_010E,B08126_011E,B08126_012E,B08126_013E,B08126_014E,B08126_015E,B08126_046E,B08126_047E,B08126_048E,B08126_049E,B08126_050E,B08126_051E,B08126_052E,B08126_053E,B08126_054E,B08126_055E,B08126_056E,B08126_057E,B08126_058E,B08126_059E,B08126_060E,B08519_001E,B08519_002E,B08519_003E,B08519_004E,B08519_005E,B08519_006E,B08519_007E,B08519_008E,B08519_009E,B08519_028E,B08519_029E,B08519_030E,B08519_031E,B08519_032E,B08519_033E,B08519_034E,B08519_035E,B08519_036')"
		parsedVar := strings.Replace(TABLE.GeoCenVar3[0],",'GAMMA","",-1)
		sql_statement := "SELECT unique_id, sequence_number FROM \"public\".table_shell JOIN \"public\".data_dictionary ON \"public\".table_shell.table_id = \"public\".data_dictionary.table_id AND \"public\".table_shell.var_order = \"public\".data_dictionary.line_number WHERE \"public\".table_shell.unique_id in "+parsedVar
		//fmt.Println(sql_statement)
		rows, err := db.Query(sql_statement)
		if err != nil {
			fmt.Println(sql_statement)
			return "B"
		}
		var rowString string
		var geoVarSQL string
		var geoVarMaps []SeqAndGeoVar
		var stateSubStrArr []string
		//arrayPos := 0
		//var hashMaps = make(map[string]int)
		//var outputArray []GeoCensusOutput
		seqNumCount := 0
		skip := 0
		for rows.Next(){
			if err := rows.Scan(&geoVarSQL,&rowString); err != nil {

        		return "C"
            }
            if len(geoVarMaps) == 0{
				var geoVarMap SeqAndGeoVar
		        geoVarMap.GeoVar = geoVarSQL
		        geoVarMap.SequenceNum = rowString
		        geoVarMaps = append(geoVarMaps,geoVarMap)
		        seqNumCount++	            	
            } else {
            	for i := 0;i<len(geoVarMaps);i++{
            		if geoVarMaps[i].SequenceNum == rowString{
            			geoVarMaps[i].GeoVar = geoVarMaps[i].GeoVar+","+geoVarSQL
            			skip = 1
            		}
            	}
            	if skip == 0{
            		var geoVarMap SeqAndGeoVar
			        geoVarMap.GeoVar = geoVarSQL
			        geoVarMap.SequenceNum = rowString
			        geoVarMaps = append(geoVarMaps,geoVarMap)
			        seqNumCount++
            	}
            }
            skip = 0
    	}
    	//fmt.Println(geoVarMaps)
    	stateSubStrArr = SubStringArray(TABLE.States)
    	var orderedResults []string
	    for si := 0; si < len(stateSubStrArr); si++ {
	    	//fmt.Println(stateSubStrArr[si])
	    	var orderedResultsF string
    		for iterator := 0;iterator<seqNumCount;iterator++{
	    	
	    	
	    	checkStr, errS := strconv.ParseInt(geoVarMaps[iterator].SequenceNum, 10, 0)
	    	seqStr := ""
	    	if checkStr < 10 && errS == nil {
	    		seqStr = "0"+geoVarMaps[iterator].SequenceNum
	    	} else{
	    		seqStr = geoVarMaps[iterator].SequenceNum
	    	}
	    	//fmt.Println(geoVarMaps[iterator].SequenceNum)


	    	//Fix below logic!

	    	sql_statement2 := "select geoid,name,"+geoVarMaps[iterator].GeoVar+" from \"acs2010_5yr\".seq00"+seqStr+" as a join \"acs2010_5yr\".geoheader as b ON a.logrecno = b.logrecno and a.stusab = b.stusab where b.sumlevel='"+TABLE.Counties+"' and b.geoid LIKE '"+TABLE.Counties+"00US"+stateSubStrArr[si]+"%'" 	
			rows2, err3 := db.Query(sql_statement2)
			if err3 != nil {
				//log.Fatal("SQL error "+err3.Error())
				fmt.Println(sql_statement2)
				return "D"
			}
			// var rowString4 sql.NullString
			// var rowString5 sql.NullString
			// var rowString6 sql.NullString
			// rowString2 := ""
			// rowString3 := ""
			// rowString7 := ""
			// newItem := 1
			// rowString2 = "0"
			for rows2.Next(){
				
				/*var temp GeoCensusOutput
				if err := rows2.Scan(&rowString5,&rowString4,&rowString6); err != nil {
		           	    return "E"
		       	    }
		       	    if rowString4.Valid{
		       	    rowString2 = rowString4.String
		       	    rowString2 = strings.Replace(rowString2, ".", "", -1)
		       	    rowString2 = strings.Replace(rowString2, "e+06", "", -1)
		       	    rowString2 = strings.Replace(rowString2, "e+07", "", -1)
		       		}
		       		if rowString5.Valid{
		       			rowString3 = rowString5.String
		       	    	rowString3 = strings.Replace(rowString3,TABLE.Counties+"00US","",-1)
		       		}
		       		if rowString6.Valid{
		       			rowString7 = rowString6.String
		       		}*/
		       		ssA := strings.Split(geoVarMaps[iterator].GeoVar,",")
		       		columnNames, err := rows2.Columns()
					if err != nil {
					    return "A"+err.Error() // or whatever error handling is appropriate
					}
					columns := make([][]byte, len(columnNames))
					columnPointers := make([]interface{}, len(columnNames))
					//result := make([]string, len(columnNames))
					for i := 0; i < len(columnNames); i++ {
					    columnPointers[i] = &columns[i]
					}
					if err := rows2.Scan(columnPointers...); err != nil {
					    return "B"+err.Error()
					}
					for i, raw := range columns {
						if i > 1{
				            if raw == nil {
				                orderedResultsF = orderedResultsF+",\""+ssA[i-2]+"\":0"
				            } else {
				                orderedResultsF = orderedResultsF+",\""+ssA[i-2]+"\":"+string(raw)
				            }
				        }
			        }

			        //fmt.Printf("%#v\n", orderedResultsF)


		       		
	// //////Hash function goes here!


	// 	       	//First element in the array
	// 	       	    if arrayPos == 0 {
	// 	       	    	temp.Geoid = rowString3
	// 	       	    	temp.Tract = rowString7
	// 	       	        var tempGeoid string = string(geoVarMaps[iterator].GeoVar)
	// 	       	 	  	tempValue, err := strconv.Atoi(rowString2)
	// 	       	 	  	if err != nil {
	//         				tempValue = 0
	//     				}
	// 	       	 	  	tempMap := map[string]int{
	// 	       	 	  		tempGeoid:tempValue,
	// 	       	 	  	}
	// 	       	    	temp.CensusVariables = append(temp.CensusVariables,tempMap)
	// 	       	    	outputArray = append(outputArray,temp)
	// 	       	    	arrayPos++
	// 	       	    	hashMaps[rowString3] = arrayPos
	// 	       	    } else {
	// 	    //continuous element in the array
	// 	       	    	newItem = 1
	// 	       	    	if hashMaps[rowString3] > 0{
	// 	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	// 	       	    			tempValue, err := strconv.Atoi(rowString2)
	// 			       	 	  	if err != nil {
	// 		        				tempValue = 0
	// 		        			}

	// 	       	    			tempMap := map[string]int{
	// 	       	 	  				tempGeoid:tempValue,
	// 	       	 	  			}
	// 	       	    			outputArray[hashMaps[rowString3]-1].CensusVariables = append(outputArray[hashMaps[rowString3]-1].CensusVariables,tempMap)
	// 	       	    			newItem = 0
	// 	       	    		}
		   
	// 	       //Add new element to the array	    	
	// 	       	    	if newItem == 1{
	// 	       	    			temp.Geoid = rowString3
	// 	       	    			temp.Tract = rowString7
	// 	       	    			tempGeoid := string(geoVarMaps[iterator].GeoVar)
	// 	       	    			tempValue, err := strconv.Atoi(rowString2)
	// 	       	       	 	  	if err != nil {
	// 		        				tempValue = 0
	//     						}

	// 	       	    			tempMap := map[string]int{
	// 	       	 	  				tempGeoid:tempValue,
	// 	       	 	  			}
	// 	       	    			temp.CensusVariables = append(temp.CensusVariables,tempMap)
	// 	       	    			outputArray = append(outputArray,temp)
	// 	       	    			arrayPos++
	// 	       	    			hashMaps[rowString3] = arrayPos
	// 	       	    		}
		       	    	
	// 	       	    }
	// ///////Hash function stops here


		       	}
		       	

		    }
		    orderedResultsF = strings.Replace(orderedResultsF,",","",1)
		    orderedResults = append(orderedResults,orderedResultsF)
		    //fmt.Println(orderedResultsF)
		}
			//fmt.Println(orderedResults)


// 	/*b, err3 := json.Marshal(outputArray)
//     if err3 != nil {
// 		//log.Fatal("Marshal error")
// 		return ""
// 	}*/
// 	//fmt.Println("Below")
// 	//fmt.Println(b)
// 	appendString := ""
// 	var apprendStrArr []string
// 	//fmt.Println(outputArray)
// 	//fmt.Println(geoVarMaps)
// 	for count := 0;count<len(stateSubStrArr);count++{
// 		for counter := 0;counter<len(geoVarMaps);counter++{
// 			//fmt.Println(parsedArr[counter])
// 			//fmt.Println(strings.Replace(parsedArr[counter],",","",-1))
// 			//fmt.Println(strconv.Itoa(outputArray[0].CensusVariables[counter][strings.Replace(parsedArr[counter],",","",-1)]))
// 			if counter + 1 !=  len(geoVarMaps){
// 				appendString = appendString+"\""+geoVarMaps[counter].GeoVar+"\":"+strconv.Itoa(outputArray[count].CensusVariables[counter][geoVarMaps[counter].GeoVar])+","
// 			} else{
// 				appendString = appendString+"\""+geoVarMaps[counter].GeoVar+"\":"+strconv.Itoa(outputArray[count].CensusVariables[counter][geoVarMaps[counter].GeoVar])
// 			}
// 		}
// 		apprendStrArr = append(apprendStrArr,appendString)
// 		appendString = ""
		
// 	}
// 	//fmt.Println(apprendStrArr)
// 	//fmt.Println(outputArray[0].CensusVariables[0][parsedVar])
// 	//fmt.Println(len(outputArray[0].CensusVariables))
	

	//Below is where string to be turned into a json object is created. This will fuse the given values together 
	//and return the desired result

	/*Statement for type of sum level goes:

	tracts
	counties or states

	*/
	tracts :=[]map[string]string{}
	for count := 0;count<len(stateSubStrArr);count++{
		if TABLE.Tract == "tract"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+TABLE.StateFips+"_tract WHERE geoid = '"+stateSubStrArr[count]+"'"
		} else if TABLE.Tract == "bg"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+TABLE.StateFips+"_bg WHERE geoid = '"+stateSubStrArr[count]+"'"
		} else if TABLE.Tract == "county"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county WHERE geoid = '"+stateSubStrArr[count]+"'"
		} else if TABLE.Counties == "140"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+TABLE.States+"_tract"
		} else if TABLE.Counties == "050"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_us_county WHERE geoid = '"+TABLE.States+"'"
		} else if TABLE.Counties == "150"{
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,namelsad,geoid FROM tl_2013_"+TABLE.States+"_bg"
		} else {
			sql_statement = "SELECT ST_AsGeoJSON(the_geom) as geom,name,geoid FROM tl_2013_us_state WHERE geoid = '"+TABLE.States+"'"
		}
		//fmt.Println(sql_statement)	
		rows3, err := db.Query(sql_statement)
		if err != nil {
			fmt.Println(sql_statement)
			return "F "+err.Error()
		}	

		var namelsad string
		var geom string
		var geoid string

		for rows3.Next() {
	        	if err := rows3.Scan(&geom, &namelsad, &geoid); err != nil {
	            	return "G"+err.Error()
	        	}
	        	tract := map[string]string{
	            	"geometry": geom ,
	            	"properties": "{\"geoid\": \""+geoid+"\", \"namelsad\": \""+namelsad+"\","+orderedResults[count]+"}",
	            	"type":"Feature",
	            }
	        	tracts = append(tracts, tract)
	        }
	    }
	        c, err := json.Marshal(tracts)
	        d := strings.Replace(string(c), "\\","",-1)
	        d = strings.Replace(d, "properties\":\"", "properties\":", -1)
			d = strings.Replace(d, "\",\"geometry", ",\"geometry", -1)
			d = strings.Replace(d, "geometry\":\"{", "geometry\":{", -1)
			d = strings.Replace(d, "\",\"type", ",\"type", -1)
			d = strings.Replace(d,"]]]]}\",\"properties\":{\"","]]]]},\"properties\":{\"",-1)
	        e := "{\n  \"type\": \"FeatureCollection\",\n  \"features\": \n"+d+"}"
	        if params["filetype"] == "shape" {													//shapefile requested
		        filename := "downloads/geography.geojson"
		 
			    f, err5 := os.Create(filename)
			    if err5 != nil {
			        fmt.Println(err5)
			    }
			    n, err6 := io.WriteString(f, e)
			    if err6 != nil {
			        fmt.Println(n, err6)
			    }

			    cmd := exec.Command("ogr2ogr","-F","ESRI Shapefile","output.shp", filename, "OGRGeoJSON") //First arg is command. Later args are arguments to be given. Returns cmd struct

				output, err := cmd.CombinedOutput()
				if err != nil {
				    fmt.Println(fmt.Sprint(err) + ": " + string(output))
				    return "A"
				}
				cmd2 := exec.Command("zip","downloads/outPut.zip","output.dbf","output.prj","output.shp","output.shx")
				output2, err2 := cmd2.CombinedOutput()
				if err2 != nil {
					fmt.Println(fmt.Sprint(err2) + ": " + string(output2))
					return "B"
				}
				cmd2 = exec.Command("rm","output.dbf","output.prj","output.shp","output.shx")
				output, err2 = cmd2.CombinedOutput()
				if err2 != nil {
					fmt.Println(fmt.Sprint(err2) + ": " + string(output))
					return "C"
				}
				f.Close()
				return "{\"stuff\":\"Shapefile Requested\"}"
			}
		
		return e
	
	}
return "F"

}