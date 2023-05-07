import React, { FormEvent, useEffect, useState } from 'react';
import { useRecoilValue } from 'recoil';
import { currentUserAtom } from '../../states/auth.state';
import { Link, useNavigate } from 'react-router-dom';
import authService, { AuthResponse } from '../../services/auth.service';
import { Notify } from 'notiflix';
import { AxiosError } from 'axios';

function RegisterPage() {
  const [regEmail, setRegEmail] = useState('');
  const [regPassword, setRegPassword] = useState('');
  const [regPasswordRepeat, setRegPasswordRepeat] = useState('');

  const currentUser = useRecoilValue(currentUserAtom);
  const navigate = useNavigate();

  useEffect(() => {
    if (currentUser.token) navigate('/');
  }, [currentUser, navigate]);

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();

    const regex = /^\S+@\S+\.\S+$/;
    const valid = regex.test(regEmail);
    if (!valid) {
      Notify.failure('Email is invalid');
      return;
    }

    if (regPassword.length < 8) {
      Notify.failure('Password is too simple');
      return;
    }

    if (regPassword != regPasswordRepeat) {
      Notify.failure("Passwords doesn't match");
      return;
    }

    authService
      .register(regEmail, regPassword, regPasswordRepeat)
      .then(() => {
        Notify.success('Account successfuly registered!');
        navigate('/');
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
    </div>
  );
}

export default RegisterPage;
