# daf-wails-template

```{=html}
<p align="center">
```
`<b>`{=html}A Production-Grade Wails v2 Desktop Starter
Template`</b>`{=html}`<br/>`{=html} Built with Go 1.23 · Wails v2.11.0 ·
React 18 · Mantine 7
```{=html}
</p>
```
```{=html}
<p align="center">
```
`<img src="https://img.shields.io/badge/Wails-v2.11.0-blue" />`{=html}
`<img src="https://img.shields.io/badge/Go-1.23-00ADD8" />`{=html}
`<img src="https://img.shields.io/badge/React-18-61DAFB" />`{=html}
`<img src="https://img.shields.io/badge/Mantine-7-339af0" />`{=html}
`<img src="https://img.shields.io/badge/License-MIT-green" />`{=html}
```{=html}
</p>
```

------------------------------------------------------------------------

## 🚀 What Is This?

**daf-wails-template** is a production-ready Wails desktop application
template designed for real-world software delivery.

It eliminates repetitive infrastructure work so you can focus on
building your product.

------------------------------------------------------------------------

## ⭐ Why This Template Exists

The official Wails template is great for demos.

But production apps always need:

-   Logging system
-   Build metadata injection
-   System configuration
-   User config persistence
-   Internationalization
-   Clean architecture
-   Minimal UI foundation
-   Custom titlebar support
-   Health / About system pages

This template standardizes those patterns into a reusable architecture.

------------------------------------------------------------------------

## 🆚 Official Template vs daf-wails-template

  Feature                    Official Template   daf-wails-template
  -------------------------- ------------------- --------------------
  Structured Logger          ❌                  ✅
  Build Metadata Injection   ❌                  ✅
  System Config Layer        ❌                  ✅
  User Config Persistence    ❌                  ✅
  i18n Ready                 ❌                  ✅
  Modular Backend Core       ❌                  ✅
  Production UI Layout       ❌                  ✅
  Custom Titlebar Ready      ❌                  ✅

------------------------------------------------------------------------

## ✨ Features

### 🧱 Backend (Go)

-   Go 1.23
-   Wails v2.11.0
-   Structured logging system
-   Build-time version injection (`ldflags`)
-   System metadata management
-   User configuration persistence
-   Modular core architecture
-   Cross-platform ready

### 🎨 Frontend (React + Mantine)

-   React 18
-   Mantine 7
-   AppShell layout
-   Sidebar navigation
-   Health Check page
-   About / System Info page
-   i18n (EN / CN)
-   Clean desktop-focused design

------------------------------------------------------------------------

## 🖼 Screenshot

> Replace this with your actual screenshot

docs/screenshot.png

------------------------------------------------------------------------

## 📦 Tech Stack

Backend: - Go 1.23 - Wails v2.11.0

Frontend: - React 18 - Mantine 7 - React Router - i18next - Vite

Package Manager: - pnpm (recommended)

------------------------------------------------------------------------

## 📂 Project Structure

backend/ core/ config/ logger/ system/ app/

frontend/ src/ layout/ pages/ i18n/ api/

build/

------------------------------------------------------------------------

## ⚡ Quick Start

### Install Wails CLI

go install github.com/wailsapp/wails/v2/cmd/wails@latest

### Install dependencies

pnpm install

### Run development

wails dev

### Build production

wails build

------------------------------------------------------------------------

## 🧬 Inject Build Metadata

wails build -ldflags " -X main.Version=1.0.0 -X main.BuildTime=\$(date
-u '+%Y-%m-%dT%H:%M:%SZ') "

------------------------------------------------------------------------

## 🎯 Use Cases

-   SaaS desktop companion apps
-   Enterprise internal tools
-   Monitoring agents
-   System utilities
-   Admin dashboards

------------------------------------------------------------------------

## 🛣 Roadmap

-   [ ] Auto updater module
-   [ ] Plugin system
-   [ ] License management
-   [ ] Telemetry integration
-   [ ] CLI scaffolding generator

------------------------------------------------------------------------

## 🤝 Contributing

PRs are welcome.

If you build something useful on top of this template, feel free to
share.

------------------------------------------------------------------------

## 📜 License

MIT License

------------------------------------------------------------------------

## ⭐ Support

If this template saves you time, please give it a Star.

It helps grow the Wails ecosystem.
