# import psycopg2

# # PostgreSQL serveriga ulanish
# conn = psycopg2.connect(
#     host="13.233.155.60",
#     port=5432,
#     database="learn",
#     user="postgres",
#     password="postgres"
# )

# # SQL so'rovni yaratish
# sql_query = "SELECT * FROM mineuser;"

# # SQL so'rovni bajargan natijani olish
# with conn.cursor() as cursor:
#     cursor.execute(sql_query)
#     rows = cursor.fetchall()

# # Natijani chiqarish
# for row in rows:
#     print(row)

# # Ulanishni yopish
# conn.close()


import psycopg2

#establishing the connection
conn = psycopg2.connect(
   database="learn", user='postgres', password='postgres', host='13.233.155.60', port= '5432'
)

#Setting auto commit false
conn.autocommit = True

#Creating a cursor object using the cursor() method
cursor = conn.cursor()

#Retrieving data
cursor.execute('''SELECT * from mineuser''')

#Fetching 1st row from the table
result = cursor.fetchone();
print(result)

#Fetching 1st row from the table
result = cursor.fetchall();
print(result)

#Commit your changes in the database
conn.commit()

#Closing the connection
conn.close()