#include <stdio.h>
#include <string.h>

#define PASSWORD "cH4l3nG3"
#define FLAG_LENGTH 43

const char xor_key = 0x5A;

unsigned char encoded_flag[] = {
0x16,
0x11,
0x9,
0xa,
0x17,
0x16,
0x21,
0x3d,
0x3d,
0x5,
0x34,
0x35,
0x2d,
0x5,
0x23,
0x35,
0x2f,
0x5,
0x3b,
0x28,
0x3f,
0x5,
0x2e,
0x32,
0x3f,
0x5,
0x28,
0x3f,
0x3b,
0x36,
0x5,
0x28,
0x3f,
0x2c,
0x5,
0x37,
0x3b,
0x29,
0x2e,
0x3f,
0x28,
0x27,
'\0'
};

int check_password(char *input) {
    int valid = 1;
    for (int i = 0; i < strlen(PASSWORD); i++) {
        if ((input[i] ) != PASSWORD[i]) {
            valid = 0;
            break;
        }
    }
    return valid;
}

void reveal_flag() {
    char decoded_flag[FLAG_LENGTH];
    
    for (int i = 0; i < FLAG_LENGTH - 1; i++) {
        decoded_flag[i] = encoded_flag[i] ^ xor_key;
    }
    
    decoded_flag[FLAG_LENGTH - 1] = '\0';
    printf("%s\n", decoded_flag);
}

int main() {
    char input[50];
    
    printf("Enter the password: ");
    scanf("%49s", input);

    if (strlen(input) != strlen(PASSWORD)) {
        printf("Incorrect password length!\n");
        return 1;
    }

    if (check_password(input)) {
        reveal_flag();
    } else {
        printf("Wrong password. Try again.\n");
    }

    return 0;
}
