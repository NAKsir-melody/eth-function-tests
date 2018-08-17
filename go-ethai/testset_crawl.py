import requests
from bs4 import BeautifulSoup
import re
import binascii as bina
import struct

def gather_blockinfo(nblock):
#    print nblock
    url = "https://etherscan.io/block/" + str(nblock)
    source = requests.get(url)
    #print source.text
    soup = BeautifulSoup(source.text, 'lxml')
    table = soup.find(id="ContentPlaceHolder1_maintable")
    item = table.find_all('div')
    count = 1;
    result = ""


    for i in item:
        #diff
        if count is 16:
            #print int(re.sub(',','', i.text))
            #print hex(int(re.sub(',','', i.text)))
            result = str(hex(int(re.sub(',','', i.text))))
            hex_int = int(result[2:], 16)
            result_binary = "{0:b}".format(hex_int).zfill(64)
            for i in range(0,63):
                training_data.write(result_binary[i])
                training_data.write(", ")
            training_data.write(result_binary[63])
            training_data.write("\r\n")
            #print result_binary
        #nonce
        if count is 26:
            result = str(i.text)
            #print result
            hex_int = int(result[2:], 16)
            result_binary = "{0:b}".format(hex_int).zfill(64)
            for i in range(0,63):
                label_data.write(result_binary[i])
                label_data.write(", ")
            label_data.write(result_binary[63])
            label_data.write("\r\n")
            #print result_binary

        count += 1


training_data = open("test_input_data.txt",'w')
label_data = open("test_label_data.txt",'w')
#for i in range(5942634, 5942736):
for i in range(5942724, 5942736):
    gather_blockinfo(i)

training_data.close()
label_data.close()
