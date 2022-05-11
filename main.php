<?php

$tz = new DateTimeZone('America/New_York');

$t = (new DateTime('now', $tz))
	->setDate(2022, 3, 13)
	->setTime(2, 15, 0, 0);
printf("%s\n", $t->format(DateTime::RFC3339));
