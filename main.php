<?php

$tz = new DateTimeZone('America/New_York');

$t = (new DateTime('now', $tz))
	->setDate(2021, 3, 14)
	->setTime(2, 0, 0, 0);
printf("%s\n", $t->format(DateTime::RFC3339));
