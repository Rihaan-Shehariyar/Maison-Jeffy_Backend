console.log("login.js loaded");

function login(){
  console.log("login clicked");
 const email = document.getElementById("email").value;
 const password = document.getElementById("password").value;


 fetch("http://localhost:8080/admin/login",{
  method:"POST",
  headers:{ "Content-type":"application/json"},
  body : JSON.stringify({email,password})
})
.then(res=>res.json())
.then(data=>{
  if (data.token){
  localStorage.setItem("adminToken",data.token);
  window.location.href = "dashboard.html"
}else{
 alert("Login Failed")
} 
})
}