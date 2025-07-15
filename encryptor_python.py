from cryptography.fernet import Fernet
import sys

def generate_key():
    key = Fernet.generate_key()
    with open("secret.key", "wb") as key_file:
        key_file.write(key)

def load_key():
    return open("secret.key", "rb").read()

def encrypt_file(filename):
    key = load_key()
    f = Fernet(key)
    with open(filename, "rb") as file:
        data = file.read()
    encrypted = f.encrypt(data)
    with open(filename + ".enc", "wb") as file:
        file.write(encrypted)

def decrypt_file(filename):
    key = load_key()
    f = Fernet(key)
    with open(filename, "rb") as file:
        encrypted_data = file.read()
    decrypted = f.decrypt(encrypted_data)
    with open(filename.replace(".enc", ".dec"), "wb") as file:
        file.write(decrypted)

if __name__ == "__main__":
    action = sys.argv[1]
    file = sys.argv[2]
    if action == "generate":
        generate_key()
    elif action == "encrypt":
        encrypt_file(file)
    elif action == "decrypt":
        decrypt_file(file)
    else:
        print("Usage: python encryptor.py [generate|encrypt|decrypt] filename")
