#include <time.h>

#include <iomanip>
#include <iostream>
#include <sstream>
#include <string>

bool to_rfc3339(std::string *dst, const struct tm *t) {
    dst->clear();

    if (t == nullptr) {
        return false;
    }

    char buf[255];
    if (strftime(buf, sizeof(buf), "%FT%T", t) == 0) {
        return false;
    }

    std::stringstream ss;
    ss << buf;
    if (t->tm_gmtoff == 0) {
        ss << "Z";
    } else {
        int v = abs(t->tm_gmtoff);
        int h = v / 3600;
        int m = v % 3600;
        ss << ((t->tm_gmtoff > 0) ? "+" : "-");
        ss << std::setfill('0') << std::setw(2) << h;
        ss << ':';
        ss << std::setfill('0') << std::setw(2) << m;
    }
    dst->append(ss.str());
    return true;
}

int main(int argc, char *argv[]) {
    struct tm a = {
        0,
        0,
        2,
        14,
        2,
        2021 - 1900,
        0,
        0,
        0,
    };

    time_t ta = mktime(&a);

    std::string loc, utc;
    to_rfc3339(&loc, localtime(&ta));
    to_rfc3339(&utc, gmtime(&ta));
    std::cout << "Mar 14 2021 02:00" << std::endl;
    std::cout << loc << '\t' << utc << std::endl;

    a = {0, 0, 2, 14, 2, 2021 - 1900, 0, 0, 1};
    ta = mktime(&a);
    to_rfc3339(&loc, localtime(&ta));
    to_rfc3339(&utc, gmtime(&ta));
    std::cout << loc << '\t' << utc << std::endl;

    struct tm b = {
        0,
        0,
        1,
        7,
        10,
        2021 - 1900,
        0,
        0,
        0,
    };

    time_t tb = mktime(&b);
    to_rfc3339(&loc, localtime(&tb));
    to_rfc3339(&utc, gmtime(&tb));
    std::cout << "Nov 7 2021 01:00" << std::endl;
    std::cout << loc << '\t' << utc << std::endl;

    return 0;
}