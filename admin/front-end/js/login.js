function login() {
  fetch("http://localhost:8080/admin/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      email: email.value,
      password: password.value
    })
  })
  .then(async res => {
    const data = await res.json();

    if (!res.ok) {
      throw new Error(data.error || "Login failed");
    }

    return data;
  })
  .then(data => {
    localStorage.setItem("adminToken", data.token);

    showToast("Login successful", "success");

    setTimeout(() => {
      window.location.replace("dashboard.html");
    }, 800);
  })
  .catch(err => {
    showToast(err.message, "error");
  });
}
