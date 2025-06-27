import React, { useEffect, useState } from 'react'
import axios from 'axios'

const gatewayUrl = import.meta.env.VITE_GATEWAY_URL || 'http://localhost:8080'

export default function ProductList() {
  const [products, setProducts] = useState([])

  useEffect(() => {
    axios.get(`${gatewayUrl}/products`)
      .then(res => setProducts(res.data))
      .catch(err => console.error('Failed to load products', err))
  }, [])

  return (
    <div>
      <h2>Products</h2>
      {products.length === 0 ? (
        <p>No products available.</p>
      ) : (
        <ul>
          {products.map(p => (
            <li key={p.id}>
              <strong>{p.name}</strong> — ${p.price}
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
