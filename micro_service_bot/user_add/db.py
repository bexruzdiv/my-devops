from pymongo import MongoClient

def db_connect():
    # Replace 'your_username' and 'your_password' with your MongoDB credentials
    mongo = MongoClient("mongodb://root:example@mongo/?retryWrites=true&w=majority", port=27017)

    print(f"Connected to the MongoDB database!")
    return mongo

def db_close(db):
    db.close()

def insert_one_to_db(database, collection, data):
    connection = db_connect()
    db = connection[database]
    response = db[collection].insert_one(data)
    db_close(connection)
    return response.inserted_id



