import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import CreateSheet from './components/CreateSheet';
import ExploreSheets from './components/ExploreSheets';
import Home from './components/Home';

function App() {
  const useStateCallbackWrapper = (initilValue, callBack) => {
    const [state, setState] = useState(initilValue);

    useEffect(() => callBack(state), [state]);
    return [state, setState];
  };

  const log = (key, state) => console.log(`${key}`, state);

  const [error, setError] = useStateCallbackWrapper(null, state => log('error', state)),
        [isLoaded, setIsLoaded] = useStateCallbackWrapper(false,  state => log('isLoaded', state)),
        [sheets, setSheets] = useStateCallbackWrapper({}, state => log('sheets', state));

  const handleAjaxError = error => {
          setIsLoaded(true);
          setError(error);
        };

    // useEffect(() => {
    //   fetch('/sheets/')
    //   .then(sheets => sheets.json())
    //   .then(sheets => {
    //     setSheets(sheets);
    //     setIsLoaded(true);
    //     }, handleAjaxError)
    // }, []);

  useEffect(() => {
    setIsLoaded(true);
  }, []);

  if (error) {
    return <p>error</p>;
  } else if (!isLoaded) {
    return <p>loading</p>;
  } else {
    return (
      <Router>
        <div className="App">
          <header className="App-header">
            <nav>
              <ul>
                <li>
                  <Link to="/">Home</Link>
                </li>
                <li>
                  <Link to="/create">Create new</Link>
                </li>
                <li>
                  <Link to="/explore">Explore sheets</Link>
                </li>
              </ul>
            </nav>
          </header>
        </div>

        <Switch>
          <Route path="/create">
            <CreateSheet />
          </Route>
          <Route path="/explore">
            <ExploreSheets />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </Router>
    );
  }
}

export default App;
