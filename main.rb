require 'time';
require 'active_support/time';

tz = Time.find_zone('America/New_York')
a = tz.local(2021, 3, 14, 2)
puts "#{a.rfc3339}\t#{a.utc.rfc3339}"
b = tz.local(2021, 11, 7, 1)
puts "#{b.rfc3339}\t#{b.utc.rfc3339}"