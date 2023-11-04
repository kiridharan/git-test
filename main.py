from flask import Flask, render_template_string, request, jsonify

app = Flask(__name__)

# Data storage for the todos
todos = []

html_template = """
<!DOCTYPE html>
<html>
<head>
    <title>Todo List</title>
</head>
<body>
    <h1>Todo List</h1>
    # <form id="todoForm" action="/todos" method="post">
    #     <input type="text" name="title" placeholder="Enter a new todo" />
    #     <button type="submit">Add</button>
    # </form>
    
    # <h2>New Todos:</h2>
    # <ul id="newTodos">
    #     {% for todo in todos %}
    #     <li>{{ todo['title'] }}</li>
    #     {% endfor %}
    # </ul>
</body>
</html>
"""


@app.route('/', methods=['GET'])
def index():
    return render_template_string(html_template, todos=todos)


# @app.route('/todos', methods=['GET'])
# def get_todos():
#     return jsonify(todos)


# @app.route('/todos', methods=['POST'])
# def create_todo():
#     data = request.get_json()
#     title = data.get('title')
#     if title:
#         todo = {
#             'id': len(todos) + 1,
#             'title': title,
#             'done': False
#         }
#         todos.append(todo)
#         return jsonify(todo), 201
#     else:
#         return 'Bad Request', 400


if __name__ == '__main__':
    app.run(debug=True)
