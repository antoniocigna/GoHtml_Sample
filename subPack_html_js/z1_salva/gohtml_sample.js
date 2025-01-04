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