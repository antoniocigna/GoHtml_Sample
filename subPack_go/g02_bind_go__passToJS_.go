package subPack_go

	import (
		"fmt"
		"strings"		
		"strconv"
	)
//-------------------------------------------
func bind_go_func_to_js() {		

		ui.Bind("goFunCalledByJS_mgr", 
			func(goFunc string, js_function string, var1 string, var2 string, var3 string, var4 string, var5 string ) { 				
				switch goFunc {
					case "funCalledByJs_html_is_ready":  funCalledByJs_html_is_ready(  var1,  js_function)  
					case "funCalledByJs_sommaDueNumeri": funCalledByJs_sommaDueNumeri( getInt(var1), getInt(var2), js_function) 						
					default: return;
				}
			} )
			
		//--------------------------------------		
		
		
		return 
}

//---------------------------------------------
func funCalledByJs_html_is_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func bind_go_passToJs_html_is_ready () " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   ", green("html has been loaded"), "   XXXXXXXXXXXX")
	
	begin() 
	
	fmt.Println( green( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" + 
		"xxxxxxxxxxxxxx you can use the tool xxxxxxxxxxxxxxxxxx\n"  + 
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n\n"  ) ) 
	
} // end of bind_go_passToJs_html_is_ready
 
//--------------------------------------
func funCalledByJs_sommaDueNumeri( num1 int, num2 int, js_function string) {
	outS1 :=  strconv.Itoa(  num1+num2 );  
	go_exec_js_function( js_function, outS1 );
}
//---------------------------------------- 

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