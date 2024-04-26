import psycopg2

def change_class(first_name, new_class):
    # Establish a connection to the PostgreSQL database
    connection = psycopg2.connect(host="127.0.0.1", dbname="student", user="postgres", password="postgres")

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

# Example usage
change_class("Bexruz", 1)

