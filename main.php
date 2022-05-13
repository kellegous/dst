<?php

$tz = new DateTimeZone('America/New_York');
$utc = new DateTimeZone('UTC');

$t = (new DateTimeImmutable('now', $tz))
	->setDate(2021, 3, 14)
	->setTime(2, 0, 0, 0);
printf(
	"%s\t%s\n",
	$t->format(DateTime::RFC3339),
	$t->setTimezone($utc)->format(DateTime::RFC3339)
);

$t = (new DateTimeImmutable('now', $tz))
	->setDate(2021, 11, 7)
	->setTime(1, 0, 0, 0);
printf(
	"%s\t%s\n",
	$t->format(DateTime::RFC3339),
	$t->setTimezone($utc)->format(DateTime::RFC3339)
);