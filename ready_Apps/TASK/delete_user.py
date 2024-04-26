import psycopg2

def delete_with_name(first_name):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="127.0.0.1", dbname="student", user="postgres", password="postgres")

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