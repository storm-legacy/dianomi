import React, { useEffect, useState, FormEvent } from 'react';
import { Notify } from 'notiflix/build/notiflix-notify-aio';
import { Link, useNavigate } from 'react-router-dom';
import { useRecoilState } from 'recoil';
import { currentUserAtom } from '../../states/auth.state';
import authService, { AuthResponse } from '../../services/auth.service';
import { AxiosError } from 'axios';

const LoginPage = () => {
  const [currentUser, setCurrentUser] = useRecoilState(currentUserAtom);
  const navigate = useNavigate();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  useEffect(() => {
    if (currentUser.token) navigate('/');
  }, [currentUser, navigate]);

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    authService
      .login(email, password)
      .then(({ data: data }) => {
        setCurrentUser(data.data);
        localStorage.setItem('currentUser', JSON.stringify(data.data));
      })
      .catch((err: AxiosError) => {
        const message = err.response?.data as AuthResponse;
        console.error(err);
        if (err.response?.status == 500) Notify.failure('Service is currently unavailable');
        Notify.failure(`${message.data}`);
      });
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
            value={email}
            onChange={(event) => setEmail(event.target.value)}
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
            value={password}
            onChange={(event) => setPassword(event.target.value)}
          />
        </label>
        <br />
        <button type="submit" className="btn btn-primary mt-10">
          Login
        </button>
      </form>
      <Link to={'/Register'}>Rejestracja</Link>
    </div>
  );
};

export default LoginPage;
