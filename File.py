import numpy as np

def read_file():
    all_data = np.loadtxt("a_example.in", dtype=int, delimiter = " ", skiprows = 0)
    first_row = all_data[0,:]
    all_data = all_data[1:all_data.shape[0],:]
    return first_row, all_data