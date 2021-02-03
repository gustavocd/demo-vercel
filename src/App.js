import React from 'react'
import { useQuery } from 'react-query'
import { useState } from 'react'

import { formatDate } from './util'
import { useUserQueries, fetchUsers } from './api'

function App() {
  const { updateMutation, deleteMutation, createMutation } = useUserQueries()
  const { data: users = [], isLoading } = useQuery('users-list', fetchUsers, {
    refetchOnWindowFocus: false,
  })
  const [user, setUser] = useState({ name: '', email: '' })

  async function handleGreet(user) {
    const res = await fetch(`/api/users/${user.id}`)
    const data = await res.json()
    alert(JSON.stringify(data, null, '  '))
  }

  async function handleUpdate(id) {
    updateMutation.mutate(id)
  }

  async function handleDelete(id) {
    deleteMutation.mutate(id)
  }

  function handleChange(event) {
    const { name, value } = event.target
    setUser(state => ({
      ...state,
      [name]: value,
    }))
  }

  async function handleSubmit(event) {
    event.preventDefault()
    createMutation.mutate(user, {
      onSuccess: () => {
        setUser({ name: '', email: '' })
      },
    })
  }

  return (
    <main className="mx-auto max-w-4xl px-4 py-8">
      <form
        onSubmit={handleSubmit}
        className="max-w-lg mx-auto bg-white p-8 shadow-lg rounded-lg mb-16"
      >
        <h1
          className="text-center
        text-3xl mb-8 uppercase tracking-wide font-bold text-blue-900"
        >
          Crear usuario
        </h1>
        <div className="mb-8">
          <label className="block mb-2 font-bold text-lg" htmlFor="name">
            Tu nombre:
          </label>
          <input
            type="text"
            name="name"
            id="name"
            className="py-2 px-4 w-full border border-gray-400 rounded-lg"
            placeholder="Gustavo Castillo"
            value={user.name}
            onChange={handleChange}
          />
        </div>
        <div className="mb-8">
          <label className="block mb-2 font-bold text-lg" htmlFor="email">
            Tu correo:
          </label>
          <input
            type="email"
            name="email"
            id="email"
            className="py-2 px-4 w-full border border-gray-400 rounded-lg"
            placeholder="email@hey.com"
            value={user.email}
            onChange={handleChange}
          />
        </div>
        <button
          className="bg-blue-400 text-white font-bold tracking-wide uppercase py-2 px-8 rounded-lg w-full hover:shadow-lg hover:bg-blue-500 transition"
          type="submit"
        >
          Enviar
        </button>
      </form>
      {isLoading && (
        <p className="text-center">Cargando lista de usuarios...</p>
      )}
      {users.length > 0 ? (
        <table className="w-full bg-white p-4">
          <thead>
            <tr>
              <th className="p-2 border border-gray-200 uppercase">Nombre</th>
              <th className="p-2 border border-gray-200 uppercase">Correo</th>
              <th className="p-2 border border-gray-200 uppercase">Fecha</th>
              <th className="p-2 border border-gray-200 uppercase">Acciones</th>
            </tr>
          </thead>
          <tbody>
            {users.map(user => (
              <tr key={user.id}>
                <td className="p-2 border border-gray-200">{user.name}</td>
                <td className="p-2 border border-gray-200">{user.email}</td>
                <td className="p-2 border border-gray-200">
                  {formatDate(user.created_at)}
                </td>
                <td className="p-2 border border-gray-200">
                  <div className="flex justify-between flex-col md:flex-row">
                    <button onClick={() => handleGreet(user)}>
                      <span role="img" aria-label="add details">
                        👀
                      </span>
                    </button>
                    <button onClick={() => handleDelete(user.id)}>
                      <span role="img" aria-label="delete user">
                        🗑
                      </span>
                    </button>
                    <button onClick={() => handleUpdate(user.id)}>
                      <span role="img" aria-label="edit user">
                        📝
                      </span>
                    </button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : null}
    </main>
  )
}

export default App
