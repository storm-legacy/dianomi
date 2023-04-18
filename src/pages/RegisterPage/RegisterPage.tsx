import React, { useState } from 'react';
import { Link } from 'react-router-dom';

function RegisterPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [RepeatPassword, setRepeatPassword] = useState('');
  const [email, setEmail] = useState('');

  function handleSubmit(event: any) {
    event.preventDefault();
    if (RepeatPassword == password) console.log('Zarejestrowany użytkownik: ', username, email, password);
    else console.log('błędne hasło');
  }

  return (
    <div className="text-center float-start" style={{ marginLeft: 50 }}>
      <h3>Rejestracja</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">Nazwa użytkownika:</p>
          <input
            type="text"
            className="form-control"
            placeholder="Username"
            aria-label="Username"
            aria-describedby="basic-addon1"
            value={username}
            onChange={(event) => setUsername(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Adres email:</p>
          <input
            type="email"
            className="form-control"
            aria-describedby="emailHelp"
            placeholder="Enter Email"
            value={email}
            onChange={(event) => setEmail(event.target.value)}
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
            value={password}
            onChange={(event) => setPassword(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Powtórz Hasło:</p>
          <input
            type="RepeatPassword"
            className="form-control"
            placeholder="RepeatPassword"
            aria-label="RepeatPassword"
            aria-describedby="basic-addon1"
            value={RepeatPassword}
            onChange={(event) => setRepeatPassword(event.target.value)}
          />
        </label>
        <br />
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
