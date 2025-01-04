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
	console.log("html is ready") 
	
	/*
	goFunCalledByJS_mgr is not a js function, but a link to a go function (see go func 'bind_go_func_to_js()' ) 
		parameters: 1 : the actual go func name to run 
		            2 : the js fuction to be run when the go func has ended
					3,4,5,6,7: the parameters for the go func ( all of them must be passed as strings )  	
	*/
	goFunCalledByJS_mgr( "funCalledByJs_0_html_is_ready", "", msg1, "2","3","4","5" ) ; 
	
}

//-------------------------------------		

function js_go_ready() {
	// this function is run by go
	console.log("go is ready"); 
	
	document.getElementById("id_start001").style.display = "none";
	document.getElementById("id_myPage01").style.display = "flex"; 
	
} // end of js_go_ready
//-------------------------------
//=====================================
function onclick_sum2Num() {
	var var1 = document.getElementById("inp11").value;
	var var2 = document.getElementById("inp12").value;	
	/*
	goFunCalledByJS_mgr is not a js function, but a link to a go function (see go func 'bind_go_func_to_js()' ) 
		parameters: 1 : the actual go func name to run 
		            2 : the js fuction to be run when the go func has ended
					3,4,5,6,7: the parameters for the go func ( all of them must be passed as strings )  
	*/
	goFunCalledByJS_mgr( "funCalledByJs_1_sum2numbers", "js_go_1_returnSum", ""+var1, ""+var2 , "3","4","5");  
}
//----------------------------------------------------
function js_go_1_returnSum(go_sommaVal, js_parm, jsFunc,goFunc) { 
	// this function is run by go,  its name is the second parameter of goFunCalledByJS_mgr    
	document.getElementById("out1").innerHTML = "GO = " + go_sommaVal;	

} // end of js_go_ritornaSomma

//--------------------------------

function onclick_multiply2Num() {
	var var1 = document.getElementById("inp21").value;
	var var2 = document.getElementById("inp22").value;	
	/*
	goFunCalledByJS_mgr is not a js function, but a link to a go function (see go func 'bind_go_func_to_js()' ) 
		parameters: 1 : the actual go func name to run 
		            2 : the js fuction to be run when the go func has ended
					3,4,5,6,7: the parameters for the go func ( all of them must be passed as strings )  
	*/
	goFunCalledByJS_mgr( "funCalledByJs_2_multiply2numbers", "js_go_2_returnProd", ""+var1, ""+var2 , "3","4","5");  
}
//----------------------------------------------------
function js_go_2_returnProd(go_prodVal, js_parm, jsFunc,goFunc) { 
	// this function is run by go,  its name is the second parameter of goFunCalledByJS_mgr    
	document.getElementById("out2").innerHTML = "GO = " + go_prodVal;	

} // end of js_go_ritornaSomma
//======================================================