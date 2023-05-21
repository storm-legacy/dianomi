import React, { useState, FormEvent } from 'react';
import { Link } from 'react-router-dom';
import authService from '../../services/auth.service';

export const PasswordResPage = () => {
  const [isValid, setIsValid] = useState(true);
  const [emailExist, setEmailExist] = useState<boolean>(false);
  const [email, setEmail] = useState<string>('');
  const [isDisabled, setIsDisabled] = useState<boolean>(false);
  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    if (!/^\S+@\S+\.\S+$/.test(email)) {
      setIsValid(false);
    }
    const { request } = authService.genreset(email);
    request
      .then((res) => {
        console.log(res);
        setEmailExist(true);
      })
      .catch((err) => {
        console.error(err);
      });
    setIsDisabled(true);
    console.log(email);
  };
  return (
    <>
      <div className="text-center mylogin">
        <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
          <form className="row" onSubmit={handleSubmit}>
            {' '}
            <label>
              <h3>Password Reset</h3>
              All your base are belong to us
              <div className="text-left m-3">
                {' '}
                Enter your email address below <br />
                to send a password reset email{' '}
              </div>
              <input
                disabled={isDisabled}
                className="form-control"
                type="text"
                value={email}
                onChange={(event) => setEmail(event.target.value)}
              />
            </label>
            <p></p>
            <button type="submit" className="btn btn-primary " disabled={isDisabled} style={{ marginBottom: '10px' }}>
              Send
            </button>
          </form>
          {!isValid && (
            <p className="alert alert-danger" role="alert">
              Please enter a valid email.
            </p>
          )}
          {emailExist && (
            <div className="alert alert-success" role="alert">
              E-mail has been sent to your address if he exist <br /> Go back to <Link to={'/login'}>Login page</Link>
            </div>
          )}
        </div>
      </div>
    </>
  );
};
