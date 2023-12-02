total = 0
with open("AdventOfCode/2023/1-input.txt", 'r') as textIn:
    for line in textIn:
        str = ""
        for character in line:
            try:
                int(character)
                str += character
            except:
                pass
        combo_int = str[0] + str[-1]
        total += int(combo_int)

print(total)
        