# Sentinel-X Cybersecurity Dashboard

<div align="center">
  <img src="https://img.shields.io/badge/Status-Active-success.svg" alt="Status">
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License">
  <img src="https://img.shields.io/badge/Stack-HTML%20%7C%20TailwindCSS%20%7C%20JS-00E5FF.svg" alt="Tech Stack">
  <img src="https://img.shields.io/badge/Simulation-Golang-00ADD8.svg" alt="Golang">
</div>

<br>

**Sentinel-X (System Surveillance Suite v4.2)** is a state-of-the-art, interactive cybersecurity dashboard designed for threat monitoring, Active Directory pentest visualization, and security auditing. It features a cinematic, cyberpunk-inspired UI built with pure HTML, Tailwind CSS, and Vanilla JavaScript. 

The project includes a fully responsive multi-page dashboard to visualize threats, network topology, and exploit payloads, accompanied by a built-in **Golang penetration simulation** engine.

---

## 🚀 Features

- **🛡️ Main Dashboard**: High-level system overview featuring an aggregate risk score, live terminal feed simulation, and severity distribution metrics.
- **🕸️ Network Topology Map**: Interactive visualizer mapping the relationship between Core Gateways, Active Directory Domain Controllers, and Client Workstations. Adapts dynamically to current threat levels.
- **🐞 Vulnerability Database**: Comprehensive threat tracking interface with live search, severity filtering, and CVE-level drill-downs.
- **💻 Exploitation Results**: Real-time feedback and audit logs of payloads executed against target hosts (e.g., PrintNightmare, Zerologon).
- **📄 Automated PDF Reports**: One-click generation of professional security audit and penetration test reports using `jsPDF`.
- **⚙️ System Settings**: Configuration panel for operator profiles, clearance levels, safe-mode enforcement, and alarm thresholds.
- **🥷 Golang AD Pentest Simulator**: A standalone Go application (`simulation/main.go`) that emulates a realistic Active Directory attack chain (Reconnaissance, LLMNR Poisoning, CVE-2020-1472, and Golden Ticket forgery) straight from the terminal.

## 📂 Project Structure

```text
stitch_sentinel_cybersecurity_dashboard/
├── login_screen/          # Authentication guard and entry point
├── main_dashboard/        # Primary system overview and live metrics
├── network_map/           # Interactive AD network topology
├── vulnerabilities/       # Threat database and CVE filtering
├── exploitation_results/  # Payload execution and audit logs
├── reports/               # Automated PDF report generation module
├── settings/              # Operator configurations and thresholds
├── simulation/            # Golang terminal pentest simulator
│   └── main.go
├── exploits.json          # Core JSON database driving the dashboard metrics
└── README.md              # Project documentation
```

## 🛠️ Technology Stack

- **Frontend Core**: HTML5, Vanilla JavaScript
- **Styling**: [Tailwind CSS](https://tailwindcss.com/) (via CDN) with custom CSS variables and dark-mode optimization
- **Typography & Icons**: Google Fonts (Inter, Space Grotesk) & Material Symbols
- **PDF Generation**: [jsPDF](https://github.com/parallax/jsPDF)
- **Backend/Simulation**: [Go (Golang)](https://go.dev/)

## 🚦 Getting Started

### 1. Launching the Dashboard
Because Sentinel-X uses pure HTML/JS and relies on `fetch()` to load local JSON data (`exploits.json`), running it directly from the file system (`file://`) may cause CORS errors in modern browsers. 

To run it locally:
1. Open a terminal in the root directory.
2. Serve the directory using any local web server. For example:
   - **Python**: `python -m http.server 8000`
   - **Node.js**: `npx serve .`
   - **PHP**: `php -S localhost:8000`
3. Navigate to `http://localhost:8000/login_screen/code.html` in your browser.
4. Default login (if prompted):
   - **Username**: `mahdi` (or any value)
   - **Password**: Any value to bypass the visual login screen.

### 2. Running the Pentest Simulator
To experience the immersive terminal penetration testing simulation:
1. Ensure you have [Go installed](https://go.dev/doc/install) on your system.
2. Navigate to the `simulation` folder:
   ```bash
   cd simulation
   ```
3. Run the Go script:
   ```bash
   go run main.go
   ```
4. Watch the animated Active Directory attack chain unfold in your terminal.

## 🗄️ Data Management (`exploits.json`)

The entire dashboard is fully dynamic and driven by `exploits.json`. 
- **Modifying Data**: You can add, edit, or remove CVEs directly in the JSON file. The dashboard will automatically recalculate risk scores, update the network map statuses, and populate the vulnerability tables based on this file.
- **Empty State**: If `exploits.json` is cleared or missing, the dashboard gracefully falls back to a perfect `0` vulnerability "SECURE" state, hiding all attack graphics and halting warning alerts.

## 📜 License

This project is open-source and available under the [MIT License](LICENSE).
