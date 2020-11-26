# Track 3 - Flag 3

`flag3.rs`:

```rust
use std::{env, fs};
use std::net::TcpListener;
use std::path::PathBuf;


fn main() -> Result<()> {
    let listener = TcpListener::bind("127.0.0.1:7001").unwrap(); //unwrap handle l'erreur avec un panic s'il y a lieu etant donne que bind retourn un Result<T, E> 
    let pool = ThreadPool::new(4);

    for stream in listener.incoming() {
        let stream = stream.unwrap();
        pool.execute(|| {
            handle_connection(stream);
        });
        println!("Connection established");
    }

    let current_dir = env::current_dir()?;
    println!(
        "Entries modified in the last 24 hours in {:?}:",
        current_dir
    );

    // let mut current_dir = match env::current_dir() { 
    //     Ok(dir) => dir,
    //     _ => PathBuf::new()
    // };
    // let path = String::from("filename/<>/notreallytheflagnameyouexpected.txt");
    // current_dir.push(path);
    // println!{"Path to file {:?}", &current_dir};
    // let file = fs::read_to_string(current_dir);

    for entry in fs::read_dir(current_dir)? {
        let entry = entry?;
        let path = entry.path();

        let metadata = fs::metadata(&path)?;
        let last_modified = metadata.modified()?.elapsed()?.as_secs();

        if last_modified < 24 * 3600 && metadata.is_file() {
            println!(
                "Last modified: {:?} seconds, is read only: {:?}, size: {:?} bytes, filename: {:?}",
                last_modified,
                metadata.permissions().readonly(),
                metadata.len(),
                path.file_name().ok_or("No filename")?
            );
        }
    }
    Ok(())
}

fn handle_connection(mut stream: TcpStream) {
    let mut buffer = [0;1024];
    stream.read(&mut buffer).unwrap(); 
    
    let get = b"GET / HTTP/1.1\r\n";

    let (status_line, filename) = if buffer.starts_with(get) {
        ("HTTP/1.1 200 OK\r\n\r\n", "index.html")
    } else {
        ("HTTP/1.1 404 NOT FOUND\r\n\r\n", "404.html")
    };

    let contents = fs::read_to_string(filename).unwrap();

    let response = format!("{}{}", status_line, contents);

    stream.write(response.as_bytes()).unwrap();
    stream.flush().unwrap(); //prend un &[u8] et l'envoie directement dans le stream

}
```

This one was a little more complex. Most of the code was there as decoys and red herrings.
The semi-important part was the commented code, which was hinting towards finding a way to read files.

By going to `/hf_twelve_challenge_random/flag3`, we could see that the `<input>` was named `filename`.
That led to a very straightforward LFI vulnerability. Simply put any file name in there like `../../../etc/passwd` and you could read it.

I enumerated the machine a little with this vulnerability and when I enumerated the environment variables with `../../../self/prof/environ`, I was able to read all 4 flags of the track. Fortunately/unfortunately, that was the last track we worked on so we didn't get any free flags this way, but that's how I got flag 3.

I spoke to the challenge designer afterwards, and he told me that the intended solution was to read the well known rust file `./build.rs` using the LFI. There was a directory name in the comments of that file that led to the flag.
