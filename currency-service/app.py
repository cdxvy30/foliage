import requests

res = requests.get('https://tw.rter.info/capi.php')

currency = res.json()

print(currency)