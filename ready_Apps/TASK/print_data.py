import psycopg2
import os
import pandas as pd

def get_all_data():
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="127.0.0.1", dbname="student", user="postgres", password="postgres")

    # Create a cursor object to execute SQL queries
    cur = connection.cursor()

    # SQL command to fetch all data from the "student" table
    select_all_data = """
        SELECT * FROM student;
    """
    
    # Execute the SQL command
    cur.execute(select_all_data)

    # Fetch all rows
    rows = cur.fetchall()
    df = pd.DataFrame()
    for x in rows:
        df2=pd.DataFrame(list(x)).T
        df=pd.concat([df,df2])
    df.to_html('slq-data.html')

    # Close the cursor and connection
    cur.close()
    connection.close()

    # Convert the result to a list of dictionaries for JSON response
    # result = []
    # for row in rows:
    #     result.append({
    #         'first_name': row[0],
    #         'last_name': row[1],
    #         'class': row[2]
    #     })

    #     # Print the current row
    #     # print(result[-1])
    #     # print(result)
    # return result

# Call the function and print and return data line by line
result = get_all_data()

