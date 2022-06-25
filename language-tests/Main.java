import java.text.SimpleDateFormat;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;

public class Main {
    public static void main(String[] args) {
        var f = DateTimeFormatter.ofPattern("yyyy-MM-dd'T'HH:mm:ssXXX");
        var utc = ZoneId.of("UTC");
        var et = ZoneId.of("America/New_York");
        var t = ZonedDateTime.of(
                2021,
                3,
                14,
                2,
                0,
                0,
                0,
                et);
        System.out.printf("%s\t%s\n", f.format(t), f.format(t.withZoneSameInstant(utc)));

        t = ZonedDateTime.of(
                2021,
                11,
                7,
                1,
                0,
                0,
                0,
                et);
        System.out.printf("%s\t%s\n", f.format(t), f.format(t.withZoneSameInstant(utc)));
    }
}
