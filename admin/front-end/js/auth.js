if (!localStorage.getItem("adminToken")) {
  window.location.href = "/index.html";
}

function logout() {
  localStorage.removeItem("adminToken");
  window.location.href = "/index.html";
}
