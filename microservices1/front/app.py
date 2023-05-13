from flask import Flask, render_template
app = Flask(__name__, template_folder='templates')

@app.route('/')
def auth():
    return render_template('auth.html')
@app.route('/register')
def register():
    return render_template('register.html')