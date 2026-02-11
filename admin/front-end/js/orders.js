const token = localStorage.adminToken;

fetch("http://localhost:8080/admin/orders", {
  headers: { Authorization: token }
})
.then(r => r.json())
.then(orders => {
  orderTable.innerHTML = "";

  orders.forEach(o => {
    orderTable.innerHTML += `
      <tr>
        <td>#${o.id}</td>
        <td>
          ${o.user?.name || "Unknown"}<br>
          <small>${o.user?.email || ""}</small>
        </td>
        <td>₹${o.total_amount}</td>
        <td>
          <select id="status-${o.id}">
            ${statusOption("placed", o.status)}
            ${statusOption("shipped", o.status)}
            ${statusOption("delivered", o.status)}
            ${statusOption("cancelled", o.status)}
          </select>
        </td>
        <td>
          <button onclick="updateStatus(${o.id})">Save</button>
        </td>
      </tr>
    `;
  });
});


// Status Updation
function statusOption(value, current) {
  return `<option value="${value}" ${value === current ? "selected" : ""}>
            ${value.toUpperCase()}
          </option>`;
}

function updateStatus(id) {
  const status = document.getElementById(`status-${id}`).value;

  fetch(`http://localhost:8080/admin/orders/${id}/status`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
      Authorization: token
    },
    body: JSON.stringify({ status })
  })
  .then(() => alert("Order status updated"));
}
