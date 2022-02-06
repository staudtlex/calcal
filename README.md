# calcal - Compute and convert dates

## About
_calcal_ is a command line application that computes and converts dates from 11 calendars, consisting of the Gregorian, ISO, Julian, Islamic, Hebrew, Mayan (long count, haab, tzolkin), French Revolutionary, and Old Hindu (solar, lunar) calendars.

For the calendrical computations, _calcal_ uses the functions implemented in the Go package [_libcalendar_](https://github.com/staudtlex/libcalendar). For a discussion of the original functions and Lisp code (upon which _libcalendar_ is based), see [Dershowitz, Nachum, and Edward Reingold. 1990. "Calendrical Calculations", Software - Practice and Experience, 20 (9), 899-928](https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.17.4274) and [Reingold, Edward, Nachum Dershowitz, and Stewart Clamen. 1993. "Calendrical Calculations, II: Three Historical Calendars", Software - Practice & Experience, 23 (4), 383-404.](https://citeseerx.ist.psu.edu/viewdoc/summary?doi=10.1.1.13.9215)

## Installing
```
go install staudtlex.de/calcal@latest
```
Alternatively,
```
git clone staudtlex.de/calcal calcal && cd calcal && go install
```

## Usage
To compute the dates corresponding to the current date (e.g. 2022-01-12), simply run `calcal` on the command line.
```
$ calcal
Please note:
- for dates before 1 January 1 (Gregorian), calcal may return incorrect
  ISO and Julian dates
- for dates before the beginning of the Islamic calendar (16 July 622
  Julian), calcal returns an invalid Islamic date
- for dates before the beginning of the French Revolutionary calendar
  calendar (22 September 1792), calcal returns an invalid French date
- Old Hindu (solar and lunar) calendar dates may be off by one day

Gregorian               12 January 2022                 
ISO                     2022-W02-3                      
Julian                  30 December 2021                
Islamic                 8 Jumada II 1443                
Hebrew                  10 Shevat 5782                  
Mayan Long Count        13.0.9.3.9                      
Mayan Haab              7 Muan                          
Mayan Tzolkin           11 Muluc                        
French Revolutionary    23 Nivôse an 230                
Old Hindu Solar         28 Dhanus 5122                  
Old Hindu Lunar         10 Pausha 5122        
```
Run `calcal -h` to show all available options.
```
$ calcal -h
Usage of calcal:
  -c string
        comma-separated list of calendars. Currently, calcal supports
        all             all calendars listed below (default)
        gregorian       Gregorian calendar
        iso             ISO calendar
        julian          Julian calendar
        islamic         Islamic calendar
        hebrew          Hebrew calendar
        mayanLongCount  Mayan Long Count calendar
        mayanHaab       Mayan Haab calendar
        mayanTzolkin    Mayan Tzolkin calendar
        french          French Revolutionary calendar
        oldHinduSolar   Old Hindu Solar calendar
        oldHinduLunar   Old Hindu Lunar calendar
  -d string
        date (format: yyyy-mm-dd). When omitted, '-d' defaults to the
        current date (2022-01-12)
```

## Limitations
_calcal_ currently only supports converting _Gregorian calendar dates_ into ISO, Julian, Islamic, Hebrew, Mayan (long count, haab, tzolkin), French Revolutionary, and Old Hindu (solar, lunar) dates. Converting dates _from_ any of those calendars is not supported (yet).

Using the calendrical functions implemented in [_libcalendar_](https://github.com/staudtlex/libcalendar), note that:

- for dates before 1 January 1 (Gregorian), _calcal_ may return incorrect ISO and Julian dates

- for dates before the beginning of the Islamic calendar (16 July 622 Julian/19 July 622 Gregorian), _calcal_ returns an invalid Islamic date (currently displayed as "`0  0`" on stdout)

- for dates before the beginning of the French Revolutionary calendar calendar (22 September 1792), _calcal_ returns an invalid French date (currently displayed as "`0   an 0`")

- Old Hindu (solar and lunar) calendar dates may be off by one day