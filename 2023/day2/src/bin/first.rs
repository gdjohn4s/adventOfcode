use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;

fn load_file() -> String {
    let mut file = File::open("input.txt").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    contents
}

const MAX_RED: i32 = 12;
const MAX_GREEN: i32 = 13;
const MAX_BLUE: i32 = 14;

fn get_games(content: String) -> HashMap<String, Vec<String>> {
    let mut games = HashMap::new();

    for line in content.lines() {
        let parts: Vec<&str> = line.splitn(2, ':').collect();
        if parts.len() < 2 {
            continue;
        }

        let game_id = parts[0].trim().to_string();
        let colors = parts[1]
            .split(';')
            .map(|segment| segment.trim().to_string())
            .collect();

        games.insert(game_id, colors);
    }

    games
}

fn is_set_valid(set: &str) -> bool {
    set.split(',').all(|cube| {
        let parts: Vec<&str> = cube.trim().split_whitespace().collect();
        if parts.len() < 2 {
            return false;
        }

        match parts[0].parse::<i32>() {
            Ok(number) => match parts[1] {
                "red" => number <= MAX_RED,
                "green" => number <= MAX_GREEN,
                "blue" => number <= MAX_BLUE,
                _ => false,
            },
            Err(_) => false,
        }
    })
}

fn check_valid_games(games: HashMap<String, Vec<String>>) -> i32 {
    let mut sum_of_valid_game_ids: i32 = 0;

    for (game_id, sets) in games {
        let game_id_number: i32 = game_id
            .split_whitespace()
            .nth(1)
            .and_then(|id| id.parse::<i32>().ok())
            .unwrap_or(0);

        if sets.iter().all(|set| is_set_valid(set)) {
            sum_of_valid_game_ids += game_id_number;
        }
    }

    sum_of_valid_game_ids
}

fn main() {
    let content = load_file();
    let games = get_games(content);
    println!("{}", check_valid_games(games));
}
