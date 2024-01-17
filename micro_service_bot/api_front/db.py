from pymongo import MongoClient

def db_connect():
    # Replace 'your_username' and 'your_password' with your MongoDB credentials
    mongo = MongoClient("mongodb://root:example@mongo/?retryWrites=true&w=majority", port=27017)

    print(f"Connected to the MongoDB database!")
    return mongo

def db_close(db):
    db.close()




def get_all_chat_ids(database, collection):
    connection = db_connect()
    db = connection[database]

    try:
        # Retrieve all unique chat_id values from the collection
        chat_ids = db[collection].distinct("chat_id")
        print("DEBUG - All Chat IDs:", chat_ids)
        return chat_ids
    except Exception as e:
        print(f"Error retrieving chat_ids: {e}")
    finally:
        db_close(connection)