import random
from base64 import b64encode

random.seed(1337)
secret = random.randbytes(16)
flag= input("Enter the flag: ")

encrypted = ""
for i in range(len(flag)):
    encrypted += chr(ord(flag[i]) ^ secret[i % len(secret)])
encrypted = b64encode(encrypted.encode()).decode()

print(f"Encrypted: {encrypted}")
