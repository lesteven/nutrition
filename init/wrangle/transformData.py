import psycopg2

# postgresql db connection
conn = psycopg2.connect("dbname=nutrition user=steven")
cur = conn.cursor()

# sql statement template
sql = "SELECT long_name, nutrient_name, output_value, output_uom, \
    serving_size, serving_size_uom FROM product \
    INNER JOIN nutrient ON product.ndb_no = nutrient.ndb_no \
    INNER JOIN serving ON product.ndb_no = serving.ndb_no \
    WHERE product.ndb_no = %s"

# create json from sql queries, then insert into elasticsearch
def getData(prod_num):
    num = "45127487"
    prod_num = prod_num.replace('"','')
    cur.execute(sql,(prod_num,))
    data = cur.fetchall()
    print(data[0][0])
    for each in data:
        print(each[1], each[2], each[3])
    print()

foodData = '/home/steven/Downloads/nutData/data/Products.csv'

# access file and read first csv value in each line 
i = 0
with open(foodData) as f:
    for line in f:
        prod_num = line.split(',')[0]
        getData(prod_num)
        i += 1
        if i == 2:
            break

cur.close()
conn.close()
