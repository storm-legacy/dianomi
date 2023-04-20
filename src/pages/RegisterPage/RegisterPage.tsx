import axios from 'axios';
import { Validator, useEffect } from 'react';
import React, { useState } from 'react';
import { render } from 'react-dom';

import { Link, redirect, useNavigate } from 'react-router-dom';

const backendURL = 'https://localhost/api/v1';
function RegisterPage() {
  const [regEmail, setRegEmail] = useState('');
  const [regPassword, setRegPassword] = useState('');
  const [regPasswordRepeat, setRegPasswordRepeat] = useState('');
  const [regError, setRegError] = useState(null);
  const [isValid, setIsValid] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    const validateEmail = () => {
      const regex = /^\S+@\S+\.\S+$/;
      const valid = regex.test(regEmail);
      setIsValid(valid);
    };

    validateEmail();
  }, [regEmail]);

  const handleSubmit = (event: any) => {
    event.preventDefault();

    const userData = {
      email: regEmail,
      password: regPassword,
      password_repeat: regPasswordRepeat,
    };

    axios
      .post(`${backendURL}/auth/register`, userData)
      .then((res) => {
        if (res.status === 200) {
          console.log(res);
          alert('Rejestracja powiodła się');
          navigate('/login?registerSuccess=true');
        }
      })
      .catch((err) => {
        // Info if error

        setRegError(err.response.data.error);
        console.log(regError);
      });
  };

  return (
    <div className="text-center float-start shadow-lg p-3 mb-5 bg-white rounded" style={{ marginLeft: 50 }}>
      <h3>Rejestracja</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">E-mail:</p>
          <input
            type="text"
            className="form-control"
            placeholder="E-mail"
            aria-label="E-mail"
            aria-describedby="basic-addon1"
            value={regEmail}
            onChange={(event) => setRegEmail(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Password:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            aria-label="Password"
            aria-describedby="basic-addon1"
            value={regPassword}
            onChange={(event) => setRegPassword(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Repeat password:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Repeat password"
            aria-label="PasswordRepeat"
            aria-describedby="basic-addon1"
            value={regPasswordRepeat}
            onChange={(event) => setRegPasswordRepeat(event.target.value)}
          />
        </label>

        <br />
        <button type="submit" className="btn btn-primary">
          Zarejestruj
        </button>
      </form>
      <Link to={'/login'}>Logowanie</Link>
      {!isValid && <p className="alert alert-danger">Wprowadź poprawny email.</p>}
      {regPassword != regPasswordRepeat && regPasswordRepeat != '' && (
        <p className="alert alert-danger">
          {' '}
          Hasłą nie są identyczne <br />{' '}
        </p>
      )}
    </div>
  );
}

export default RegisterPage;
