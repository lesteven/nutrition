import psycopg2
from elasticsearch import Elasticsearch

# postgresql db connection
conn = psycopg2.connect("dbname=nutrition user=steven")
cur = conn.cursor()

# elasticsearch
es = Elasticsearch()

# sql statement template
sqlProduct = "SELECT long_name, manufacturer, ingredients \
        FROM product WHERE ndb_no = %s"
sqlNutrition = "SELECT nutrient_name, output_value, output_uom\
        FROM nutrient WHERE ndb_no = %s"
sqlServing = "SELECT serving_size, serving_size_uom, household_serving,\
        household_serving_uom FROM serving WHERE ndb_no = %s"


# append errors to log
def writeErr(log, prod_num):
    with open(log, "a") as myfile:
        myfile.write(prod_num + "\n")

# transform data to put into elastic search
# b/c es wont take tuples
def transformData(product, nutrition, serving):
    nut = []
    for each in nutrition:
        nut.append({
            'name':each[0],
            'amount': each[1],
            'unit': each[2]
            })
    serv = []
    for each in serving:
        serv.append({
            'size': each[0],
            'unit': each[1],
            'hsize': each[2],
            'hunit': each[3],
        })
    return {
        'name': product[0][0],
        'manufacturer': product[0][1],
        'ingredients': product[0][2],
        'nutrients': nut,
        'serving': serv
    }

# create json from sql queries, then insert into elasticsearch
def getData(prod_num):
    cur.execute(sqlProduct,(prod_num,))
    product = cur.fetchall()

    cur.execute(sqlNutrition,(prod_num,))
    nutrition = cur.fetchall()

    cur.execute(sqlServing,(prod_num,))
    serving = cur.fetchall()

    if product and nutrition and serving:
        return transformData(product, nutrition, serving)
    else:
        writeErr("getErr.txt", prod_num)


# insert data into elasticsearch
def insertData(json, prod_num):
    res = es.index(index='product', doc_type='food', id=prod_num, body=json)
    if res['_shards']['failed'] != 0:
        writeErr("inserErr.txt", prod_num)



foodData = '/home/steven/Downloads/nutData/data/Products.csv'

# access file and read first csv value in each line 
#i = 0
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
