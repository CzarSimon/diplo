import store from 'store';

export const save = (key, value) => {
  try {
    store.set(key, value);
    console.log(`${key} was saved`);
  } catch (err) {
    console.log(err);
  }
}

export const get = key => {
  try {
    console.log(`${key} is being retrieved`);
    return store.get(key);
  } catch (err) {
    console.log(err);
    return undefined;
  }
}

export const remove = key => {
  try {
    store.remove(key);
  } catch (err) {
    console.log(err);
  }
}
