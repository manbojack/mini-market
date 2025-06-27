import React, { useState } from 'react'
import axios from 'axios'

const Login = ({ onLogin }) => {
  const [form, setForm] = useState({ username: '', password: '' })
  const [msg, setMsg] = useState(null)

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      const res = await axios.post(`${import.meta.env.VITE_GATEWAY_URL}/login`, form)
      localStorage.setItem('token', res.data.access_token)
      setMsg("Login successful")
      onLogin?.()
    } catch (err) {
      setMsg(err.response?.data?.msg || 'Login failed')
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h3>Login</h3>
      <input name="username" placeholder="Username" onChange={handleChange} required />
      <input name="password" type="password" placeholder="Password" onChange={handleChange} required />
      <button type="submit">Login</button>
      {msg && <p>{msg}</p>}
    </form>
  )
}

export default Login
