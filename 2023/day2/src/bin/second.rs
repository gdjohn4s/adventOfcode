use std::collections::HashMap;
use std::fs::File;
use std::io::prelude::*;

// the second part i wanted to try to use struct
// to define what a cube is
#[derive(Debug)]
struct Cubes {
    red: i32,
    green: i32,
    blue: i32,
}

fn load_file() -> String {
    let mut file = File::open("input.txt").unwrap();
    let mut contents = String::new();
    file.read_to_string(&mut contents).unwrap();
    contents
}

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

fn calculate_min_cubes(set: &Vec<String>) -> Cubes {
    let mut min_cubes = Cubes {
        red: 0,
        green: 0,
        blue: 0,
    };

    for cube in set {
        for color_block in cube.split(',') {
            let parts: Vec<&str> = color_block.trim().split_whitespace().collect();

            if let (Ok(number), Some(color)) = (parts[0].parse::<i32>(), parts.get(1)) {
                match *color {
                    "red" => min_cubes.red = min_cubes.red.max(number),
                    "green" => min_cubes.green = min_cubes.green.max(number),
                    "blue" => min_cubes.blue = min_cubes.blue.max(number),
                    _ => (),
                }
            }
        }
    }
    println!("{:?}", min_cubes);
    min_cubes
}

fn calculate_total_power(games: HashMap<String, Vec<String>>) -> i32 {
    let mut total_power = 0;

    for (_game_id, sets) in games {
        let min_cubes = calculate_min_cubes(&sets);
        total_power += min_cubes.red * min_cubes.green * min_cubes.blue;
    }

    total_power
}

fn main() {
    let content = load_file();
    let games = get_games(content);
    let total_power = calculate_total_power(games);

    println!("{}", total_power);
}
