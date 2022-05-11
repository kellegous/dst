function toRFC339(d) {
	function p(v) {
		return ('' + v).padStart(2, '0');
	}

	function tz(offset) {
		if (offset === 0) {
			return 'Z';
		}
		const sign = (offset > 0) ? '+' : '-',
			v = Math.abs(offset),
			h = (v / 60) | 0,
			m = v % 60;
		return `${sign}${p(h)}:${p(m)}`;
	}

	const ys = d.getFullYear() + '',
		ms = p(d.getMonth()),
		ds = p(d.getDate()),
		hs = p(d.getHours()),
		ns = p(d.getMinutes()),
		ss = p(d.getSeconds());
	return `${ys}-${ms}-${ds}T${hs}:${ns}:${ss}${tz(d.getTimezoneOffset())}`;

}

process.env.TZ = 'America/New_York';
const ta = new Date(2021, 2, 14, 2, 1, 0);
console.log(`${toRFC339(ta)}\t${ta.toISOString()}`);
const tb = new Date(2021, 10, 7, 1, 1, 0, 0);
console.log(`${toRFC339(tb)}\t${tb.toISOString()}`);