import React, { FormEvent, useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import authService from '../../services/auth.service';
import { AuthContext } from '../../context/AuthContext';

interface LoginResponse {
  status: string;
  data: {
    email: string;
    role: string;
    token: string;
    verified: boolean;
    banned: boolean;
  };
}

const LoginPage = () => {
  const { setUser } = useContext(AuthContext);

  const [loginEmail, setLoginEmail] = useState('');
  const [loginPassword, setLoginPassword] = useState('');
  const [logError, setLogError] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    const { request, cancel } = authService.login(loginEmail, loginPassword);
    request
      .then(({ data }: { data: LoginResponse }) => {
        const user = {
          email: data.data.email,
          role: data.data.role,
          verified: Boolean(data.data.verified),
          authToken: data.data.token,
          banned: data.data.banned
        };
        setUser(user);
        localStorage.setItem('user', JSON.stringify(user));
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
      {logError && <p className="alert alert-danger">Wrong password or email</p>}
    </div>
  );
};

export default LoginPage;
