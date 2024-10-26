# VORTEX
VORTEX, a world-class VSaaS solution

## Project Structure

> [!NOTE]
> **Single Codebase philosophy** inspired by [Flutter](https://flutter.dev/multi-platform).
> - [Flutter](https://flutter.dev/multi-platform) allows you to build apps for mobile, web, desktop, and embedded devices  
    — all from a single codebase.
> - Demo: [FlutLab.io - Flutter IDE online](https://flutlab.io)

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

## Why Single Codebase aka Monorepo
1. Ruddy 老師：「[專案開始之初，首重看見全貌。](https://ruddyblog.wordpress.com/2017/10/27/看見全貌/)」
2. 更容易使用 Docker 建立一整套完整的開發環境、獨立的測試環境
3. 更容易在本機端做到**完整的端到端測試**，縮短驗證迴路
4. 有了測試保護，重構和維護程式碼的成本降低
5. 程式碼即文件：用會動的程式理解整個複雜系統的運作流程
6. 促進跨職能交流、協作

> [!WARNING]
> #### 懶惰是工程師的美德
> 採用  Multi-repo 的結構，在開發的時候經常需要在多個專案視窗之間切換，
> 如果你很勤勞覺得這樣不麻煩，  
> 那你可以想不開挑戰看看小小小的微服務架構，
> 同時開 3、4 個以上的微服務專案視窗、  
> 記得喔，你還常常要開至少 1 個文件的視窗、查資料的十幾個瀏覽器分頁，再開個 ChatGPT，  
> 試問這需要多大、多少個螢幕才夠用？

> [!TIP]
> #### 動起來！現在就開始 Monorepo！   
> 創建一個 `vortex-monorepo` 檔案夾，把那些有的沒的 repo 通通丟進去，  
> 用 IDE 開啟 `/path/to/vortex-monorepo` 這個路徑，Wala 🎉

## Reference
Open Source Monorepo:
> - [Kubernetes](https://github.com/kubernetes/kubernetes)
- [Angular](https://github.com/angular/angular)
- [Elastic](https://github.com/elastic/elasticsearch)
- [NestJS](https://github.com/nestjs/nest)
- [Apollo](https://github.com/apollographql/apollo-server)
- [Expo](https://github.com/expo/expo)
- [React](https://github.com/facebook/react)
