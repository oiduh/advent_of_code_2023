import sys, os, re


def read_data(file_path: str):
    with open(file_path, "r") as f:
        data_ = f.read().strip()
    return data_


if __name__ == "__main__":
    input_file_path = sys.argv[1]
    abs_path = os.path.abspath(input_file_path)
    data = read_data(abs_path)
    lines = data.split("\n")
    result_1, result_2 = 0, 0
    copies = [1 for _ in range(len(lines))]
    for idx, line in enumerate(lines):
        _, rest = line.split(": ")
        win, have = rest.split("|")
        pattern = re.compile(r"\d+")
        win_numbers = set(pattern.findall(win))
        have_numbers = set(pattern.findall(have))
        common = len(win_numbers & have_numbers)
        for x in range(common):
            if x + 1 + idx < len(copies):
                copies[x+1+idx] += copies[idx]
        result_1 += (2**(common-1) if common>0 else 0)
    print(result_1)
    result_2 = sum(copies)
    print(result_2)


