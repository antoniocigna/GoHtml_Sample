package subPack_go

import (  
	"fmt"
    "strings"
	"runtime"
)
//------------------------
//g05_go_exec_js_function.go
//----------------------------------------------------------------
func go_exec_js_function(js_function0 string, inpstr string) {

	var goFunc string 
 	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		goFunc = strings.ReplaceAll(details.Name(), "main.","")
	} else {
		goFunc=""
	}
	js_fun        := strings.Split( (js_function0 + ",,,,") ,",") 	
	js_function   := strings.TrimSpace( js_fun[0] )
	if js_function == ""  { return }
	jsInpFunction := strings.TrimSpace( js_fun[1] )
	
	js_parm:=""
	k1:= strings.Index(js_function, "(") 
	if k1 > 0 {
		js_parm     = strings.ReplaceAll(  js_function[k1+1:], ")","")			
		js_function = strings.TrimSpace(js_function[0:k1] )
	} 	
	
	/*
	This function executes a javascript eval command 
	which must execute a function by passing string constant to it. 
	Should this string contain some new line, e syntax error would occur in eval the statement.
	
	To avoid this kind of error, the string argument (inpstr) of the javascript function (js_function) 
	is forced to be always enclosed in back ticks trasforming it in "template literal".  
	Just in case back ticks and dollars are in the string, they are replaced by " "   	
	*/
	inpstr = strings.ReplaceAll( inpstr, "`", " "   ); 	   	 
	inpstr = strings.ReplaceAll( inpstr, "$", "&dollar;"); 
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, inpstr, js_parm, "js=" + jsInpFunction, "go=" + goFunc ) ; 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function

//----------------------------------------------------------------

// end of g05_go_exec_js_function.go
//----------------------------------------------------