const token = localStorage.getItem("adminToken")

fetch("http://localhost:8080/admin/products",{
  headers : {
 "Authorization" : token
}
})
.then(res=>res.json)
.then(data=>{
 const tbody = document.getElementById("products");
 tbody.innerHTML = "";

 data.foreach(p=>{
 
 tbody.innerHTML += `
 <tr>
  <td>${p.name}</td>
  <td>${p.price}</td>
  <td>${p.stock}</td>
  </tr>
`;

})
})


document.getElementById("productForm").addEventListener("submit", function(e) {
  e.preventDefault();

  const formData = new FormData(this);

  fetch("http://localhost:8081/admin/products", {
    method: "POST",
    headers: {
      "Authorization": localStorage.getItem("adminToken")
    },
    body: formData
  })
  .then(res => res.json())
  .then(() => {
    alert("Product created");
    window.location.reload();
  });
});
