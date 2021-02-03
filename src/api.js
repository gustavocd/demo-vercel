import { useMutation, useQueryClient } from 'react-query'

export function useUserQueries({ queryKey = 'users-list' } = {}) {
  const queryClient = useQueryClient()

  const updateMutation = useMutation(
    id =>
      fetch(`/api/users/${id}`, {
        method: 'PUT',
        body: JSON.stringify({
          name: 'New name',
          email: 'New email' + Date.now(),
        }),
      }),
    {
      onSuccess: () => queryClient.fetchQuery(queryKey),
    }
  )

  const deleteMutation = useMutation(
    id =>
      fetch(`/api/users/${id}`, {
        method: 'DELETE',
      }),
    {
      onSuccess: () => queryClient.invalidateQueries(queryKey),
    }
  )

  const createMutation = useMutation(
    user =>
      fetch(`/api/users`, {
        method: 'POST',
        body: JSON.stringify(user),
      }),
    {
      onSuccess: () => queryClient.fetchQuery(queryKey),
    }
  )

  return {
    updateMutation,
    createMutation,
    deleteMutation,
  }
}

export async function fetchUsers() {
  const res = await fetch('/api/users')
  const data = await res.json()
  return data
}
