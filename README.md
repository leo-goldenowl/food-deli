# Learn Go
The description will be updated later...

## Project Structure

```sh
  └── workspaces
      ├── backend
      └── frontend
```
- workspaces
  - backend: contains all main server.
  - frontend: contains all main web app.

## Set Up Project
- [learn-go-api](workspaces/backend/api-gateway//README.md)
- [learn-go-web-app](workspaces/frontend/README.md)

## Git Branching
We use Git Flow process which means we will mainly work on `web_feature/web_fix/web_enhance/web_hotfix...` branch (if server, `server_feature/server_fix/server_enhance/server_hotfix...` branch). All PRs will be created and made against the `develop` branch. The `main` branch will only be used when we release a new update to production.

Example:  
- web_feature/integrate-home-ui
- server_feature/create-user-api

## Git Committing
When we commit, please fill in message follow rule `web: <msg>` (if server, `server: <msg>`)

Example:
- web: fix show password when filling password
- server: secure password with bcrypt

## Notice
- Please read README files carefully, and follow all rules in the project <3.
