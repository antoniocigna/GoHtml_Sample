package subPack_go
//--------------------------
//g01_declare.go
//-----------------------------------------

const appName = "gohtml_sample"
const htmlFile= appName + ".html"
const parameter_path_html string  = "/subPack_html_js"
//--------------------------------------
// color.go
//---------------------------------------
/**         ***  COLORS: got from https // www.dolthub.com/blog/2024-02-23-colors-in-golang/ ***
    var Reset   = "\033[0m" 
	var Red     = "\033[31m" 
	var Green   = "\033[32m" 
	var Yellow  = "\033[33m" 
	var Blue    = "\033[34m" 
	var Magenta = "\033[35m" 
	var Cyan    = "\033[36m" 
	var Gray    = "\033[37m" 
	var White   = "\033[97m"
**/

func red(     str1 string) string { return "\033[31m" + str1 +  "\033[0m" }
func green(   str1 string) string { return "\033[32m" + str1 +  "\033[0m" }
func yellow(  str1 string) string { return "\033[33m" + str1 +  "\033[0m" }
//func blue(  str1 string) string { return "\033[34m" + str1 +  "\033[0m" }
func magenta( str1 string) string { return "\033[35m" + str1 +  "\033[0m" }
func cyan(    str1 string) string { return "\033[36m" + str1 +  "\033[0m" }
//func gray(  str1 string) string { return "\033[37m" + str1 +  "\033[0m" }
//func white( str1 string) string { return "\033[97m" + str1 +  "\033[0m" }

//---------------------------------------------------------------------------------
var sw_stop bool = false
var errorMSG = ""; 

var sw_begin_ended = false     
var sw_HTML_ready  = false     
//var apiceInverso = `40`  //  in windows:  tasto Alt + 96 (tastierino numerico)
//---------------------------

//---------------
// end of g01_declare.go
//-----------------------------------------