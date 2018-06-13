export const isValidToken = token => (
  token !== null && token !== undefined && token.split('.').length === 3
)
