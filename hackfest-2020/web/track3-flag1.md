# Track 3 - Flag 1

In this challenge, we were landing in a page that displayed an image with a link to `/<REDACTED>/scr1pt_h3r3`.

```html
<!DOCTYPE html>
<html lang="en">

<head>
  <link rel="stylesheet" href="/files/css/style.css">
</head>

<body>
  <div class=bg>
    <title>RustWeb</title>
    <h1>Rustacean Web Application</h1>
    <h2>What a joker this Ferris! He spreads the flags at different places in the application. In order to find them, you will have to find the code files, solve the problems and send the answers to the right place in order to progress in the challenge.</h2>
    <img src="/files/images/rustacean.png">
    <a href="/<REDACTED>/scr1pt_h3r3"></a>
  </div>

</body>

</html>
```

I tried bruteforcing the `<REDACTED>` part but didn't get anything.

I then tried downloading the image and see if it contained anything hidden inside.
I was able to extract the path by using `exiftool`, or a simple strings would have worked as well.

```
$ exiftool rustacean.png
ExifTool Version Number         : 12.00
File Name                       : rustacean.png
Directory                       : .
File Size                       : 48 kB
File Modification Date/Time     : 2020:11:20 22:20:14-05:00
File Access Date/Time           : 2020:11:20 22:29:31-05:00
File Inode Change Date/Time     : 2020:11:20 22:20:14-05:00
File Permissions                : rw-r--r--
File Type                       : PNG
File Type Extension             : png
MIME Type                       : image/png
Image Width                     : 1200
Image Height                    : 800
Bit Depth                       : 8
Color Type                      : RGB with Alpha
Compression                     : Deflate/Inflate
Filter                          : Adaptive
Interlace                       : Noninterlaced
Software                        : Adobe ImageReady
Comment                         : /hf_twelve_challenge_random
Image Size                      : 1200x800
Megapixels                      : 0.960


$ strings -n 10 rustacean.png
tEXtSoftware
Adobe ImageReadyq
#tEXtComment
/hf_twelve_challenge_random

VR,VT[{mwm;
NOpP2&@?3&Y
\e7&%q]a\$
```

By navigating to `/hf_twelve_challenge_random/scr1pt_h3r3`, we could find a directory listing of all 4 challenges.

The first one was pretty simple: 

```rust
use std;

// /flag1
fn main() { 
    println!("This flag wasn't hard.");
}
```

Simply go to `/hf_twelve_challenge_random/flag1` and enter `This flag wasn't hard.` in the textbox to get the flag.
