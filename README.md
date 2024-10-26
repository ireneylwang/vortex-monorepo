# VORTEX
VORTEX, a world-class VSaaS solution

## Project Structure

```python
.
├── android/       # Android app
├── backend/       # Backend modules, APIs, serverless, cron jobs 
├── e2e/           # [Frontend --> Backend --> Infra --> IoT] turly end-to-end tests
├── frontend/      # Frontend web app
├── infra/         # IaC, CI/CD
├── ios/           # iOS app
├── iot/           # Firmware, IoT
└── README.md
```

> [!NOTE]
> **Single Codebase philosophy** inspired by [Flutter](https://flutter.dev/multi-platform).
> - [Flutter](https://flutter.dev/multi-platform) allows you to build apps for mobile, web, desktop, and embedded devices  
    — all from a single codebase.
> - Demo: [FlutLab.io - Flutter IDE online](https://flutlab.io)

## Why Monorepo
1. 更容易使用 Docker 建立一整套完整的開發環境、獨立的測試環境
2. 更容易在本機端做到**完整的端到端測試**，縮短驗證迴路
3. 有了測試保護，重構和維護程式碼的成本降低
4. 促進跨職能交流、協作
5. 不用在多個專案視窗之間切換

## Reference
Open Source Monorepo:
> - [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Angular](https://github.com/angular/angular)
- [Elastic](https://github.com/elastic/elasticsearch)
- [NestJS](https://github.com/nestjs/nest)
- [Apollo](https://github.com/apollographql/apollo-server)
- [Expo](https://github.com/expo/expo)
- [React](https://github.com/facebook/react)
