function showToast(message, type = "success") {
  let bg = "#16a34a";

  if (type === "error") bg = "#dc2626";
  if (type === "warning") bg = "#f59e0b";

  Toastify({
    text: message,
    duration: 3000,
    gravity: "top",
    position: "right",
    close: true,
    backgroundColor: bg,
  }).showToast();
}
