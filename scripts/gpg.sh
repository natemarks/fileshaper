#!/bin/bash
# usage: source scripts/gpg.sh
decrypt() {
    local target_file="${1}"
    local output_file="${1%.asc}"
    local env_var_key="SYMMETRIC_KEY"

    if [[ -z "${!env_var_key}" ]]; then
        read -rsp "Enter symmetric key: " key
        export "$env_var_key"="$key"
        echo -e "\nSymmetric key exported to environment variable."
    fi

    if ! gpg --batch --yes --decrypt --passphrase "${!env_var_key}" -o "$output_file" "$target_file"; then
        echo "Error: Decryption failed."
        exit 1
    fi

    echo "File decrypted successfully to $output_file"
}

encrypt() {
    local target_file="${1}"
    local output_file="${1}.asc"
    local env_var_key="SYMMETRIC_KEY"

    if [[ -z "${!env_var_key}" ]]; then
        read -rsp "Enter symmetric key: " key
        export "$env_var_key"="$key"
        echo -e "\nSymmetric key exported to environment variable."
    fi

    if ! gpg --batch --yes --symmetric --armor --cipher-algo AES256 --passphrase "${!env_var_key}" -o "$output_file" "$target_file"; then
        echo "Error: Decryption failed."
        exit 1
    fi

    echo "File decrypted successfully to $output_file"
}