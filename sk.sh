#!/bin/bash
echo "DO NOT RUN AGAIN"
exit

manager=$1
project=$2
brand=$3

# Create a temporary directory for Wails project initialization
temp_dir=$(mktemp -d)
trap 'rm -rf "$temp_dir"' EXIT

# Initialize Wails project with Svelte template in the temporary directory
cd "$temp_dir"
wails init -n "$project" -t svelte

# Copy the necessary files from the temporary directory to the current directory
cp -r "$temp_dir/$project/." "$OLDPWD"
cd "$OLDPWD"

# Update wails.json with the specified package manager
sed -i "s|npm|$manager|g" wails.json
sed -i 's|"auto",|"auto",\n  "wailsjsdir": "./frontend/src/lib",|' wails.json

# Update main.go for Svelte build directory
sed -i "s|all:frontend/dist|all:frontend/build|" main.go

# If brand is specified, move and update files
if [[ -n "$brand" ]]; then
    mv frontend/src/App.svelte +page.svelte
    sed -i "s|'./assets|'\$lib/assets|" +page.svelte
    sed -i "s|'../wails|'\$lib/wails|" +page.svelte
    mv frontend/src/assets .
fi

# Remove frontend directory
rm -r frontend

# Create new Svelte project
$manager create svelte@latest frontend

# Move and adjust files if brand is specified
if [[ -n "$brand" ]]; then
    mv +page.svelte frontend/src/routes/+page.svelte
    mkdir -p frontend/src/lib
    mv assets frontend/src/lib/
fi

# Navigate to frontend directory and install dependencies
cd frontend || exit
$manager i
$manager uninstall @sveltejs/adapter-auto
$manager i -D @sveltejs/adapter-static

# Update Svelte configuration for static adapter
echo -e "export const prerender = true\nexport const ssr = false" > src/routes/+layout.ts
sed -i "s|-auto';|-static';|" svelte.config.js

# Go back to project root directory and start Wails development server
cd ..
wails dev
