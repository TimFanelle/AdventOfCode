use std::fs::File;
use std::io::{BufReader, BufRead};

fn main() {
    let (ranges, ingredients) = read_input("5-input.txt");
    let p1_solution = p1(&ranges, ingredients);
    println!("Part 1 Solution: {}", p1_solution);
    let p2_solution = p2(&ranges);
    println!("Part 2 Solution: {}", p2_solution);
}

fn p1(ranges: &Vec<Range>, ingredients: Vec<i64>) -> i64{
    let mut end_val = 0;
    for ing in ingredients {
        for ran in ranges{
            if ran.low <= ing && ran.high >= ing {
                end_val += 1;
                break
            }
        }
    }
    return end_val;
}

fn p2(ranges: &Vec<Range>) -> i64{
    let mut end_val = 0;
    for ran in ranges {
        end_val += ran.high - ran.low +1;
    }
    
    return end_val;
}

#[derive(Debug, Clone)]
struct Range{
    low: i64,
    high: i64
}

fn read_input(file_name: &str) -> (Vec<Range>, Vec<i64>){
    //open the file in read-only mode
    let file = File::open(&file_name).expect("unable to open file");

    //create buffered reader
    let reader = BufReader::new(file);

    let mut range_values = Vec::new();
    let mut ingredients = Vec::new();

    let mut range_section = true;
    //Read the file line by line using the lines() iterator and save to vector after converting to number
    for line in reader.lines(){
        let temp = line.expect("unable to read line");
        if temp.is_empty(){
            range_section = false
        } else if range_section {
            let r_v: Vec<i64> = temp.split("-").map(|x| x.parse().unwrap()).collect();
            range_values.push(Range{low: r_v[0], high: r_v[1]});
        } else {
            ingredients.push(temp.parse().unwrap());
        }
    }
    range_values = merge_ranges(range_values);
    return (range_values, ingredients);
}

fn merge_ranges(mut v: Vec<Range>) -> Vec<Range> {
    if v.is_empty() {
        return v;
    }

    v.sort_by_key(|r| r.low);

    let mut out = vec![v[0].clone()];
    for r in v.into_iter().skip(1) {
        let last = out.last_mut().unwrap();

        if r.low <= last.high {
            last.high = last.high.max(r.high);
        } else {
            out.push(r);
        }
    }
    return out
}