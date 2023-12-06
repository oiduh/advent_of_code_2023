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
