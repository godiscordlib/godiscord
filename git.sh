#!/bin/bash

if git diff-index --quiet HEAD --
then
    echo "ℹ️ No changes to commit."
    exit 0
fi

echo "➕ Adding changes to commit..."
git add . || { echo "❌ Failed to add files"; exit 1; }
echo "✅ Added changes to commit!"

read -p "📝 Enter commit message: " git_commit
git commit -m "$git_commit" || { echo "❌ Commit failed"; exit 1; }
echo "✅ Commited with message $git_commit"

echo "🕓 Pushing..."
git push || { echo "❌ Push failed"; exit 1; }
echo "✅ Done pushing."

echo "✅ All done."
