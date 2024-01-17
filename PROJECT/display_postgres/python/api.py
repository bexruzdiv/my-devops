from flask import Flask, request, jsonify
from flask_cors import CORS
import psycopg2
import os
import pandas as pd
from flask import render_template, jsonify
app = Flask(__name__)
CORS(app)

#################################
############    Functions
#################################

##################
### add user 
##################
def add_user_to_database(first_name, last_name, class_number):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="65.21.241.48", dbname="student", user="postgres", password="postgres")

    # Create a cursor object to execute SQL queries
    cur = connection.cursor()

    # SQL command to insert a new user record
    insert_command = """
        INSERT INTO student (first_name, last_name, class)
        VALUES (%s, %s, %s);
    """
    
    # Execute the SQL command with parameters
    cur.execute(insert_command, (first_name, last_name, class_number))

    # Commit the changes to the database
    connection.commit()

    # Close the cursor and connection
    cur.close()
    connection.close()



##################
### delete user
##################
def delete_with_name(first_name):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="65.21.241.48", dbname="student", user="postgres", password="postgres")

    # Create a cursor object to execute SQL queries
    cur = connection.cursor()

    # SQL command to update the class for the specified student
    delete_command = """
        DELETE FROM student WHERE first_name = %s;
    """
    
    # Execute the SQL command with parameters
    cur.execute(delete_command, (first_name,))

    # Commit the changes to the database
    connection.commit()

    # Close the cursor and connection
    cur.close()
    connection.close()


##################
### chnage class  
##################
def change_class(first_name, new_class):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="65.21.241.48", dbname="student", user="postgres", password="postgres")

    # Create a cursor object to execute SQL queries
    cur = connection.cursor()

    # SQL command to update the class for the specified student
    update_command = """
        UPDATE student
        SET class = %s
        WHERE first_name = %s;
    """
    
    # Execute the SQL command with parameters
    cur.execute(update_command, (new_class, first_name))

    # Commit the changes to the database
    connection.commit()

    # Close the cursor and connection
    cur.close()
    connection.close()





##################
### data  
##################
def get_all_data():
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="65.21.241.48", dbname="student", user="postgres", password="postgres")

    # Create a cursor object to execute SQL queries
    cur = connection.cursor()

    # SQL command to fetch all data from the "student" table
    select_all_data = """
        SELECT * FROM student;
    """
    cur.execute(select_all_data)
    # Fetch all rows
    rows = cur.fetchall()
    df = pd.DataFrame()
    for x in rows:
        df2=pd.DataFrame(list(x)).T
        df=pd.concat([df,df2])
    df.to_html('slq-data.html')

    cur.close()
    connection.close()




# API endpoint to print user information

@app.route('/change_class', methods=['POST'])
def print_user():
    data = request.json

    first_name = data['first_name']
    class_number = data['class_number']

    change_class(first_name, class_number)


######################
###  delete with name
######################
@app.route('/delete', methods=['POST'])
def delete_student():
    data = request.json

    first_name = data['first_name']

    delete_with_name(first_name)

    return jsonify({"message": "Record deleted successfully"})



####################
# add user 
####################
@app.route('/add', methods=['POST'])
def add_user_route():
    data = request.json

    first_name = data['first_name']
    last_name = data['last_name']
    class_number = data['class_number']
    
    add_user_to_database(first_name, last_name, class_number)

    return jsonify({"message": "Record added successfully"})


@app.route('/data', methods=['POST'])
def get_all_datas(): 
    get_all_data()
    return jsonify({"message": "Record added successfully"})
    # return render_template('sql-data.html')








if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)