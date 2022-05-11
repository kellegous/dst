use chrono::prelude::*;
use chrono::Utc;
use chrono_tz::US::Eastern;

fn main() {
    match Eastern.ymd(2021, 3, 14).and_hms_opt(2, 1, 0) {
        Some(t) => println!("{}\t{}", t.to_rfc3339(), t.with_timezone(&Utc).to_rfc3339()),
        None => println!("none\tnone"),
    }
    match Eastern.ymd(2021, 11, 7).and_hms_opt(1, 1, 0) {
        Some(t) => println!("{}\t{}", t.to_rfc3339(), t.with_timezone(&Utc).to_rfc3339()),
        None => println!("none\tnone"),
    }
}
