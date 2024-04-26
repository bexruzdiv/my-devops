from flask import Flask, request, jsonify
from flask_cors import CORS
from postgres_changer import change_class
from delete_user import delete_with_name
from add_users import add_user_to_database
from print_data import get_all_data
from flask import Flask, render_template, jsonify
app = Flask(__name__)
CORS(app)
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