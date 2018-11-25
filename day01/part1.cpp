#include <iostream>
#include <fstream>

int main()
{
    std::ifstream infile("input");
    int n = 0;
    char c;
    while (infile >> c)
    {
        switch(c) {
            case '(' :
                n++;
                break;
            case ')' :
                n--;
                break;
        }
    }
    std::cout << n;
}