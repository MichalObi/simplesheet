import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const useStateCallbackWrapper = (initilValue, callBack) => {
    const [state, setState] = useState(initilValue);

    useEffect(() => callBack(state), [state]);
    return [state, setState];
  };

  const log = (key, state) => console.log(`${key}`, state);

  const [error, setError] = useStateCallbackWrapper(null, state => log('error', state)),
        [isLoaded, setIsLoaded] = useStateCallbackWrapper(false,  state => log('isLoaded', state)),
        [sheetsAndGroups, setSheetsAndGroups] = useStateCallbackWrapper({}, state => log('sheetsAndGroups', state));

  const handleAjaxError = error => {
          setIsLoaded(true);
          setError(error);
        };

    useEffect(() => {
      fetch('/sheets/1')
      .then(sheets => sheets.json())
      .then(sheets => {
          if (sheets) {
            fetch(`/groups/${sheets.id}`)
            .then(groups => groups.json())
            .then(groups => {
              const sheetsAndGroups = Object.assign({}, sheets, groups);
              delete sheetsAndGroups.id;

              setSheetsAndGroups(sheetsAndGroups);
              setIsLoaded(true);
            }, handleAjaxError)
          }
        }, handleAjaxError)
    }, []);

  if (error) {
    return <p>error</p>;
  } else if (!isLoaded) {
    return <p>loading</p>;
  } else {
    return (
      <div className="App">
        <header className="App-header">
          Simple Sheet
        </header>
      </div>
    );
  }
}

export default App;
