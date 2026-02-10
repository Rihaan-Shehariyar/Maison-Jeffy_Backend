// DASHBOARD DATA
apiFetch("http://localhost:8080/admin/dashboard")
  .then(res => res.json())
  .then(data => {

    // KPI COUNTS
    ordersCount.innerText = data.totalOrders;
    pendingCount.innerText = data.pendingOrders;
    successCount.innerText = data.successfulOrders;
    usersCount.innerText = data.totalUsers;
    revenueCount.innerText = data.totalRevenue.toFixed(2);

    // CHARTS
    renderOrdersChart(data.ordersChart);
    renderRevenueChart(data.revenueChart);
  })
  .catch(err => showToast(err.message, "error"));


// ORDERS CHART
function renderOrdersChart(chartData) {
  new Chart(document.getElementById("ordersChart"), {
    type: "line",
    data: {
      labels: chartData.labels,
      datasets: [{
        label: "Orders",
        data: chartData.values,
        borderColor: "#2563eb",
        fill: false,
        tension: 0.3
      }]
    }
  });
}

// REVENUE CHART
function renderRevenueChart(chartData) {
  new Chart(document.getElementById("revenueChart"), {
    type: "bar",
    data: {
      labels: chartData.labels,
      datasets: [{
        label: "Revenue",
        data: chartData.values,
        backgroundColor: "#16a34a"
      }]
    }
  });
}
