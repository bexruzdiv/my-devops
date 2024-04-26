import psycopg2

def add_user_to_database(first_name, last_name, class_number):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="127.0.0.1", dbname="student", user="postgres", password="postgres")

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