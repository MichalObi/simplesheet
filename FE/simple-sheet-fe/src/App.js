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

export default App;
