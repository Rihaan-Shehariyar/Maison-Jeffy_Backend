function login() {
  fetch("http://localhost:8080/admin/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email: email.value,
      password: password.value
    })
  })
  .then(res => res.json())
  .then(data => {
    if (data.token) {
      localStorage.setItem("adminToken", data.token);
      window.location.href = "dashboard.html";
    } else {
      alert("Invalid credentials");
    }
  });
}
