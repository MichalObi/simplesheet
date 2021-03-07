import React, {useState, useEffect} from 'react';

function CreateSheet() {

  const [type, setType] = useState(null);

  const handleSelectType = ({target: {value}}) => setType(value);

  return (
    <div>
      <h2>Create Sheet</h2>
      <form className="create-sheet">

        <div>
          <h4>Type of sheet</h4>
          <select id="type" name="type" defaultValue=""
            onChange={e => handleSelectType(e)}>
            <option hidden value="">select your option</option>
            <option name="crypto">Crypto</option>
            <option name="metals">Metals</option>
            <option name="etf">ETF</option>
          </select>
        </div>

        <div>
          <h4>Crypto fields</h4>
          <select id="crypto-fields" name="crypto-fields">
            <option name="btc">BTC</option>
            <option name="eth">ETH</option>
            <option name="link">Chainlink</option>
          </select>
        </div>

        <div>
          <h4>Metals fields</h4>
          <select id="metals-fields" name="metals-fields">
            <option name="gold">Gold</option>
            <option name="silver">Silver</option>
            <option name="platinum">Platinum</option>
          </select>
        </div>

        <div>
          <h4>ETF fields</h4>
          <select id="etf-fields" name="etf-fields">
            <option name="gdxj-uk">GDXJ.UK</option>
            <option name="ipol-uk">IPOL.UK</option>
            <option name="vix">VIX</option>
          </select>
        </div>
      </form>
    </div>
  );
}

export default CreateSheet;
