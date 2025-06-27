import React, { useState } from 'react'
import Register from './components/Auth/Register'
import Login from './components/Auth/Login'
import Profile from './components/Auth/Profile'
import ProductList from './components/ProductList'

const App = () => {
  const [reload, setReload] = useState(false)

  const handleLogin = () => setReload(!reload)

  return (
    <div style={{ padding: '2rem', fontFamily: 'sans-serif' }}>
      <h1>🛍️ Mini Market</h1>
      <Register />
      <Login onLogin={handleLogin} />
      <Profile key={reload} />
      <hr />
      <ProductList />
    </div>
  )
}

export default App
