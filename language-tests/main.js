// Never change, Javascript
class Time {
	constructor(
		year,
		month,
		day,
		hours,
		minutes,
		seconds,
		offset
	) {
		this.year = year;
		this.month = month;
		this.day = day;
		this.hours = hours;
		this.minutes = minutes;
		this.seconds = seconds;
		this.offset = offset;
	}

	toRFC3339() {
		function p(v) {
			return ('' + v).padStart(2, '0');
		}

		function tz(offset) {
			if (offset === 0) {
				return 'Z';
			}
			const sign = (offset > 0) ? '-' : '+',
				v = Math.abs(offset),
				h = (v / 60) | 0,
				m = v % 60;
			return `${sign}${p(h)}:${p(m)}`;
		}

		const { year, month, day, hours, minutes, seconds, offset } = this;
		return `${year}-${p(month)}-${p(day)}T${p(hours)}:${p(minutes)}:${p(seconds)}${tz(offset)}`;
	}

	static local(t) {
		return new this(
			t.getFullYear(),
			t.getMonth() + 1,
			t.getDate(),
			t.getHours(),
			t.getMinutes(),
			t.getSeconds(),
			t.getTimezoneOffset(),
		);
	}

	static utc(t) {
		return new this(
			t.getUTCFullYear(),
			t.getUTCMonth() + 1,
			t.getUTCDate(),
			t.getUTCHours(),
			t.getUTCMinutes(),
			t.getUTCSeconds(),
			0
		);
	}
}

process.env.TZ = 'America/New_York';
const ta = new Date(2021, 2, 14, 2, 0, 0);
console.log(`${Time.local(ta).toRFC3339()}\t${Time.utc(ta).toRFC3339()}`);

const tb = new Date(2021, 10, 7, 1, 0, 0, 0);
console.log(`${Time.local(tb).toRFC3339()}\t${Time.utc(tb).toRFC3339()}`);