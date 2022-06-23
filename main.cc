#include <getopt.h>
#include <stdint.h>
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

struct Flags {
    Flags() : is_dst(false), verbose(false) {}
    bool is_dst;
    bool verbose;
};

void ParseArgs(int argc, char *argv[], Flags *flags) {
    static struct option long_options[] = {
        {"dst", no_argument, 0, 'd'},
        {"verbose", no_argument, 0, 'v'},
        {0, 0, 0, 0},
    };

    while (true) {
        int option_index = 0;
        int c = getopt_long(argc, argv, "dv", long_options, &option_index);
        if (c == -1) {
            break;
        }

        switch (c) {
            case 'd':
                flags->is_dst = true;
                break;
            case 'v':
                flags->verbose = true;
            case '?':
                break;
            default:
                abort();
        }
    }
}

int main(int argc, char *argv[]) {
    setenv("TZ", "America/New_York", 1);

    Flags flags;

    ParseArgs(argc, argv, &flags);

    if (flags.verbose) {
        std::cout << "Daylight Savings Time: " << (flags.is_dst ? "Yes" : "No") << std::endl;
    }

    struct tm a = {
        .tm_sec = 0,
        .tm_min = 0,
        .tm_hour = 2,
        .tm_mday = 14,
        .tm_mon = 2,
        .tm_year = 2021 - 1900,
        .tm_yday = 0,
        .tm_wday = 0,
        .tm_isdst = flags.is_dst,
    };

    time_t ta = mktime(&a);

    std::string loc, utc;
    to_rfc3339(&loc, localtime(&ta));
    to_rfc3339(&utc, gmtime(&ta));
    std::cout << loc << '\t' << utc << std::endl;

    struct tm b = {
        .tm_sec = 0,
        .tm_min = 0,
        .tm_hour = 1,
        .tm_mday = 7,
        .tm_mon = 10,
        .tm_year = 2021 - 1900,
        .tm_yday = 0,
        .tm_wday = 0,
        .tm_isdst = flags.is_dst,
    };

    time_t tb = mktime(&b);
    to_rfc3339(&loc, localtime(&tb));
    to_rfc3339(&utc, gmtime(&tb));
    std::cout << loc << '\t' << utc << std::endl;

    return 0;
}