"use strict";
/*  
gohtml_sample - Antonio Cigna 2025
license MIT: you can share and modify the software, but you must include the license file 
*/
/* jshint strict: true */
/* jshint esversion: 6 */
/* jshint undef: true, unused: true */
//-----------------------------------------------
function js_call_go_on_html_body_load() {  // called by  html page body onload
	var msg1 = "html loaded"; 
	console.log("html function js_call_go: ", msg1, "\n") 
	
	go_passToJs_html_is_ready(msg1,  "") ;  //  "js_go_go_is_ready"); 

}

//-------------------------------------		

function js_go_ready( prevRun00) {
	console.log("************************* js_go_ready(" + prevRun00.trim() + ")" ); 
	
	document.getElementById("id_start001").style.display = "none";
	document.getElementById("id_myPage01").style.display = "flex"; 
	
} // end of js_go_ready
//-------------------------------
//=====================================
function sommaORIG() {
	var var1 = document.getElementById("inp1").value;
	var var2 = document.getElementById("inp2").value;
	var addedValue = var1*1 + var2*1; 
	document.getElementById("out1").innerHTML = addedValue;
}

function somma() {
	var var1 = document.getElementById("inp1").value;
	var var2 = document.getElementById("inp2").value;
	go_passToJs_sommaDueNumeri(""+var1, ""+var2, "js_go_ritornaSomma"); // ask 'go' to give sum  by js_... function 
}
//----------------------------------------------------
function js_go_ritornaSomma(go_sommaVal, js_parm, jsFunc,goFunc) { 
	document.getElementById("out1").innerHTML = "GO = " + go_sommaVal;	

} // end of js_go_showPrefixWordList

//=============

/*****

// ijn go 
		
func get_all_bindsRESTO() {		
		
		//----------------------------
		ui.Bind("go_passToJs_wordList", func( isChange bool, s_fromWord string, s_numWords string, sel_level string, 
															sel_extrRow string, sel_toBeLearned string,  js_function string) { 		
				bind_go_passToJs_wordList( isChange, getInt(s_fromWord), getInt(s_numWords), sel_level, sel_extrRow, sel_toBeLearned,  js_function)  })
		//--------------------------------------		
		ui.Bind("go_passToJs_prefixWordList", func(s_numWords string, wordPrefix string, js_function string) {
				bind_go_passToJs_prefixWordList( getInt(s_numWords), wordPrefix, js_function) } )
		//--------------------------------------		
}		

func bind_go_passToJs_prefixWordList( numWords int, wordPrefix string, js_function string) {
	
	bind_go_passToJs_betweenWordList( numWords, wordPrefix, js_function) 
			
} // end of bind_go_passToJs_prefixWordList

//-----------------------------------------

func bind_go_passToJs_betweenWordList( maxNumWords int, fromWordPref string, js_function string) {
	
	var onlyThisLevel string = "any" ; // "A0"  // questo deve arrivare da parametro  
	var outS1 string; 	
	//----------------------------------------------------------------------
	fromWord   := strings.ToLower(strings.TrimSpace( fromWordPref));  
	lenFrom := len(fromWord) 	
	sw_oneWord := false 
	
	//-----------
	go_exec_js_function( js_function, outS1 ); 		
	
			
} // end of bind_go_passToJs_betweenWordList
//------------------

	
//----------------------------------------------------------------
func go_exec_js_function(js_function0 string, inpstr string) {
	
	evalStr := fmt.Sprintf( "%s(`%s`,`%s`,`%s`,`%s`);",  js_function, inpstr, js_parm, "js=" + jsInpFunction, "go=" + goFunc ) ; 
	
	ui.Eval(evalStr)
	
} // end of go_exec_js_function
//---------------------------------------------------

***/