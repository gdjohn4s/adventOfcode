use core::borrow::Borrow;
use std::fs::File;
use std::io::prelude::*;

fn load_file() -> String {
    let mut file = File::open("input.txt").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    contents
}

fn first_digit(line: impl Iterator<Item = impl Borrow<u8>>) -> (usize, i32) {
    line.enumerate()
        .find(|(_, c)| c.borrow().is_ascii_digit())
        .map(|(idx, c)| (idx, (c.borrow() - b'0') as i32))
        .unwrap_or((usize::MAX, 0))
}

const WORDS: [&[u8]; 9] = [
    b"one", b"two", b"three", b"four", b"five", b"six", b"seven", b"eight", b"nine",
];

fn index_of_word(word: &[u8], line: &[u8]) -> usize {
    line.windows(word.len())
        .position(|w| w == word)
        .unwrap_or(usize::MAX)
}

fn first_word(line: &[u8], transform: &dyn Fn(&[u8]) -> Vec<u8>) -> (usize, i32) {
    WORDS
        .iter()
        .enumerate()
        .map(|(idx, word)| {
            (
                index_of_word(transform(word).as_slice(), transform(line).as_slice()),
                (idx + 1) as i32,
            )
        })
        .min_by_key(|(idx, _)| *idx)
        .unwrap()
}

fn reverse_word(word: &[u8]) -> Vec<u8> {
    word.iter().rev().copied().collect::<Vec<_>>()
}

fn solve(input: &str) -> i32 {
    input
        .split('\n')
        .map(str::as_bytes)
        .map(|line| {
            10 * [first_word(line, &|w| w.to_vec()), first_digit(line.iter())]
                .iter()
                .min_by_key(|(idx, _)| *idx)
                .unwrap()
                .1
                + [
                    first_word(line, &reverse_word),
                    first_digit(line.iter().rev()),
                ]
                .iter()
                .min_by_key(|(idx, _)| *idx)
                .unwrap()
                .1
        })
        .sum()
}

const TEST_INPUT: &str = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

fn main() {
    assert_eq!(solve(TEST_INPUT), 281);
    println!("Solution: {}", solve(load_file().trim_end()));
}
