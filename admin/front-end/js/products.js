let editingId = null;
let productsCache = [];


// product Fetching

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
        </tr>`;
    });
  })
  .catch(err => showToast(err.message, "error"));

// Product Editing

function editProductById(id) {
  const p = productsCache.find(x => x.id === id);
  if (!p) return;

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


// Product Deletion

function deleteProduct(id) {
  if (!confirm("Delete product?")) return;

  apiFetch(`http://localhost:8080/admin/products/${id}`, {
    method: "DELETE"
  })
    .then(() => {
      showToast("Product deleted");
      setTimeout(() => location.reload(), 500);
    })
    .catch(err => showToast(err.message, "error"));
}


// Sku auto-generate

function generateSku() {
  const name = productName.value;

  if (!name) {
    showToast("Enter product name first", "warning");
    return;
  }

  const prefix = "JC";
  const cleanName = name
    .toUpperCase()
    .replace(/[^A-Z0-9]/g, "")
    .slice(0, 6);

  const random = Math.floor(1000 + Math.random() * 9000);
  sku.value = `${prefix}-${cleanName}-${random}`;
}
