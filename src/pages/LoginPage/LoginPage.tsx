import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { useAuthHelper } from '../../helpers/authHelper';
import { authAtom } from '../../states/auth';
import { useRecoilState } from 'recoil';
import axios from 'axios';

const backendURL = 'https://localhost/api/v1';

const LoginPage = () => {
  const [auth, setAuth] = useRecoilState(authAtom);
  const navigate = useNavigate();
  const [loginEmail, setLoginEmail] = useState('');
  const [loginPassword, setLoginPassword] = useState('');
  const [logError, setLogError] = useState(null);

  useEffect(() => {
    if (auth) navigate('/');
  });

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
          setAuth(res.data['token']);
          navigate('/');
        }
      })
      .catch((err) => {
        // Info if error
        setLogError(err.response.data.error);
      });
  };

  return (
    <div className="text-center float-start">
      <h3>Login</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">E-mail:</p>
          <br />
          <input
            type="text"
            className="form-control"
            placeholder="email"
            aria-label="email"
            aria-describedby="email-field"
            value={loginEmail}
            onChange={(event) => setLoginEmail(event.target.value)}
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
            aria-describedby="password-field"
            value={loginPassword}
            onChange={(event) => setLoginPassword(event.target.value)}
          />
        </label>
        <br />
        <button type="submit" className="btn btn-primary" onSubmit={handleSubmit}>
          Login
        </button>
      </form>
      <Link to={'/Register'}>Rejestracja</Link>
      <p className="text-danger">{logError}</p>
    </div>
  );
};

export default LoginPage;
