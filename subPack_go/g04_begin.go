package subPack_go

	import (
		"fmt"
	)

//g04_begin.go
//------------------------------

func begin() { 	

	if sw_stop { 	
		fmt.Println("UI is ready, but run stopped because of some error")		
		
		go_exec_js_function("UI is ready, but run stopped because of some error","")	
		return
		
	} else {		
		fmt.Println(cyan("\nREADY"), "\n") 
		
		go_exec_js_function("js_go_ready","")	
	}
	
}// end of begin	

//--------------------------------

// end g04_begin.go
//-------------------------------------- 