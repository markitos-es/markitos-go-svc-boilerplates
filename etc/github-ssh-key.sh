#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Uso: $0 <solo-letras-minus> <email>"
    exit 1
fi

NOMBRE=$1
EMAIL=$2
SSH_DIR="$HOME/.ssh"
KEY_FILE="$SSH_DIR/$NOMBRE"

mkdir -p "$SSH_DIR"
ssh-keygen -t rsa -b 4096 -C "$EMAIL" -f "$KEY_FILE" -N ""
echo "Clave pública generada:"
cat "${KEY_FILE}.pub"