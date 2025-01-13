# Track 3 - Flag 2

`flag2.rs`:

```rust
fn main() {

    const RUST_CONS: u32 = 13_33_777;
    
    let first_method  = add_method(RUST_CONS);

    let x: u8 = 2 * b'A';
    let x = x + 7;
    
    let second_method: u32 = first_method + add_method_a(x, RUST_CONS);
}

fn add_method (num: u32) -> u32 {
    return num + 33710_11; 
}

fn add_method_a(num: u8, rust_const: u32) -> u32 { 
    num as u32 + rust_const
}
```

This one was also very simple. All that was needed was to print `second_method` at the end of `main()`:

```rust
    println!("{}", second_method);
```

Then input the result in `/hf_twelve_challenge_random/flag2` to get the flag.
