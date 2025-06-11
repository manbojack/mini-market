from flask import Flask, jsonify, request
import psycopg2
import os

app = Flask(__name__)

DB_HOST = os.getenv('DB_HOST', 'postgres')
DB_PORT = os.getenv('DB_PORT', '5432')
DB_NAME = os.getenv('DB_NAME', 'mini_market')
DB_USER = os.getenv('DB_USER', 'user')
DB_PASSWORD = os.getenv('DB_PASSWORD', 'password')

def get_db_connection():
    conn = psycopg2.connect(
        host=DB_HOST,
        port=DB_PORT,
        dbname=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD
    )
    return conn

@app.route('/analytics/orders', methods=['GET'])
def get_order_analytics():
    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("""
        SELECT product_id, COUNT(*) AS total_orders
        FROM order_analytics
        GROUP BY product_id
        ORDER BY total_orders DESC;
    """)
    rows = cur.fetchall()
    cur.close()
    conn.close()

    analytics = [{'product_id': row[0], 'total_orders': row[1]} for row in rows]
    return jsonify(analytics)

@app.route('/analytics/orders', methods=['POST'])
def add_order_analytics():
    data = request.json
    product_id = data.get('product_id')
    if not product_id:
        return jsonify({'error': 'Missing product_id'}), 400

    conn = get_db_connection()
    cur = conn.cursor()
    cur.execute("""
        INSERT INTO order_analytics (product_id)
        VALUES (%s);
    """, (product_id,))
    conn.commit()
    cur.close()
    conn.close()

    return jsonify({'message': 'Analytics data inserted successfully'}), 201

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5004)
