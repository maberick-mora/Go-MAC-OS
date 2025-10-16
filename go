#!/bin/bash

# Name of the file containing the encrypted password
password_file="encrypted_pass.txt"

# Check if the password file exists
if [ ! -f "$password_file" ]; then
    echo "The password file does not exist: $password_file"

    # Prompt the user to enter the SSH password
    read -sp "Enter your SSH password: " ssh_password
    echo

    # Prompt the user to enter the encryption password
    read -sp "Enter a password to encrypt the SSH password file: " encryption_password
    echo

    # Encrypt the SSH password and save it to the password file
    echo "$ssh_password" | gpg --symmetric --cipher-algo AES256 --pinentry-mode loopback --output "$password_file" --passphrase "$encryption_password"
    echo "The encrypted password has been saved to: $password_file"
    echo "Please keep this file secure and do not share it with others."
    exit 0
fi

# Prompt the user to enter the decryption password
read -sp "Enter the password to decrypt the encrypted file: " decryption_password
echo

# Decrypt the password file
echo "Trying to decrypt the password file with the provided passphrase..."
password=$(gpg --decrypt --quiet --pinentry-mode loopback --passphrase "$decryption_password" "$password_file" 2>/dev/null)

# Prompt the user to enter the correct decryption password if decryption fails
while [ -z "$password" ]; do
    echo "Error: Failed to decrypt the password file. Please enter the correct password."
    read -sp "Enter the password to decrypt the encrypted file: " decryption_password
    echo
    echo "Trying to decrypt the password file with the provided passphrase..."
    password=$(gpg --decrypt --quiet --pinentry-mode loopback --passphrase "$decryption_password" "$password_file" 2>/dev/null)
done

echo "Password decrypted successfully."

# IP address provided as an argument
ip_address="$1"

# Check if an IP address was provided
if [ -z "$ip_address" ]; then
    echo "Usage: ./go <ip_address>"
    exit 1
fi

# Function to add the host key to the known_hosts file
add_to_known_hosts () {
    echo "Getting host key for $ip_address..."
    ssh-keyscan "$ip_address" >> ~/.ssh/known_hosts
    echo "The host key has been added to the known_hosts file."
}

# Check if the IP address is already in the known_hosts file
if ! grep -q "$ip_address" ~/.ssh/known_hosts; then
    add_to_known_hosts "$ip_address"
fi

# Connect using SSH with the decrypted password | Change your username by modifying user.name 
sshpass -p "$password" ssh -o StrictHostKeyChecking=yes user.name@"$ip_address"
