import React, {useState, useEffect} from 'react';

export const log = (key, state) => console.log(`${key}`, state);

export const useStateCallbackWrapper = (initilValue, callBack) => {
  const [state, setState] = useState(initilValue);

  useEffect(() => callBack(state), [state]);
  return [state, setState];
};
