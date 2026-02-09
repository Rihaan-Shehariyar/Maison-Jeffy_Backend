fetch("http://localhost:8080/admin/products", {
  headers: { Authorization: localStorage.getItem("adminToken") }
})
.then(res => res.json())
.then(d => document.getElementById("pCount").innerText = d.length);

fetch("http://localhost:8080/admin/users", {
  headers: { Authorization: localStorage.getItem("adminToken") }
})
.then(res => res.json())
.then(d => document.getElementById("uCount").innerText = d.length);

fetch("http://localhost:8080/admin/orders", {
  headers: { Authorization: localStorage.getItem("adminToken") }
})
.then(res => res.json())
.then(d => document.getElementById("oCount").innerText = d.length);
