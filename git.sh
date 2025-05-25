#!/bin/bash

echo "➕ Adding changes to commit..."
git add .
echo "✅ Added changes to commit!"

read -p "📝 Enter commit message: " git_commit
git commit -m "$git_commit"
echo "✅ Commited with message $git_commit"

echo "🕓 Pushing..."
git push
echo "✅ Done pushing."

echo "✅ Done."
