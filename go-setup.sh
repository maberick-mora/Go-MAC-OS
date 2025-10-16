#!/bin/bash

# Function to install Homebrew if not already installed
install_homebrew () {
    if ! command -v brew &>/dev/null; then
        echo "Homebrew is not installed. Installing..."
        /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
        echo "Homebrew installation completed."
    else
        echo "Homebrew is already installed."
    fi
}

# Function to set up Homebrew environment variables
setup_homebrew_env () {
    echo "Setting up Homebrew environment variables..."
    eval "$(/opt/homebrew/bin/brew shellenv)"
}

# Function to install sshpass using Homebrew
install_sshpass () {
    if ! command -v brew &>/dev/null; then
        echo "Homebrew is not installed. Skipping sshpass installation."
        return
    fi

    if ! brew list | grep -q "sshpass"; then
        echo "Installing sshpass using Homebrew..."
        brew install sshpass
    else
        echo "sshpass is already installed."
    fi
}

# Function to check if GPG is installed and install it if necessary
install_gpg() {
    if ! command -v gpg &>/dev/null; then
        echo "GPG is not installed. Installing GPG using Homebrew..."
        if ! command -v brew &>/dev/null; then
            echo "Homebrew is not installed. Please install Homebrew to continue."
            exit 1
        fi
        brew install gnupg
        echo "GPG has been installed successfully."
    fi
}

# Function to add alias entry to ~/.zshrc
add_alias_to_zshrc() {
    if ! grep -q "alias go='~/go'" ~/.zshrc; then
        echo "Adding alias entry to ~/.zshrc..."
        echo "alias go='~/go'" >> ~/.zshrc
        source ~/.zshrc
        echo "Alias added successfully."
    else
        echo "Alias 'go' already exists in ~/.zshrc. Skipping."
    fi
}

# Install Homebrew
install_homebrew

# Set up Homebrew environment variables
setup_homebrew_env

# Install sshpass using Homebrew
install_sshpass

# Install GPG using Homebrew
install_gpg

# Add alias entry to ~/.zshrc if not already exists
add_alias_to_zshrc
