# 🚀 GLog

**GLog** is a lightweight real-time log analysis engine that helps developers detect anomalies and hidden bugs in microservices.

---

## 🧠 Overview

Modern applications generate massive logs, making debugging difficult and time-consuming.

BugSense simplifies this by:

* Processing logs in real-time
* Detecting anomalies instantly
* Providing a simple and scalable architecture

---

## ⚙️ How It Works

```
User → API → Buffer → Worker → Analyzer → 🚨 Detection
```

### Explanation:

1. **User sends logs** via API
2. Logs are placed into a **buffer (queue)**
3. **Workers** process logs asynchronously
4. The **analyzer** evaluates patterns
5. If abnormal → system triggers an anomaly alert

---

## 🧪 Example

### Input

```json
{
  "service": "auth",
  "message": "timeout error"
}
```

### Output

```
🚨 ANOMALY DETECTED
Score: 12
Service: auth
```

---

## 🧠 Detection Logic

BugSense uses a simple scoring system:

| Pattern | Score |
| ------- | ----- |
| error   | +5    |
| timeout | +7    |
| fail    | +4    |

If total score exceeds a threshold → anomaly detected.

---

## 🌐 Features

* ⚡ Real-time log processing
* 🔁 Buffer-based architecture (queue system)
* 🧵 Concurrent workers (high performance)
* 🧠 Pattern-based anomaly detection
* 🎯 Simple and extensible design

---

## 🏗️ Architecture

```
[ Client ]
     ↓
[ API Layer ]
     ↓
[ Buffer / Queue ]
     ↓
[ Worker Pool ]
     ↓
[ Analyzer Engine ]
     ↓
[ Result / Alert ]
```

---

## 📁 Project Structure

```
GLog/
├── backend/
│   ├── main.go
│   ├── api.go
│   ├── buffer.go
│   └── analyzer.go
└── README.md
```

---

## 🚀 Getting Started

```bash
https://github.com/fenq26/GLog.git
cd Glog
go run .
```

---

## 🔮 Future Improvements

* 📊 Real-time dashboard
* 🤖 Machine learning anomaly detection
* 🔐 API key authentication system
* 🌍 Distributed tracing support

---

## 🤝 Contributing

Contributions are welcome!
Feel free to open issues or submit pull requests.

---

## 💡 Vision

To make debugging microservices faster, simpler, and more accessible for every developer.
