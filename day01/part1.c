#include <stdio.h>

int main () {
    char c;
    int n=0;
    while((c=fgetc(stdin))>0) {
        n+=(c=='('?1:(c==')'?-1:0));
    }
    printf("%d\n",n);
}
