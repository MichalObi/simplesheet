import React, {useState, useEffect} from 'react';
import Sheet from './Sheet'

function ExploreSheets() {

  const useStateCallbackWrapper = (initilValue, callBack) => {
    const [state, setState] = useState(initilValue);

    useEffect(() => callBack(state), [state]);
    return [state, setState];
  };

  const log = (key, state) => console.log(`${key}`, state);

  const [error, setError] = useStateCallbackWrapper(null, state => log('error', state)),
        [isLoaded, setIsLoaded] = useStateCallbackWrapper(false,  state => log('isLoaded', state)),
        [sheets, setSheets] = useStateCallbackWrapper([], state => log('sheets', state));

  const handleAjaxError = error => {
          setIsLoaded(true);
          setError(error);
        };

    useEffect(() => {
      fetch('/sheets')
      .then(sheets => sheets.json())
      .then(({sheets}) => {
        setSheets(sheets);
        setIsLoaded(true);
        }, handleAjaxError)
    }, []);

  useEffect(() => {
    setIsLoaded(true);
  }, []);

  if (error) {
    return <p>error</p>;
  } else if (!isLoaded) {
    return <p>loading</p>;
  } else {
    return (
      <div>
        <h2>Explore sheets</h2>
        {sheets.map((sheet, index) => (
          <Sheet key={index}
                 id={sheet.id}
                 type={sheet.type}
                 fields={sheet.fields} />))}
      </div>
    );
  }
}

export default ExploreSheets;
