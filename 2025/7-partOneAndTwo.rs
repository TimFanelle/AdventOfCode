use std::fs::File;
use std::io::{BufReader, BufRead};
use std::collections::{HashSet, HashMap};

fn main() {
    let grid = read_input("7-input.txt");
    let p1_solution = p1(&grid);
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2(&grid);
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1(manifold: &Vec<Vec<char>>) -> i64{
    let mut beams = HashSet::<Downward>::new();
    beams.insert(Downward {
        x: {
            if let Some(index) = manifold.iter().position(|inner| inner.first() == Some(&'S')) {
                index
            } else {
                0
            }
        },
        y: 0
    });
    
    let mut end_val = 0;
    
    while !beams.is_empty() {
        let mut beams_to_remove = HashSet::<Downward>::new();
        let mut beams_to_add = vec!();
        for beam in beams.iter() {
            if beam.y +1 < manifold[0].len() && manifold[beam.x][beam.y + 1] == '^' {
                end_val += 1;
                if beam.x > 0{
                    beams_to_add.push(Downward{
                        x: beam.x - 1,
                        y: beam.y + 1
                    });
                }
                if beam.x < manifold.len() - 1 {
                    beams_to_add.push(Downward{
                        x: beam.x + 1,
                        y: beam.y + 1
                    });
                }
            }
            else if manifold[0].len() > beam.y + 1{
                beams_to_add.push(Downward{
                    x: beam.x,
                    y: beam.y + 1
                });
            }
            beams_to_remove.insert(beam.clone());
        }
        beams.retain(|b| !beams_to_remove.contains(b));
        for b in beams_to_add {
            beams.insert(b);
        }
    }

    return end_val;
}

fn p2(manifold: &Vec<Vec<char>>) -> i64{
    let start_x = manifold
        .iter()
        .position(|row| row.first() == Some(&'S'))
        .unwrap();

    let start_pos = Downward { x: start_x, y: 0 };

    let mut memo: HashMap<Downward, i64> = HashMap::new();

    fn dfs(pos: Downward, manifold: &Vec<Vec<char>>, memo: &mut HashMap<Downward, i64>) -> i64 {
        let width = manifold.len();
        let height = manifold[0].len();

        if let Some(&cached) = memo.get(&pos) {
            return cached;
        }

        if pos.y >= height {
            return 1;
        }

        let mut total = 0;

        if manifold[pos.x][pos.y] == '^' {
            if pos.x > 0 {
                total += dfs(Downward { x: pos.x - 1, y: pos.y }, manifold, memo);
            }
            if pos.x + 1 < width {
                total += dfs(Downward { x: pos.x + 1, y: pos.y }, manifold, memo);
            }
        } else {
            total += dfs(Downward { x: pos.x, y: pos.y + 1 }, manifold, memo);
        }

        memo.insert(pos, total);

        total
    }

    dfs(start_pos, manifold, &mut memo)
}

fn read_input(file_name: &str) -> Vec<Vec<char>>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);
    let mut manifold: Vec<Vec<char>> = vec!();
    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        manifold.push(line.unwrap().chars().collect());  
    }
    let rows = manifold.len();
    let cols = manifold[0].len();

    let transposed: Vec<Vec<char>> = (0..cols).map(|col| {
        (0..rows)
            .map(|row| manifold[row][col])
            .collect()
    }).collect();
    return transposed;
}

#[derive(Debug, Clone, Eq, Hash, PartialEq)]
struct Downward {
    x: usize,
    y: usize
}
