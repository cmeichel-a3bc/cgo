#include "foo.h"
#include <math.h>

int hasLetter(char* str, char c) {
    int i;
    for(i=0; str[i]; i++) {
        if (str[i] == c) {
            return 1;
        }
    }

    return 0;
}