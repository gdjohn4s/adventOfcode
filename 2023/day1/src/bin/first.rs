use core::borrow::Borrow;
use std::fs::File;
use std::io::prelude::*;

fn load_file() -> String {
    let mut file = File::open("input.txt").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    contents
}

fn first_digit(mut line: impl Iterator<Item = impl Borrow<u8>>) -> i32 {
    line.find(|c| c.borrow().is_ascii_digit())
        .map(|c| (c.borrow() - b'0') as i32)
        .unwrap()
}

fn solve(input: &str) -> i32 {
    input
        .split('\n')
        .map(str::as_bytes)
        .map(|line| 10 * first_digit(line.iter()) + first_digit(line.iter().rev()))
        .sum()
}

const TEST_INPUT: &str = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

fn main() {
    assert_eq!(solve(TEST_INPUT), 142);
    println!("Solution: {}", solve(load_file().trim_end()));
}
