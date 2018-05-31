// postRequest() Creates and executes a post request
export const postRequest = (route, token, body = {}) => {
  return fetch(route, postRequestObject(body, token))
  .then(checkReponse)
}

// postRequestJSON() Creates and executes a post request and parses JSON response
export const postRequestJSON = (route, token, body = {}) => (
  postRequest(route, token, body)
  .then(res => res.json())
);

// getRequest() Creates and executes a get request
export const getRequest = (route, token) => (
  fetch(route, makeRequestObject('GET', token))
  .then(checkReponse)
  .then(res => res.json())
);

// deleteRequest() Creates and executes a delete request
export const deleteRequest = (route, token, body = {}) => (
  fetch(route, makeRequestObjectWithBody('DELETE', token, body))
  .then(checkReponse)
);

/**
* postRequestObject() Returns a reqest object to be passed to fetch in
* order to make a post request
*/
export const postRequestObject = (body, token) => makeRequestObjectWithBody('POST', token, body);

/**
* makeRequestObjectWithBody() Returns a reqest object to be passed to fetch in
* order to make a request of the supplied method
*/
export const makeRequestObjectWithBody = (method, token, body) => ({
  ...makeRequestObject(method, token),
  body: JSON.stringify(body)
})

/**
* makeRequestObject() Creates a request object with HTTP method and headers.
*/
export const makeRequestObject = (method, token) => ({
  method,
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${token}`,
  },
})

// checkReponse() Checks whether a fetch response was ok, throws an error if not
export const checkReponse = response => {
  if (response.ok) {
    console.log(response);
    return response
  } else {
    let error = new Error(response.statusText);
    error.response = response;
    throw error;
  }
};
