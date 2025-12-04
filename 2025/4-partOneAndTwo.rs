use std::fs::File;
use std::io::{BufReader, BufRead};
use std::collections::HashSet;

fn main() {
    let p1_solution = p1();
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1() -> i64{
    let roll_list = read_input("4-input.txt");

    let mut end_val = 0;
    let roll_set: HashSet<Roll> = roll_list.iter().cloned().collect();
    for r in roll_list {
        if num_surrounding(r, &roll_set) < 4 {
            end_val += 1;
        }
    }
    return end_val;
}

fn p2() -> i64{
    let mut roll_list = read_input("4-input.txt");

    let mut end_val = 0;
    let mut roll_set: HashSet<Roll> = roll_list.iter().cloned().collect();
    let mut removed_rolls = true;
    while removed_rolls {
        let mut rolls_removed = vec!();
        for r in &roll_list {
            if num_surrounding(r.clone(), &roll_set) < 4 {
                end_val += 1;
                rolls_removed.push(r.clone());
            }
        }
        removed_rolls = rolls_removed.len() > 0;
        for rem in &rolls_removed {
            roll_set.remove(rem);
        }
        roll_list = roll_list.into_iter().filter(|x| !rolls_removed.contains(x)).collect();
    }
    
    
    return end_val;
}


#[derive(Debug, Clone, Eq, Hash, PartialEq)]
struct Roll {
    x: i32,
    y: i32,
}

fn read_input(file_name: &str) -> Vec<Roll>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v= Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    let mut row = 0;
    for line in reader.lines(){
        let temp = line.expect("unable to read line");
        temp.chars()
            .enumerate()
            .for_each(|(col, val)| {
                if val == '@' {
                    v.push(Roll { y: row, x: col as i32});
                }
            });

        row += 1;
    }
    return v;
}

fn num_surrounding(r: Roll, all: &HashSet<Roll>) -> i32 {
    let surrounding = vec!{
        Roll{y: r.y -1, x:r.x-1}, Roll{y: r.y -1, x:r.x},Roll{y: r.y -1, x:r.x+1},
        Roll{y: r.y   , x:r.x-1},                        Roll{y: r.y,    x:r.x+1},
        Roll{y: r.y +1, x:r.x-1}, Roll{y: r.y +1, x:r.x},Roll{y: r.y +1, x:r.x+1}
    };
    let count = surrounding.iter().filter(|&x| all.contains(x)).count();
    return count as i32;
}