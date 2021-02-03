export const formatDate = (date = Date.now()) => {
  const dtf = new Intl.DateTimeFormat('es', {
    timeStyle: 'short',
    hour12: true,
    dateStyle: 'long',
  })

  return dtf.format(new Date(date))
}
