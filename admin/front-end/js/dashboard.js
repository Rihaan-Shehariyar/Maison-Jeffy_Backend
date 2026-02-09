const token = localStorage.adminToken;

fetch("http://localhost:8080/admin/dashboard", {
  headers: { Authorization: token }
})
.then(r => r.json())
.then(data => {
  // KPI
  ordersCount.innerText = data.total_orders;
  usersCount.innerText = data.total_users;
  revenueCount.innerText = data.total_revenue.toFixed(2);

  // Prepare chart data
  const labels = data.daily_stats.map(d => d.date);
  const orderCounts = data.daily_stats.map(d => d.count);
  const revenueData = data.daily_stats.map(d => d.sum);

  // Orders Chart
  new Chart(document.getElementById("ordersChart"), {
    type: "line",
    data: {
      labels,
      datasets: [{
        label: "Orders",
        data: orderCounts,
        borderColor: "#2563eb",
        tension: 0.4
      }]
    }
  });

  // Revenue Chart
  new Chart(document.getElementById("revenueChart"), {
    type: "bar",
    data: {
      labels,
      datasets: [{
        label: "Revenue",
        data: revenueData,
        backgroundColor: "#16a34a"
      }]
    }
  });
});
