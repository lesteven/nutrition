import psycopg2
from elasticsearch import Elasticsearch

# postgresql db connection
conn = psycopg2.connect("dbname=nutrition user=steven")
cur = conn.cursor()

# elasticsearch
es = Elasticsearch()

# sql statement template
sql = "SELECT long_name, nutrient_name, output_value, output_uom, \
    serving_size, serving_size_uom, household_serving, \
    household_serving_uom, preparation_state FROM product \
    INNER JOIN nutrient ON product.ndb_no = nutrient.ndb_no \
    INNER JOIN serving ON product.ndb_no = serving.ndb_no \
    WHERE product.ndb_no = %s"


# append errors to log
def writeErr(log, prod_num):
    with open(log, "a") as myfile:
        myfile.write(prod_num + "\n")

def makeJson(data):
    json = {
        'Name': data[0][0], 
        'Serving Size': {
            'Value': data[0][4], 
            'Unit': data[0][5],
            'HValue': data[0][6],
            'HUnit': data[0][7],
            'Prep_state': data[0][8]
        }
    }
    for each in data:
        json[each[1]] = {'Value': each[2], 'Unit': each[3]}
    return json

# create json from sql queries, then insert into elasticsearch
def getData(prod_num):
    cur.execute(sql,(prod_num,))
    data = cur.fetchall()
    if data and data[0] and data[0][0]:
        return makeJson(data)
    else:
        writeErr("getErr.txt", prod_num)


# insert data into elasticsearch
def insertData(json, prod_num):
    res = es.index(index='product', doc_type='food', id=prod_num, body=json)
    if res['_shards']['failed'] != 0:
        writeErr("inserErr.txt", prod_num)



foodData = '/home/steven/Downloads/nutData/data/Products.csv'
# foodData = '/home/steven/Downloads/nutData/data/test.csv'


# access file and read first csv value in each line 
# i = 0
with open(foodData) as f:
    for line in f:
        prod_num = line.split(',')[0].replace('"','')
        json = getData(prod_num)
        if json:
            insertData(json, prod_num)
        '''
        i += 1
        if i == 2:
            break
            '''

# close postgresql
cur.close()
conn.close()
