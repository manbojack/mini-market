from flask import Flask, request, jsonify
import psycopg2
import os

app = Flask(__name__)

# Читаем переменные окружения из Docker Compose
DB_HOST = os.getenv("DB_HOST", "postgres")
DB_PORT = os.getenv("DB_PORT", 5432)
DB_NAME = os.getenv("DB_NAME", "mini_market")
DB_USER = os.getenv("DB_USER", "user")
DB_PASSWORD = os.getenv("DB_PASSWORD", "password")

def get_db_connection():
    conn = psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        dbname=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD
    )
    return conn

@app.route('/cart/<int:user_id>', methods=['GET'])
def get_cart(user_id):
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute('SELECT product_id, quantity FROM cart WHERE user_id = %s', (user_id,))
    items = cur.fetchall()
    cur.close()
    conn.close()
    cart = [{"product_id": row[0], "quantity": row[1]} for row in items]
    return jsonify(cart)

@app.route('/cart/<int:user_id>', methods=['POST'])
def add_to_cart(user_id):
    data = request.json
    product_id = data['product_id']
    quantity = data['quantity']

    conn = get_db_connection()
    cur = conn.cursor()
    # Проверим, есть ли запись
    cur.execute('SELECT id FROM cart WHERE user_id = %s AND product_id = %s', (user_id, product_id))
    row = cur.fetchone()
    if row:
        cur.execute('UPDATE cart SET quantity = quantity + %s WHERE id = %s', (quantity, row[0]))
    else:
        cur.execute('INSERT INTO cart (user_id, product_id, quantity) VALUES (%s, %s, %s)', (user_id, product_id, quantity))
    conn.commit()
    cur.close()
    conn.close()

    return jsonify({"message": "Product added to cart"}), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
