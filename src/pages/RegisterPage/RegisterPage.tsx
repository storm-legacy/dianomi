import React, { useState, FormEvent } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import authService from '../../services/auth.service';

function RegisterPage() {
  const [regEmail, setRegEmail] = useState('');
  const [regPassword, setRegPassword] = useState('');
  const [regPasswordRepeat, setRegPasswordRepeat] = useState('');
  const [isValid, setIsValid] = useState(true);
  const navigate = useNavigate();

  const handleSubmit = (event: FormEvent) => {
    event.preventDefault();
    if (!/^\S+@\S+\.\S+$/.test(regEmail)) {
      setIsValid(false);
    }

    const { request } = authService.register(regEmail, regPassword, regPasswordRepeat);
    request
      .then(() => {
        alert('Account registered');
        navigate('/');
      })
      .catch((err) => {
        console.error(err.message);
        setIsValid(false);
      });
  };

  return (
    <div className="Mylogin position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
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
          Create
        </button>
      </form>
      Already have an account? <Link to={'/login'}>Log in</Link>
      {!isValid && <p className="alert alert-danger">Please enter a valid email.</p>}
      {regPassword != regPasswordRepeat && regPasswordRepeat != '' && (
        <p className="alert alert-danger" style={{ marginTop: '5px' }}>
          {' '}
          The passwords are not identical <br />{' '}
        </p>
      )}
    </div>
  );
}

export default RegisterPage;
