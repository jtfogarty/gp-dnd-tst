FROM gitpod/workspace-full

# Install Go
RUN sudo apt-get update -y && \
    sudo apt-get install -y golang-go

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | sudo -E bash - && \
    sudo apt-get install -y nodejs

# Install pnpm
RUN npm install -g pnpm

# Install Wails CLI
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
