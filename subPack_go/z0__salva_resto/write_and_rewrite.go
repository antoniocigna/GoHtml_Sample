package portaSubPack

import (  
	"fmt"
	"os"	
    "bufio"	
	//"sort"
	"time"
	//"strings"
)
//--------------------------------
//write_and_rewrite.go
//----------------------------------
//--------------------------------

func rewrite_LemmaTranDict_file() {

	//fmt.Println( "GO ", green("rewrite_LemmaTranDict_file" ))
	
	//outFile := FOLDER_IO_lastTRAN +  string(os.PathSeparator) + FILE_ outLemmaTranDict;
	
	outFile := FOLDER_IO_lastTRAN  +  string(os.PathSeparator) + FILE_last_updated_dict_words 

	lines:= make([]string, 0, 10)
	lines = append(lines,  "__" + outFile + "\n" + "_lemma	_traduzione")
	/**
	for z:=0; z < len(dictLemmaTran); z++ {
	
		pkey=key
		
		
		
		key = dictLemmaTran[z].dL_lemma2 + "|"  + dictLemmaTran[z].dL_tran  		////cigna1_3
		
		//if strings.Index(key,"eindhoven")>=0 { fmt.Println("rewrite_LemmaTranDict_file ",  dictLemmaTran[z],  " key=",key, " pkey=", pkey) }
		
		if pkey == key { continue}
		
		lines = append(lines, key ) 
	}
	**/
	writeList( outFile, lines )
	//--------------------

	currentTime := time.Now()		
	outF1 		:= FOLDER_O_arc_TRAN_words +  string(os.PathSeparator) + "dictL"  		
	outFile2 := outF1 + currentTime.Format("20060102150405") + ".txt"
	
	writeList( outFile2, lines )
	
	
} // end of rewrite_LemmaTranDict_file

//----------------------
func writeList( fileName string, lines []string)  {
	// create file
    f, err := os.Create( fileName )
    if err != nil {
        fmt.Println( red("error")," in writeList file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)

    for _, line := range lines {
        _, err := buffer.WriteString(line + "\n")
        if err != nil {
           fmt.Println( red("error"), " in buffer.WriteString file=", fileName,"\n\t" , err ) //log.Fatal(err)
        }
    }
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
        fmt.Println( red("error"), " in buffer.Flush()cls file=", fileName,"\n\t" , err ) //  log.Fatal(err)
    }
} 
//----------------------------------------
//end of write_and_rewrite.go
//----------------------------------