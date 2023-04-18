import axios from 'axios';
import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';

const backendURL = 'https://localhost/api/v1';

const LoginPage = () => {
  const navigate = useNavigate();
  const [loginEmail, setLoginEmail] = useState('');
  const [loginPassword, setLoginPassword] = useState('');

  const handleSubmit = (event: any) => {
    event.preventDefault();

    const userData = {
      email: loginEmail,
      password: loginPassword,
    };

    axios
      .post(`${backendURL}/auth/login`, userData)
      .then((res) => {
        if (res.status === 200) {
          localStorage.setItem('accessToken', res.data['token']);
          navigate('/');
        }
      })
      .catch((err) => {
        // Info if error
        console.error(err.response.data.error);
        console.error(err.response);
      });
  };

  return (
    <div className="text-center float-start">
      <h3>Logowanie</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">Nazwa użytkownika:</p>
          <br />
          <input
            type="text"
            className="form-control"
            placeholder="UserName"
            aria-label="UserName"
            aria-describedby="basic-addon1"
            value={loginEmail}
            onChange={(event) => setLoginEmail(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Hasło:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            aria-label="Password"
            aria-describedby="basic-addon1"
            value={loginPassword}
            onChange={(event) => setLoginPassword(event.target.value)}
          />
        </label>
        <br />
        <button type="submit" className="btn btn-primary" onSubmit={handleSubmit}>
          Zaloguj
        </button>
      </form>
      <Link to={'/Register'}>Rejestracja</Link>
    </div>
  );
};

export default LoginPage;
