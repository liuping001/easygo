// g++ -std=c++11 -o test test.cpp
#include "test.h"

int main() {
    auto t = test::FromFile("test.json");
    if (t != nullptr) {
        std::cout << test::ToJson(*t);
    }
    return 0;
}
