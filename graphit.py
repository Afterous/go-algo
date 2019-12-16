# Using python3, graph the performance of Test[^]+SortTime
import pandas as pd
import matplotlib.pyplot as plt
import numpy as np
import csv
# Example
df = pd.read_csv('/PATH/QuickSortOutput.csv',sep=',')
df.plot(kind = 'line')
df.plot(x='N',y='nanoseconds',kind= 'line')
plt.style.use('ggplot')
plt.title("quickSort")
plt.show()
