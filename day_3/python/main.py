import re
import sys, os
from copy import deepcopy
from functools import reduce

def read_data(file_path: str):
    with open(file_path, "r") as f:
        data_ = f.read().strip()
    return data_


if __name__ == "__main__":
    input_file_path = sys.argv[1]
    abs_file_path = os.path.abspath(input_file_path)
    data = read_data(abs_file_path)
    lines = data.split("\n")

    symbols = set(data)
    symbols = [x for x in symbols if not x.isdigit()]
    symbols.remove("\n")
    symbols.remove(".")

    new_data = deepcopy(data)
    for x in symbols:
        new_data = new_data.replace(x, "+")

    padding = []
    for x in new_data.split("\n"):
        padding.append(f".{x}.")

    empty_line = "."*len(padding[0])
    padding.insert(0, empty_line)
    padding.append(empty_line)

    is_part = []
    for row in range(1, len(padding)):
        line = padding[row]
        numbers = re.finditer(r"\d+", line)
        for number in numbers:
            start, end = number.span()
            if line[start-1] == "+" or line[end] == "+":
                is_part.append(number.group(0))
                continue
            for idx in range(start-1, end+1):
                if padding[row-1][idx] == "+" or padding[row+1][idx] == "+":
                    is_part.append(number.group(0))
                    break

    print(sum(int(x) for x in is_part))

    # part 2
    symbols.remove("*")
    gear_data = deepcopy(data)
    for x in symbols:
        gear_data = gear_data.replace(x, ".")

    padding = []
    for x in gear_data.split("\n"):
        padding.append(f".{x}.")

    empty_line = "."*len(padding[0])
    padding.insert(0, empty_line)
    padding.append(empty_line)

    gear_dict = {}

    for row in range(1, len(padding)):
        line = padding[row]
        numbers = re.finditer(r"\d+", line)
        for number in numbers:
            start, end = number.span()
            if line[start-1] == "*":
                x = gear_dict.get(f"{row}-{start-1}")
                if x:
                    gear_dict[f"{row}-{start-1}"].append(number.group(0))
                else:
                    gear_dict[f"{row}-{start-1}"] = [number.group(0)]
            if line[end] == "*":
                x = gear_dict.get(f"{row}-{end}")
                if x:
                    gear_dict[f"{row}-{end}"].append(number.group(0))
                else:
                    gear_dict[f"{row}-{end}"] = [number.group(0)]
            for idx in range(start-1, end+1):
                if padding[row-1][idx] == "*":
                    x = gear_dict.get(f"{row-1}-{idx}")
                    if x:
                        gear_dict[f"{row-1}-{idx}"].append(number.group(0))
                    else:
                        gear_dict[f"{row-1}-{idx}"] = [number.group(0)]
                if padding[row+1][idx] == "*":
                    x = gear_dict.get(f"{row+1}-{idx}")
                    if x:
                        gear_dict[f"{row+1}-{idx}"].append(number.group(0))
                    else:
                        gear_dict[f"{row+1}-{idx}"] = [number.group(0)]

    result = 0
    for x in gear_dict.values():
        if len(x) == 2:
            result += reduce(lambda a, b: int(a)*int(b) , x)

    print(result)



