// g++ -std=gnu++11 -o test test.cpp
#include <iostream>
#include <fstream>
#include "test.h"
#include <rapidjson/istreamwrapper.h>

int main() {
    using namespace rapidjson;

    std::ifstream ifs { "test.json" };
    if ( !ifs.is_open() ) {
        std::cerr << "Could not open file for reading!\n";
        return EXIT_FAILURE;
    }
    IStreamWrapper isw { ifs };
    Document doc {};
    doc.ParseStream( isw );

    auto t = test::FromJson(doc);
    std::cout << test::ToJson(t);
    return 0;
}
