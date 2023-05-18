import React, { FormEvent, useState } from 'react';
import { Link } from 'react-router-dom';

export const CodePage = () => {
  const [isValid, setIsValid] = useState(true);
  const [isDisabled, setIsDisabled] = useState<boolean>(false);
  const [value, setValue] = useState('');

  const handleChange = (e: { target: { value: any } }) => {
    const inputValue = e.target.value;
    const numericValue = inputValue.replace(/[^0-9]/g, ''); // Usuwa wszystko oprócz cyfr
    setValue(numericValue);
  };
  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    setIsDisabled(true);
  };
  return (
    <>
      <div className="d-flex justify-content-center">
        <div className="position-absolute top-50 start-50 translate-middle text-center shadow-lg p-3 mb-5 bg-white rounded">
          <form className="" onSubmit={handleSubmit}>
            {' '}
            <label>
              <h3>Enter Code</h3>
              <p> Please enter your verification code</p>
              <p style={{ float: 'left', marginTop: '3px', marginLeft: '3vw', fontSize: '18px' }}> Code:</p>
              <input
                disabled={isDisabled}
                className="form-control p-1 Myone"
                style={{ width: '5vw' }}
                type="text"
                maxLength={5}
                value={value}
                onChange={handleChange}
              />
            </label>
            <p></p>
            <button type="submit" className="btn btn-primary " disabled={isDisabled} style={{ marginBottom: '10px' }}>
              Verified
            </button>
          </form>
        </div>
      </div>
    </>
  );
};
