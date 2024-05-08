import requests
import json
import pandas as pd

STOCK_API_URL = "https://mis.twse.com.tw/stock/api/getStockInfo.jsp"
ODD_API_URL = "https://mis.twse.com.tw/stock/api/getOddInfo.jsp"
tse_list = ['0050', '0056', '2330', '079831']
otc_list = ['6547']

list1 = '|'.join('tse_{}.tw'.format(stock) for stock in tse_list)
list2 = '|'.join('otc_{}.tw'.format(stock) for stock in otc_list)

total_list = list1 + '|' + list2

res = requests.get(STOCK_API_URL+"?ex_ch="+total_list)

if res.status_code != 200:
  raise Exception('Failed to get TWSE stack data.')
else:
  print(res.text)

# Wash
# data = json.loads(res.text)

# cols = ['n', 'c', 'z']
# df = pd.DataFrame(data['msgArray'], columns=cols)
# df.columns = ['公司名稱', '股票代號', '成交價']

# print(df)