package subPack_go

	import (
		"fmt"	
		"strconv"
	)
//-------------------------------------------
func bind_go_func_to_js() {		

		ui.Bind("goFunCalledByJS_mgr", 
			func(goFunc string, js_function string, var1 string, var2 string, var3 string, var4 string, var5 string ) { 				
				switch goFunc {
					case "funCalledByJs_0_html_is_ready":    funCalledByJs_0_html_is_ready(  var1,  js_function)  
					case "funCalledByJs_1_sum2numbers":      funCalledByJs_1_sum2numbers( getInt(var1), getInt(var2), js_function) 	
					case "funCalledByJs_2_multiply2numbers": funCalledByJs_2_multiply2numbers( getInt(var1), getInt(var2), js_function) 		
					default: 
						fmt.Println( red("error in bind_go_func_to_js"), " the string '", green(goFunc) , "' is unknown, It cannot be to any go func")    
						return;
				}
			} )
			
		//--------------------------------------		
		
		
		return 
}

//---------------------------------------------
func funCalledByJs_0_html_is_ready( msg1 string,  js_function string) {
	fmt.Println("\n", "go func bind_go_passToJs_html_is_ready () " , "\n\t msg from html: ", msg1 )  
	
	fmt.Println("XXXXXXXXXX   ", green("html has been loaded"), "   XXXXXXXXXXXX")
	
	begin() 
	
	fmt.Println( green( "\n\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n" + 
		"xxxxxxxxxxxxxx you can use the tool xxxxxxxxxxxxxxxxxx\n"  + 
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx\n\n"  ) ) 
	
} // end of bind_go_passToJs_html_is_ready
 
//--------------------------------------

func funCalledByJs_1_sum2numbers( num1 int, num2 int, js_function string) {
	outS1 :=  strconv.Itoa(  num1+num2 );  
	go_exec_js_function( js_function, outS1 );
}

//--------------------------------------

func funCalledByJs_2_multiply2numbers( num1 int, num2 int, js_function string) {
	outS1 :=  strconv.Itoa(  num1*num2 );  
	go_exec_js_function( js_function, outS1 );
}
//---------------------------------------- 
