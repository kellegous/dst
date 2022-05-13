# date time ambiguities

I was writing a test case one night and wrote the following Go code to create a `time.Time`.

```go
loc, _ := time.LoadLocation("America/New_York")
t := time.Date(2021, time.March, 14, 2, 0, 0, 0, loc)
```

In the location `America/New_York`, the date March 14, 2021 02:00 does not exist. At precisely that time the US eastern timezone switches from EST (which is five hours behind GMT) to EDT (which is four hours behind GMT). If you were staring at a microwave clock (one that actually updates automatically) you would see that after `01:59` the clock would display `03:00`. So asking the `time` library for 2am is ambigous. There is no 2am on March 14 2021 in the eastern time zone, so the time library has to return you one of a few options that you might be looking for. There does not appear to be any standard on what should be returned in this case and, indeed, we get a variety of different answers from the standard libraries of a lot of runtimes. This is a little project to explore what each language/library returns because I'm fascinated how this obvious abiguity has had such a long life in computing.