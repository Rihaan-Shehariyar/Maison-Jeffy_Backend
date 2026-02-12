

let editingId = null;
let productsCache = [];


  //  DOM ELEMENTS


const modal = document.getElementById("productModal");
const modalTitle = document.getElementById("modalTitle");

const productTable = document.getElementById("productTable");

const productName = document.getElementById("productName"); 
const price = document.getElementById("price");
const stock = document.getElementById("stock");
const category = document.getElementById("category");
const sku = document.getElementById("sku");
const description = document.getElementById("description");
const image = document.getElementById("image");
const preview = document.getElementById("preview");


  //  LOAD PRODUCTS


apiFetch("http://localhost:8080/admin/products")
  .then(r => r.json())
  .then(products => {

    productsCache = products;
    productTable.innerHTML = "";

    products.forEach(p => {

      productTable.innerHTML += `
        <tr>
          <td><img src="/${p.image_url}" class="thumb"></td>
          <td>${p.name}</td>
          <td>${p.price}</td>
          <td>${p.stock}</td>
          <td>${p.sku}</td>
          <td>
            <button onclick="editProductById(${p.id})">Edit</button>
            <button class="btn-danger" onclick="deleteProduct(${p.id})">Delete</button>
          </td>
        </tr>
      `;
    });

  })
  .catch(err => showToast(err.message, "error"));



  //  OPEN ADD MODAL


function openModal() {

  editingId = null;
  modalTitle.innerText = "Add Product";

  resetForm();

  modal.style.display = "block";
}


  //  CLOSE MODAL


function closeModal() {

  editingId = null;
  modal.style.display = "none";

  resetForm();

  modalTitle.innerText = "Add Product";
}


  //  RESET FORM


function resetForm(){

  productName.value = "";
  price.value = "";
  stock.value = "";
  category.value = "";
  sku.value = "";
  description.value = "";

  if(preview) preview.src = "";
}



// edit Product

function editProductById(id){

  const p = productsCache.find(x => x.id === id);
  if(!p) return;

  editingId = p.id;
  modalTitle.innerText = "Edit Product";

  productName.value = p.name;
  price.value = p.price;
  stock.value = p.stock;
  category.value = p.category;
  sku.value = p.sku;
  description.value = p.description;
  preview.src = "/" + p.image_url;

  modal.style.display = "block";
}



  //  SAVE PRODUCT 


function saveProduct(){

  const formData = new FormData();

  formData.append("name", productName.value);
  formData.append("price", price.value);
  formData.append("stock", stock.value);
  formData.append("category", category.value);
  formData.append("description", description.value);
  formData.append("sku", sku.value);

  if(image && image.files[0]){
    formData.append("image", image.files[0]);
  }

  let url = "http://localhost:8080/admin/products";
  let method = "POST";

  if(editingId){
    url = `http://localhost:8080/admin/products/${editingId}`;
    method = "PUT";
  }

  apiFetch(url,{
    method,
    body: formData
  })
  .then(()=>{
    showToast("Saved successfully","success");
    setTimeout(()=>location.reload(),600);
  })
  .catch(err=>showToast(err.message,"error"));
}



  //  DELETE PRODUCT


function deleteProduct(id){

  if(!confirm("Delete product?")) return;

  apiFetch(`http://localhost:8080/admin/products/${id}`,{
    method:"DELETE"
  })
  .then(()=>{
    showToast("Product deleted","success");
    setTimeout(()=>location.reload(),500);
  })
  .catch(err=>showToast(err.message,"error"));
}



  //  IMAGE PREVIEW


if(image){
  image.addEventListener("change", ()=>{
    const file = image.files[0];
    if(file){
      preview.src = URL.createObjectURL(file);
    }
  });
}



  //  SKU AUTO GENERATE

function generateSku(){

  const name = productName.value;

  if(!name){
    showToast("Enter product name first","warning");
    return;
  }

  const prefix = "JC";
  const cleanName = name
    .toUpperCase()
    .replace(/[^A-Z0-9]/g,"")
    .slice(0,6);

  const random = Math.floor(1000 + Math.random()*9000);

  sku.value = `${prefix}-${cleanName}-${random}`;
}
