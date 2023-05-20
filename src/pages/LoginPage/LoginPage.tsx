import React, { FormEvent, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import authService from '../../services/auth.service';

interface LoginResponse {
  status: string;
  data: {
    email: string;
    role: string;
    token: string;
    verified: boolean;
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
        console.log(data.data.token);
        localStorage.setItem('token', data.data.token);
        localStorage.setItem('role', data.data.role);
        localStorage.setItem('email', data.data.email);
        localStorage.setItem('verified', String(data.data.verified));
        navigate('/');
      })
      .catch((err) => {
        console.error(err.message);
        setLogError(true);
      });

    return () => cancel();
  };

  return (
    <div className="Mylogin position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
      <form onSubmit={handleSubmit}>
        <h3>Login</h3>
        <label>
          <p className="h5" style={{ marginTop: '15px' }}>
            E-mail:
          </p>
          <input
            type="text"
            style={{ marginTop: '15px' }}
            className="form-control "
            placeholder="email"
            aria-label="email"
            aria-describedby="email-field"
            value={loginEmail}
            onChange={(event) => setLoginEmail(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5" style={{ marginTop: '15px' }}>
            Password:
          </p>
          <input
            type="password"
            style={{ marginTop: '15px' }}
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
      <br />
      You do not have an account ? <Link to={'/Register'}>Create it</Link>
      <br />
      <Link to={'/password/reset'}>Forgot your password?</Link>
      <Link to="/verificate">code test</Link>
      {logError && <p className="alert alert-danger">Błędne hasło lub Email</p>}
    </div>
  );
};

export default LoginPage;
