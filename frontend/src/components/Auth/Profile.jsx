import React, { useEffect, useState } from 'react'
import axios from 'axios'

const Profile = () => {
  const [user, setUser] = useState(null)

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (!token) return
    axios.get(`${import.meta.env.VITE_GATEWAY_URL}/me`, {
      headers: { Authorization: `Bearer ${token}` }
    }).then(res => setUser(res.data))
      .catch(err => console.error(err))
  }, [])

  if (!user) return <p>Not logged in</p>

  return (
    <div>
      <h3>Welcome, {user.username}</h3>
      <p>User ID: {user.id}</p>
    </div>
  )
}

export default Profile
