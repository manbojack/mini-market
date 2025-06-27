import React, { useState } from 'react'
import axios from 'axios'

const Register = () => {
  const [form, setForm] = useState({ username: '', password: '' })
  const [msg, setMsg] = useState(null)

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      const res = await axios.post(`${import.meta.env.VITE_GATEWAY_URL}/register`, form)
      setMsg(res.data.msg)
    } catch (err) {
      setMsg(err.response?.data?.msg || 'Error registering')
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <h3>Register</h3>
      <input name="username" placeholder="Username" onChange={handleChange} required />
      <input name="password" type="password" placeholder="Password" onChange={handleChange} required />
      <button type="submit">Register</button>
      {msg && <p>{msg}</p>}
    </form>
  )
}

export default Register
