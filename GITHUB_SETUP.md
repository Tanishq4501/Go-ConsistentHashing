# GitHub Repository Setup Guide

Follow these steps to push your Go Consistent Hashing project to GitHub.

---

## Step 1: Create a GitHub Repository

1. Go to [GitHub](https://github.com) and sign in
2. Click the **"+"** icon in the top right → **"New repository"**
3. Fill in the repository details:
   - **Repository name**: `go-hash` (or your preferred name)
   - **Description**: `Comprehensive Go implementation of consistent hashing with three progressive approaches`
   - **Visibility**: Public (recommended for portfolio) or Private
   - **DO NOT** initialize with README, .gitignore, or license (we already have them)
4. Click **"Create repository"**

---

## Step 2: Create Essential Files

### .gitignore

Run this in PowerShell from your project directory:

```powershell
cd d:\GoLang\go-hash
```

Create `.gitignore`:

```powershell
@"
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out

# Go workspace file
go.work

# Dependency directories
vendor/

# IDE specific files
.idea/
.vscode/
*.swp
*.swo
*~

# OS specific files
.DS_Store
Thumbs.db

# Build artifacts
bin/
build/
dist/

# Temporary files
tmp/
temp/
"@ | Out-File -FilePath .gitignore -Encoding utf8
```

### LICENSE (MIT License)

```powershell
@"
MIT License

Copyright (c) 2026 Tanishq4501

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"@ | Out-File -FilePath LICENSE -Encoding utf8
```

---

## Step 3: Initialize Git Repository

```powershell
# Initialize git repository
git init

# Check git status
git status
```

---

## Step 4: Stage All Files

```powershell
# Add all files to staging
git add .

# Verify what will be committed
git status
```

You should see:
```
Changes to be committed:
  new file:   .gitignore
  new file:   LICENSE
  new file:   README.md
  new file:   go.mod
  new file:   hashing-1/hash_ring.go
  new file:   main.go
  new file:   redundant-hashing/hash_ring.go
  new file:   replication-hashing/hash_ring.go
  new file:   screenshots/...
```

---

## Step 5: Create Initial Commit

```powershell
git commit -m "Initial commit: Consistent hashing implementation with three approaches"
```

---

## Step 6: Connect to GitHub Remote

Replace `Tanishq4501` with your actual GitHub username:

```powershell
git remote add origin https://github.com/Tanishq4501/go-hash.git
```

Verify the remote:

```powershell
git remote -v
```

---

## Step 7: Push to GitHub

```powershell
# Rename branch to main (if needed)
git branch -M main

# Push to GitHub
git push -u origin main
```

### If you encounter authentication issues:

#### Option 1: Personal Access Token (Recommended)

1. Go to GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Click "Generate new token (classic)"
3. Give it a name: `go-hash-repo-access`
4. Select scopes: `repo` (full control)
5. Click "Generate token"
6. **Copy the token immediately** (you won't see it again!)
7. Use it as your password when pushing

#### Option 2: GitHub CLI (Modern approach)

```powershell
# Install GitHub CLI (if not already installed)
winget install --id GitHub.cli

# Authenticate
gh auth login

# Push using GitHub CLI
gh repo create go-hash --public --source=. --push
```

---

## Step 8: Verify on GitHub

1. Go to `https://github.com/Tanishq4501/go-hash`
2. You should see:
   - ✅ README.md rendered beautifully
   - ✅ All your code files
   - ✅ Screenshots folder with images
   - ✅ Go badge showing version
   - ✅ License badge

---

## Step 9: Add Repository Topics (Optional but Recommended)

On your GitHub repository page:

1. Click the **⚙️ gear icon** next to "About"
2. Add topics:
   ```
   go, golang, consistent-hashing, distributed-systems, hash-ring, 
   load-balancing, caching, virtual-nodes, redundancy, scalability
   ```
3. Add website (if you have documentation hosted)
4. Click "Save changes"

---

## Step 10: Enable GitHub Pages (Optional)

To host documentation:

1. Go to repository **Settings** → **Pages**
2. Source: Deploy from a branch
3. Branch: `main` → `/docs` or `/root`
4. Click "Save"

---

## Common Git Commands for Future Updates

### Making Changes

```powershell
# Check status
git status

# Add specific files
git add main.go README.md

# Or add all changes
git add .

# Commit with message
git commit -m "Add feature: improved error handling"

# Push to GitHub
git push origin main
```

### Creating Branches

```powershell
# Create and switch to new branch
git checkout -b feature/new-algorithm

# Make changes, commit them
git add .
git commit -m "Implement new consistent hashing algorithm"

# Push branch to GitHub
git push origin feature/new-algorithm

# Switch back to main
git checkout main

# Merge branch (after testing)
git merge feature/new-algorithm
```

### Syncing with Remote

```powershell
# Fetch latest changes
git fetch origin

# Pull latest changes
git pull origin main

# View commit history
git log --oneline --graph --all
```

---

## Step 11: Create GitHub Repository Sections

### Add Repository Description

In your GitHub repository:
1. Click the **⚙️ gear** next to "About"
2. Description: `🔄 Production-ready consistent hashing implementation in Go featuring basic hashing, virtual nodes, and redundancy support`
3. Add website URL (if any)
4. Add topics (as mentioned in Step 9)

### Enable Issues

Settings → Features → ✅ Issues

### Add Labels

Go to Issues → Labels, create:
- `enhancement` - New features
- `bug` - Something isn't working  
- `documentation` - Documentation improvements
- `good first issue` - Good for newcomers
- `help wanted` - Extra attention needed

---

## Step 12: Add Repository Badges to README (Already included!)

Your README already has:
- Go version badge
- License badge

To add more badges in the future, check [shields.io](https://shields.io)

---

## Step 13: Star Your Own Repository ⭐

Don't forget to star your own repository to show it in your profile!

---

## Quick Reference Card

```powershell
# Status check
git status

# Add all changes
git add .

# Commit
git commit -m "Your message"

# Push to GitHub  
git push origin main

# Pull latest
git pull origin main

# View remotes
git remote -v

# View branches
git branch -a

# Create branch
git checkout -b branch-name

# Delete local branch
git branch -d branch-name

# Delete remote branch
git push origin --delete branch-name
```

---

## Troubleshooting

### Error: "fatal: remote origin already exists"

```powershell
git remote remove origin
git remote add origin https://github.com/Tanishq4501/go-hash.git
```

### Error: "failed to push some refs"

```powershell
# Pull first, then push
git pull origin main --rebase
git push origin main
```

### Error: "Permission denied (publickey)"

Use HTTPS instead of SSH:
```powershell
git remote set-url origin https://github.com/Tanishq4501/go-hash.git
```

### Large files warning

If screenshots are too large:
```powershell
# Install Git LFS
winget install --id GitHub.GitLFS

# Track large files
git lfs track "*.png"
git lfs track "*.svg"
git add .gitattributes
git commit -m "Add Git LFS tracking"
```

---

## Next Steps

1. ✅ Repository is live on GitHub
2. 📝 Share the link on LinkedIn, Twitter
3. 📊 Add to your resume/portfolio
4. 🔄 Keep updating with improvements
5. 🌟 Encourage others to star and contribute
6. 📖 Consider writing a blog post about it
7. 🎥 Create a demo video

---

## Share Your Repository

Share with this template:

```
🚀 Just published my Go Consistent Hashing implementation on GitHub!

🔄 Three progressive approaches:
✅ Basic Consistent Hashing
✅ Virtual Nodes for better distribution
✅ Redundancy for fault tolerance

Perfect for distributed systems, caching, and load balancing!

🔗 https://github.com/Tanishq4501/go-hash

#golang #distributedsystems #opensource
```

---

**Good luck with your GitHub repository! 🎉**
