if (!localStorage.getItem("adminToken")) {
  window.location.href = "/index.html";
}

function logout() {
  localStorage.removeItem("adminToken");
  window.location.href = "index.html";
}

function apiFetch(url, options = {}) {
  return fetch(url, {
    ...options,
    headers: {
      ...(options.headers || {}),
      Authorization: localStorage.getItem("adminToken"),
    },
  }).then(async res => {
    if (res.status === 401) {
      localStorage.removeItem("adminToken");

      showToast("Session expired. Please login again.", "error");

      setTimeout(() => {
        window.location.replace("/index.html");
      }, 1500);

      throw new Error("Unauthorized");
    }
    return res;
  });
}

