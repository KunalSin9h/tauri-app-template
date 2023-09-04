# Tauri App Template

Features

1. Vite + React
2. Tailwind Css
3. pnpm
4. Prettier
5. Eslint
6. [Rspc](https://www.rspc.dev/) with Tauri + React Integration
7. GoLang CLI for scripting
8. [Husky](https://typicode.github.io/husky/) for pre commit
9. GitHub Workflow
10. Tauri vscode debug
11. [Shadcn UI](https://ui.shadcn.com/)

## Local Setup

Clone the repository

```bash
git clone https://github.com/KunalSin9h/tauri-app-template
```

Download Dependencies

```bash
pnpm install # for frontend

cd src-tauri
cargo fetch # for backend

cd ../ # come back to home repo
```

Run the development app

```bash
pnpm tauri dev
```

Application will be build by GitHub Action

Tauri [Docs](https://tauri.app/v1/guides/)