# Tauri App Template

[![publish](https://github.com/KunalSin9h/tauri-app-template/actions/workflows/release.yml/badge.svg)](https://github.com/KunalSin9h/tauri-app-template/actions/workflows/release.yml)

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

### Auto Updater Info

Replace the `{username}` and `{repo}` in the `./src-tauri/tauri.conf.json`'s updater endpoint with you username and repo name to get latest version info rom the github releases.

```json
"endpoints": [
    "https://github.com/{username}/{repo}/releases/latest/download/latest.json"
],

```

In this template **Auto Updater** Is enabled by default. So you need to setup `keys` for auto-updater from the [Tauri Docs](https://tauri.app/v1/guides/distribution/updater).

And Update the public key in `./src-tauri/tauri.conf.json` and put `TAURI_PRIVATE_KEY` and `TAURI_KEY_PASSWORD` in GitHub Secrets.

These will be used in build process, like this.

```yaml
- uses: tauri-apps/tauri-action@v0
    env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        TAURI_PRIVATE_KEY: ${{ secrets.TAURI_PRIVATE_KEY }}
        TAURI_KEY_PASSWORD: ${{ secrets.TAURI_KEY_PASSWORD }}
```

> [!WARNING]
> You have to manually trigger `publish` workflow for publishing builds, as the workflow has `workflow_dispatch` trigger.

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
