import React from 'react';

function Sheet({id, type, fields}) {
    return (
      <div>
        <div>Numer arkusza {id}</div>
        <div>Typ pozycji {type}</div>
        <ul>
          {fields.map(({field, value}, index) =>
            <li key={index}>{field}: {value}</li>)}
        </ul>
      </div>
    )
}

export default Sheet;
