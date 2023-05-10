import React, { FormEvent, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import authService from '../../services/auth.service';

interface LoginResponse {
  status: string;
  data: {
    email: string;
    role: string;
    token: string;
  };
}

const LoginPage = () => {
  const [loginEmail, setLoginEmail] = useState('');
  const [loginPassword, setLoginPassword] = useState('');
  const [logError, setLogError] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    const { request, cancel } = authService.login(loginEmail, loginPassword);
    request
      .then(({ data }: { data: LoginResponse }) => {
        localStorage.setItem('token', data.data.token);
        localStorage.setItem('role', data.data.role);
        localStorage.setItem('email', data.data.email);
        navigate('/');
      })
      .catch((err) => {
        console.error(err.message);
        setLogError(true);
      });

    return () => cancel();
  };

  return (
    <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
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
        <button type="submit" className="btn btn-primary mt-10" onSubmit={handleSubmit}>
          Login
        </button>
      </form>
      <Link to={'/Register'}>Rejestracja</Link>
      {logError && <p className="alert alert-danger">Błędne hasło lub Email</p>}
    </div>
  );
};

export default LoginPage;
