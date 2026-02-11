const token = localStorage.adminToken;

fetch("http://localhost:8080/admin/users", {
  headers: { Authorization: token }
})
.then(r => r.json())
.then(users => {
  userTable.innerHTML = "";

  users.forEach(u => {
    userTable.innerHTML += `
      <tr>
        <td>${u.name}</td>
        <td>${u.email}</td>
        <td>${u.is_blocked ? "Blocked" : "Active"}</td>
        <td>
          <button
            class="${u.is_blocked ? "btn-primary" : "btn-danger"}"
            onclick="toggleBlock(${u.id})">
            ${u.is_blocked ? "Unblock" : "Block"}
          </button>
        </td>
      </tr>
    `;
  });
});


// User-Block
function toggleBlock(id) {
  fetch(`http://localhost:8080/admin/users/${id}/block`, {
    method: "PUT",
    headers: {
      Authorization: token
    }
  })
  .then(() => location.reload());
}
