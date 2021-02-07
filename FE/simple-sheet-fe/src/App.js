import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  const [error, setError] = useState(null),
        [isLoaded, setIsLoaded] = useState(false),
        [items, setItems] = useState([]),
        handleAjaxError = error => {
          setIsLoaded(true);
          setError(error);

          console.log('err', error);
        };

        useEffect(() => {
          fetch('/sheets/1')
          .then(res => res.json())
          .then(
            res => {
              console.log('sheets', res);

              fetch(`/groups/${res.id}`)
              .then(res => res.json())
              .then(res => {
                setIsLoaded(true);
                console.log('groups', res);
              }, handleAjaxError)
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
