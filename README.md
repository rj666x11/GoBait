# GoBait 🪤

**GoBait** is a lightweight, single-binary fake login page server built in Go. It allows red teamers and penetration testers to quickly deploy credential harvesting traps during authorized engagements.

---

## 🔥 Features

- ⚡ Instant deployment with zero dependencies
- 🎭 Fake login templates (e.g., Office 365)
- 📜 Logs username, password, IP, and User-Agent
- 📦 All-in-one binary using `embed.FS`
- 📩 Optional redirect to legitimate site after form submission
- 🧪 CLI-configurable for stealthy use

---

## 🚀 Usage

### 🔧 Command-line Options

```bash
go run main.go --port=8080 --log=logs/credentials.csv --template=login.html

| Flag         | Description                                |
| ------------ | ------------------------------------------ |
| `--port`     | Port to run the HTTP server on             |
| `--log`      | Path to the credential log CSV file        |
| `--template` | Template file to serve (from `/templates`) |


🔐 Captured Logs
Each form submission is saved to a log file in CSV format:

timestamp,username,password,ip_address,user_agent
2024-05-03T12:34:56Z,test@example.com,SuperSecret123,192.168.1.10,Mozilla/5.0 ...


🛠️ File Structure

gobait/
├── main.go
├── go.mod
├── templates/
│   └── login.html
└── logs/
    └── credentials.csv

🧪 Example Login Page
The default template mimics a basic Office 365 login page. You can create and use your own templates by placing them in the templates/ directory.

🛡️ Legal & Ethical Use
This tool is intended only for use in authorized penetration testing and red team assessments.
Unauthorized use of this tool to collect credentials or impersonate services may violate laws and ethical guidelines.

📜 License
MIT License — see LICENSE

✨ TODOs / Future Features
 Webhook support for alerting (Slack/Discord)

 Email alert when trap is triggered

 Multiple login page templates

 Tor .onion support for hidden services

