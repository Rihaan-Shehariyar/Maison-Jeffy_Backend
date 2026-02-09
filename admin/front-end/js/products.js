let editingId = null;
const token = localStorage.adminToken;

// LOAD PRODUCTS
fetch("http://localhost:8080/admin/products", {
  headers: { Authorization: token }
})
.then(r => r.json())
.then(products => {
  productTable.innerHTML = "";
  products.forEach(p => {
    productTable.innerHTML += `
      <tr>
        <td>
          <img src="/${p.image_url}" class="thumb">
        </td>
        <td>${p.name}</td>
        <td>${p.price}</td>
        <td>${p.stock}</td>
        <td>
          <button onclick='editProduct(${JSON.stringify(p)})'>Edit</button>
          <button class="btn-danger" onclick="deleteProduct(${p.id})">Delete</button>
        </td>
      </tr>`;
  });
});

// MODAL CONTROL
function openModal() {
  editingId = null;
  modalTitle.innerText = "Add Product";
  modal.style.display = "block";
}

function closeModal() {
  modal.style.display = "none";
}

// IMAGE PREVIEW
function previewImage() {
  const file = image.files[0];
  preview.src = URL.createObjectURL(file);
}

// EDIT
function editProduct(p) {
  editingId = p.id;
  modalTitle.innerText = "Edit Product";

  name.value = p.name;
  price.value = p.price;
  stock.value = p.stock;
  category.value = p.category;
  description.value = p.description;
  preview.src = "/" + p.image_url;

  modal.style.display = "block";
}

// SAVE (CREATE / UPDATE)
function saveProduct() {
  const formData = new FormData();
  formData.append("name", name.value);
  formData.append("price", price.value);
  formData.append("stock", stock.value);
  formData.append("category", category.value);
  formData.append("description", description.value);

  if (image.files[0]) {
    formData.append("image", image.files[0]);
  }

  let url = "http://localhost:8080/admin/products";
  let method = "POST";

  if (editingId) {
    url = `http://localhost:8080/admin/products/${editingId}`;
    method = "PUT";
  }

  fetch(url, {
    method,
    headers: { Authorization: token },
    body: formData
  }).then(() => location.reload());
}

// DELETE
function deleteProduct(id) {
  if (!confirm("Delete product?")) return;

  fetch(`http://localhost:8080/admin/products/${id}`, {
    method: "DELETE",
    headers: { Authorization: token }
  }).then(() => location.reload());
}
