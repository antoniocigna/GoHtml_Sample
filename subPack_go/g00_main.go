package subPack_go

	import (
		"fmt"
		"os"
		"os/signal"
		"strings"		
		"strconv"
		"runtime"	
	)
//---------------------
//g00_main.go
//------------------------------------------------------
func Main() {

	fmt.Println("\n======================\n         My Main()  INIZIO di mainPack \n===============================\n")
	fmt.Println(  red( appName + " - Main") )
	
	//fmt.Println( "\ncolori:", red("rosso"), green("verde"), yellow("giallo"),  magenta("magenta"), cyan("ciano") , "\n"  )  
	
	//---------------
	//val0, val1, val2, val3, val4 := getPgmArgs("-html",  "-input" , "-countNumLines" , "-maxNumLinesToWrite", "-lemmaformat")	
	
	//-----------------------------------
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	//  err := lorca.New("", "", 480, 320, args...) moved out of main so that ui is available outside main()
	if err != nil {
		fmt.Println( red( "errore in lorca "), err )  //  //log.Fatal(err)
	}
	defer ui.Close()
	
	bind_go_func_to_js()  //  bind inside is executed asynchronously after calling from js function (html/js are ready) 
	
	begin_GO_HTML_Talk();  // this function is  executed firstily before html/js is ready  
	
	// the following in main() is executed at the end when the browser is close 
	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
		case <-sigc:
		case <-ui.Done():
	}
	fmt.Println("exiting") // log.Println("exiting...")
}
//-----------------------------------------

func endBegin(wh string) {

	//fmt.Println("func endBegin (", wh,")")
	if sw_stop { 
		fmt.Println("\nXXXXXXXX  error found XXXXXXXXXXXXXX\n"); 
	}	
	sw_begin_ended = true 		
}
//--------------------------------

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//-----------------------------------------

func getInt(x string) int {	
	y1, e1 := strconv.Atoi( x ) 
	if e1 == nil { 
		return y1
	} 
	y2, e2 := strconv.Atoi(  "0"+strings.TrimSpace(x)  ) 
	if e2 == nil {
		return y2
	} else {
		fmt.Println("error in getInt(",x,") ", e2) 
	}
	return 0
}

//-----------------------------------------------------------
//end of g00_main.go
//------------------------------------------------------
