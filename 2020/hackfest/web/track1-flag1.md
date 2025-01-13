# Track 1 - Flag 1

By digging around the app a little, we find that `robots.txt` contained an entry for `/dev`.
That `/dev` path gave a HTTP Forbidden error message so we moved on to bruteforce the app.

The bruteforce found a git folder (`.git/HEAD`), which was then extracted using the Dumper tool from `https://github.com/internetwache/GitTools`.

After restoring the repo, there was a `flag.txt` file waiting for us.
