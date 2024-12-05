use std::fs::File;
use std::io::{BufReader, BufRead};


fn main(){
    let p1_solution = p1();
    println!("Part 1 solution: {}", p1_solution);
    let p2_solution = p2();
    println!("Part 2 solution: {}", p2_solution);
}

fn p1() -> i64 {
    let grid = read_input("3-input.txt");
    let coords_and_updates : Vec<usize> = vec![0,0,3,1];
    return sled_trees(&grid, coords_and_updates);
}

fn p2() -> i64 {
    let grid = read_input("3-input.txt");
    let coords_and_updates : Vec<Vec<usize>> = vec![vec![0,0,1,1], vec![0,0,3,1], vec![0,0,5,1], vec![0,0,7,1], vec![0,0,1,2]];
    let mut tot_trees : i64 = 1;
    for set in coords_and_updates{
        tot_trees *= sled_trees(&grid, set);
    }
    return tot_trees;
}

fn read_input(file_name: &str) -> Vec<Vec<char>>{
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut v: Vec<Vec<char>> = Vec::new();

    //Read the file line by line using the lines() iterator and save to vector after converting to number
    let mut temp: Vec<char>;
    for line in reader.lines(){
        temp = line.unwrap().chars().collect();
        v.push(temp);
    }

    return v;
}

fn sled_trees(grid : &Vec<Vec<char>>, vals : Vec<usize>) -> i64 {
    let mut row = vals[0];
    let mut col = vals[1];
    let hor_shift = vals[2];
    let vert_shift = vals[3];

    let mut trees : i64 = 0;

    while row < grid.len(){
        if grid[row][col] == '#'{
            trees += 1;
        }
        row += vert_shift;
        col += hor_shift;
        if row < grid.len(){
            if col >= grid[row].len(){
                col = col % grid[row].len();
            }
        }
    }
    return trees;
}