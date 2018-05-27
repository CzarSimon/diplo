// postRequest() Creates and executes a post request
export const postRequest = (route, body = {}) => {
  console.log(postRequestObject(body));
  return fetch(route, postRequestObject(body))
  .then(checkReponse)
}

// postRequestJSON() Creates and executes a post request and parses JSON response
export const postRequestJSON = (route, body = {}) => (
  postRequest(route, body)
  .then(res => res.json())
);

// getRequest() Creates and executes a get request
export const getRequest = route => (
  fetch(route)
  .then(checkReponse)
  .then(res => res.json())
);

// deleteRequest() Creates and executes a delete request
export const deleteRequest = (route, body = {}) => (
  fetch(route, makeRequestObject('DELETE', body))
  .then(checkReponse)
);

/**
* postRequestObject() Returns a reqest object to be passed to fetch in
* order to make a post request
*/
export const postRequestObject = body => makeRequestObject('POST', body);

/**
* makeRequestObject() Returns a reqest object to be passed to fetch in
* order to make a request of the supplied method
*/
export const makeRequestObject = (method, body) => ({
  method,
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
  },
  body: JSON.stringify(body)
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