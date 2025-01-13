# Track 3 - Flag 4

`flag4.rs`:

```rust
use std::u8;
use std::str;
use base64::{encode};
use data_encoding::{HEXUPPER};
use url::form_urlencoded::{byte_serialize};

fn do_some_stuff (data: String) -> String {
    fs::create_dir("..\\flag31")?;
    fs::File::create("..\\flag3\\flag10.txt")?;
    fs::write("..\\flag31\\flag31.txt", &FLAG.flag3)?;
}

fn do_other_stuff (data: String) -> String { 
    return
}

//static file in binary https://lib.rs/crates/actix-web-static-files

/* #[get("/printhis")]
async fn printhis() -> impl Responder {
    let decoded: String = parse(urlencoded.as_bytes())
        .map(|(key, val)| [key, val].concat())
        .collect();
    HttpResponse::Ok().body("Hello world!")
} */

/* #[get("/hf/flag{id}")]
async fn hello(web::Path((id)): web::Path<(u32)>) -> impl Responder {
    println!("Hello from flag {}" id)
} */
//https://serde.rs/#data-formats

fn main() {
    
    let value = "what is the value to insert ?";

    let decoded_value = "";

    let stuff = do_some_stuff("rust_do_something");

    let encode_step_2: String = byte_serialize(value.as_bytes()).collect();

    let encode_step_1: String  = HEXUPPER.encode(encode_step_2.as_bytes());
   
    do_other_stuff("other stuff");
    
    let encode_step_3: String =  encode(&encode_step_1) 

    println!("Value :", encode_step_3);
    
    //Find the encoded value of :Rust tricky trick 🦀
}
```

This one was done by one of my teammates, but the solution was to remove the `do stuff` lines, replace the `value` with the one in the comment, and fix a few syntax errors.

```rust
use base64::{encode};
use data_encoding::{HEXUPPER};
use url::form_urlencoded::{byte_serialize};

fn main() {

    let value = "Rust tricky trick 🦀";

    let decoded_value = "";

    let encode_step_2: String = byte_serialize(value.as_bytes()).collect();

    let encode_step_1: String  = HEXUPPER.encode(encode_step_2.as_bytes());
   
    let encode_step_3: String =  encode(&encode_step_1);

    println!("Value : {}", encode_step_3);
    
    //Find the encoded value of :Rust tricky trick 🦀
}
```

Then input the result in `/hf_twelve_challenge_random/flag4` to get the flag.
